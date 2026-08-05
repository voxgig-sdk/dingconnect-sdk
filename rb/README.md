# Dingconnect Ruby SDK



The Ruby SDK for the Dingconnect API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.AccountLookup` — with named operations (`list`/`create`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/dingconnect-sdk/releases](https://github.com/voxgig-sdk/dingconnect-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "Dingconnect_sdk"

client = DingconnectSDK.new({
  "apikey" => ENV["DINGCONNECT_APIKEY"],
})
```

### 2. List accountlookup records

```ruby
begin
  # list returns an Array of AccountLookup records — iterate directly.
  accountlookups = client.AccountLookup.list
  accountlookups.each do |item|
    puts "#{item["account_number_normalized"]}"
  end
rescue => err
  warn "list failed: #{err}"
end
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  currencys = client.Currency.list()
rescue => err
  warn "list failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required:

```ruby
client = DingconnectSDK.test

# Entity ops return the bare mock record (raises on error).
currency = client.Currency.list()
puts currency
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = DingconnectSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
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
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### DingconnectSDK

```ruby
require_relative "Dingconnect_sdk"
client = DingconnectSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `String` | API key for authentication. |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = DingconnectSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### DingconnectSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
| `AccountLookup` | `(data) -> AccountLookupEntity` | Create an AccountLookup entity instance. |
| `Balance` | `(data) -> BalanceEntity` | Create a Balance entity instance. |
| `CancelResult` | `(data) -> CancelResultEntity` | Create a CancelResult entity instance. |
| `Country` | `(data) -> CountryEntity` | Create a Country entity instance. |
| `Currency` | `(data) -> CurrencyEntity` | Create a Currency entity instance. |
| `ErrorCodeDescription` | `(data) -> ErrorCodeDescriptionEntity` | Create an ErrorCodeDescription entity instance. |
| `Estimate` | `(data) -> EstimateEntity` | Create an Estimate entity instance. |
| `Product` | `(data) -> ProductEntity` | Create a Product entity instance. |
| `ProductDescription` | `(data) -> ProductDescriptionEntity` | Create a ProductDescription entity instance. |
| `Promotion` | `(data) -> PromotionEntity` | Create a Promotion entity instance. |
| `PromotionDescription` | `(data) -> PromotionDescriptionEntity` | Create a PromotionDescription entity instance. |
| `Provider` | `(data) -> ProviderEntity` | Create a Provider entity instance. |
| `ProviderStatus` | `(data) -> ProviderStatusEntity` | Create a ProviderStatus entity instance. |
| `Region` | `(data) -> RegionEntity` | Create a Region entity instance. |
| `SendTransfer` | `(data) -> SendTransferEntity` | Create a SendTransfer entity instance. |
| `TransferRecord` | `(data) -> TransferRecordEntity` | Create a TransferRecord entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `list` | `(reqmatch = nil, ctrl) -> Array` | List entities matching the criteria (call with no argument to list all). Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `DingconnectError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

### Entities

#### AccountLookup

| Field | Description |
| --- | --- |
| `account_number_normalized` |  |
| `country_iso` |  |
| `error_code` |  |
| `item` |  |
| `result_code` |  |

Operations: List.

API path: `/GetAccountLookup`

#### Balance

| Field | Description |
| --- | --- |
| `code` |  |
| `context` |  |

Operations: List.

API path: `/GetBalance`

#### CancelResult

| Field | Description |
| --- | --- |
| `error_code` |  |
| `item` |  |
| `result_code` |  |

Operations: Create.

API path: `/CancelTransfers`

#### Country

| Field | Description |
| --- | --- |
| `error_code` |  |
| `item` |  |
| `result_code` |  |

Operations: List.

API path: `/GetCountries`

#### Currency

| Field | Description |
| --- | --- |
| `error_code` |  |
| `item` |  |
| `result_code` |  |

Operations: List.

API path: `/GetCurrencies`

#### ErrorCodeDescription

| Field | Description |
| --- | --- |
| `error_code` |  |
| `item` |  |
| `result_code` |  |

