# Dingconnect SDK configuration


# The sekreto plugin DEFINITIONS the model selected per feature, imported
# above by name from the modules the catalogue's active `plugin.def`
# entries declare. Handed to each feature (secrets builds its Sekreto
# with them): a provider kind not listed here is unknown to that SDK.
FEATURE_PLUGINS = {
}


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
            "debug": {
        "options": {
          "active": False,
          "max": 100,
          "redact": [
            "authorization",
            "cookie",
            "set-cookie",
            "api-key",
            "apikey",
            "x-api-key",
            "idempotency-key",
          ],
        },
        "optspec": {
          "now": "`$FUNCTION`",
          "onEntry": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "none",
      },
            "idempotency": {
        "options": {
          "active": False,
          "header": "Idempotency-Key",
          "methods": [
            "POST",
            "PUT",
            "PATCH",
            "DELETE",
          ],
          "ops": [
            "create",
            "update",
            "remove",
          ],
        },
        "optspec": {
          "keygen": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "none",
      },
            "metrics": {
        "options": {
          "active": False,
        },
        "optspec": {
          "now": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "none",
      },
            "paging": {
        "options": {
          "active": False,
          "afterVar": "after",
          "cursorParam": "cursor",
          "firstVar": "first",
          "limitParam": "limit",
          "pageParam": "page",
          "startPage": 1,
        },
        "optspec": {
          "limit": "`$NUMBER`",
          "ops": "`$LIST`",
        },
        "strict": False,
        "transport": "none",
      },
            "ratelimit": {
        "options": {
          "active": False,
          "burst": 5,
          "rate": 5,
        },
        "optspec": {
          "now": "`$FUNCTION`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "retry": {
        "options": {
          "active": False,
          "factor": 2,
          "maxDelay": 2000,
          "minDelay": 50,
          "retries": 2,
          "statuses": [
            408,
            425,
            429,
            500,
            502,
            503,
            504,
          ],
        },
        "optspec": {
          "jitter": "`$BOOLEAN`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "test": {
        "options": {
          "active": False,
        },
        "optspec": {
          "entity": "`$MAP`",
          "net": "`$MAP`",
        },
        "strict": False,
        "transport": "base",
      },
            "timeout": {
        "options": {
          "active": False,
          "ms": 30000,
        },
        "optspec": {
          "clearTimer": "`$FUNCTION`",
          "setTimer": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
        },
        "options": {
            "base": "https://api.dingconnect.com",
            "auth": {
                "prefix": "",
                "name": "api_key",
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
            "title": "Account Number Normalized",
            "type": "`$STRING`",
            "short": "We attempt to normalize phone numbers following the public telecommunication numbering plan <a href=\"https://en.wikipedia.org/wiki/E.164\" target=\"_blank\">E.164</a>, if we succeed the normalized number will be returned in this field formatt…",
          },
          {
            "name": "CountryIso",
            "title": "Country Iso",
            "type": "`$STRING`",
            "short": "The country of the account number",
          },
          {
            "name": "ErrorCodes",
            "title": "Error Codes",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "Items",
            "title": "Items",
            "type": "`$ARRAY`",
            "req": True,
            "short": "This will contain provider information associated to the account number.",
          },
          {
            "name": "ResultCode",
            "title": "Result Code",
            "type": "`$INTEGER`",
            "req": True,
            "format": "int32",
          },
        ],
        "name": "account_lookup",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetAccountLookup",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "GetAccountLookup",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "GetAccountLookup",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "account_number",
                      "orig": "account_number",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "account_number",
                    "x_correlation_id",
                  ],
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
            "title": "Code",
            "type": "`$STRING`",
            "req": True,
            "short": "The code that can be used to lookup the explanatory message associated with the error",
          },
          {
            "name": "Context",
            "title": "Context",
            "type": "`$STRING`",
            "short": "API specific context as to the reason for the specific code",
          },
        ],
        "name": "balance",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetBalance",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "GetBalance",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "GetBalance",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.ErrorCodes`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_correlation_id",
                  ],
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
            "title": "Error Codes",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "Items",
            "title": "Items",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "ResultCode",
            "title": "Result Code",
            "type": "`$INTEGER`",
            "req": True,
            "format": "int32",
          },
        ],
        "name": "cancel_transfer",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "kind": "http",
                "method": "POST",
                "orig": "/api/V1/CancelTransfers",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "CancelTransfers",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "CancelTransfers",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "cancellation_request",
                      "orig": "cancellation_request",
                      "type": "`$ARRAY`",
                      "kind": "query",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "cancellation_request",
                    "x_correlation_id",
                  ],
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
            "title": "Error Codes",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "Items",
            "title": "Items",
            "type": "`$ARRAY`",
            "req": True,
            "short": "The list of countries that our system is aware of.",
          },
          {
            "name": "ResultCode",
            "title": "Result Code",
            "type": "`$INTEGER`",
            "req": True,
            "format": "int32",
          },
        ],
        "name": "country",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetCountries",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "GetCountries",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "GetCountries",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_correlation_id",
                  ],
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
            "title": "Error Codes",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "Items",
            "title": "Items",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "ResultCode",
            "title": "Result Code",
            "type": "`$INTEGER`",
            "req": True,
            "format": "int32",
          },
        ],
        "name": "currency",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetCurrencies",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "GetCurrencies",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "GetCurrencies",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_correlation_id",
                  ],
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
            "title": "Error Codes",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "Items",
            "title": "Items",
            "type": "`$ARRAY`",
            "req": True,
            "short": "A list of ErrorCodes and their localized descriptions",
          },
          {
            "name": "ResultCode",
            "title": "Result Code",
            "type": "`$INTEGER`",
            "req": True,
            "format": "int32",
          },
        ],
        "name": "error_code_description",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetErrorCodeDescriptions",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "GetErrorCodeDescriptions",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "GetErrorCodeDescriptions",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_correlation_id",
                  ],
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
            "title": "Error Codes",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "Items",
            "title": "Items",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "ResultCode",
            "title": "Result Code",
            "type": "`$INTEGER`",
            "req": True,
            "format": "int32",
          },
        ],
        "name": "estimate_price",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "kind": "http",
                "method": "POST",
                "orig": "/api/V1/EstimatePrices",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "EstimatePrices",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "EstimatePrices",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "requested_estimation",
                      "orig": "requested_estimation",
                      "type": "`$ARRAY`",
                      "kind": "query",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "requested_estimation",
                    "x_correlation_id",
                  ],
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
            "title": "Error Codes",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "Items",
            "title": "Items",
            "type": "`$ARRAY`",
            "req": True,
            "short": "The list of items satisfying the transfer query.",
          },
          {
            "name": "ResultCode",
            "title": "Result Code",
            "type": "`$INTEGER`",
            "req": True,
            "format": "int32",
          },
          {
            "name": "ThereAreMoreItems",
            "title": "There Are More Items",
            "type": "`$BOOLEAN`",
            "req": True,
            "short": "Indicates if the caller should execute the query again.",
          },
        ],
        "name": "list_transfer_record",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "kind": "http",
                "method": "POST",
                "orig": "/api/V1/ListTransferRecords",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "ListTransferRecords",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "ListTransferRecords",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "request",
                      "orig": "request",
                      "type": "`$OBJECT`",
                      "kind": "query",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "request",
                    "x_correlation_id",
                  ],
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
            "title": "Error Codes",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "Items",
            "title": "Items",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "ResultCode",
            "title": "Result Code",
            "type": "`$INTEGER`",
            "req": True,
            "format": "int32",
          },
        ],
        "name": "lookup_bill",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "kind": "http",
                "method": "POST",
                "orig": "/api/V1/LookupBills",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "LookupBills",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "LookupBills",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "request",
                      "orig": "request",
                      "type": "`$OBJECT`",
                      "kind": "query",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "request",
                    "x_correlation_id",
                  ],
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
            "title": "Error Codes",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "Items",
            "title": "Items",
            "type": "`$ARRAY`",
            "req": True,
            "short": "A list of products that fulfil the submitted criteria.",
          },
          {
            "name": "ResultCode",
            "title": "Result Code",
            "type": "`$INTEGER`",
            "req": True,
            "format": "int32",
          },
        ],
        "name": "product",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetProducts",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "GetProducts",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "GetProducts",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "account_number",
                      "orig": "account_number",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                    {
                      "name": "benefit",
                      "orig": "benefit",
                      "type": "`$ANY`",
                      "kind": "query",
                    },
                    {
                      "name": "country_iso",
                      "orig": "country_iso",
                      "type": "`$ANY`",
                      "kind": "query",
                    },
                    {
                      "name": "provider_code",
                      "orig": "provider_code",
                      "type": "`$ANY`",
                      "kind": "query",
                    },
                    {
                      "name": "region_code",
                      "orig": "region_code",
                      "type": "`$ANY`",
                      "kind": "query",
                    },
                    {
                      "name": "sku_code",
                      "orig": "sku_code",
                      "type": "`$ANY`",
                      "kind": "query",
                    },
                  ],
                },
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
            "title": "Error Codes",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "Items",
            "title": "Items",
            "type": "`$ARRAY`",
            "req": True,
            "short": "A localized list of product descriptions.",
          },
          {
            "name": "ResultCode",
            "title": "Result Code",
            "type": "`$INTEGER`",
            "req": True,
            "format": "int32",
          },
        ],
        "name": "product_description",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetProductDescriptions",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "GetProductDescriptions",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "GetProductDescriptions",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "language_code",
                      "orig": "language_code",
                      "type": "`$ANY`",
                      "kind": "query",
                    },
                    {
                      "name": "sku_code",
                      "orig": "sku_code",
                      "type": "`$ANY`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "language_code",
                    "sku_code",
                    "x_correlation_id",
                  ],
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
            "title": "Error Codes",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "Items",
            "title": "Items",
            "type": "`$ARRAY`",
            "req": True,
            "short": "List of available promotions",
          },
          {
            "name": "ResultCode",
            "title": "Result Code",
            "type": "`$INTEGER`",
            "req": True,
            "format": "int32",
          },
        ],
        "name": "promotion",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetPromotions",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "GetPromotions",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "GetPromotions",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "account_number",
                      "orig": "account_number",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                    {
                      "name": "country_iso",
                      "orig": "country_iso",
                      "type": "`$ANY`",
                      "kind": "query",
                    },
                    {
                      "name": "provider_code",
                      "orig": "provider_code",
                      "type": "`$ANY`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "account_number",
                    "country_iso",
                    "provider_code",
                    "x_correlation_id",
                  ],
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
            "title": "Error Codes",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "Items",
            "title": "Items",
            "type": "`$ARRAY`",
            "req": True,
            "short": "A localized list of promotions.",
          },
          {
            "name": "ResultCode",
            "title": "Result Code",
            "type": "`$INTEGER`",
            "req": True,
            "format": "int32",
          },
        ],
        "name": "promotion_description",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetPromotionDescriptions",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "GetPromotionDescriptions",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "GetPromotionDescriptions",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "language_code",
                      "orig": "language_code",
                      "type": "`$ANY`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "language_code",
                    "x_correlation_id",
                  ],
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
            "title": "Error Codes",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "Items",
            "title": "Items",
            "type": "`$ARRAY`",
            "req": True,
            "short": "A list of providers that the distributor has Products for.",
          },
          {
            "name": "ResultCode",
            "title": "Result Code",
            "type": "`$INTEGER`",
            "req": True,
            "format": "int32",
          },
        ],
        "name": "provider",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetProviders",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "GetProviders",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "GetProviders",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "account_number",
                      "orig": "account_number",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                    {
                      "name": "country_iso",
                      "orig": "country_iso",
                      "type": "`$ANY`",
                      "kind": "query",
                    },
                    {
                      "name": "provider_code",
                      "orig": "provider_code",
                      "type": "`$ANY`",
                      "kind": "query",
                    },
                    {
                      "name": "region_code",
                      "orig": "region_code",
                      "type": "`$ANY`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "account_number",
                    "country_iso",
                    "provider_code",
                    "region_code",
                    "x_correlation_id",
                  ],
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
            "title": "Error Codes",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "Items",
            "title": "Items",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "ResultCode",
            "title": "Result Code",
            "type": "`$INTEGER`",
            "req": True,
            "format": "int32",
          },
        ],
        "name": "provider_status",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetProviderStatus",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "GetProviderStatus",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "GetProviderStatus",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "provider_code",
                      "orig": "provider_code",
                      "type": "`$ANY`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "provider_code",
                    "x_correlation_id",
                  ],
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
            "title": "Error Codes",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "Items",
            "title": "Items",
            "type": "`$ARRAY`",
            "req": True,
            "short": "The list of regions that the system uses.",
          },
          {
            "name": "ResultCode",
            "title": "Result Code",
            "type": "`$INTEGER`",
            "req": True,
            "format": "int32",
          },
        ],
        "name": "region",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/V1/GetRegions",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "GetRegions",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "GetRegions",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "country_iso",
                      "orig": "country_iso",
                      "type": "`$ANY`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "country_iso",
                    "x_correlation_id",
                  ],
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
            "title": "Error Codes",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "ResultCode",
            "title": "Result Code",
            "type": "`$INTEGER`",
            "req": True,
            "format": "int32",
          },
          {
            "name": "TransferRecord",
            "title": "Transfer Record",
            "type": "`$OBJECT`",
            "req": True,
          },
        ],
        "name": "send_transfer",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "kind": "http",
                "method": "POST",
                "orig": "/api/V1/SendTransfer",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "V1",
                  },
                  {
                    "lit": "SendTransfer",
                  },
                ],
                "parts": [
                  "api",
                  "V1",
                  "SendTransfer",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "request",
                      "orig": "request",
                      "type": "`$OBJECT`",
                      "kind": "query",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "request",
                    "x_correlation_id",
                  ],
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
