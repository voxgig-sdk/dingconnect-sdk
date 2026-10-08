# Dingconnect PHP SDK Reference

Complete API reference for the Dingconnect PHP SDK.


## DingconnectSDK

### Constructor

```php
require_once __DIR__ . '/dingconnect_sdk.php';

$client = new DingconnectSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `DingconnectSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = DingconnectSDK::test();
```


### Instance Methods

#### `AccountLookup($data = null)`

Create a new `AccountLookupEntity` instance. Pass `null` for no initial data.

#### `Balance($data = null)`

Create a new `BalanceEntity` instance. Pass `null` for no initial data.

#### `CancelTransfer($data = null)`

Create a new `CancelTransferEntity` instance. Pass `null` for no initial data.

#### `Country($data = null)`

Create a new `CountryEntity` instance. Pass `null` for no initial data.

#### `Currency($data = null)`

Create a new `CurrencyEntity` instance. Pass `null` for no initial data.

#### `ErrorCodeDescription($data = null)`

Create a new `ErrorCodeDescriptionEntity` instance. Pass `null` for no initial data.

#### `EstimatePrice($data = null)`

Create a new `EstimatePriceEntity` instance. Pass `null` for no initial data.

#### `ListTransferRecord($data = null)`

Create a new `ListTransferRecordEntity` instance. Pass `null` for no initial data.

#### `LookupBill($data = null)`

Create a new `LookupBillEntity` instance. Pass `null` for no initial data.

#### `Product($data = null)`

Create a new `ProductEntity` instance. Pass `null` for no initial data.

#### `ProductDescription($data = null)`

Create a new `ProductDescriptionEntity` instance. Pass `null` for no initial data.

#### `Promotion($data = null)`

Create a new `PromotionEntity` instance. Pass `null` for no initial data.

#### `PromotionDescription($data = null)`

Create a new `PromotionDescriptionEntity` instance. Pass `null` for no initial data.

#### `Provider($data = null)`

Create a new `ProviderEntity` instance. Pass `null` for no initial data.

#### `ProviderStatus($data = null)`

Create a new `ProviderStatusEntity` instance. Pass `null` for no initial data.

#### `Region($data = null)`

Create a new `RegionEntity` instance. Pass `null` for no initial data.

#### `SendTransfer($data = null)`

Create a new `SendTransferEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): DingconnectUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## AccountLookupEntity

```php
$account_lookup = $client->AccountLookup();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `AccountNumberNormalized` | `string` | No | We attempt to normalize phone numbers following the public telecommunication numbering plan <a href="https://en.wikipedia.org/wiki/E.164" target="_blank">E.164</a>, if we succeed the normalized number will be returned in this field formatt… |
| `CountryIso` | `string` | No | The country of the account number |
| `ErrorCodes` | `array` | Yes |  |
| `Items` | `array` | Yes | This will contain provider information associated to the account number. |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record, and throws on error.

```php
$results = $client->AccountLookup()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): AccountLookupEntity`

Create a new `AccountLookupEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## BalanceEntity

```php
$balance = $client->Balance();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `Code` | `string` | Yes | The code that can be used to lookup the explanatory message associated with the error |
| `Context` | `string` | No | API specific context as to the reason for the specific code |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record, and throws on error.

```php
$results = $client->Balance()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): BalanceEntity`

Create a new `BalanceEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CancelTransferEntity

