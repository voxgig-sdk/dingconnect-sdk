-- Dingconnect SDK error

local DingconnectError = {}
DingconnectError.__index = DingconnectError


function DingconnectError.new(code, msg, ctx)
  local self = setmetatable({}, DingconnectError)
  self.is_sdk_error = true
  self.sdk = "Dingconnect"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function DingconnectError:error()
  return self.msg
end


function DingconnectError:__tostring()
  return self.msg
end


return DingconnectError
