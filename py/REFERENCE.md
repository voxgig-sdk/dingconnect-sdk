# Dingconnect Python SDK Reference

Complete API reference for the Dingconnect Python SDK.


## DingconnectSDK

### Constructor

```python
from dingconnect_sdk import DingconnectSDK

client = DingconnectSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `DingconnectSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = DingconnectSDK.test()
```


### Instance Methods

#### `AccountLookup(data=None)`

Create a new `AccountLookupEntity` instance. Pass `None` for no initial data.

#### `Balance(data=None)`

Create a new `BalanceEntity` instance. Pass `None` for no initial data.

#### `CancelTransfer(data=None)`

Create a new `CancelTransferEntity` instance. Pass `None` for no initial data.

#### `Country(data=None)`

Create a new `CountryEntity` instance. Pass `None` for no initial data.

#### `Currency(data=None)`

Create a new `CurrencyEntity` instance. Pass `None` for no initial data.

#### `ErrorCodeDescription(data=None)`

Create a new `ErrorCodeDescriptionEntity` instance. Pass `None` for no initial data.

#### `EstimatePrice(data=None)`

Create a new `EstimatePriceEntity` instance. Pass `None` for no initial data.

#### `ListTransferRecord(data=None)`

Create a new `ListTransferRecordEntity` instance. Pass `None` for no initial data.

#### `LookupBill(data=None)`

Create a new `LookupBillEntity` instance. Pass `None` for no initial data.

#### `Product(data=None)`

Create a new `ProductEntity` instance. Pass `None` for no initial data.

#### `ProductDescription(data=None)`

Create a new `ProductDescriptionEntity` instance. Pass `None` for no initial data.

#### `Promotion(data=None)`

Create a new `PromotionEntity` instance. Pass `None` for no initial data.

#### `PromotionDescription(data=None)`

Create a new `PromotionDescriptionEntity` instance. Pass `None` for no initial data.

#### `Provider(data=None)`

Create a new `ProviderEntity` instance. Pass `None` for no initial data.

#### `ProviderStatus(data=None)`

Create a new `ProviderStatusEntity` instance. Pass `None` for no initial data.

#### `Region(data=None)`

Create a new `RegionEntity` instance. Pass `None` for no initial data.

#### `SendTransfer(data=None)`

Create a new `SendTransferEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## AccountLookupEntity

