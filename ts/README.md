# Dingconnect TypeScript SDK



The TypeScript SDK for the Dingconnect API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.AccountLookup()` — each with a small set of operations (`list`, `create`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/dingconnect-sdk/releases](https://github.com/voxgig-sdk/dingconnect-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { DingconnectSDK } from '@voxgig-sdk/dingconnect'

const client = new DingconnectSDK({
  apikey: process.env.DINGCONNECT_APIKEY,
})
```

### 2. List accountlookup records

`list()` resolves to an array of AccountLookup ENTITIES — every operation
resolves to entities, not raw records. Iterate them directly, and call
`.data()` on one for the record it holds:

```ts
const accountlookups = await client.AccountLookup().list()

for (const accountlookup of accountlookups) {
  console.log(accountlookup)
}
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const currencys = await client.Currency().list()
  console.log(currencys)
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = DingconnectSDK.test()

const currency = await client.Currency().list()
// currency is the entity, populated with mock response data
// — call currency.data() for the record itself
console.log(currency)
```

You can also use the instance method:

```ts
const client = new DingconnectSDK({ apikey: '...' })
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.Currency()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new DingconnectSDK({
  apikey: '...',
  extend: [logger],
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
cd ts && npm test
```


## Reference

### DingconnectSDK

#### Constructor

```ts
new DingconnectSDK(options?: {
  apikey?: string
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `AccountLookup(data?)` | `AccountLookupEntity` | Create an AccountLookup entity instance. |
| `Balance(data?)` | `BalanceEntity` | Create a Balance entity instance. |
| `CancelTransfer(data?)` | `CancelTransferEntity` | Create a CancelTransfer entity instance. |
| `Country(data?)` | `CountryEntity` | Create a Country entity instance. |
| `Currency(data?)` | `CurrencyEntity` | Create a Currency entity instance. |
| `ErrorCodeDescription(data?)` | `ErrorCodeDescriptionEntity` | Create an ErrorCodeDescription entity instance. |
| `EstimatePrice(data?)` | `EstimatePriceEntity` | Create an EstimatePrice entity instance. |
| `ListTransferRecord(data?)` | `ListTransferRecordEntity` | Create a ListTransferRecord entity instance. |
| `LookupBill(data?)` | `LookupBillEntity` | Create a LookupBill entity instance. |
| `Product(data?)` | `ProductEntity` | Create a Product entity instance. |
| `ProductDescription(data?)` | `ProductDescriptionEntity` | Create a ProductDescription entity instance. |
| `Promotion(data?)` | `PromotionEntity` | Create a Promotion entity instance. |
| `PromotionDescription(data?)` | `PromotionDescriptionEntity` | Create a PromotionDescription entity instance. |
| `Provider(data?)` | `ProviderEntity` | Create a Provider entity instance. |
| `ProviderStatus(data?)` | `ProviderStatusEntity` | Create a ProviderStatus entity instance. |
| `Region(data?)` | `RegionEntity` | Create a Region entity instance. |
| `SendTransfer(data?)` | `SendTransferEntity` | Create a SendTransfer entity instance. |
| `tester(testopts?, sdkopts?)` | `DingconnectSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `DingconnectSDK.test(testopts?, sdkopts?)` | `DingconnectSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `create` | `create(reqdata?, ctrl?): Promise<Entity>` | Create a new entity. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): DingconnectSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `create` resolves to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

### Entities

#### AccountLookup

| Field | Description |
| --- | --- |
| `AccountNumberNormalized` |  |
| `CountryIso` |  |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: list.

API path: `/api/V1/GetAccountLookup`

#### Balance

| Field | Description |
| --- | --- |
| `Code` |  |
| `Context` |  |

Operations: list.

API path: `/api/V1/GetBalance`

#### CancelTransfer

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: create.

API path: `/api/V1/CancelTransfers`

#### Country

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: list.

API path: `/api/V1/GetCountries`

#### Currency

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: list.

API path: `/api/V1/GetCurrencies`

#### ErrorCodeDescription

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: list.

API path: `/api/V1/GetErrorCodeDescriptions`

#### EstimatePrice

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: create.

API path: `/api/V1/EstimatePrices`

#### ListTransferRecord

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |
| `ThereAreMoreItems` |  |

Operations: create.

API path: `/api/V1/ListTransferRecords`

#### LookupBill

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: create.

API path: `/api/V1/LookupBills`

#### Product

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: list.

API path: `/api/V1/GetProducts`

#### ProductDescription

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: list.

API path: `/api/V1/GetProductDescriptions`

#### Promotion

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: list.

API path: `/api/V1/GetPromotions`

#### PromotionDescription

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: list.

API path: `/api/V1/GetPromotionDescriptions`

#### Provider

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: list.

API path: `/api/V1/GetProviders`

#### ProviderStatus

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: list.

API path: `/api/V1/GetProviderStatus`

#### Region

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `Items` |  |
| `ResultCode` |  |

Operations: list.

API path: `/api/V1/GetRegions`

#### SendTransfer

| Field | Description |
| --- | --- |
| `ErrorCodes` |  |
| `ResultCode` |  |
| `TransferRecord` |  |

Operations: create.

API path: `/api/V1/SendTransfer`



## Entities


### AccountLookup

Create an instance: `const account_lookup = client.AccountLookup()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `AccountNumberNormalized` | `string` |  |
| `CountryIso` | `string` |  |
| `ErrorCodes` | `any[]` |  |
| `Items` | `any[]` |  |
| `ResultCode` | `number` |  |

