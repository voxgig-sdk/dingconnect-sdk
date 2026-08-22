# Dingconnect Lua SDK



The Lua SDK for the Dingconnect API — an entity-oriented client using Lua conventions.

It exposes the API as capitalised, semantic **Entities** — e.g. `client:AccountLookup()` — each with the same small set of operations (`list`, `create`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to LuaRocks. Install it from the
GitHub release tag (`lua/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/dingconnect-sdk/releases)),
or add the source directory to your `LUA_PATH`:

```bash
export LUA_PATH="path/to/lua/?.lua;path/to/lua/?/init.lua;;"
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```lua
local sdk = require("dingconnect_sdk")

local client = sdk.new({
  apikey = os.getenv("DINGCONNECT_APIKEY"),
})
```

### 2. List accountlookup records

Entity operations return `(value, err)`. For `list`, `value` is the
array of records itself — iterate it directly (there is no wrapper).

```lua
local accountlookups, err = client:AccountLookup():list()
if err then error(err) end

for _, item in ipairs(accountlookups) do
  print(item["AccountNumberNormalized"])
end
```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local currencys, err = client:Currency():list()
if err then error(err) end
```

`direct` follows the same `(value, err)` convention:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example_id" },
})
if err then error(err) end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
if err then error(err) end

if result["ok"] then
  print(result["status"])  -- 200
  print(result["data"])    -- response body
end
```

### Prepare a request without sending it

```lua
local fetchdef, err = client:prepare({
  path = "/api/resource/{id}",
  method = "DELETE",
  params = { id = "example" },
})
if err then error(err) end

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```lua
local client = sdk.test()

local result, err = client:Currency():list()
-- result is the returned data; err is set on failure
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```lua
local function mock_fetch(url, init)
  return {
    status = 200,
    statusText = "OK",
    headers = {},
    json = function()
      return { id = "mock01" }
    end,
  }, nil
end

local client = sdk.new({
  base = "http://localhost:8080",
  system = {
    fetch = mock_fetch,
  },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
DINGCONNECT_TEST_LIVE=TRUE
DINGCONNECT_APIKEY=<your-key>
```

Then run:

```bash
cd lua && busted test/
```


## Reference

### DingconnectSDK

```lua
local sdk = require("dingconnect_sdk")
local client = sdk.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `table` | Feature activation flags. |
| `extend` | `table` | Additional Feature instances to load. |
| `system` | `table` | System overrides (e.g. custom `fetch` function). |

### test

```lua
local client = sdk.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### DingconnectSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> table` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> table, err` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> table, err` | Build and send an HTTP request. |
| `AccountLookup` | `(data) -> AccountLookupEntity` | Create an AccountLookup entity instance. |
| `Balance` | `(data) -> BalanceEntity` | Create a Balance entity instance. |
| `CancelTransfer` | `(data) -> CancelTransferEntity` | Create a CancelTransfer entity instance. |
| `Country` | `(data) -> CountryEntity` | Create a Country entity instance. |
| `Currency` | `(data) -> CurrencyEntity` | Create a Currency entity instance. |
| `ErrorCodeDescription` | `(data) -> ErrorCodeDescriptionEntity` | Create an ErrorCodeDescription entity instance. |
| `EstimatePrice` | `(data) -> EstimatePriceEntity` | Create an EstimatePrice entity instance. |
| `ListTransferRecord` | `(data) -> ListTransferRecordEntity` | Create a ListTransferRecord entity instance. |
| `LookupBill` | `(data) -> LookupBillEntity` | Create a LookupBill entity instance. |
| `Product` | `(data) -> ProductEntity` | Create a Product entity instance. |
| `ProductDescription` | `(data) -> ProductDescriptionEntity` | Create a ProductDescription entity instance. |
| `Promotion` | `(data) -> PromotionEntity` | Create a Promotion entity instance. |
| `PromotionDescription` | `(data) -> PromotionDescriptionEntity` | Create a PromotionDescription entity instance. |
| `Provider` | `(data) -> ProviderEntity` | Create a Provider entity instance. |
| `ProviderStatus` | `(data) -> ProviderStatusEntity` | Create a ProviderStatus entity instance. |
| `Region` | `(data) -> RegionEntity` | Create a Region entity instance. |
| `SendTransfer` | `(data) -> SendTransferEntity` | Create a SendTransfer entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `list` | `(reqmatch, ctrl) -> any, err` | List entities matching the criteria. |
| `create` | `(reqdata, ctrl) -> any, err` | Create a new entity. |
| `data_get` | `() -> table` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> table` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> string` | Return the entity name. |

### Result shape

Entity operations return `(value, err)`. The `value` is the operation's
data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `create` | the entity record (a `table`) |
| `list` | an array (`table`) of entity records |

Check `err` first (it is non-`nil` on failure), then use `value`:

    local account_lookup, err = client:AccountLookup():list()
    if err then error(err) end
    -- account_lookup is the record list

Only `direct()` returns a response envelope — a `table` with `ok`,
`status`, `headers`, and `data` keys.

### Entities

#### AccountLookup

| Field | Description |
| --- | --- |
| `AccountNumberNormalized` | We attempt to normalize phone numbers following the public telecommunication numbering plan <a href="https://en.wikipedia.org/wiki/E.164" target="_blank">E.164</a>, if we succeed the normalized number will be returned in this field formatt… |
| `CountryIso` | The country of the account number |
| `ErrorCodes` |  |
| `Items` | This will contain provider information associated to the account number. |
| `ResultCode` |  |

Operations: List.

API path: `/api/V1/GetAccountLookup`

#### Balance

| Field | Description |
| --- | --- |
| `Code` | The code that can be used to lookup the explanatory message associated with the error |
| `Context` | API specific context as to the reason for the specific code |

Operations: List.

API path: `/api/V1/GetBalance`

#### CancelTransfer

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: Create.

API path: `/api/V1/CancelTransfers`

#### Country

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` | The list of countries that our system is aware of. |
| `ResultCode` |  |

Operations: List.

API path: `/api/V1/GetCountries`

#### Currency

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: List.

API path: `/api/V1/GetCurrencies`

#### ErrorCodeDescription

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` | A list of ErrorCodes and their localized descriptions |
| `ResultCode` |  |

Operations: List.

API path: `/api/V1/GetErrorCodeDescriptions`

#### EstimatePrice

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: Create.

API path: `/api/V1/EstimatePrices`

#### ListTransferRecord

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` | The list of items satisfying the transfer query. |
| `ResultCode` |  |
| `ThereAreMoreItems` | Indicates if the caller should execute the query again. |

Operations: Create.

API path: `/api/V1/ListTransferRecords`

#### LookupBill

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: Create.

API path: `/api/V1/LookupBills`

#### Product

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` | A list of products that fulfil the submitted criteria. |
| `ResultCode` |  |

Operations: List.

API path: `/api/V1/GetProducts`

#### ProductDescription

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` | A localized list of product descriptions. |
| `ResultCode` |  |

Operations: List.

API path: `/api/V1/GetProductDescriptions`

#### Promotion

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` | List of available promotions |
| `ResultCode` |  |

Operations: List.

API path: `/api/V1/GetPromotions`

#### PromotionDescription

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` | A localized list of promotions. |
| `ResultCode` |  |

Operations: List.

API path: `/api/V1/GetPromotionDescriptions`

#### Provider

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` | A list of providers that the distributor has Products for. |
| `ResultCode` |  |

Operations: List.

API path: `/api/V1/GetProviders`

#### ProviderStatus

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: List.

API path: `/api/V1/GetProviderStatus`

#### Region

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` | The list of regions that the system uses. |
| `ResultCode` |  |

Operations: List.

API path: `/api/V1/GetRegions`

#### SendTransfer

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `ResultCode` |  |
| `TransferRecord` |  |

Operations: Create.

API path: `/api/V1/SendTransfer`



## Entities


### AccountLookup

Create an instance: `local account_lookup = client:AccountLookup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `AccountNumberNormalized` | `string` | We attempt to normalize phone numbers following the public telecommunication numbering plan <a href="https://en.wikipedia.org/wiki/E.164" target="_blank">E.164</a>, if we succeed the normalized number will be returned in this field formatt… |
| `CountryIso` | `string` | The country of the account number |
| `ErrorCodes` | `table` |  |
| `Items` | `table` | This will contain provider information associated to the account number. |
| `ResultCode` | `number` |  |

#### Example: List

```lua
local account_lookups, err = client:AccountLookup():list()
```


### Balance

Create an instance: `local balance = client:Balance(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `Code` | `string` | The code that can be used to lookup the explanatory message associated with the error |
| `Context` | `string` | API specific context as to the reason for the specific code |

#### Example: List

```lua
local balances, err = client:Balance():list()
```


### CancelTransfer

Create an instance: `local cancel_transfer = client:CancelTransfer(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `table` |  |
| `Items` | `table` |  |
| `ResultCode` | `number` |  |

#### Example: Create

```lua
local cancel_transfer, err = client:CancelTransfer():create({
  ErrorCodes = {}, -- table
  Items = {}, -- table
  ResultCode = 1, -- number
})
```


### Country

Create an instance: `local country = client:Country(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `table` |  |
| `Items` | `table` | The list of countries that our system is aware of. |
| `ResultCode` | `number` |  |

#### Example: List

```lua
local countrys, err = client:Country():list()
```


### Currency

Create an instance: `local currency = client:Currency(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `table` |  |
| `Items` | `table` |  |
| `ResultCode` | `number` |  |

#### Example: List

```lua
local currencys, err = client:Currency():list()
```


### ErrorCodeDescription

Create an instance: `local error_code_description = client:ErrorCodeDescription(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `table` |  |
| `Items` | `table` | A list of ErrorCodes and their localized descriptions |
| `ResultCode` | `number` |  |

#### Example: List

```lua
local error_code_descriptions, err = client:ErrorCodeDescription():list()
```


### EstimatePrice

Create an instance: `local estimate_price = client:EstimatePrice(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `table` |  |
| `Items` | `table` |  |
| `ResultCode` | `number` |  |

#### Example: Create

```lua
local estimate_price, err = client:EstimatePrice():create({
  ErrorCodes = {}, -- table
  Items = {}, -- table
  ResultCode = 1, -- number
})
```


### ListTransferRecord

Create an instance: `local list_transfer_record = client:ListTransferRecord(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `table` |  |
| `Items` | `table` | The list of items satisfying the transfer query. |
| `ResultCode` | `number` |  |
| `ThereAreMoreItems` | `boolean` | Indicates if the caller should execute the query again. |

#### Example: Create

```lua
local list_transfer_record, err = client:ListTransferRecord():create({
  ErrorCodes = {}, -- table
  Items = {}, -- table
  ResultCode = 1, -- number
  ThereAreMoreItems = true, -- boolean
})
```


### LookupBill

Create an instance: `local lookup_bill = client:LookupBill(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `table` |  |
| `Items` | `table` |  |
| `ResultCode` | `number` |  |

#### Example: Create

```lua
local lookup_bill, err = client:LookupBill():create({
  ErrorCodes = {}, -- table
  Items = {}, -- table
  ResultCode = 1, -- number
})
```


### Product

Create an instance: `local product = client:Product(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `table` |  |
| `Items` | `table` | A list of products that fulfil the submitted criteria. |
| `ResultCode` | `number` |  |

#### Example: List

```lua
local products, err = client:Product():list()
```


### ProductDescription

Create an instance: `local product_description = client:ProductDescription(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `table` |  |
| `Items` | `table` | A localized list of product descriptions. |
| `ResultCode` | `number` |  |

#### Example: List

```lua
local product_descriptions, err = client:ProductDescription():list()
```


### Promotion

Create an instance: `local promotion = client:Promotion(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `table` |  |
| `Items` | `table` | List of available promotions |
| `ResultCode` | `number` |  |

#### Example: List

```lua
local promotions, err = client:Promotion():list()
```


### PromotionDescription

Create an instance: `local promotion_description = client:PromotionDescription(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `table` |  |
| `Items` | `table` | A localized list of promotions. |
| `ResultCode` | `number` |  |

#### Example: List

```lua
local promotion_descriptions, err = client:PromotionDescription():list()
```


### Provider

Create an instance: `local provider = client:Provider(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `table` |  |
| `Items` | `table` | A list of providers that the distributor has Products for. |
| `ResultCode` | `number` |  |

#### Example: List

```lua
local providers, err = client:Provider():list()
```


### ProviderStatus

Create an instance: `local provider_status = client:ProviderStatus(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `table` |  |
| `Items` | `table` |  |
| `ResultCode` | `number` |  |

#### Example: List

```lua
local provider_statuss, err = client:ProviderStatus():list()
```


### Region

Create an instance: `local region = client:Region(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `table` |  |
| `Items` | `table` | The list of regions that the system uses. |
| `ResultCode` | `number` |  |

#### Example: List

```lua
local regions, err = client:Region():list()
```


### SendTransfer

Create an instance: `local send_transfer = client:SendTransfer(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `table` |  |
| `ResultCode` | `number` |  |
| `TransferRecord` | `table` |  |

#### Example: Create

```lua
local send_transfer, err = client:SendTransfer():create({
  ErrorCodes = {}, -- table
  ResultCode = 1, -- number
  TransferRecord = {}, -- table
})
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Lua table
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as tables

The Lua SDK uses plain Lua tables throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a table.

### Module structure

```
lua/
├── dingconnect_sdk.lua    -- Main SDK module
├── config.lua               -- Configuration
├── features.lua             -- Feature factory
├── core/                    -- Core types and context
├── entity/                  -- Entity implementations
├── feature/                 -- Built-in features (Base, Test, Log)
├── utility/                 -- Utility functions and struct library
└── test/                    -- Test suites
```

The main module (`dingconnect_sdk`) exports the SDK constructor
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```lua
local currency = client:Currency()
currency:list()

-- currency:data_get() now returns the currency data from the last list
-- currency:match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