```python
account_lookup = client.AccountLookup()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `AccountNumberNormalized` | `str` | No | We attempt to normalize phone numbers following the public telecommunication numbering plan <a href="https://en.wikipedia.org/wiki/E.164" target="_blank">E.164</a>, if we succeed the normalized number will be returned in this field formatt… |
| `CountryIso` | `str` | No | The country of the account number |
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes | This will contain provider information associated to the account number. |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[AccountLookupEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.AccountLookup().list()
for account_lookup in results:
    print(account_lookup.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AccountLookupEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## BalanceEntity

```python
balance = client.Balance()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `Code` | `str` | Yes | The code that can be used to lookup the explanatory message associated with the error |
| `Context` | `str` | No | API specific context as to the reason for the specific code |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[BalanceEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Balance().list()
for balance in results:
    print(balance.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BalanceEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CancelTransferEntity

```python
cancel_transfer = client.CancelTransfer()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes |  |
| `ResultCode` | `int` | Yes |  |
| `cancellations` | `list` | No | An explicit list of records to cancel. |

### Operations

#### `create(reqdata, ctrl=None) -> CancelTransferEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.CancelTransfer().create({
    "ErrorCodes": [],  # list
    "Items": [],  # list
    "ResultCode": 1,  # int
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CancelTransferEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CountryEntity

```python
country = client.Country()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes | The list of countries that our system is aware of. |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[CountryEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Country().list()
for country in results:
    print(country.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CountryEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CurrencyEntity

```python
currency = client.Currency()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes |  |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[CurrencyEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Currency().list()
for currency in results:
    print(currency.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CurrencyEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ErrorCodeDescriptionEntity

```python
error_code_description = client.ErrorCodeDescription()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes | A list of ErrorCodes and their localized descriptions |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ErrorCodeDescriptionEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ErrorCodeDescription().list()
for error_code_description in results:
    print(error_code_description.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ErrorCodeDescriptionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EstimatePriceEntity

```python
estimate_price = client.EstimatePrice()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes |  |
| `ResultCode` | `int` | Yes |  |
| `estimations` | `list` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> EstimatePriceEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.EstimatePrice().create({
    "ErrorCodes": [],  # list
    "Items": [],  # list
    "ResultCode": 1,  # int
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EstimatePriceEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ListTransferRecordEntity

```python
list_transfer_record = client.ListTransferRecord()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `AccountNumber` | `str` | No | Filter transfers by AccountNumber |
| `DistributorRef` | `str` | No | Filter transfers by DistributorRef. |
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes | The list of items satisfying the transfer query. |
| `ResultCode` | `int` | Yes |  |
| `Skip` | `int` | No | The amount of records to by-pass before returning the remaining records |
| `Take` | `int` | Yes | The amount of records to return |
| `ThereAreMoreItems` | `bool` | Yes | Indicates if the caller should execute the query again. |
| `TransferRef` | `str` | No | Filter by Ding TransferRef |

### Operations

#### `create(reqdata, ctrl=None) -> ListTransferRecordEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.ListTransferRecord().create({
    "ErrorCodes": [],  # list
    "Items": [],  # list
    "ResultCode": 1,  # int
    "Take": 1,  # int
    "ThereAreMoreItems": True,  # bool
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ListTransferRecordEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## LookupBillEntity

```python
lookup_bill = client.LookupBill()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `AccountNumber` | `str` | Yes | The account number to target |
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes |  |
| `ResultCode` | `int` | Yes |  |
| `Settings` | `list` | No | Product specific name/value pairs to be associated with the lookup bills request |
| `SkuCode` | `str` | Yes | Code provided by GetProducts API |

### Operations

#### `create(reqdata, ctrl=None) -> LookupBillEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.LookupBill().create({
    "AccountNumber": "example_AccountNumber",  # str
    "ErrorCodes": [],  # list
    "Items": [],  # list
    "ResultCode": 1,  # int
    "SkuCode": "example_SkuCode",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `LookupBillEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ProductEntity

```python
product = client.Product()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes | A list of products that fulfil the submitted criteria. |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ProductEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Product().list()
for product in results:
    print(product.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProductEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ProductDescriptionEntity

```python
product_description = client.ProductDescription()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes | A localized list of product descriptions. |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ProductDescriptionEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ProductDescription().list()
for product_description in results:
    print(product_description.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProductDescriptionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PromotionEntity

```python
promotion = client.Promotion()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes | List of available promotions |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[PromotionEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Promotion().list()
for promotion in results:
    print(promotion.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PromotionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PromotionDescriptionEntity

```python
promotion_description = client.PromotionDescription()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes | A localized list of promotions. |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[PromotionDescriptionEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.PromotionDescription().list()
for promotion_description in results:
    print(promotion_description.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PromotionDescriptionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ProviderEntity

```python
provider = client.Provider()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes | A list of providers that the distributor has Products for. |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ProviderEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Provider().list()
for provider in results:
    print(provider.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProviderEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ProviderStatusEntity

```python
provider_status = client.ProviderStatus()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes |  |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ProviderStatusEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ProviderStatus().list()
for provider_status in results:
    print(provider_status.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProviderStatusEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RegionEntity

```python
region = client.Region()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes | The list of regions that the system uses. |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[RegionEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Region().list()
for region in results:
    print(region.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RegionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SendTransferEntity

```python
send_transfer = client.SendTransfer()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `AccountNumber` | `str` | Yes | The account number to target |
| `BillRef` | `str` | No | Bill reference. |
| `DistributorRef` | `str` | Yes | Unique identifier in the distributor system to be associated with the transfer |
| `ErrorCodes` | `list` | Yes |  |
| `ResultCode` | `int` | Yes |  |
| `SendCurrencyIso` | `str` | No | The currency of the `SendValue`. |
| `SendValue` | `float` | Yes | The transfer value to be sent. |
| `Settings` | `list` | No | Product specific name/value pairs to be associated with the transfer request |
| `SkuCode` | `str` | Yes | Code provided by GetProducts API |
| `TransferRecord` | `dict` | Yes |  |
| `ValidateOnly` | `bool` | Yes | Validate the request with the provider without doing a transfer |

### Operations

#### `create(reqdata, ctrl=None) -> SendTransferEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.SendTransfer().create({
    "AccountNumber": "example_AccountNumber",  # str
    "DistributorRef": "example_DistributorRef",  # str
    "ErrorCodes": [],  # list
    "ResultCode": 1,  # int
    "SendValue": 1,  # float
    "SkuCode": "example_SkuCode",  # str
    "TransferRecord": {},  # dict
    "ValidateOnly": True,  # bool
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SendTransferEntity` instance with the same options.

#### `get_name() -> str`

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

```python
client = DingconnectSDK({
    "feature": {
        "debug": {"active": True},
        "idempotency": {"active": True},
        "metrics": {"active": True},
        "paging": {"active": True},
        "ratelimit": {"active": True},
        "retry": {"active": True},
        "test": {"active": True},
        "timeout": {"active": True},
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

