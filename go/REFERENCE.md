# Dingconnect Golang SDK Reference

Complete API reference for the Dingconnect Golang SDK.


## DingconnectSDK

### Constructor

```go
func NewDingconnectSDK(options map[string]any) *DingconnectSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *DingconnectSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *DingconnectSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `AccountLookup(data map[string]any) DingconnectEntity`

Create a new `AccountLookup` entity instance. Pass `nil` for no initial data.

#### `Balance(data map[string]any) DingconnectEntity`

Create a new `Balance` entity instance. Pass `nil` for no initial data.

#### `CancelResult(data map[string]any) DingconnectEntity`

Create a new `CancelResult` entity instance. Pass `nil` for no initial data.

#### `Country(data map[string]any) DingconnectEntity`

Create a new `Country` entity instance. Pass `nil` for no initial data.

#### `Currency(data map[string]any) DingconnectEntity`

Create a new `Currency` entity instance. Pass `nil` for no initial data.

#### `ErrorCodeDescription(data map[string]any) DingconnectEntity`

Create a new `ErrorCodeDescription` entity instance. Pass `nil` for no initial data.

#### `Estimate(data map[string]any) DingconnectEntity`

Create a new `Estimate` entity instance. Pass `nil` for no initial data.

#### `Product(data map[string]any) DingconnectEntity`

Create a new `Product` entity instance. Pass `nil` for no initial data.

#### `ProductDescription(data map[string]any) DingconnectEntity`

Create a new `ProductDescription` entity instance. Pass `nil` for no initial data.

#### `Promotion(data map[string]any) DingconnectEntity`

Create a new `Promotion` entity instance. Pass `nil` for no initial data.

#### `PromotionDescription(data map[string]any) DingconnectEntity`

Create a new `PromotionDescription` entity instance. Pass `nil` for no initial data.

#### `Provider(data map[string]any) DingconnectEntity`

Create a new `Provider` entity instance. Pass `nil` for no initial data.

#### `ProviderStatus(data map[string]any) DingconnectEntity`

Create a new `ProviderStatus` entity instance. Pass `nil` for no initial data.

#### `Region(data map[string]any) DingconnectEntity`

Create a new `Region` entity instance. Pass `nil` for no initial data.

#### `SendTransfer(data map[string]any) DingconnectEntity`

Create a new `SendTransfer` entity instance. Pass `nil` for no initial data.

#### `TransferRecord(data map[string]any) DingconnectEntity`

Create a new `TransferRecord` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## AccountLookupEntity

```go
accountLookup := client.AccountLookup(nil)
fmt.Println(accountLookup.GetName()) // "account_lookup"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_number_normalized` | `string` | No |  |
| `country_iso` | `string` | No |  |
| `error_code` | `[]any` | No |  |
| `item` | `[]any` | No |  |
| `result_code` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.AccountLookup(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AccountLookupEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## BalanceEntity

```go
balance := client.Balance(nil)
fmt.Println(balance.GetName()) // "balance"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `code` | `string` | No |  |
| `context` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Balance(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `BalanceEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CancelResultEntity

```go
cancelResult := client.CancelResult(nil)
fmt.Println(cancelResult.GetName()) // "cancel_result"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error_code` | `[]any` | No |  |
| `item` | `[]any` | No |  |
| `result_code` | `int` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CancelResult(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CancelResultEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CountryEntity

```go
country := client.Country(nil)
fmt.Println(country.GetName()) // "country"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error_code` | `[]any` | No |  |
| `item` | `[]any` | No |  |
| `result_code` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Country(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CountryEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CurrencyEntity

```go
currency := client.Currency(nil)
fmt.Println(currency.GetName()) // "currency"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error_code` | `[]any` | No |  |
| `item` | `[]any` | No |  |
| `result_code` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Currency(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CurrencyEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ErrorCodeDescriptionEntity

```go
errorCodeDescription := client.ErrorCodeDescription(nil)
fmt.Println(errorCodeDescription.GetName()) // "error_code_description"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error_code` | `[]any` | No |  |
| `item` | `[]any` | No |  |
| `result_code` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.ErrorCodeDescription(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ErrorCodeDescriptionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EstimateEntity

```go
estimate := client.Estimate(nil)
fmt.Println(estimate.GetName()) // "estimate"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error_code` | `[]any` | No |  |
| `item` | `[]any` | No |  |
| `result_code` | `int` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Estimate(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EstimateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ProductEntity

```go
product := client.Product(nil)
fmt.Println(product.GetName()) // "product"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error_code` | `[]any` | No |  |
| `item` | `[]any` | No |  |
| `result_code` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Product(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ProductEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ProductDescriptionEntity

```go
productDescription := client.ProductDescription(nil)
fmt.Println(productDescription.GetName()) // "product_description"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error_code` | `[]any` | No |  |
| `item` | `[]any` | No |  |
| `result_code` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.ProductDescription(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ProductDescriptionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PromotionEntity

```go
promotion := client.Promotion(nil)
fmt.Println(promotion.GetName()) // "promotion"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error_code` | `[]any` | No |  |
| `item` | `[]any` | No |  |
| `result_code` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Promotion(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PromotionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PromotionDescriptionEntity

```go
promotionDescription := client.PromotionDescription(nil)
fmt.Println(promotionDescription.GetName()) // "promotion_description"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error_code` | `[]any` | No |  |
| `item` | `[]any` | No |  |
| `result_code` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.PromotionDescription(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PromotionDescriptionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ProviderEntity

```go
provider := client.Provider(nil)
fmt.Println(provider.GetName()) // "provider"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error_code` | `[]any` | No |  |
| `item` | `[]any` | No |  |
| `result_code` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Provider(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ProviderEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ProviderStatusEntity

```go
providerStatus := client.ProviderStatus(nil)
fmt.Println(providerStatus.GetName()) // "provider_status"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error_code` | `[]any` | No |  |
| `item` | `[]any` | No |  |
| `result_code` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.ProviderStatus(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ProviderStatusEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## RegionEntity

```go
region := client.Region(nil)
fmt.Println(region.GetName()) // "region"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error_code` | `[]any` | No |  |
| `item` | `[]any` | No |  |
| `result_code` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Region(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `RegionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SendTransferEntity

```go
sendTransfer := client.SendTransfer(nil)
fmt.Println(sendTransfer.GetName()) // "send_transfer"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_number` | `string` | Yes |  |
| `distributor_ref` | `string` | Yes |  |
| `error_code` | `[]any` | No |  |
| `result_code` | `int` | No |  |
| `send_currency_iso` | `string` | No |  |
| `send_value` | `float64` | Yes |  |
| `setting` | `[]any` | No |  |
| `sku_code` | `string` | Yes |  |
| `transfer_record` | `map[string]any` | No |  |
| `validate_only` | `bool` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SendTransfer(nil).Create(map[string]any{
    "account_number": "example_account_number",
    "distributor_ref": "example_distributor_ref",
    "send_value": 1,
    "sku_code": "example_sku_code",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SendTransferEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TransferRecordEntity

```go
transferRecord := client.TransferRecord(nil)
fmt.Println(transferRecord.GetName()) // "transfer_record"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_number` | `string` | No |  |
| `distributor_ref` | `[]any` | No |  |
| `ended_at_utc` | `string` | No |  |
| `error_code` | `[]any` | No |  |
| `item` | `[]any` | No |  |
| `result_code` | `int` | No |  |
| `skip` | `int` | Yes |  |
| `sku_code` | `[]any` | No |  |
| `started_at_utc` | `string` | No |  |
| `take` | `int` | Yes |  |
| `there_are_more_item` | `bool` | No |  |
| `transfer_ref` | `[]any` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.TransferRecord(nil).Create(map[string]any{
    "skip": 1,
    "take": 1,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TransferRecordEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewDingconnectSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