#### Example: List

```ts
const account_lookups = await client.AccountLookup().list()
```


### Balance

Create an instance: `const balance = client.Balance()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `Code` | `string` |  |
| `Context` | `string` |  |

#### Example: List

```ts
const balances = await client.Balance().list()
```


### CancelTransfer

Create an instance: `const cancel_transfer = client.CancelTransfer()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `any[]` |  |
| `Items` | `any[]` |  |
| `ResultCode` | `number` |  |

#### Example: Create

```ts
const cancel_transfer = await client.CancelTransfer().create({
  ErrorCodes: [],
  Items: [],
  ResultCode: 1,
})
```


### Country

Create an instance: `const country = client.Country()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `any[]` |  |
| `Items` | `any[]` |  |
| `ResultCode` | `number` |  |

#### Example: List

```ts
const countrys = await client.Country().list()
```


### Currency

Create an instance: `const currency = client.Currency()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `any[]` |  |
| `Items` | `any[]` |  |
| `ResultCode` | `number` |  |

#### Example: List

```ts
const currencys = await client.Currency().list()
```


### ErrorCodeDescription

Create an instance: `const error_code_description = client.ErrorCodeDescription()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `any[]` |  |
| `Items` | `any[]` |  |
| `ResultCode` | `number` |  |

#### Example: List

```ts
const error_code_descriptions = await client.ErrorCodeDescription().list()
```


### EstimatePrice

Create an instance: `const estimate_price = client.EstimatePrice()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `any[]` |  |
| `Items` | `any[]` |  |
| `ResultCode` | `number` |  |

#### Example: Create

```ts
const estimate_price = await client.EstimatePrice().create({
  ErrorCodes: [],
  Items: [],
  ResultCode: 1,
})
```


### ListTransferRecord

Create an instance: `const list_transfer_record = client.ListTransferRecord()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `any[]` |  |
| `Items` | `any[]` |  |
| `ResultCode` | `number` |  |
| `ThereAreMoreItems` | `boolean` |  |

#### Example: Create

```ts
const list_transfer_record = await client.ListTransferRecord().create({
  ErrorCodes: [],
  Items: [],
  ResultCode: 1,
  ThereAreMoreItems: true,
})
```


### LookupBill

Create an instance: `const lookup_bill = client.LookupBill()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `any[]` |  |
| `Items` | `any[]` |  |
| `ResultCode` | `number` |  |

#### Example: Create

```ts
const lookup_bill = await client.LookupBill().create({
  ErrorCodes: [],
  Items: [],
  ResultCode: 1,
})
```


### Product

Create an instance: `const product = client.Product()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `any[]` |  |
| `Items` | `any[]` |  |
| `ResultCode` | `number` |  |

#### Example: List

```ts
const products = await client.Product().list()
```


### ProductDescription

Create an instance: `const product_description = client.ProductDescription()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `any[]` |  |
| `Items` | `any[]` |  |
| `ResultCode` | `number` |  |

#### Example: List

```ts
const product_descriptions = await client.ProductDescription().list()
```


### Promotion

Create an instance: `const promotion = client.Promotion()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `any[]` |  |
| `Items` | `any[]` |  |
| `ResultCode` | `number` |  |

#### Example: List

```ts
const promotions = await client.Promotion().list()
```


### PromotionDescription

Create an instance: `const promotion_description = client.PromotionDescription()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `any[]` |  |
| `Items` | `any[]` |  |
| `ResultCode` | `number` |  |

#### Example: List

```ts
const promotion_descriptions = await client.PromotionDescription().list()
```


### Provider

Create an instance: `const provider = client.Provider()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `any[]` |  |
| `Items` | `any[]` |  |
| `ResultCode` | `number` |  |

#### Example: List

```ts
const providers = await client.Provider().list()
```


### ProviderStatus

Create an instance: `const provider_status = client.ProviderStatus()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `any[]` |  |
| `Items` | `any[]` |  |
| `ResultCode` | `number` |  |

#### Example: List

```ts
const provider_statuss = await client.ProviderStatus().list()
```


### Region

Create an instance: `const region = client.Region()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `any[]` |  |
| `Items` | `any[]` |  |
| `ResultCode` | `number` |  |

#### Example: List

```ts
const regions = await client.Region().list()
```


### SendTransfer

Create an instance: `const send_transfer = client.SendTransfer()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ErrorCodes` | `any[]` |  |
| `ResultCode` | `number` |  |
| `TransferRecord` | `Record<string, any>` |  |

#### Example: Create

```ts
const send_transfer = await client.SendTransfer().create({
  ErrorCodes: [],
  ResultCode: 1,
  TransferRecord: {},
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

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
dingconnect/
├── src/
│   ├── DingconnectSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { DingconnectSDK } from '@voxgig-sdk/dingconnect'
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const currency = client.Currency()
await currency.list()

// currency.data() now returns the currency data from the last `list`
// currency.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
