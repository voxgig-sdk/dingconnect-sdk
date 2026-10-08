# Ding API

The Ding API is a Level 0 REST web service. We have used the &lt;a href=&quot;http://swagger.io&quot;&gt;swagger&lt;/a&gt; standard to describe this service. As a result, we are able to provide this interactive documentation page. For further information, you may view the &lt;a href=&quot;/api/description&quot;&gt;additional documentation&lt;/a&gt;; read our &lt;a href=&quot;/api/faq&quot;&gt;FAQ&lt;/a&gt; or contact partnersupport@ding.com

## Start here

This guide introduces the API, the client libraries, and the companion tools in this repository. Start with the API capabilities, choose a client for your application, and use the linked reference when you need exact request and response details.

The selected API surface contains 17 entities and 17 HTTP routes. There are 6 SDK targets and 2 companion tools.

An entity groups related API operations. An operation can have several routes with different inputs or authentication requirements. The SDK exposes the entity and its operations using the conventions of the selected language.

## What the API provides

### AccountLookup

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `AccountNumberNormalized`: We attempt to normalize phone numbers following the public telecommunication numbering plan &lt;a href=&quot;https://en.wikipedia.org/wiki/E.164&quot; target=&quot;`_blank`&quot;&gt;E.164&lt;/a&gt;, if we succeed the normalized number will be returned in this field formatted as E164 without leading &#39;+&#39;
- `CountryIso`: The country of the account number
- `Items`: This will contain provider information associated to the account number. If we can succesfully lookup the account number the list will contain the info for products associated to it.

### Balance

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `Code`: The code that can be used to lookup the explanatory message associated with the error
- `Context`: API specific context as to the reason for the specific code

### CancelTransfer

Results: OK.

SDK operations: `create`.

Key fields to recognise:

- `ErrorCodes`: ErrorCodes (if any) for processing the batch item with the given BatchItemRef
- `ResultCode`: The individual result code for processing the batch item with the given BatchItemRef.
- `cancellations`: An explicit list of records to cancel.

### Country

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `Items`: The list of countries that our system is aware of.

### Currency

Results: OK.

SDK operations: `list`.

### ErrorCodeDescription

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `Items`: A list of ErrorCodes and their localized descriptions

### EstimatePrice

Results: OK.

SDK operations: `create`.

Key fields to recognise:

- `ErrorCodes`: ErrorCodes (if any) for processing the batch item with the given BatchItemRef
- `ResultCode`: The individual result code for processing the batch item with the given BatchItemRef.

### ListTransferRecord

Results: OK.

SDK operations: `create`.

Key fields to recognise:

- `AccountNumber`: The account number targeted in the transfer
- `DistributorRef`: The distributor&#39;s identifier for the transfer.
- `ErrorCodes`: Any error codes that were returned as part of the SendTransfer or in the case of a Batch `ProcessingMode` that may have occurred later after the batch was submitted to the Provider.
- `Items`: The list of items satisfying the transfer query.
- `Skip`: The amount of records to by-pass before returning the remaining records

### LookupBill

Results: OK.

SDK operations: `create`.

Key fields to recognise:

- `AccountNumber`: The account number to target
- `Settings`: Product specific name/value pairs to be associated with the lookup bills request
- `SkuCode`: Code provided by GetProducts API

### Product

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `Items`: A list of products that fulfil the submitted criteria.

### ProductDescription

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `Items`: A localized list of product descriptions.

### Promotion

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `Items`: List of available promotions

### PromotionDescription

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `Items`: A localized list of promotions.

### Provider

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `Items`: A list of providers that the distributor has Products for.

### ProviderStatus

Results: OK.

SDK operations: `list`.

### Region

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `Items`: The list of regions that the system uses.

### SendTransfer

Results: OK.

SDK operations: `create`.

Key fields to recognise:

- `AccountNumber`: The account number targeted in the transfer
- `BillRef`: Bill reference.
- `DistributorRef`: The distributor&#39;s identifier for the transfer.
- `SendCurrencyIso`: The currency of the SendValue field
- `SendValue`: The value that is submitted to SendTransfer

### Route map

Use this map to locate a capability. Consult the entity reference before supplying request data; routes for the same operation can require different fields.