```php
$cancel_transfer = $client->CancelTransfer();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `array` | Yes |  |
| `Items` | `array` | Yes |  |
| `ResultCode` | `int` | Yes |  |
| `cancellations` | `array` | No | An explicit list of records to cancel. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Returns the created entity and throws on error.

```php
$result = $client->CancelTransfer()->create([
  "ErrorCodes" => null, // array
  "Items" => null, // array
  "ResultCode" => null, // int
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CancelTransferEntity`

Create a new `CancelTransferEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CountryEntity

```php
$country = $client->Country();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `array` | Yes |  |
| `Items` | `array` | Yes | The list of countries that our system is aware of. |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record, and throws on error.

```php
$results = $client->Country()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CountryEntity`

Create a new `CountryEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CurrencyEntity

```php
$currency = $client->Currency();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `array` | Yes |  |
| `Items` | `array` | Yes |  |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record, and throws on error.

```php
$results = $client->Currency()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CurrencyEntity`

Create a new `CurrencyEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ErrorCodeDescriptionEntity

```php
$error_code_description = $client->ErrorCodeDescription();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `array` | Yes |  |
| `Items` | `array` | Yes | A list of ErrorCodes and their localized descriptions |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record, and throws on error.

```php
$results = $client->ErrorCodeDescription()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ErrorCodeDescriptionEntity`

Create a new `ErrorCodeDescriptionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EstimatePriceEntity

```php
$estimate_price = $client->EstimatePrice();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `array` | Yes |  |
| `Items` | `array` | Yes |  |
| `ResultCode` | `int` | Yes |  |
| `estimations` | `array` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Returns the created entity and throws on error.

```php
$result = $client->EstimatePrice()->create([
  "ErrorCodes" => null, // array
  "Items" => null, // array
  "ResultCode" => null, // int
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EstimatePriceEntity`

Create a new `EstimatePriceEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ListTransferRecordEntity

```php
$list_transfer_record = $client->ListTransferRecord();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `AccountNumber` | `string` | No | Filter transfers by AccountNumber |
| `DistributorRef` | `string` | No | Filter transfers by DistributorRef. |
| `ErrorCodes` | `array` | Yes |  |
| `Items` | `array` | Yes | The list of items satisfying the transfer query. |
| `ResultCode` | `int` | Yes |  |
| `Skip` | `int` | No | The amount of records to by-pass before returning the remaining records |
| `Take` | `int` | Yes | The amount of records to return |
| `ThereAreMoreItems` | `bool` | Yes | Indicates if the caller should execute the query again. |
| `TransferRef` | `string` | No | Filter by Ding TransferRef |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Returns the created entity and throws on error.

```php
$result = $client->ListTransferRecord()->create([
  "ErrorCodes" => null, // array
  "Items" => null, // array
  "ResultCode" => null, // int
  "Take" => null, // int
  "ThereAreMoreItems" => null, // bool
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ListTransferRecordEntity`

Create a new `ListTransferRecordEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## LookupBillEntity

```php
$lookup_bill = $client->LookupBill();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `AccountNumber` | `string` | Yes | The account number to target |
| `ErrorCodes` | `array` | Yes |  |
| `Items` | `array` | Yes |  |
| `ResultCode` | `int` | Yes |  |
| `Settings` | `array` | No | Product specific name/value pairs to be associated with the lookup bills request |
| `SkuCode` | `string` | Yes | Code provided by GetProducts API |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Returns the created entity and throws on error.

```php
$result = $client->LookupBill()->create([
  "AccountNumber" => null, // string
  "ErrorCodes" => null, // array
  "Items" => null, // array
  "ResultCode" => null, // int
  "SkuCode" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): LookupBillEntity`

Create a new `LookupBillEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ProductEntity

```php
$product = $client->Product();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `array` | Yes |  |
| `Items` | `array` | Yes | A list of products that fulfil the submitted criteria. |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record, and throws on error.

```php
$results = $client->Product()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ProductEntity`

Create a new `ProductEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ProductDescriptionEntity

```php
$product_description = $client->ProductDescription();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `array` | Yes |  |
| `Items` | `array` | Yes | A localized list of product descriptions. |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record, and throws on error.

```php
$results = $client->ProductDescription()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ProductDescriptionEntity`

Create a new `ProductDescriptionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PromotionEntity

```php
$promotion = $client->Promotion();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `array` | Yes |  |
| `Items` | `array` | Yes | List of available promotions |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record, and throws on error.

```php
$results = $client->Promotion()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PromotionEntity`

Create a new `PromotionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PromotionDescriptionEntity

```php
$promotion_description = $client->PromotionDescription();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `array` | Yes |  |
| `Items` | `array` | Yes | A localized list of promotions. |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record, and throws on error.

```php
$results = $client->PromotionDescription()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PromotionDescriptionEntity`

Create a new `PromotionDescriptionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ProviderEntity

```php
$provider = $client->Provider();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `array` | Yes |  |
| `Items` | `array` | Yes | A list of providers that the distributor has Products for. |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record, and throws on error.

```php
$results = $client->Provider()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ProviderEntity`

Create a new `ProviderEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ProviderStatusEntity

```php
$provider_status = $client->ProviderStatus();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `array` | Yes |  |
| `Items` | `array` | Yes |  |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record, and throws on error.

```php
$results = $client->ProviderStatus()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ProviderStatusEntity`

Create a new `ProviderStatusEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## RegionEntity

```php
$region = $client->Region();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `array` | Yes |  |
| `Items` | `array` | Yes | The list of regions that the system uses. |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array of entities, one per record, and throws on error.

```php
$results = $client->Region()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): RegionEntity`

Create a new `RegionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SendTransferEntity

```php
$send_transfer = $client->SendTransfer();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `AccountNumber` | `string` | Yes | The account number to target |
| `BillRef` | `string` | No | Bill reference. |
| `DistributorRef` | `string` | Yes | Unique identifier in the distributor system to be associated with the transfer |
| `ErrorCodes` | `array` | Yes |  |
| `ResultCode` | `int` | Yes |  |
| `SendCurrencyIso` | `string` | No | The currency of the `SendValue`. |
| `SendValue` | `float` | Yes | The transfer value to be sent. |
| `Settings` | `array` | No | Product specific name/value pairs to be associated with the transfer request |
| `SkuCode` | `string` | Yes | Code provided by GetProducts API |
| `TransferRecord` | `array` | Yes |  |
| `ValidateOnly` | `bool` | Yes | Validate the request with the provider without doing a transfer |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Returns the created entity and throws on error.

```php
$result = $client->SendTransfer()->create([
  "AccountNumber" => null, // string
  "DistributorRef" => null, // string
  "ErrorCodes" => null, // array
  "ResultCode" => null, // int
  "SendValue" => null, // float
  "SkuCode" => null, // string
  "TransferRecord" => null, // array
  "ValidateOnly" => null, // bool
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SendTransferEntity`

Create a new `SendTransferEntity` instance with the same client and
options.

#### `get_name(): string`

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

```php
$client = new DingconnectSDK([
  "feature" => [
    "debug" => ["active" => true],
    "idempotency" => ["active" => true],
    "metrics" => ["active" => true],
    "paging" => ["active" => true],
    "ratelimit" => ["active" => true],
    "retry" => ["active" => true],
    "test" => ["active" => true],
    "timeout" => ["active" => true],
  ],
]);
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

