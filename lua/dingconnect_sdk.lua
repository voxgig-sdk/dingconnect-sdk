-- Dingconnect SDK

local vs = require("utility.struct.struct")
local Utility = require("core.utility_type")
local Spec = require("core.spec")
local helpers = require("core.helpers")

-- Load utility registration (populates Utility._registrar)
require("utility.register")

-- Typed-model annotations (LuaLS ---@class); empty at runtime.
require("dingconnect_types")

-- Load features
local BaseFeature = require("feature.base_feature")
local features_factory = require("features")


local DingconnectSDK = {}
DingconnectSDK.__index = DingconnectSDK


local function _make_feature(name)
  local factory = features_factory[name]
  if factory ~= nil then
    return factory()
  end
  return features_factory.base()
end

DingconnectSDK._make_feature = _make_feature


function DingconnectSDK.new(options)
  local self = setmetatable({}, DingconnectSDK)
  self.mode = "live"
  self.features = {}
  self.options = nil

  local utility = Utility.new()
  self._utility = utility

  local config = require("config")()

  self._rootctx = utility.make_context({
    client = self,
    utility = utility,
    config = config,
    options = options or {},
    shared = {},
  }, nil)

  self.options = utility.make_options(self._rootctx)

  if vs.getpath(self.options, "feature.test.active") == true then
    self.mode = "test"
  end

  self._rootctx.options = self.options

  -- Add features in the resolved order (make_options puts an explicit list
  -- order first, else defaults to test-first). Ordering matters: the `test`
  -- feature installs the base mock transport and the transport features
  -- (retry/cache/netsim/proxy/ratelimit) wrap whatever is current, so `test`
  -- must be added before them to sit at the base of the chain.
  local feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
  if feature_opts ~= nil then
    local featureorder = vs.getpath(self.options, "__derived__.featureorder")
    if type(featureorder) == "table" then
      for _, fname in ipairs(featureorder) do
        local fopts = helpers.to_map(feature_opts[fname])
        if fopts ~= nil and fopts["active"] == true then
          utility.feature_add(self._rootctx, _make_feature(fname))
        end
      end
    end
  end

  -- Add extension features.
  local extend = vs.getprop(self.options, "extend")
  if type(extend) == "table" then
    for _, f in ipairs(extend) do
      if type(f) == "table" and type(f.get_name) == "function" then
        utility.feature_add(self._rootctx, f)
      end
    end
  end

  -- Initialize features.
  for _, f in ipairs(self.features) do
    utility.feature_init(self._rootctx, f)
  end

  utility.feature_hook(self._rootctx, "PostConstruct")

    -- feature: test


  return self
end


function DingconnectSDK:options_map()
  local out = vs.clone(self.options)
  if type(out) == "table" then
    return out
  end
  return {}
end


function DingconnectSDK:get_utility()
  return Utility.copy(self._utility)
end


function DingconnectSDK:get_root_ctx()
  return self._rootctx
end


function DingconnectSDK:prepare(fetchargs)
  local utility = self._utility

  fetchargs = fetchargs or {}

  local ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl")) or {}

  local ctx = utility.make_context({
    opname = "prepare",
    ctrl = ctrl,
  }, self._rootctx)

  local options = self.options

  local path = vs.getprop(fetchargs, "path") or ""
  if type(path) ~= "string" then path = "" end

  local method = vs.getprop(fetchargs, "method") or "GET"
  if type(method) ~= "string" then method = "GET" end

  local params = helpers.to_map(vs.getprop(fetchargs, "params")) or {}
  local query = helpers.to_map(vs.getprop(fetchargs, "query")) or {}

  local headers = utility.prepare_headers(ctx)

  local base = vs.getprop(options, "base") or ""
  if type(base) ~= "string" then base = "" end
  local prefix = vs.getprop(options, "prefix") or ""
  if type(prefix) ~= "string" then prefix = "" end
  local suffix = vs.getprop(options, "suffix") or ""
  if type(suffix) ~= "string" then suffix = "" end

  ctx.spec = Spec.new({
    base = base,
    prefix = prefix,
    suffix = suffix,
    path = path,
    method = method,
    params = params,
    query = query,
    headers = headers,
    body = vs.getprop(fetchargs, "body"),
    step = "start",
  })

  -- Merge user-provided headers.
  local uh = vs.getprop(fetchargs, "headers")
  if type(uh) == "table" then
    for k, v in pairs(uh) do
      ctx.spec.headers[k] = v
    end
  end

  local _, err = utility.prepare_auth(ctx)
  if err ~= nil then
    return nil, err
  end

  return utility.make_fetch_def(ctx)
end


function DingconnectSDK:direct(fetchargs)
  local utility = self._utility

  local fetchdef, err = self:prepare(fetchargs)
  if err ~= nil then
    return { ok = false, err = err }, nil
  end

  fetchargs = fetchargs or {}
  local ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl")) or {}

  local ctx = utility.make_context({
    opname = "direct",
    ctrl = ctrl,
  }, self._rootctx)

  local url = fetchdef["url"] or ""
  local fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

  if fetch_err ~= nil then
    return { ok = false, err = fetch_err }, nil
  end

  if fetched == nil then
    return {
      ok = false,
      err = ctx:make_error("direct_no_response", "response: undefined"),
    }, nil
  end

  if type(fetched) == "table" then
    local status = helpers.to_int(vs.getprop(fetched, "status"))
    local headers = vs.getprop(fetched, "headers") or {}

    -- No-body responses (204, 304) and explicit zero content-length
    -- must skip JSON parsing — calling json() on an empty body errors.
    local content_length = nil
    if type(headers) == "table" then
      content_length = headers["content-length"]
    end
    local no_body = status == 204 or status == 304 or tostring(content_length) == "0"

    local json_data = nil
    if not no_body then
      local jf = vs.getprop(fetched, "json")
      if type(jf) == "function" then
        local ok, result = pcall(jf)
        if ok then
          json_data = result
        end
        -- Non-JSON body: json_data stays nil, status/headers preserved.
      end
    end

    return {
      ok = status >= 200 and status < 300,
      status = status,
      headers = headers,
      data = json_data,
    }, nil
  end

  return {
    ok = false,
    err = ctx:make_error("direct_invalid", "invalid response type"),
  }, nil