| Entity | SDK operation | HTTP route | Authentication |
| --- | --- | --- | --- |
| AccountLookup | `list` | `GET /api/V1/GetAccountLookup` | See reference |
| Balance | `list` | `GET /api/V1/GetBalance` | See reference |
| CancelTransfer | `create` | `POST /api/V1/CancelTransfers` | See reference |
| Country | `list` | `GET /api/V1/GetCountries` | See reference |
| Currency | `list` | `GET /api/V1/GetCurrencies` | See reference |
| ErrorCodeDescription | `list` | `GET /api/V1/GetErrorCodeDescriptions` | See reference |
| EstimatePrice | `create` | `POST /api/V1/EstimatePrices` | See reference |
| ListTransferRecord | `create` | `POST /api/V1/ListTransferRecords` | See reference |
| LookupBill | `create` | `POST /api/V1/LookupBills` | See reference |
| Product | `list` | `GET /api/V1/GetProducts` | See reference |
| ProductDescription | `list` | `GET /api/V1/GetProductDescriptions` | See reference |
| Promotion | `list` | `GET /api/V1/GetPromotions` | See reference |
| PromotionDescription | `list` | `GET /api/V1/GetPromotionDescriptions` | See reference |
| Provider | `list` | `GET /api/V1/GetProviders` | See reference |
| ProviderStatus | `list` | `GET /api/V1/GetProviderStatus` | See reference |
| Region | `list` | `GET /api/V1/GetRegions` | See reference |
| SendTransfer | `create` | `POST /api/V1/SendTransfer` | See reference |

## Connect to the API

- API server: `https://api.dingconnect.com`

The default credential is sent in the `api_key` header.

Use the API key that was issued to you.

Check authentication for the route you plan to call. A route that declares no authentication can be used without credentials; this does not change the requirements of other routes. Keep credentials in environment variables or a configured secret provider, and keep them out of source control and logs.

## Make a first request

1. Choose the API server and an operation that matches your task.
2. Check the operation’s required input and authentication. Use values valid for your account and environment.
3. Send one request and inspect the returned data before adding retries, concurrency, or a larger batch.

For an SDK call, install or build the chosen client, create a client instance with its documented configuration, and call the required entity operation. Language references describe the argument shape, asynchronous behaviour, and returned values.

## Choose an SDK

Choose the language already used by your application or service. The clients represent the same API model, while package setup, naming, and return types follow each language. Check the selected client’s reference and tests before integrating it into an existing application.

| Client | Repository directory | Distribution |
| --- | --- | --- |
| Golang | `go/` | Build from source |
| Lua | `lua/` | Build from source |
| PHP | `php/` | Build from source |
| Python | `py/` | Build from source |
| Ruby | `rb/` | Build from source |
| TypeScript | `ts/` | Build from source |

Build-from-source entries are not marked as published in the project model. Follow the build instructions in that target’s README, then consume the resulting package using your language’s local dependency mechanism. Published entries give the installation command recorded for that client.

## Companion tools

These targets provide another way to use the API. Their available commands or tools can cover a smaller set of operations than the client libraries.

### Go CLI

Use the command-line interface for shell-based tasks and scripts.

Repository directory: `go-cli/`. Not published. Build from the go-cli directory.


### Go MCP server

Use the MCP server to expose supported API operations to an MCP client.

Repository directory: `go-mcp/`. Not published. Build from the go-mcp directory.

- `dingconnect_list`: List records for an entity. Supported entities: `account_lookup`, `balance`, `country`, `currency`, `error_code_description`, `product`, `product_description`, `promotion`, `promotion_description`, `provider`, `provider_status`, `region`.
- `dingconnect_load`: Load one record for an entity. No active entity supports this operation.

## Operational features

Features supply behaviour around API calls, such as request handling, diagnostics, or local testing. Inclusion in this project does not mean a feature is enabled at runtime. Check the selected SDK’s supported features and configuration defaults, then enable the behaviour your application needs.

- `debug`: Request/response capture ring buffer for debugging
- `idempotency`: Idempotency keys for safe retries of mutating operations
- `metrics`: Statistics capture: per-operation counters and latency
- `paging`: Pagination signals for list operations
- `ratelimit`: Client-side rate limiting via a token bucket
- `retry`: Automatic retry of transient failures with exponential backoff
- `test`: In-memory mock transport for testing without a live server
- `timeout`: Per-request timeout with transport abort

Start with the default client configuration. Add request limits and diagnostics as needed, test error paths, and review retry behaviour before using operations that change data. A retry can repeat an operation unless the API provides a suitable guarantee.

## Continue with the documentation

- Follow the first-call guide for the setup sequence.
- Read the authentication guide before using protected routes.
- Use the API reference for request schemas, response formats, and status codes.
- Check the chosen SDK or companion tool reference for its configuration and supported operations.