Operations: List.

API path: `/GetErrorCodeDescriptions`

#### Estimate

| Field | Description |
| --- | --- |
| `error_code` |  |
| `item` |  |
| `result_code` |  |

Operations: Create.

API path: `/EstimatePrices`

#### Product

| Field | Description |
| --- | --- |
| `error_code` |  |
| `item` |  |
| `result_code` |  |

Operations: List.

API path: `/GetProducts`

#### ProductDescription

| Field | Description |
| --- | --- |
| `error_code` |  |
| `item` |  |
| `result_code` |  |

Operations: List.

API path: `/GetProductDescriptions`

#### Promotion

| Field | Description |
| --- | --- |
| `error_code` |  |
| `item` |  |
| `result_code` |  |

Operations: List.

API path: `/GetPromotions`

#### PromotionDescription

| Field | Description |
| --- | --- |
| `error_code` |  |
| `item` |  |
| `result_code` |  |

Operations: List.

API path: `/GetPromotionDescriptions`

#### Provider

| Field | Description |
| --- | --- |
| `error_code` |  |
| `item` |  |
| `result_code` |  |

Operations: List.

API path: `/GetProviders`

#### ProviderStatus

| Field | Description |
| --- | --- |
| `error_code` |  |
| `item` |  |
| `result_code` |  |

Operations: List.

API path: `/GetProviderStatus`

#### Region

| Field | Description |
| --- | --- |
| `error_code` |  |
| `item` |  |
| `result_code` |  |

Operations: List.

API path: `/GetRegions`

#### SendTransfer

| Field | Description |
| --- | --- |
| `account_number` |  |
| `distributor_ref` |  |
| `error_code` |  |
| `result_code` |  |
| `send_currency_iso` |  |
| `send_value` |  |
| `setting` |  |
| `sku_code` |  |
| `transfer_record` |  |
| `validate_only` |  |

Operations: Create.

API path: `/SendTransfer`

#### TransferRecord

| Field | Description |
| --- | --- |
| `account_number` |  |
| `distributor_ref` |  |
| `ended_at_utc` |  |
| `error_code` |  |
| `item` |  |
| `result_code` |  |
| `skip` |  |
| `sku_code` |  |
| `started_at_utc` |  |
| `take` |  |
| `there_are_more_item` |  |
| `transfer_ref` |  |

Operations: Create.

API path: `/ListTransferRecords`



## Entities


### AccountLookup

