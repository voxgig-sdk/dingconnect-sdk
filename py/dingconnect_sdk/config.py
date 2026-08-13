# Dingconnect SDK configuration


def make_config():
    return {
        "main": {
            "name": "Dingconnect",
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
            "active": True,
            "name": "AccountNumberNormalized",
            "req": False,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "CountryIso",
            "req": False,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 4,
          },
        ],
        "name": "account_lookup",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "account_number",
                      "orig": "account_number",
                      "reqd": False,
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
                "index$": 0,
              },
            ],
            "key$": "list",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "balance": {
        "fields": [
          {
            "active": True,
            "name": "Code",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "Context",
            "req": False,
            "type": "`$STRING`",
            "index$": 1,
          },
        ],
        "name": "balance",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
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
                "index$": 0,
              },
            ],
            "key$": "list",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "cancel_transfer": {
        "fields": [
          {
            "active": True,
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 2,
          },
        ],
        "name": "cancel_transfer",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
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
                "index$": 0,
              },
            ],
            "key$": "create",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "country": {
        "fields": [
          {
            "active": True,
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 2,
          },
        ],
        "name": "country",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
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
                "index$": 0,
              },
            ],
            "key$": "list",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "currency": {
        "fields": [
          {
            "active": True,
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 2,
          },
        ],
        "name": "currency",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
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
                "index$": 0,
              },
            ],
            "key$": "list",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "error_code_description": {
        "fields": [
          {
            "active": True,
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 2,
          },
        ],
        "name": "error_code_description",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
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
                "index$": 0,
              },
            ],
            "key$": "list",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "estimate_price": {
        "fields": [
          {
            "active": True,
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 2,
          },
        ],
        "name": "estimate_price",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
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
                "index$": 0,
              },
            ],
            "key$": "create",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "list_transfer_record": {
        "fields": [
          {
            "active": True,
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "ThereAreMoreItems",
            "req": True,
            "type": "`$BOOLEAN`",
            "index$": 3,
          },
        ],
        "name": "list_transfer_record",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
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
                "index$": 0,
              },
            ],
            "key$": "create",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "lookup_bill": {
        "fields": [
          {
            "active": True,
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 2,
          },
        ],
        "name": "lookup_bill",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
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
                "index$": 0,
              },
            ],
            "key$": "create",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "product": {
        "fields": [
          {
            "active": True,
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 2,
          },
        ],
        "name": "product",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "account_number",
                      "orig": "account_number",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "benefit",
                      "orig": "benefit",
                      "reqd": False,
                      "type": "`$ANY`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "country_iso",
                      "orig": "country_iso",
                      "reqd": False,
                      "type": "`$ANY`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "provider_code",
                      "orig": "provider_code",
                      "reqd": False,
                      "type": "`$ANY`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "region_code",
                      "orig": "region_code",
                      "reqd": False,
                      "type": "`$ANY`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "sku_code",
                      "orig": "sku_code",
                      "reqd": False,
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
                "index$": 0,
              },
            ],
            "key$": "list",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "product_description": {
        "fields": [
          {
            "active": True,
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 2,
          },
        ],
        "name": "product_description",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "language_code",
                      "orig": "language_code",
                      "reqd": False,
                      "type": "`$ANY`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "sku_code",
                      "orig": "sku_code",
                      "reqd": False,
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
                "index$": 0,
              },
            ],
            "key$": "list",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "promotion": {
        "fields": [
          {
            "active": True,
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 2,
          },
        ],
        "name": "promotion",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "account_number",
                      "orig": "account_number",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "country_iso",
                      "orig": "country_iso",
                      "reqd": False,
                      "type": "`$ANY`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "provider_code",
                      "orig": "provider_code",
                      "reqd": False,
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
                "index$": 0,
              },
            ],
            "key$": "list",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "promotion_description": {
        "fields": [
          {
            "active": True,
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 2,
          },
        ],
        "name": "promotion_description",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "language_code",
                      "orig": "language_code",
                      "reqd": False,
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
                "index$": 0,
              },
            ],
            "key$": "list",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "provider": {
        "fields": [
          {
            "active": True,
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 2,
          },
        ],
        "name": "provider",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "account_number",
                      "orig": "account_number",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "country_iso",
                      "orig": "country_iso",
                      "reqd": False,
                      "type": "`$ANY`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "provider_code",
                      "orig": "provider_code",
                      "reqd": False,
                      "type": "`$ANY`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "region_code",
                      "orig": "region_code",
                      "reqd": False,
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
                "index$": 0,
              },
            ],
            "key$": "list",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "provider_status": {
        "fields": [
          {
            "active": True,
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 2,
          },
        ],
        "name": "provider_status",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "provider_code",
                      "orig": "provider_code",
                      "reqd": False,
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
                "index$": 0,
              },
            ],
            "key$": "list",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "region": {
        "fields": [
          {
            "active": True,
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "Items",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 2,
          },
        ],
        "name": "region",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "country_iso",
                      "orig": "country_iso",
                      "reqd": False,
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
                "index$": 0,
              },
            ],
            "key$": "list",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "send_transfer": {
        "fields": [
          {
            "active": True,
            "name": "ErrorCodes",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "ResultCode",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "TransferRecord",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 2,
          },
        ],
        "name": "send_transfer",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {
                  "header": [
                    {
                      "active": True,
                      "kind": "header",
                      "name": "x_correlation_id",
                      "orig": "x_correlation_id",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
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
                "index$": 0,
              },
            ],
            "key$": "create",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
