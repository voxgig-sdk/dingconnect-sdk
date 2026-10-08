# Dingconnect Ruby SDK Reference

Complete API reference for the Dingconnect Ruby SDK.


## DingconnectSDK

### Constructor

```ruby
require_relative 'Dingconnect_sdk'

client = DingconnectSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["apikey"]` | `String` | API key for authentication. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `DingconnectSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = DingconnectSDK.test
```


### Instance Methods

#### `AccountLookup(data = nil)`

Create a new `AccountLookup` entity instance. Pass `nil` for no initial data.

#### `Balance(data = nil)`

Create a new `Balance` entity instance. Pass `nil` for no initial data.

#### `CancelTransfer(data = nil)`

Create a new `CancelTransfer` entity instance. Pass `nil` for no initial data.

#### `Country(data = nil)`

Create a new `Country` entity instance. Pass `nil` for no initial data.

#### `Currency(data = nil)`

Create a new `Currency` entity instance. Pass `nil` for no initial data.

#### `ErrorCodeDescription(data = nil)`

Create a new `ErrorCodeDescription` entity instance. Pass `nil` for no initial data.

#### `EstimatePrice(data = nil)`

Create a new `EstimatePrice` entity instance. Pass `nil` for no initial data.

#### `ListTransferRecord(data = nil)`

Create a new `ListTransferRecord` entity instance. Pass `nil` for no initial data.

#### `LookupBill(data = nil)`

Create a new `LookupBill` entity instance. Pass `nil` for no initial data.

#### `Product(data = nil)`

Create a new `Product` entity instance. Pass `nil` for no initial data.

#### `ProductDescription(data = nil)`

Create a new `ProductDescription` entity instance. Pass `nil` for no initial data.

#### `Promotion(data = nil)`

Create a new `Promotion` entity instance. Pass `nil` for no initial data.

#### `PromotionDescription(data = nil)`

Create a new `PromotionDescription` entity instance. Pass `nil` for no initial data.

#### `Provider(data = nil)`

Create a new `Provider` entity instance. Pass `nil` for no initial data.

#### `ProviderStatus(data = nil)`

Create a new `ProviderStatus` entity instance. Pass `nil` for no initial data.

#### `Region(data = nil)`

Create a new `Region` entity instance. Pass `nil` for no initial data.

#### `SendTransfer(data = nil)`

Create a new `SendTransfer` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## AccountLookupEntity

```ruby
account_lookup = client.AccountLookup
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `AccountNumberNormalized` | `String` | No | We attempt to normalize phone numbers following the public telecommunication numbering plan <a href="https://en.wikipedia.org/wiki/E.164" target="_blank">E.164</a>, if we succeed the normalized number will be returned in this field formatt… |
| `CountryIso` | `String` | No | The country of the account number |
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes | This will contain provider information associated to the account number. |
| `ResultCode` | `Integer` | Yes |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record; `data_get` reads each record. Raises on error.

```ruby
results = client.AccountLookup.list
results.each { |item| puts item.data_get }
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `AccountLookupEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## BalanceEntity

```ruby
balance = client.Balance
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `Code` | `String` | Yes | The code that can be used to lookup the explanatory message associated with the error |
| `Context` | `String` | No | API specific context as to the reason for the specific code |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record; `data_get` reads each record. Raises on error.

```ruby
results = client.Balance.list
results.each { |item| puts item.data_get }
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `BalanceEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CancelTransferEntity

