# Dingconnect TypeScript SDK Reference

Complete API reference for the Dingconnect TypeScript SDK.


## DingconnectSDK

### Constructor

```ts
new DingconnectSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `DingconnectSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = DingconnectSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `DingconnectSDK` instance in test mode.


### Instance Methods

#### `AccountLookup(data?: object)`

Create a new `AccountLookup` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AccountLookupEntity` instance.

#### `Balance(data?: object)`

Create a new `Balance` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `BalanceEntity` instance.

#### `CancelTransfer(data?: object)`

Create a new `CancelTransfer` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CancelTransferEntity` instance.

#### `Country(data?: object)`

Create a new `Country` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CountryEntity` instance.

#### `Currency(data?: object)`

Create a new `Currency` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CurrencyEntity` instance.

#### `ErrorCodeDescription(data?: object)`

Create a new `ErrorCodeDescription` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ErrorCodeDescriptionEntity` instance.

#### `EstimatePrice(data?: object)`

Create a new `EstimatePrice` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EstimatePriceEntity` instance.

#### `ListTransferRecord(data?: object)`

Create a new `ListTransferRecord` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ListTransferRecordEntity` instance.

#### `LookupBill(data?: object)`

Create a new `LookupBill` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `LookupBillEntity` instance.

#### `Product(data?: object)`

Create a new `Product` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ProductEntity` instance.

#### `ProductDescription(data?: object)`

Create a new `ProductDescription` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ProductDescriptionEntity` instance.

#### `Promotion(data?: object)`

Create a new `Promotion` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PromotionEntity` instance.

#### `PromotionDescription(data?: object)`

Create a new `PromotionDescription` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PromotionDescriptionEntity` instance.

#### `Provider(data?: object)`

Create a new `Provider` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ProviderEntity` instance.

#### `ProviderStatus(data?: object)`

Create a new `ProviderStatus` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ProviderStatusEntity` instance.

#### `Region(data?: object)`

Create a new `Region` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `RegionEntity` instance.

#### `SendTransfer(data?: object)`

Create a new `SendTransfer` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SendTransferEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `DingconnectSDK.test()`.

**Returns:** `DingconnectSDK` instance in test mode.


---

## AccountLookupEntity

```ts
const account_lookup = client.AccountLookup()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `AccountNumberNormalized` | `string` | No | We attempt to normalize phone numbers following the public telecommunication numbering plan <a href="https://en.wikipedia.org/wiki/E.164" target="_blank">E.164</a>, if we succeed the normalized number will be returned in this field formatt… |
| `CountryIso` | `string` | No | The country of the account number |
| `ErrorCodes` | `any[]` | Yes |  |
| `Items` | `any[]` | Yes | This will contain provider information associated to the account number. |
| `ResultCode` | `number` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.AccountLookup().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AccountLookupEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## BalanceEntity

```ts
const balance = client.Balance()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `Code` | `string` | Yes | The code that can be used to lookup the explanatory message associated with the error |
| `Context` | `string` | No | API specific context as to the reason for the specific code |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Balance().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `BalanceEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CancelTransferEntity

```ts
const cancel_transfer = client.CancelTransfer()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `any[]` | Yes |  |
| `Items` | `any[]` | Yes |  |
| `ResultCode` | `number` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CancelTransfer().create({
  ErrorCodes: [],
  Items: [],
  ResultCode: 1,
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CancelTransferEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CountryEntity

```ts
const country = client.Country()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `any[]` | Yes |  |
| `Items` | `any[]` | Yes | The list of countries that our system is aware of. |
| `ResultCode` | `number` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Country().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CountryEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CurrencyEntity

```ts
const currency = client.Currency()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `any[]` | Yes |  |
| `Items` | `any[]` | Yes |  |
| `ResultCode` | `number` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Currency().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CurrencyEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ErrorCodeDescriptionEntity

```ts
const error_code_description = client.ErrorCodeDescription()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `any[]` | Yes |  |
| `Items` | `any[]` | Yes | A list of ErrorCodes and their localized descriptions |
| `ResultCode` | `number` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.ErrorCodeDescription().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ErrorCodeDescriptionEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EstimatePriceEntity

```ts
const estimate_price = client.EstimatePrice()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `any[]` | Yes |  |
| `Items` | `any[]` | Yes |  |
| `ResultCode` | `number` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EstimatePrice().create({
  ErrorCodes: [],
  Items: [],
  ResultCode: 1,
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EstimatePriceEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ListTransferRecordEntity

```ts
const list_transfer_record = client.ListTransferRecord()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `any[]` | Yes |  |
| `Items` | `any[]` | Yes | The list of items satisfying the transfer query. |
| `ResultCode` | `number` | Yes |  |
| `ThereAreMoreItems` | `boolean` | Yes | Indicates if the caller should execute the query again. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ListTransferRecord().create({
  ErrorCodes: [],
  Items: [],
  ResultCode: 1,
  ThereAreMoreItems: true,
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ListTransferRecordEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## LookupBillEntity

```ts
const lookup_bill = client.LookupBill()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `any[]` | Yes |  |
| `Items` | `any[]` | Yes |  |
| `ResultCode` | `number` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.LookupBill().create({
  ErrorCodes: [],
  Items: [],
  ResultCode: 1,
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `LookupBillEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ProductEntity

```ts
const product = client.Product()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `any[]` | Yes |  |
| `Items` | `any[]` | Yes | A list of products that fulfil the submitted criteria. |
| `ResultCode` | `number` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Product().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ProductEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ProductDescriptionEntity

```ts
const product_description = client.ProductDescription()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `any[]` | Yes |  |
| `Items` | `any[]` | Yes | A localized list of product descriptions. |
| `ResultCode` | `number` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.ProductDescription().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ProductDescriptionEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PromotionEntity

```ts
const promotion = client.Promotion()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `any[]` | Yes |  |
| `Items` | `any[]` | Yes | List of available promotions |
| `ResultCode` | `number` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Promotion().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PromotionEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PromotionDescriptionEntity

```ts
const promotion_description = client.PromotionDescription()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `any[]` | Yes |  |
| `Items` | `any[]` | Yes | A localized list of promotions. |
| `ResultCode` | `number` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PromotionDescription().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PromotionDescriptionEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ProviderEntity

```ts
const provider = client.Provider()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `any[]` | Yes |  |
| `Items` | `any[]` | Yes | A list of providers that the distributor has Products for. |
| `ResultCode` | `number` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Provider().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ProviderEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ProviderStatusEntity

```ts
const provider_status = client.ProviderStatus()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `any[]` | Yes |  |
| `Items` | `any[]` | Yes |  |
| `ResultCode` | `number` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.ProviderStatus().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ProviderStatusEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## RegionEntity

```ts
const region = client.Region()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `any[]` | Yes |  |
| `Items` | `any[]` | Yes | The list of regions that the system uses. |
| `ResultCode` | `number` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Region().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `RegionEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SendTransferEntity

```ts
const send_transfer = client.SendTransfer()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `any[]` | Yes |  |
| `ResultCode` | `number` | Yes |  |
| `TransferRecord` | `Record<string, any>` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SendTransfer().create({
  ErrorCodes: [],
  ResultCode: 1,
  TransferRecord: {},
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SendTransferEntity` instance with the same client and
options.

#### `client()`

Return the parent `DingconnectSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new DingconnectSDK({
  feature: {
    test: { active: true },
  }
})
```

