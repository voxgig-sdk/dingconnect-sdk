# Dingconnect Python SDK



The Python SDK for the Dingconnect API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.AccountLookup()` — each
carrying a small, uniform set of operations (`list`, `create`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/dingconnect-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from dingconnect_sdk import DingconnectSDK

client = DingconnectSDK({
    "apikey": os.environ.get("DINGCONNECT_APIKEY"),
})
```

### 2. List accountlookup records

`list()` returns a `list` of records (each a `dict`) and raises on
error — iterate it directly.

```python
try:
    accountlookups = client.AccountLookup().list()
    for accountlookup in accountlookups:
        print(accountlookup)
except Exception as err:
    print(f"list failed: {err}")
```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    currencys = client.Currency().list()
    print(currencys)
except Exception as err:
    print(f"list failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = DingconnectSDK.test()

# Entity ops return the bare record and raise on error.
currency = client.Currency().list()
# currency contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = DingconnectSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### DingconnectSDK

```python
from dingconnect_sdk import DingconnectSDK

client = DingconnectSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = DingconnectSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### DingconnectSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
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
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the bare result data (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

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

Create an instance: `account_lookup = client.AccountLookup()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_number_normalized` | `str` |  |
| `country_iso` | `str` |  |
| `error_code` | `list` |  |
| `item` | `list` |  |
| `result_code` | `int` |  |

#### Example: List

```python
account_lookups = client.AccountLookup().list()
```


### Balance

Create an instance: `balance = client.Balance()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `code` | `str` |  |
| `context` | `str` |  |

#### Example: List

```python
balances = client.Balance().list()
```


### CancelResult

Create an instance: `cancel_result = client.CancelResult()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `list` |  |
| `item` | `list` |  |
| `result_code` | `int` |  |

#### Example: Create

```python
cancel_result = client.CancelResult().create({
})
```


### Country

Create an instance: `country = client.Country()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `list` |  |
| `item` | `list` |  |
| `result_code` | `int` |  |

#### Example: List

```python
countrys = client.Country().list()
```


### Currency

Create an instance: `currency = client.Currency()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `list` |  |
| `item` | `list` |  |
| `result_code` | `int` |  |

#### Example: List

```python
currencys = client.Currency().list()
```


### ErrorCodeDescription

Create an instance: `error_code_description = client.ErrorCodeDescription()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `list` |  |
| `item` | `list` |  |
| `result_code` | `int` |  |

#### Example: List

```python
error_code_descriptions = client.ErrorCodeDescription().list()
```


### Estimate

Create an instance: `estimate = client.Estimate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `list` |  |
| `item` | `list` |  |
| `result_code` | `int` |  |

#### Example: Create

```python
estimate = client.Estimate().create({
})
```


### Product

Create an instance: `product = client.Product()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `list` |  |
| `item` | `list` |  |
| `result_code` | `int` |  |

#### Example: List

```python
products = client.Product().list()
```


### ProductDescription

Create an instance: `product_description = client.ProductDescription()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `list` |  |
| `item` | `list` |  |
| `result_code` | `int` |  |

#### Example: List

```python
product_descriptions = client.ProductDescription().list()
```


### Promotion

Create an instance: `promotion = client.Promotion()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `list` |  |
| `item` | `list` |  |
| `result_code` | `int` |  |

#### Example: List

```python
promotions = client.Promotion().list()
```


### PromotionDescription

Create an instance: `promotion_description = client.PromotionDescription()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `list` |  |
| `item` | `list` |  |
| `result_code` | `int` |  |

#### Example: List

```python
promotion_descriptions = client.PromotionDescription().list()
```


### Provider

Create an instance: `provider = client.Provider()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `list` |  |
| `item` | `list` |  |
| `result_code` | `int` |  |

#### Example: List

```python
providers = client.Provider().list()
```


### ProviderStatus

Create an instance: `provider_status = client.ProviderStatus()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `list` |  |
| `item` | `list` |  |
| `result_code` | `int` |  |

#### Example: List

```python
provider_statuss = client.ProviderStatus().list()
```


### Region

Create an instance: `region = client.Region()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_code` | `list` |  |
| `item` | `list` |  |
| `result_code` | `int` |  |

#### Example: List

```python
regions = client.Region().list()
```


### SendTransfer

Create an instance: `send_transfer = client.SendTransfer()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_number` | `str` |  |
| `distributor_ref` | `str` |  |
| `error_code` | `list` |  |
| `result_code` | `int` |  |
| `send_currency_iso` | `str` |  |
| `send_value` | `float` |  |
| `setting` | `list` |  |
| `sku_code` | `str` |  |
| `transfer_record` | `dict` |  |
| `validate_only` | `bool` |  |

#### Example: Create

```python
send_transfer = client.SendTransfer().create({
    "account_number": "example_account_number",  # str
    "distributor_ref": "example_distributor_ref",  # str
    "send_value": 1,  # float
    "sku_code": "example_sku_code",  # str
})
```


### TransferRecord

Create an instance: `transfer_record = client.TransferRecord()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_number` | `str` |  |
| `distributor_ref` | `list` |  |
| `ended_at_utc` | `str` |  |
| `error_code` | `list` |  |
| `item` | `list` |  |
| `result_code` | `int` |  |
| `skip` | `int` |  |
| `sku_code` | `list` |  |
| `started_at_utc` | `str` |  |
| `take` | `int` |  |
| `there_are_more_item` | `bool` |  |
| `transfer_ref` | `list` |  |

#### Example: Create

```python
transfer_record = client.TransferRecord().create({
    "skip": 1,  # int
    "take": 1,  # int
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

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── dingconnect_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`dingconnect_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```python
currency = client.Currency()
currency.list()

# currency.data_get() now returns the currency data from the last list
# currency.match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