```ruby
cancel_transfer = client.CancelTransfer
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes |  |
| `ResultCode` | `Integer` | Yes |  |
| `cancellations` | `Array` | No | An explicit list of records to cancel. |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Returns the created entity and raises on error.

```ruby
result = client.CancelTransfer.create({
  "ErrorCodes" => [], # Array
  "Items" => [], # Array
  "ResultCode" => 1, # Integer
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CancelTransferEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CountryEntity

```ruby
country = client.Country
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes | The list of countries that our system is aware of. |
| `ResultCode` | `Integer` | Yes |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record; `data_get` reads each record. Raises on error.

```ruby
results = client.Country.list
results.each { |item| puts item.data_get }
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CountryEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CurrencyEntity

```ruby
currency = client.Currency
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes |  |
| `ResultCode` | `Integer` | Yes |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record; `data_get` reads each record. Raises on error.

```ruby
results = client.Currency.list
results.each { |item| puts item.data_get }
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CurrencyEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ErrorCodeDescriptionEntity

```ruby
error_code_description = client.ErrorCodeDescription
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes | A list of ErrorCodes and their localized descriptions |
| `ResultCode` | `Integer` | Yes |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record; `data_get` reads each record. Raises on error.

```ruby
results = client.ErrorCodeDescription.list
results.each { |item| puts item.data_get }
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ErrorCodeDescriptionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EstimatePriceEntity

```ruby
estimate_price = client.EstimatePrice
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes |  |
| `ResultCode` | `Integer` | Yes |  |
| `estimations` | `Array` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Returns the created entity and raises on error.

```ruby
result = client.EstimatePrice.create({
  "ErrorCodes" => [], # Array
  "Items" => [], # Array
  "ResultCode" => 1, # Integer
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EstimatePriceEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ListTransferRecordEntity

```ruby
list_transfer_record = client.ListTransferRecord
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `AccountNumber` | `String` | No | Filter transfers by AccountNumber |
| `DistributorRef` | `String` | No | Filter transfers by DistributorRef. |
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes | The list of items satisfying the transfer query. |
| `ResultCode` | `Integer` | Yes |  |
| `Skip` | `Integer` | No | The amount of records to by-pass before returning the remaining records |
| `Take` | `Integer` | Yes | The amount of records to return |
| `ThereAreMoreItems` | `Boolean` | Yes | Indicates if the caller should execute the query again. |
| `TransferRef` | `String` | No | Filter by Ding TransferRef |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Returns the created entity and raises on error.

```ruby
result = client.ListTransferRecord.create({
  "ErrorCodes" => [], # Array
  "Items" => [], # Array
  "ResultCode" => 1, # Integer
  "Take" => 1, # Integer
  "ThereAreMoreItems" => true, # Boolean
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ListTransferRecordEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## LookupBillEntity

```ruby
lookup_bill = client.LookupBill
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `AccountNumber` | `String` | Yes | The account number to target |
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes |  |
| `ResultCode` | `Integer` | Yes |  |
| `Settings` | `Array` | No | Product specific name/value pairs to be associated with the lookup bills request |
| `SkuCode` | `String` | Yes | Code provided by GetProducts API |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Returns the created entity and raises on error.

```ruby
result = client.LookupBill.create({
  "AccountNumber" => "example_AccountNumber", # String
  "ErrorCodes" => [], # Array
  "Items" => [], # Array
  "ResultCode" => 1, # Integer
  "SkuCode" => "example_SkuCode", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `LookupBillEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ProductEntity

```ruby
product = client.Product
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes | A list of products that fulfil the submitted criteria. |
| `ResultCode` | `Integer` | Yes |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record; `data_get` reads each record. Raises on error.

```ruby
results = client.Product.list
results.each { |item| puts item.data_get }
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ProductEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ProductDescriptionEntity

```ruby
product_description = client.ProductDescription
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes | A localized list of product descriptions. |
| `ResultCode` | `Integer` | Yes |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record; `data_get` reads each record. Raises on error.

```ruby
results = client.ProductDescription.list
results.each { |item| puts item.data_get }
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ProductDescriptionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PromotionEntity

```ruby
promotion = client.Promotion
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes | List of available promotions |
| `ResultCode` | `Integer` | Yes |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record; `data_get` reads each record. Raises on error.

```ruby
results = client.Promotion.list
results.each { |item| puts item.data_get }
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PromotionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PromotionDescriptionEntity

```ruby
promotion_description = client.PromotionDescription
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes | A localized list of promotions. |
| `ResultCode` | `Integer` | Yes |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record; `data_get` reads each record. Raises on error.

```ruby
results = client.PromotionDescription.list
results.each { |item| puts item.data_get }
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PromotionDescriptionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ProviderEntity

```ruby
provider = client.Provider
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes | A list of providers that the distributor has Products for. |
| `ResultCode` | `Integer` | Yes |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record; `data_get` reads each record. Raises on error.

```ruby
results = client.Provider.list
results.each { |item| puts item.data_get }
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ProviderEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ProviderStatusEntity

```ruby
provider_status = client.ProviderStatus
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes |  |
| `ResultCode` | `Integer` | Yes |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record; `data_get` reads each record. Raises on error.

```ruby
results = client.ProviderStatus.list
results.each { |item| puts item.data_get }
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ProviderStatusEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## RegionEntity

```ruby
region = client.Region
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes | The list of regions that the system uses. |
| `ResultCode` | `Integer` | Yes |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record; `data_get` reads each record. Raises on error.

```ruby
results = client.Region.list
results.each { |item| puts item.data_get }
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `RegionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SendTransferEntity

```ruby
send_transfer = client.SendTransfer
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `AccountNumber` | `String` | Yes | The account number to target |
| `BillRef` | `String` | No | Bill reference. |
| `DistributorRef` | `String` | Yes | Unique identifier in the distributor system to be associated with the transfer |
| `ErrorCodes` | `Array` | Yes |  |
| `ResultCode` | `Integer` | Yes |  |
| `SendCurrencyIso` | `String` | No | The currency of the `SendValue`. |
| `SendValue` | `Float` | Yes | The transfer value to be sent. |
| `Settings` | `Array` | No | Product specific name/value pairs to be associated with the transfer request |
| `SkuCode` | `String` | Yes | Code provided by GetProducts API |
| `TransferRecord` | `Hash` | Yes |  |
| `ValidateOnly` | `Boolean` | Yes | Validate the request with the provider without doing a transfer |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Returns the created entity and raises on error.

```ruby
result = client.SendTransfer.create({
  "AccountNumber" => "example_AccountNumber", # String
  "DistributorRef" => "example_DistributorRef", # String
  "ErrorCodes" => [], # Array
  "ResultCode" => 1, # Integer
  "SendValue" => 1, # Float
  "SkuCode" => "example_SkuCode", # String
  "TransferRecord" => {}, # Hash
  "ValidateOnly" => true, # Boolean
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SendTransferEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `debug` | 0.0.1 | Debug capture |
| `idempotency` | 0.0.1 | Idempotency |
| `metrics` | 0.0.1 | Metrics |
| `paging` | 0.0.1 | Paging |
| `ratelimit` | 0.0.1 | Rate limiting |
| `retry` | 0.0.1 | Retry |
| `test` | 0.0.1 | Test transport |
| `timeout` | 0.0.1 | Timeout |


Features are activated via the `feature` option:

```ruby
client = DingconnectSDK.new({
  "feature" => {
    "debug" => { "active" => true },
    "idempotency" => { "active" => true },
    "metrics" => { "active" => true },
    "paging" => { "active" => true },
    "ratelimit" => { "active" => true },
    "retry" => { "active" => true },
    "test" => { "active" => true },
    "timeout" => { "active" => true },
  },
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### Ordering

`ratelimit`, `retry`, `timeout` wrap the transport. Each
wraps whatever is already installed, so **activation order is nesting order**:
a feature activated later sits OUTSIDE one activated earlier, and sees the call
first.

That decides behaviour, not just sequence: a feature that short-circuits the
call, such as a cache serving a hit, stops every feature nested inside it from
ever seeing that call.

`debug`, `idempotency`, `metrics`, `paging`, `test` attach to pipeline hooks
rather than the transport, so their order does not affect what they observe.

#### `debug`

Debug capture.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `max` | `100` |
| `redact` | `['authorization', 'cookie', 'set-cookie', 'api-key', 'apikey', 'x-api-key', 'idempotency-key']` |

| Option | Type |
|---|---|
| `now` | function |
| `onEntry` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.debug.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `idempotency`

Idempotency.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `header` | `'Idempotency-Key'` |
| `methods` | `['POST', 'PUT', 'PATCH', 'DELETE']` |
| `ops` | `['create', 'update', 'remove']` |

| Option | Type |
|---|---|
| `keygen` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.idempotency.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `metrics`

Metrics.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `now` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.metrics.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `paging`

Paging.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `afterVar` | `'after'` |
| `cursorParam` | `'cursor'` |
| `firstVar` | `'first'` |
| `limitParam` | `'limit'` |
| `pageParam` | `'page'` |
| `startPage` | `1` |

| Option | Type |
|---|---|
| `limit` | number |
| `ops` | list |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.paging.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `ratelimit`

Rate limiting.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

| Option | Type |
|---|---|
| `now` | function |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.ratelimit.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `retry`

Retry.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

| Option | Type |
|---|---|
| `jitter` | boolean |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.retry.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `test`

Test transport.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `entity` | map |
| `net` | map |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

#### `timeout`

Timeout.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

| Option | Type |
|---|---|
| `clearTimer` | function |
| `now` | function |
| `setTimer` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.timeout.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

