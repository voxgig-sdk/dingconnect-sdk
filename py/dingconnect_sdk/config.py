# Dingconnect SDK configuration


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "Dingconnect",
            "slug": "dingconnect",
            "version": "0.1.1",
            "target": "py",
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
      },
        },
        "options": {
            "base": "https://api.dingconnect.com",
            "auth": {
                "prefix": "",
            },
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "account_lookup": {},
                "balance": {},
                "cancel_transfer": {},
                "country": {},
                "currency": {},
                "error_code_description": {},
                "estimate_price": {},
                "list_transfer_record": {},
                "lookup_bill": {},
                "product": {},
                "product_description": {},
                "promotion": {},
                "promotion_description": {},
                "provider": {},
                "provider_status": {},
                "region": {},
                "send_transfer": {},
            },
        },
        "entity": {
      "account_lookup": {
        "fields": [
          {
            "name": "AccountNumberNormalized",
            "short": "We attempt to normalize phone numbers following the public telecommunication numbering plan <a href=\"https://en.wikipedia.org/wiki/E.164\" target=\"_blank\">E.164</a>, if we succeed the normalized number will be returned in this field formatt…",
            "type": "`$STRING`",
          },
          {
            "name": "CountryIso",
            "short": "The country of the account number",
            "type": "`$STRING`",
          },
          {
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "Items",
            "req": True,
            "short": "This will contain provider information associated to the account number.",
            "type": "`$ARRAY`",
          },
          {
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
          },
        ],
        "name": "account_lookup",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "kind": "query",
                      "name": "account_number",
                      "orig": "account_number",
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetAccountLookup",
                "parts": [
                  "api",
                  "V1",
                  "GetAccountLookup",
                ],
                "select": {
                  "exist": [
                    "account_number",
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "balance": {
        "fields": [
          {
            "name": "Code",
            "req": True,
            "short": "The code that can be used to lookup the explanatory message associated with the error",
            "type": "`$STRING`",
          },
          {
            "name": "Context",
            "short": "API specific context as to the reason for the specific code",
            "type": "`$STRING`",
          },
        ],
        "name": "balance",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetBalance",
                "parts": [
                  "api",
                  "V1",
                  "GetBalance",
                ],
                "select": {
                  "exist": [
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.ErrorCodes`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "cancel_transfer": {
        "fields": [
          {
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
          },
        ],
        "name": "cancel_transfer",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "kind": "query",
                      "name": "cancellation_request",
                      "orig": "cancellation_request",
                      "reqd": True,
                      "type": "`$ARRAY`",
                    },
                  ],
                },
                "kind": "http",
                "method": "POST",
                "orig": "/api/V1/CancelTransfers",
                "parts": [
                  "api",
                  "V1",
                  "CancelTransfers",
                ],
                "select": {
                  "exist": [
                    "cancellation_request",
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "country": {
        "fields": [
          {
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "Items",
            "req": True,
            "short": "The list of countries that our system is aware of.",
            "type": "`$ARRAY`",
          },
          {
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
          },
        ],
        "name": "country",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetCountries",
                "parts": [
                  "api",
                  "V1",
                  "GetCountries",
                ],
                "select": {
                  "exist": [
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "currency": {
        "fields": [
          {
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
          },
        ],
        "name": "currency",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetCurrencies",
                "parts": [
                  "api",
                  "V1",
                  "GetCurrencies",
                ],
                "select": {
                  "exist": [
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "error_code_description": {
        "fields": [
          {
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "Items",
            "req": True,
            "short": "A list of ErrorCodes and their localized descriptions",
            "type": "`$ARRAY`",
          },
          {
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
          },
        ],
        "name": "error_code_description",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetErrorCodeDescriptions",
                "parts": [
                  "api",
                  "V1",
                  "GetErrorCodeDescriptions",
                ],
                "select": {
                  "exist": [
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "estimate_price": {
        "fields": [
          {
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
          },
        ],
        "name": "estimate_price",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "kind": "query",
                      "name": "requested_estimation",
                      "orig": "requested_estimation",
                      "reqd": True,
                      "type": "`$ARRAY`",
                    },
                  ],
                },
                "kind": "http",
                "method": "POST",
                "orig": "/api/V1/EstimatePrices",
                "parts": [
                  "api",
                  "V1",
                  "EstimatePrices",
                ],
                "select": {
                  "exist": [
                    "requested_estimation",
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "list_transfer_record": {
        "fields": [
          {
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "Items",
            "req": True,
            "short": "The list of items satisfying the transfer query.",
            "type": "`$ARRAY`",
          },
          {
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
          },
          {
            "name": "ThereAreMoreItems",
            "req": True,
            "short": "Indicates if the caller should execute the query again.",
            "type": "`$BOOLEAN`",
          },
        ],
        "name": "list_transfer_record",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "kind": "query",
                      "name": "request",
                      "orig": "request",
                      "reqd": True,
                      "type": "`$OBJECT`",
                    },
                  ],
                },
                "kind": "http",
                "method": "POST",
                "orig": "/api/V1/ListTransferRecords",
                "parts": [
                  "api",
                  "V1",
                  "ListTransferRecords",
                ],
                "select": {
                  "exist": [
                    "request",
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "lookup_bill": {
        "fields": [
          {
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
          },
        ],
        "name": "lookup_bill",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "kind": "query",
                      "name": "request",
                      "orig": "request",
                      "reqd": True,
                      "type": "`$OBJECT`",
                    },
                  ],
                },
                "kind": "http",
                "method": "POST",
                "orig": "/api/V1/LookupBills",
                "parts": [
                  "api",
                  "V1",
                  "LookupBills",
                ],
                "select": {
                  "exist": [
                    "request",
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "product": {
        "fields": [
          {
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "Items",
            "req": True,
            "short": "A list of products that fulfil the submitted criteria.",
            "type": "`$ARRAY`",
          },
          {
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
          },
        ],
        "name": "product",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "kind": "query",
                      "name": "account_number",
                      "orig": "account_number",
                      "type": "`$INTEGER`",
                    },
                    {
                      "kind": "query",
                      "name": "benefit",
                      "orig": "benefit",
                      "type": "`$ANY`",
                    },
                    {
                      "kind": "query",
                      "name": "country_iso",
                      "orig": "country_iso",
                      "type": "`$ANY`",
                    },
                    {
                      "kind": "query",
                      "name": "provider_code",
                      "orig": "provider_code",
                      "type": "`$ANY`",
                    },
                    {
                      "kind": "query",
                      "name": "region_code",
                      "orig": "region_code",
                      "type": "`$ANY`",
                    },
                    {
                      "kind": "query",
                      "name": "sku_code",
                      "orig": "sku_code",
                      "type": "`$ANY`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetProducts",
                "parts": [
                  "api",
                  "V1",
                  "GetProducts",
                ],
                "select": {
                  "exist": [
                    "account_number",
                    "benefit",
                    "country_iso",
                    "provider_code",
                    "region_code",
                    "sku_code",
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "product_description": {
        "fields": [
          {
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "Items",
            "req": True,
            "short": "A localized list of product descriptions.",
            "type": "`$ARRAY`",
          },
          {
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
          },
        ],
        "name": "product_description",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "kind": "query",
                      "name": "language_code",
                      "orig": "language_code",
                      "type": "`$ANY`",
                    },
                    {
                      "kind": "query",
                      "name": "sku_code",
                      "orig": "sku_code",
                      "type": "`$ANY`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetProductDescriptions",
                "parts": [
                  "api",
                  "V1",
                  "GetProductDescriptions",
                ],
                "select": {
                  "exist": [
                    "language_code",
                    "sku_code",
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "promotion": {
        "fields": [
          {
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "Items",
            "req": True,
            "short": "List of available promotions",
            "type": "`$ARRAY`",
          },
          {
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
          },
        ],
        "name": "promotion",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "kind": "query",
                      "name": "account_number",
                      "orig": "account_number",
                      "type": "`$INTEGER`",
                    },
                    {
                      "kind": "query",
                      "name": "country_iso",
                      "orig": "country_iso",
                      "type": "`$ANY`",
                    },
                    {
                      "kind": "query",
                      "name": "provider_code",
                      "orig": "provider_code",
                      "type": "`$ANY`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetPromotions",
                "parts": [
                  "api",
                  "V1",
                  "GetPromotions",
                ],
                "select": {
                  "exist": [
                    "account_number",
                    "country_iso",
                    "provider_code",
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "promotion_description": {
        "fields": [
          {
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "Items",
            "req": True,
            "short": "A localized list of promotions.",
            "type": "`$ARRAY`",
          },
          {
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
          },
        ],
        "name": "promotion_description",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "kind": "query",
                      "name": "language_code",
                      "orig": "language_code",
                      "type": "`$ANY`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetPromotionDescriptions",
                "parts": [
                  "api",
                  "V1",
                  "GetPromotionDescriptions",
                ],
                "select": {
                  "exist": [
                    "language_code",
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "provider": {
        "fields": [
          {
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "Items",
            "req": True,
            "short": "A list of providers that the distributor has Products for.",
            "type": "`$ARRAY`",
          },
          {
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
          },
        ],
        "name": "provider",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "kind": "query",
                      "name": "account_number",
                      "orig": "account_number",
                      "type": "`$INTEGER`",
                    },
                    {
                      "kind": "query",
                      "name": "country_iso",
                      "orig": "country_iso",
                      "type": "`$ANY`",
                    },
                    {
                      "kind": "query",
                      "name": "provider_code",
                      "orig": "provider_code",
                      "type": "`$ANY`",
                    },
                    {
                      "kind": "query",
                      "name": "region_code",
                      "orig": "region_code",
                      "type": "`$ANY`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetProviders",
                "parts": [
                  "api",
                  "V1",
                  "GetProviders",
                ],
                "select": {
                  "exist": [
                    "account_number",
                    "country_iso",
                    "provider_code",
                    "region_code",
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "provider_status": {
        "fields": [
          {
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
          },
        ],
        "name": "provider_status",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "kind": "query",
                      "name": "provider_code",
                      "orig": "provider_code",
                      "type": "`$ANY`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetProviderStatus",
                "parts": [
                  "api",
                  "V1",
                  "GetProviderStatus",
                ],
                "select": {
                  "exist": [
                    "provider_code",
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "region": {
        "fields": [
          {
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "Items",
            "req": True,
            "short": "The list of regions that the system uses.",
            "type": "`$ARRAY`",
          },
          {
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
          },
        ],
        "name": "region",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "kind": "query",
                      "name": "country_iso",
                      "orig": "country_iso",
                      "type": "`$ANY`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetRegions",
                "parts": [
                  "api",
                  "V1",
                  "GetRegions",
                ],
                "select": {
                  "exist": [
                    "country_iso",
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "send_transfer": {
        "fields": [
          {
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
          },
          {
            "name": "TransferRecord",
            "req": True,
            "type": "`$OBJECT`",
          },
        ],
        "name": "send_transfer",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "args": {
                  "header": [
                    {
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "kind": "query",
                      "name": "request",
                      "orig": "request",
                      "reqd": True,
                      "type": "`$OBJECT`",
                    },
                  ],
                },
                "kind": "http",
                "method": "POST",
                "orig": "/api/V1/SendTransfer",
                "parts": [
                  "api",
                  "V1",
                  "SendTransfer",
                ],
                "select": {
                  "exist": [
                    "request",
                    "x_correlation_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