end



-- Idiomatic facade: client:AccountLookup():list() / client:AccountLookup():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function DingconnectSDK:AccountLookup(data)
  local EntityMod = require("entity.account_lookup_entity")
  if data == nil then
    if self._account_lookup == nil then
      self._account_lookup = EntityMod.new(self, nil)
    end
    return self._account_lookup
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Balance():list() / client:Balance():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function DingconnectSDK:Balance(data)
  local EntityMod = require("entity.balance_entity")
  if data == nil then
    if self._balance == nil then
      self._balance = EntityMod.new(self, nil)
    end
    return self._balance
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CancelResult():list() / client:CancelResult():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function DingconnectSDK:CancelResult(data)
  local EntityMod = require("entity.cancel_result_entity")
  if data == nil then
    if self._cancel_result == nil then
      self._cancel_result = EntityMod.new(self, nil)
    end
    return self._cancel_result
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Country():list() / client:Country():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function DingconnectSDK:Country(data)
  local EntityMod = require("entity.country_entity")
  if data == nil then
    if self._country == nil then
      self._country = EntityMod.new(self, nil)
    end
    return self._country
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Currency():list() / client:Currency():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function DingconnectSDK:Currency(data)
  local EntityMod = require("entity.currency_entity")
  if data == nil then
    if self._currency == nil then
      self._currency = EntityMod.new(self, nil)
    end
    return self._currency
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ErrorCodeDescription():list() / client:ErrorCodeDescription():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function DingconnectSDK:ErrorCodeDescription(data)
  local EntityMod = require("entity.error_code_description_entity")
  if data == nil then
    if self._error_code_description == nil then
      self._error_code_description = EntityMod.new(self, nil)
    end
    return self._error_code_description
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Estimate():list() / client:Estimate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function DingconnectSDK:Estimate(data)
  local EntityMod = require("entity.estimate_entity")
  if data == nil then
    if self._estimate == nil then
      self._estimate = EntityMod.new(self, nil)
    end
    return self._estimate
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Product():list() / client:Product():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function DingconnectSDK:Product(data)
  local EntityMod = require("entity.product_entity")
  if data == nil then
    if self._product == nil then
      self._product = EntityMod.new(self, nil)
    end
    return self._product
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ProductDescription():list() / client:ProductDescription():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function DingconnectSDK:ProductDescription(data)
  local EntityMod = require("entity.product_description_entity")
  if data == nil then
    if self._product_description == nil then
      self._product_description = EntityMod.new(self, nil)
    end
    return self._product_description
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Promotion():list() / client:Promotion():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function DingconnectSDK:Promotion(data)
  local EntityMod = require("entity.promotion_entity")
  if data == nil then
    if self._promotion == nil then
      self._promotion = EntityMod.new(self, nil)
    end
    return self._promotion
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PromotionDescription():list() / client:PromotionDescription():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function DingconnectSDK:PromotionDescription(data)
  local EntityMod = require("entity.promotion_description_entity")
  if data == nil then
    if self._promotion_description == nil then
      self._promotion_description = EntityMod.new(self, nil)
    end
    return self._promotion_description
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Provider():list() / client:Provider():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function DingconnectSDK:Provider(data)
  local EntityMod = require("entity.provider_entity")
  if data == nil then
    if self._provider == nil then
      self._provider = EntityMod.new(self, nil)
    end
    return self._provider
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ProviderStatus():list() / client:ProviderStatus():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function DingconnectSDK:ProviderStatus(data)
  local EntityMod = require("entity.provider_status_entity")
  if data == nil then
    if self._provider_status == nil then
      self._provider_status = EntityMod.new(self, nil)
    end
    return self._provider_status
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Region():list() / client:Region():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function DingconnectSDK:Region(data)
  local EntityMod = require("entity.region_entity")
  if data == nil then
    if self._region == nil then
      self._region = EntityMod.new(self, nil)
    end
    return self._region
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:SendTransfer():list() / client:SendTransfer():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function DingconnectSDK:SendTransfer(data)
  local EntityMod = require("entity.send_transfer_entity")
  if data == nil then
    if self._send_transfer == nil then
      self._send_transfer = EntityMod.new(self, nil)
    end
    return self._send_transfer
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:TransferRecord():list() / client:TransferRecord():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function DingconnectSDK:TransferRecord(data)
  local EntityMod = require("entity.transfer_record_entity")
  if data == nil then
    if self._transfer_record == nil then
      self._transfer_record = EntityMod.new(self, nil)
    end
    return self._transfer_record
  end
  return EntityMod.new(self, data)
end




function DingconnectSDK.test(testopts, sdkopts)
  sdkopts = sdkopts or {}
  sdkopts = vs.clone(sdkopts)
  if type(sdkopts) ~= "table" then
    sdkopts = {}
  end

  testopts = testopts or {}
  testopts = vs.clone(testopts)
  if type(testopts) ~= "table" then
    testopts = {}
  end
  testopts["active"] = true

  vs.setpath(sdkopts, "feature.test", testopts)

  local sdk = DingconnectSDK.new(sdkopts)
  sdk.mode = "test"

  return sdk
end


return DingconnectSDK
