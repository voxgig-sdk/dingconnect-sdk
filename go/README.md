# Dingconnect Golang SDK



The Golang SDK for the Dingconnect API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.AccountLookup(nil)` — each with the same small set of operations (`List`, `Create`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Also generated from this model: `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb`, `ts` — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/dingconnect-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Tags](https://github.com/voxgig-sdk/dingconnect-sdk/tags) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/dingconnect-sdk/go=../dingconnect-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record. The client sends the API key in the `api_key` header.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the entity, and for
`List` a `[]any` of entities, one per record (there is no `{ok, data}`
wrapper), so check `err` and read a record through the entity's
`Data()`.

```go
package main

import (
    "fmt"
    "os"
    sdk "github.com/voxgig-sdk/dingconnect-sdk/go"
)

func main() {
    client := sdk.NewDingconnectSDK(map[string]any{
        "apikey": os.Getenv("DINGCONNECT_APIKEY"),
    })

    // List accountLookup records — the value is a []any of entities, one per record.
    accountLookups, err := client.AccountLookup(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }
    for _, item := range accountLookups.([]any) {
        fmt.Println(item.(sdk.Entity).Data())
    }
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
currencys, err := client.Currency(nil).List(nil, nil)
if err != nil {
    // handle err
    return
}
_ = currencys
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

currencys, err := client.Currency(nil).List(
    nil, nil,
)
if err != nil {
    panic(err)
}
// A []any of entities, one per mock record.
for _, item := range currencys.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewDingconnectSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
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
cd go && go test ./test/...
```


## Reference

### NewDingconnectSDK

```go
func NewDingconnectSDK(options map[string]any) *DingconnectSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"apikey"` | `string` | API key for authentication. |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *DingconnectSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### DingconnectSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `AccountLookup` | `(data map[string]any) DingconnectEntity` | Create an AccountLookup entity instance. |
| `Balance` | `(data map[string]any) DingconnectEntity` | Create a Balance entity instance. |
| `CancelTransfer` | `(data map[string]any) DingconnectEntity` | Create a CancelTransfer entity instance. |
| `Country` | `(data map[string]any) DingconnectEntity` | Create a Country entity instance. |
| `Currency` | `(data map[string]any) DingconnectEntity` | Create a Currency entity instance. |
| `ErrorCodeDescription` | `(data map[string]any) DingconnectEntity` | Create an ErrorCodeDescription entity instance. |
| `EstimatePrice` | `(data map[string]any) DingconnectEntity` | Create an EstimatePrice entity instance. |
| `ListTransferRecord` | `(data map[string]any) DingconnectEntity` | Create a ListTransferRecord entity instance. |
| `LookupBill` | `(data map[string]any) DingconnectEntity` | Create a LookupBill entity instance. |
| `Product` | `(data map[string]any) DingconnectEntity` | Create a Product entity instance. |
| `ProductDescription` | `(data map[string]any) DingconnectEntity` | Create a ProductDescription entity instance. |
| `Promotion` | `(data map[string]any) DingconnectEntity` | Create a Promotion entity instance. |
| `PromotionDescription` | `(data map[string]any) DingconnectEntity` | Create a PromotionDescription entity instance. |
| `Provider` | `(data map[string]any) DingconnectEntity` | Create a Provider entity instance. |
| `ProviderStatus` | `(data map[string]any) DingconnectEntity` | Create a ProviderStatus entity instance. |
| `Region` | `(data map[string]any) DingconnectEntity` | Create a Region entity instance. |
| `SendTransfer` | `(data map[string]any) DingconnectEntity` | Create a SendTransfer entity instance. |

### Entity interface (DingconnectEntity)

All entities implement the `DingconnectEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria, one per record. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity, and return it. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the entity
itself — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Create` | the entity, whose `Data()` reads its record (`map[string]any`) |
| `List` | a `[]any` of entities, one per record |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    accountLookup, err := client.AccountLookup(nil).List(nil, nil)
    if err != nil { /* handle */ }
    // accountLookup is a []any of entities, one per record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### AccountLookup

| Field | Description |
| --- | --- |
| `"AccountNumberNormalized"` | We attempt to normalize phone numbers following the public telecommunication numbering plan <a href="https://en.wikipedia.org/wiki/E.164" target="_blank">E.164</a>, if we succeed the normalized number will be returned in this field formatt… |
| `"CountryIso"` | The country of the account number |
| `"ErrorCodes"` |  |
| `"Items"` | This will contain provider information associated to the account number. |
| `"ResultCode"` |  |

Operations: List.

API path: `/api/V1/GetAccountLookup`

#### Balance

| Field | Description |
| --- | --- |
| `"Code"` | The code that can be used to lookup the explanatory message associated with the error |
| `"Context"` | API specific context as to the reason for the specific code |

Operations: List.

API path: `/api/V1/GetBalance`

#### CancelTransfer

| Field | Description |
| --- | --- |
| `"ErrorCodes"` |  |
| `"Items"` |  |
| `"ResultCode"` |  |
| `"cancellations"` | An explicit list of records to cancel. |

Operations: Create.

API path: `/api/V1/CancelTransfers`

#### Country

| Field | Description |
| --- | --- |
| `"ErrorCodes"` |  |
| `"Items"` | The list of countries that our system is aware of. |
| `"ResultCode"` |  |

Operations: List.

API path: `/api/V1/GetCountries`

#### Currency

| Field | Description |
| --- | --- |
| `"ErrorCodes"` |  |
| `"Items"` |  |
| `"ResultCode"` |  |

Operations: List.

API path: `/api/V1/GetCurrencies`

#### ErrorCodeDescription

| Field | Description |
| --- | --- |
| `"ErrorCodes"` |  |
| `"Items"` | A list of ErrorCodes and their localized descriptions |
| `"ResultCode"` |  |

Operations: List.

API path: `/api/V1/GetErrorCodeDescriptions`

#### EstimatePrice

| Field | Description |
| --- | --- |
| `"ErrorCodes"` |  |
| `"Items"` |  |
| `"ResultCode"` |  |
| `"estimations"` |  |

Operations: Create.

API path: `/api/V1/EstimatePrices`

#### ListTransferRecord

| Field | Description |
| --- | --- |
| `"AccountNumber"` | Filter transfers by AccountNumber |
| `"DistributorRef"` | Filter transfers by DistributorRef. |
| `"ErrorCodes"` |  |
| `"Items"` | The list of items satisfying the transfer query. |
| `"ResultCode"` |  |
| `"Skip"` | The amount of records to by-pass before returning the remaining records |
| `"Take"` | The amount of records to return |
| `"ThereAreMoreItems"` | Indicates if the caller should execute the query again. |
| `"TransferRef"` | Filter by Ding TransferRef |

Operations: Create.

API path: `/api/V1/ListTransferRecords`

#### LookupBill

| Field | Description |
| --- | --- |
| `"AccountNumber"` | The account number to target |
| `"ErrorCodes"` |  |
| `"Items"` |  |
| `"ResultCode"` |  |
| `"Settings"` | Product specific name/value pairs to be associated with the lookup bills request |
| `"SkuCode"` | Code provided by GetProducts API |

Operations: Create.

API path: `/api/V1/LookupBills`

#### Product

| Field | Description |
| --- | --- |
| `"ErrorCodes"` |  |
| `"Items"` | A list of products that fulfil the submitted criteria. |
| `"ResultCode"` |  |

Operations: List.

API path: `/api/V1/GetProducts`

#### ProductDescription

| Field | Description |
| --- | --- |
| `"ErrorCodes"` |  |
| `"Items"` | A localized list of product descriptions. |
| `"ResultCode"` |  |

Operations: List.

API path: `/api/V1/GetProductDescriptions`

#### Promotion

| Field | Description |
| --- | --- |
| `"ErrorCodes"` |  |
| `"Items"` | List of available promotions |
| `"ResultCode"` |  |

Operations: List.

API path: `/api/V1/GetPromotions`

#### PromotionDescription

| Field | Description |
| --- | --- |
| `"ErrorCodes"` |  |
| `"Items"` | A localized list of promotions. |
| `"ResultCode"` |  |

Operations: List.

API path: `/api/V1/GetPromotionDescriptions`

#### Provider

| Field | Description |
| --- | --- |
| `"ErrorCodes"` |  |
| `"Items"` | A list of providers that the distributor has Products for. |
| `"ResultCode"` |  |

Operations: List.

API path: `/api/V1/GetProviders`

#### ProviderStatus

| Field | Description |
| --- | --- |
| `"ErrorCodes"` |  |
| `"Items"` |  |
| `"ResultCode"` |  |

Operations: List.

API path: `/api/V1/GetProviderStatus`

#### Region

| Field | Description |
| --- | --- |
| `"ErrorCodes"` |  |
| `"Items"` | The list of regions that the system uses. |
| `"ResultCode"` |  |

Operations: List.

API path: `/api/V1/GetRegions`

#### SendTransfer

| Field | Description |
| --- | --- |
| `"AccountNumber"` | The account number to target |
| `"BillRef"` | Bill reference. |
| `"DistributorRef"` | Unique identifier in the distributor system to be associated with the transfer |
| `"ErrorCodes"` |  |
| `"ResultCode"` |  |
| `"SendCurrencyIso"` | The currency of the `SendValue`. |
| `"SendValue"` | The transfer value to be sent. |
| `"Settings"` | Product specific name/value pairs to be associated with the transfer request |
| `"SkuCode"` | Code provided by GetProducts API |
| `"TransferRecord"` |  |
| `"ValidateOnly"` | Validate the request with the provider without doing a transfer |

Operations: Create.

API path: `/api/V1/SendTransfer`



## Entities


### AccountLookup

Create an instance: `accountLookup := client.AccountLookup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `AccountNumberNormalized` | `string` | We attempt to normalize phone numbers following the public telecommunication numbering plan <a href="https://en.wikipedia.org/wiki/E.164" target="_blank">E.164</a>, if we succeed the normalized number will be returned in this field formatt… |
| `CountryIso` | `string` | The country of the account number |
| `ErrorCodes` | `[]any` |  |
| `Items` | `[]any` | This will contain provider information associated to the account number. |
| `ResultCode` | `int` |  |

#### Example: List

```go
accountLookups, err := client.AccountLookup(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range accountLookups.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### Balance

Create an instance: `balance := client.Balance(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `Code` | `string` | The code that can be used to lookup the explanatory message associated with the error |
| `Context` | `string` | API specific context as to the reason for the specific code |

#### Example: List

```go
balances, err := client.Balance(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range balances.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### CancelTransfer

Create an instance: `cancelTransfer := client.CancelTransfer(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `[]any` |  |
| `Items` | `[]any` |  |
| `ResultCode` | `int` |  |
| `cancellations` | `[]any` | An explicit list of records to cancel. |

#### Example: Create

```go
result, err := client.CancelTransfer(nil).Create(map[string]any{
    "ErrorCodes": []any{},
    "Items": []any{},
    "ResultCode": 1,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Country

Create an instance: `country := client.Country(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `[]any` |  |
| `Items` | `[]any` | The list of countries that our system is aware of. |
| `ResultCode` | `int` |  |

#### Example: List

```go
countrys, err := client.Country(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range countrys.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### Currency

Create an instance: `currency := client.Currency(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `[]any` |  |
| `Items` | `[]any` |  |
| `ResultCode` | `int` |  |

#### Example: List

```go
currencys, err := client.Currency(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range currencys.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### ErrorCodeDescription

Create an instance: `errorCodeDescription := client.ErrorCodeDescription(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `[]any` |  |
| `Items` | `[]any` | A list of ErrorCodes and their localized descriptions |
| `ResultCode` | `int` |  |

#### Example: List

```go
errorCodeDescriptions, err := client.ErrorCodeDescription(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range errorCodeDescriptions.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### EstimatePrice

Create an instance: `estimatePrice := client.EstimatePrice(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `[]any` |  |
| `Items` | `[]any` |  |
| `ResultCode` | `int` |  |
| `estimations` | `[]any` |  |

#### Example: Create

```go
result, err := client.EstimatePrice(nil).Create(map[string]any{
    "ErrorCodes": []any{},
    "Items": []any{},
    "ResultCode": 1,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### ListTransferRecord

Create an instance: `listTransferRecord := client.ListTransferRecord(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `AccountNumber` | `string` | Filter transfers by AccountNumber |
| `DistributorRef` | `string` | Filter transfers by DistributorRef. |
| `ErrorCodes` | `[]any` |  |
| `Items` | `[]any` | The list of items satisfying the transfer query. |
| `ResultCode` | `int` |  |
| `Skip` | `int` | The amount of records to by-pass before returning the remaining records |
| `Take` | `int` | The amount of records to return |
| `ThereAreMoreItems` | `bool` | Indicates if the caller should execute the query again. |
| `TransferRef` | `string` | Filter by Ding TransferRef |

#### Example: Create

```go
result, err := client.ListTransferRecord(nil).Create(map[string]any{
    "ErrorCodes": []any{},
    "Items": []any{},
    "ResultCode": 1,
    "Take": 1,
    "ThereAreMoreItems": true,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### LookupBill

Create an instance: `lookupBill := client.LookupBill(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `AccountNumber` | `string` | The account number to target |
| `ErrorCodes` | `[]any` |  |
| `Items` | `[]any` |  |
| `ResultCode` | `int` |  |
| `Settings` | `[]any` | Product specific name/value pairs to be associated with the lookup bills request |
| `SkuCode` | `string` | Code provided by GetProducts API |

#### Example: Create

```go
result, err := client.LookupBill(nil).Create(map[string]any{
    "AccountNumber": "example_AccountNumber",
    "ErrorCodes": []any{},
    "Items": []any{},
    "ResultCode": 1,
    "SkuCode": "example_SkuCode",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Product

Create an instance: `product := client.Product(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `[]any` |  |
| `Items` | `[]any` | A list of products that fulfil the submitted criteria. |
| `ResultCode` | `int` |  |

#### Example: List

```go
products, err := client.Product(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range products.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### ProductDescription

Create an instance: `productDescription := client.ProductDescription(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `[]any` |  |
| `Items` | `[]any` | A localized list of product descriptions. |
| `ResultCode` | `int` |  |

#### Example: List

```go
productDescriptions, err := client.ProductDescription(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range productDescriptions.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### Promotion

Create an instance: `promotion := client.Promotion(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `[]any` |  |
| `Items` | `[]any` | List of available promotions |
| `ResultCode` | `int` |  |

#### Example: List

```go
promotions, err := client.Promotion(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range promotions.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### PromotionDescription

Create an instance: `promotionDescription := client.PromotionDescription(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `[]any` |  |
| `Items` | `[]any` | A localized list of promotions. |
| `ResultCode` | `int` |  |

#### Example: List

```go
promotionDescriptions, err := client.PromotionDescription(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range promotionDescriptions.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### Provider

Create an instance: `provider := client.Provider(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `[]any` |  |
| `Items` | `[]any` | A list of providers that the distributor has Products for. |
| `ResultCode` | `int` |  |

#### Example: List

```go
providers, err := client.Provider(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range providers.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### ProviderStatus

Create an instance: `providerStatus := client.ProviderStatus(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `[]any` |  |
| `Items` | `[]any` |  |
| `ResultCode` | `int` |  |

#### Example: List

```go
providerStatuss, err := client.ProviderStatus(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range providerStatuss.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### Region

Create an instance: `region := client.Region(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `[]any` |  |
| `Items` | `[]any` | The list of regions that the system uses. |
| `ResultCode` | `int` |  |

#### Example: List

```go
regions, err := client.Region(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range regions.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### SendTransfer

Create an instance: `sendTransfer := client.SendTransfer(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `AccountNumber` | `string` | The account number to target |
| `BillRef` | `string` | Bill reference. |
| `DistributorRef` | `string` | Unique identifier in the distributor system to be associated with the transfer |
| `ErrorCodes` | `[]any` |  |
| `ResultCode` | `int` |  |
| `SendCurrencyIso` | `string` | The currency of the `SendValue`. |
| `SendValue` | `float64` | The transfer value to be sent. |
| `Settings` | `[]any` | Product specific name/value pairs to be associated with the transfer request |
| `SkuCode` | `string` | Code provided by GetProducts API |
| `TransferRecord` | `map[string]any` |  |
| `ValidateOnly` | `bool` | Validate the request with the provider without doing a transfer |

#### Example: Create

```go
result, err := client.SendTransfer(nil).Create(map[string]any{
    "AccountNumber": "example_AccountNumber",
    "DistributorRef": "example_DistributorRef",
    "ErrorCodes": []any{},
    "ResultCode": 1,
    "SendValue": 1,
    "SkuCode": "example_SkuCode",
    "TransferRecord": map[string]any{},
    "ValidateOnly": true,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```

## Features

This SDK ships 8 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`debug`](#debug) | Debug capture |
| [`idempotency`](#idempotency) | Idempotency |
| [`metrics`](#metrics) | Metrics |
| [`paging`](#paging) | Paging |
| [`ratelimit`](#ratelimit) | Rate limiting |
| [`retry`](#retry) | Retry |
| [`test`](#test) | Test transport |
| [`timeout`](#timeout) | Timeout |

> **Order matters for `ratelimit`, `retry`, `timeout`.** These wrap the
> transport, so each one wraps whatever is already installed: the order you
> activate them in IS the nesting order. Activating them as an ordered list
> rather than a map is what fixes that order.

### debug

Debug capture.

| Option | Default |
|---|---|
| `active` | `false` |
| `max` | `100` |
| `redact` | `['authorization', 'cookie', 'set-cookie', 'api-key', 'apikey', 'x-api-key', 'idempotency-key']` |

Set `feature.debug.active` to enable it, then override any of the options above.

### idempotency

Idempotency.

| Option | Default |
|---|---|
| `active` | `false` |
| `header` | `'Idempotency-Key'` |
| `methods` | `['POST', 'PUT', 'PATCH', 'DELETE']` |
| `ops` | `['create', 'update', 'remove']` |

Set `feature.idempotency.active` to enable it, then override any of the options above.

### metrics

Metrics.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.metrics.active` to enable it, then override any of the options above.

### paging

Paging.

| Option | Default |
|---|---|
| `active` | `false` |
| `afterVar` | `'after'` |
| `cursorParam` | `'cursor'` |
| `firstVar` | `'first'` |
| `limitParam` | `'limit'` |
| `pageParam` | `'page'` |
| `startPage` | `1` |

Set `feature.paging.active` to enable it, then override any of the options above.

### ratelimit

Rate limiting.

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

Set `feature.ratelimit.active` to enable it, then override any of the options above.

`ratelimit` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### retry

Retry.

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

Set `feature.retry.active` to enable it, then override any of the options above.

`retry` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### test

Test transport.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.

### timeout

Timeout.

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

Set `feature.timeout.active` to enable it, then override any of the options above.

`timeout` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.


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

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **DebugFeature**: Debug capture
- **IdempotencyFeature**: Idempotency
- **MetricsFeature**: Metrics
- **PagingFeature**: Paging
- **RatelimitFeature**: Rate limiting
- **RetryFeature**: Retry
- **TestFeature**: Test transport
- **TimeoutFeature**: Timeout

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

An operation returns the entity, and its `Data()` returns the record. Use
`core.ToMapAny()` to safely cast that record, or data nested in it, to
`map[string]any`: it returns `nil` for anything else, an entity included.

### Package structure

```
github.com/voxgig-sdk/dingconnect-sdk/go/
├── dingconnect.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/dingconnect-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `List`, the entity
stores the returned data and match criteria internally.

```go
currency := client.Currency(nil)
currency.List(nil, nil)

// currency.Data() now returns the currency data from the last list
// currency.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
