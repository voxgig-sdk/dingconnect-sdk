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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.AccountLookup().list()
for account_lookup in results:
    print(account_lookup)
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Balance().list()
for balance in results:
    print(balance)
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

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Country().list()
for country in results:
    print(country)
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Currency().list()
for currency in results:
    print(currency)
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.ErrorCodeDescription().list()
for error_code_description in results:
    print(error_code_description)
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

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

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
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes | The list of items satisfying the transfer query. |
| `ResultCode` | `int` | Yes |  |
| `ThereAreMoreItems` | `bool` | Yes | Indicates if the caller should execute the query again. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ListTransferRecord().create({
    "ErrorCodes": [],  # list
    "Items": [],  # list
    "ResultCode": 1,  # int
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
| `ErrorCodes` | `list` | Yes |  |
| `Items` | `list` | Yes |  |
| `ResultCode` | `int` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.LookupBill().create({
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Product().list()
for product in results:
    print(product)
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.ProductDescription().list()
for product_description in results:
    print(product_description)
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Promotion().list()
for promotion in results:
    print(promotion)
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PromotionDescription().list()
for promotion_description in results:
    print(promotion_description)
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Provider().list()
for provider in results:
    print(provider)
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.ProviderStatus().list()
for provider_status in results:
    print(provider_status)
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Region().list()
for region in results:
    print(region)
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
| `ErrorCodes` | `list` | Yes |  |
| `ResultCode` | `int` | Yes |  |
| `TransferRecord` | `dict` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.SendTransfer().create({
    "ErrorCodes": [],  # list
    "ResultCode": 1,  # int
    "TransferRecord": {},  # dict
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
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = DingconnectSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