Create an instance: `account_lookup = client.AccountLookup`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_number_normalized` | `String` |  |
| `country_iso` | `String` |  |
| `error_code` | `Array` |  |
| `item` | `Array` |  |
| `result_code` | `Integer` |  |

#### Example: List

```ruby
# list returns an Array of AccountLookup records (raises on error).
account_lookups = client.AccountLookup.list
```


### Balance

Create an instance: `balance = client.Balance`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `code` | `String` |  |
| `context` | `String` |  |

#### Example: List

```ruby
# list returns an Array of Balance records (raises on error).
balances = client.Balance.list
```


### CancelResult

Create an instance: `cancel_result = client.CancelResult`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `Array` |  |
| `item` | `Array` |  |
| `result_code` | `Integer` |  |

#### Example: Create

```ruby
cancel_result = client.CancelResult.create({
})
```


### Country

Create an instance: `country = client.Country`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `Array` |  |
| `item` | `Array` |  |
| `result_code` | `Integer` |  |

#### Example: List

```ruby
# list returns an Array of Country records (raises on error).
countrys = client.Country.list
```


### Currency

Create an instance: `currency = client.Currency`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `Array` |  |
| `item` | `Array` |  |
| `result_code` | `Integer` |  |

#### Example: List

```ruby
# list returns an Array of Currency records (raises on error).
currencys = client.Currency.list
```


### ErrorCodeDescription

Create an instance: `error_code_description = client.ErrorCodeDescription`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `Array` |  |
| `item` | `Array` |  |
| `result_code` | `Integer` |  |

#### Example: List

```ruby
# list returns an Array of ErrorCodeDescription records (raises on error).
error_code_descriptions = client.ErrorCodeDescription.list
```


### Estimate

Create an instance: `estimate = client.Estimate`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `Array` |  |
| `item` | `Array` |  |
| `result_code` | `Integer` |  |

#### Example: Create

```ruby
estimate = client.Estimate.create({
})
```


### Product

Create an instance: `product = client.Product`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `Array` |  |
| `item` | `Array` |  |
| `result_code` | `Integer` |  |

#### Example: List

```ruby
# list returns an Array of Product records (raises on error).
products = client.Product.list
```


### ProductDescription

Create an instance: `product_description = client.ProductDescription`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `Array` |  |
| `item` | `Array` |  |
| `result_code` | `Integer` |  |

#### Example: List

```ruby
# list returns an Array of ProductDescription records (raises on error).
product_descriptions = client.ProductDescription.list
```


### Promotion

Create an instance: `promotion = client.Promotion`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `Array` |  |
| `item` | `Array` |  |
| `result_code` | `Integer` |  |

#### Example: List

```ruby
# list returns an Array of Promotion records (raises on error).
promotions = client.Promotion.list
```


### PromotionDescription

Create an instance: `promotion_description = client.PromotionDescription`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `Array` |  |
| `item` | `Array` |  |
| `result_code` | `Integer` |  |

#### Example: List

```ruby
# list returns an Array of PromotionDescription records (raises on error).
promotion_descriptions = client.PromotionDescription.list
```


### Provider

Create an instance: `provider = client.Provider`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `Array` |  |
| `item` | `Array` |  |
| `result_code` | `Integer` |  |

#### Example: List

```ruby
# list returns an Array of Provider records (raises on error).
providers = client.Provider.list
```


### ProviderStatus

Create an instance: `provider_status = client.ProviderStatus`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `Array` |  |
| `item` | `Array` |  |
| `result_code` | `Integer` |  |

#### Example: List

```ruby
# list returns an Array of ProviderStatus records (raises on error).
provider_statuss = client.ProviderStatus.list
```


### Region

Create an instance: `region = client.Region`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `Array` |  |
| `item` | `Array` |  |
| `result_code` | `Integer` |  |

#### Example: List

```ruby
# list returns an Array of Region records (raises on error).
regions = client.Region.list
```


### SendTransfer

Create an instance: `send_transfer = client.SendTransfer`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_number` | `String` |  |
| `distributor_ref` | `String` |  |
| `error_code` | `Array` |  |
| `result_code` | `Integer` |  |
| `send_currency_iso` | `String` |  |
| `send_value` | `Float` |  |
| `setting` | `Array` |  |
| `sku_code` | `String` |  |
| `transfer_record` | `Hash` |  |
| `validate_only` | `Boolean` |  |

#### Example: Create

```ruby
send_transfer = client.SendTransfer.create({
  "account_number" => "example_account_number", # String
  "distributor_ref" => "example_distributor_ref", # String
  "send_value" => 1, # Float
  "sku_code" => "example_sku_code", # String
})
```


### TransferRecord

Create an instance: `transfer_record = client.TransferRecord`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_number` | `String` |  |
| `distributor_ref` | `Array` |  |
| `ended_at_utc` | `String` |  |
| `error_code` | `Array` |  |
| `item` | `Array` |  |
| `result_code` | `Integer` |  |
| `skip` | `Integer` |  |
| `sku_code` | `Array` |  |
| `started_at_utc` | `String` |  |
| `take` | `Integer` |  |
| `there_are_more_item` | `Boolean` |  |
| `transfer_ref` | `Array` |  |

#### Example: Create

```ruby
transfer_record = client.TransferRecord.create({
  "skip" => 1, # Integer
  "take" => 1, # Integer
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

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── Dingconnect_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`Dingconnect_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```ruby
currency = client.Currency
currency.list()

# currency.data_get now returns the currency data from the last list
# currency.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
