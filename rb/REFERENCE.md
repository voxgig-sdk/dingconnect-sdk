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

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.AccountLookup.list
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

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Balance.list
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

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

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

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Country.list
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

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Currency.list
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

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.ErrorCodeDescription.list
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

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

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
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes | The list of items satisfying the transfer query. |
| `ResultCode` | `Integer` | Yes |  |
| `ThereAreMoreItems` | `Boolean` | Yes | Indicates if the caller should execute the query again. |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.ListTransferRecord.create({
  "ErrorCodes" => [], # Array
  "Items" => [], # Array
  "ResultCode" => 1, # Integer
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
| `ErrorCodes` | `Array` | Yes |  |
| `Items` | `Array` | Yes |  |
| `ResultCode` | `Integer` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.LookupBill.create({
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

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Product.list
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

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.ProductDescription.list
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

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Promotion.list
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

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.PromotionDescription.list
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

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Provider.list
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

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.ProviderStatus.list
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

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Region.list
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
| `ErrorCodes` | `Array` | Yes |  |
| `ResultCode` | `Integer` | Yes |  |
| `TransferRecord` | `Hash` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.SendTransfer.create({
  "ErrorCodes" => [], # Array
  "ResultCode" => 1, # Integer
  "TransferRecord" => {}, # Hash
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
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = DingconnectSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```

