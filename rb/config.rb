# Dingconnect SDK configuration

module DingconnectConfig
  # Return the process-wide config, built once on first use. The SDK reads
  # the config on every request and never writes to it, so one instance is
  # shared by every client rather than rebuilt per client.
  #
  # The returned hash is shared: treat it as read-only. Callers that need to
  # mutate should use make_config, which always returns a fresh copy.
  def self.shared_config
    @shared_config ||= make_config
  end


  # Build a fresh, fully materialised config hash. Every call rebuilds the
  # whole structure, so prefer shared_config unless you need a private copy
  # you intend to mutate.
  def self.make_config
    {
      "main" => {
        "name" => "Dingconnect",
        "slug" => "dingconnect",
        "version" => "0.1.1",
        "target" => "rb",
      },
      "feature" => {
        "debug" => {
          "options" => {
            "active" => false,
            "max" => 100,
            "redact" => [
              "authorization",
              "cookie",
              "set-cookie",
              "api-key",
              "apikey",
              "x-api-key",
              "idempotency-key",
            ],
          },
          "optspec" => {
            "now" => "`$FUNCTION`",
            "onEntry" => "`$FUNCTION`",
          },
          "strict" => false,
          "transport" => "none",
        },
        "idempotency" => {
          "options" => {
            "active" => false,
            "header" => "Idempotency-Key",
            "methods" => [
              "POST",
              "PUT",
              "PATCH",
              "DELETE",
            ],
            "ops" => [
              "create",
              "update",
              "remove",
            ],
          },
          "optspec" => {
            "keygen" => "`$FUNCTION`",
          },
          "strict" => false,
          "transport" => "none",
        },
        "metrics" => {
          "options" => {
            "active" => false,
          },
          "optspec" => {
            "now" => "`$FUNCTION`",
          },
          "strict" => false,
          "transport" => "none",
        },
        "paging" => {
          "options" => {
            "active" => false,
            "afterVar" => "after",
            "cursorParam" => "cursor",
            "firstVar" => "first",
            "limitParam" => "limit",
            "pageParam" => "page",
            "startPage" => 1,
          },
          "optspec" => {
            "limit" => "`$NUMBER`",
            "ops" => "`$LIST`",
          },
          "strict" => false,
          "transport" => "none",
        },
        "ratelimit" => {
          "options" => {
            "active" => false,
            "burst" => 5,
            "rate" => 5,
          },
          "optspec" => {
            "now" => "`$FUNCTION`",
            "sleep" => "`$FUNCTION`",
          },
          "strict" => false,
          "transport" => "wrap",
        },
        "retry" => {
          "options" => {
            "active" => false,
            "factor" => 2,
            "maxDelay" => 2000,
            "minDelay" => 50,
            "retries" => 2,
            "statuses" => [
              408,
              425,
              429,
              500,
              502,
              503,
              504,
            ],
          },
          "optspec" => {
            "jitter" => "`$BOOLEAN`",
            "sleep" => "`$FUNCTION`",
          },
          "strict" => false,
          "transport" => "wrap",
        },
        "test" => {
          "options" => {
            "active" => false,
          },
          "optspec" => {
            "entity" => "`$MAP`",
            "net" => "`$MAP`",
          },
          "strict" => false,
          "transport" => "base",
        },
        "timeout" => {
          "options" => {
            "active" => false,
            "ms" => 30000,
          },
          "optspec" => {
            "clearTimer" => "`$FUNCTION`",
            "now" => "`$FUNCTION`",
            "setTimer" => "`$FUNCTION`",
          },
          "strict" => false,
          "transport" => "wrap",
        },
      },
      "options" => {
        "base" => "https://api.dingconnect.com",
        "auth" => {
          "prefix" => "",
          "name" => "api_key",
        },
        "headers" => {
          "content-type" => "application/json",
        },
        "entity" => {
          "account_lookup" => {},
          "balance" => {},
          "cancel_transfer" => {},
          "country" => {},
          "currency" => {},
          "error_code_description" => {},
          "estimate_price" => {},
          "list_transfer_record" => {},
          "lookup_bill" => {},
          "product" => {},
          "product_description" => {},
          "promotion" => {},
          "promotion_description" => {},
          "provider" => {},
          "provider_status" => {},
          "region" => {},
          "send_transfer" => {},
        },
      },
      "entity" => {
        "account_lookup" => {
          "fields" => [
            {
              "name" => "AccountNumberNormalized",
              "title" => "Account Number Normalized",
              "type" => "`$STRING`",
              "short" => "We attempt to normalize phone numbers following the public telecommunication numbering plan <a href=\"https://en.wikipedia.org/wiki/E.164\" target=\"_blank\">E.164</a>, if we succeed the normalized number will be returned in this field formatt…",
            },
            {
              "name" => "CountryIso",
              "title" => "Country Iso",
              "type" => "`$STRING`",
              "short" => "The country of the account number",
            },
            {
              "name" => "ErrorCodes",
              "title" => "Error Codes",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "Items",
              "title" => "Items",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "This will contain provider information associated to the account number.",
            },
            {
              "name" => "ResultCode",
              "title" => "Result Code",
              "type" => "`$INTEGER`",
              "req" => true,
              "format" => "int32",
            },
          ],
          "name" => "account_lookup",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/api/V1/GetAccountLookup",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "GetAccountLookup",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "GetAccountLookup",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "account_number",
                        "orig" => "accountNumber",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "balance" => {
          "fields" => [
            {
              "name" => "Code",
              "title" => "Code",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "The code that can be used to lookup the explanatory message associated with the error",
            },
            {
              "name" => "Context",
              "title" => "Context",
              "type" => "`$STRING`",
              "short" => "API specific context as to the reason for the specific code",
            },
          ],
          "name" => "balance",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/api/V1/GetBalance",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "GetBalance",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "GetBalance",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body.ErrorCodes`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "cancel_transfer" => {
          "fields" => [
            {
              "name" => "ErrorCodes",
              "title" => "Error Codes",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "Items",
              "title" => "Items",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "ResultCode",
              "title" => "Result Code",
              "type" => "`$INTEGER`",
              "req" => true,
              "format" => "int32",
            },
            {
              "name" => "cancellations",
              "title" => "Cancellations",
              "type" => "`$ARRAY`",
              "short" => "An explicit list of records to cancel.",
            },
          ],
          "name" => "cancel_transfer",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/api/V1/CancelTransfers",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "CancelTransfers",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "CancelTransfers",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata.cancellations`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "cancellations",
                    ],
                  },
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "country" => {
          "fields" => [
            {
              "name" => "ErrorCodes",
              "title" => "Error Codes",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "Items",
              "title" => "Items",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "The list of countries that our system is aware of.",
            },
            {
              "name" => "ResultCode",
              "title" => "Result Code",
              "type" => "`$INTEGER`",
              "req" => true,
              "format" => "int32",
            },
          ],
          "name" => "country",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/api/V1/GetCountries",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "GetCountries",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "GetCountries",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "currency" => {
          "fields" => [
            {
              "name" => "ErrorCodes",
              "title" => "Error Codes",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "Items",
              "title" => "Items",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "ResultCode",
              "title" => "Result Code",
              "type" => "`$INTEGER`",
              "req" => true,
              "format" => "int32",
            },
          ],
          "name" => "currency",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/api/V1/GetCurrencies",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "GetCurrencies",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "GetCurrencies",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "error_code_description" => {
          "fields" => [
            {
              "name" => "ErrorCodes",
              "title" => "Error Codes",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "Items",
              "title" => "Items",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "A list of ErrorCodes and their localized descriptions",
            },
            {
              "name" => "ResultCode",
              "title" => "Result Code",
              "type" => "`$INTEGER`",
              "req" => true,
              "format" => "int32",
            },
          ],
          "name" => "error_code_description",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/api/V1/GetErrorCodeDescriptions",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "GetErrorCodeDescriptions",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "GetErrorCodeDescriptions",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "estimate_price" => {
          "fields" => [
            {
              "name" => "ErrorCodes",
              "title" => "Error Codes",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "Items",
              "title" => "Items",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "ResultCode",
              "title" => "Result Code",
              "type" => "`$INTEGER`",
              "req" => true,
              "format" => "int32",
            },
            {
              "name" => "estimations",
              "title" => "Estimations",
              "type" => "`$ARRAY`",
            },
          ],
          "name" => "estimate_price",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/api/V1/EstimatePrices",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "EstimatePrices",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "EstimatePrices",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata.estimations`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "estimations",
                    ],
                  },
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "list_transfer_record" => {
          "fields" => [
            {
              "name" => "AccountNumber",
              "title" => "Account Number",
              "type" => "`$STRING`",
              "short" => "Filter transfers by AccountNumber",
            },
            {
              "name" => "DistributorRef",
              "title" => "Distributor Ref",
              "type" => "`$STRING`",
              "short" => "Filter transfers by DistributorRef.",
            },
            {
              "name" => "ErrorCodes",
              "title" => "Error Codes",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "Items",
              "title" => "Items",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "The list of items satisfying the transfer query.",
            },
            {
              "name" => "ResultCode",
              "title" => "Result Code",
              "type" => "`$INTEGER`",
              "req" => true,
              "format" => "int32",
            },
            {
              "name" => "Skip",
              "title" => "Skip",
              "type" => "`$INTEGER`",
              "short" => "The amount of records to by-pass before returning the remaining records",
              "format" => "int32",
            },
            {
              "name" => "Take",
              "title" => "Take",
              "type" => "`$INTEGER`",
              "req" => true,
              "short" => "The amount of records to return",
              "format" => "int32",
            },
            {
              "name" => "ThereAreMoreItems",
              "title" => "There Are More Items",
              "type" => "`$BOOLEAN`",
              "req" => true,
              "short" => "Indicates if the caller should execute the query again.",
            },
            {
              "name" => "TransferRef",
              "title" => "Transfer Ref",
              "type" => "`$STRING`",
              "short" => "Filter by Ding TransferRef",
            },
          ],
          "name" => "list_transfer_record",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/api/V1/ListTransferRecords",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "ListTransferRecords",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "ListTransferRecords",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "lookup_bill" => {
          "fields" => [
            {
              "name" => "AccountNumber",
              "title" => "Account Number",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "The account number to target",
            },
            {
              "name" => "ErrorCodes",
              "title" => "Error Codes",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "Items",
              "title" => "Items",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "ResultCode",
              "title" => "Result Code",
              "type" => "`$INTEGER`",
              "req" => true,
              "format" => "int32",
            },
            {
              "name" => "Settings",
              "title" => "Settings",
              "type" => "`$ARRAY`",
              "short" => "Product specific name/value pairs to be associated with the lookup bills request",
            },
            {
              "name" => "SkuCode",
              "title" => "Sku Code",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "Code provided by GetProducts API",
            },
          ],
          "name" => "lookup_bill",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/api/V1/LookupBills",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "LookupBills",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "LookupBills",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "product" => {
          "fields" => [
            {
              "name" => "ErrorCodes",
              "title" => "Error Codes",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "Items",
              "title" => "Items",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "A list of products that fulfil the submitted criteria.",
            },
            {
              "name" => "ResultCode",
              "title" => "Result Code",
              "type" => "`$INTEGER`",
              "req" => true,
              "format" => "int32",
            },
          ],
          "name" => "product",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/api/V1/GetProducts",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "GetProducts",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "GetProducts",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "account_number",
                        "orig" => "accountNumber",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "",
                      },
                      {
                        "name" => "benefit",
                        "orig" => "benefits",
                        "type" => "`$ARRAY`",
                        "kind" => "query",
                      },
                      {
                        "name" => "country_iso",
                        "orig" => "countryIsos",
                        "type" => "`$ARRAY`",
                        "kind" => "query",
                      },
                      {
                        "name" => "provider_code",
                        "orig" => "providerCodes",
                        "type" => "`$ARRAY`",
                        "kind" => "query",
                      },
                      {
                        "name" => "region_code",
                        "orig" => "regionCodes",
                        "type" => "`$ARRAY`",
                        "kind" => "query",
                      },
                      {
                        "name" => "sku_code",
                        "orig" => "skuCodes",
                        "type" => "`$ARRAY`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "product_description" => {
          "fields" => [
            {
              "name" => "ErrorCodes",
              "title" => "Error Codes",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "Items",
              "title" => "Items",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "A localized list of product descriptions.",
            },
            {
              "name" => "ResultCode",
              "title" => "Result Code",
              "type" => "`$INTEGER`",
              "req" => true,
              "format" => "int32",
            },
          ],
          "name" => "product_description",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/api/V1/GetProductDescriptions",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "GetProductDescriptions",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "GetProductDescriptions",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "language_code",
                        "orig" => "languageCodes",
                        "type" => "`$ARRAY`",
                        "kind" => "query",
                      },
                      {
                        "name" => "sku_code",
                        "orig" => "skuCodes",
                        "type" => "`$ARRAY`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "promotion" => {
          "fields" => [
            {
              "name" => "ErrorCodes",
              "title" => "Error Codes",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "Items",
              "title" => "Items",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "List of available promotions",
            },
            {
              "name" => "ResultCode",
              "title" => "Result Code",
              "type" => "`$INTEGER`",
              "req" => true,
              "format" => "int32",
            },
          ],
          "name" => "promotion",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/api/V1/GetPromotions",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "GetPromotions",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "GetPromotions",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "account_number",
                        "orig" => "accountNumber",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "",
                      },
                      {
                        "name" => "country_iso",
                        "orig" => "countryIsos",
                        "type" => "`$ARRAY`",
                        "kind" => "query",
                      },
                      {
                        "name" => "provider_code",
                        "orig" => "providerCodes",
                        "type" => "`$ARRAY`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "promotion_description" => {
          "fields" => [
            {
              "name" => "ErrorCodes",
              "title" => "Error Codes",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "Items",
              "title" => "Items",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "A localized list of promotions.",
            },
            {
              "name" => "ResultCode",
              "title" => "Result Code",
              "type" => "`$INTEGER`",
              "req" => true,
              "format" => "int32",
            },
          ],
          "name" => "promotion_description",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/api/V1/GetPromotionDescriptions",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "GetPromotionDescriptions",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "GetPromotionDescriptions",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "language_code",
                        "orig" => "languageCodes",
                        "type" => "`$ARRAY`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "provider" => {
          "fields" => [
            {
              "name" => "ErrorCodes",
              "title" => "Error Codes",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "Items",
              "title" => "Items",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "A list of providers that the distributor has Products for.",
            },
            {
              "name" => "ResultCode",
              "title" => "Result Code",
              "type" => "`$INTEGER`",
              "req" => true,
              "format" => "int32",
            },
          ],
          "name" => "provider",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/api/V1/GetProviders",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "GetProviders",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "GetProviders",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "account_number",
                        "orig" => "accountNumber",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "",
                      },
                      {
                        "name" => "country_iso",
                        "orig" => "countryIsos",
                        "type" => "`$ARRAY`",
                        "kind" => "query",
                      },
                      {
                        "name" => "provider_code",
                        "orig" => "providerCodes",
                        "type" => "`$ARRAY`",
                        "kind" => "query",
                      },
                      {
                        "name" => "region_code",
                        "orig" => "regionCodes",
                        "type" => "`$ARRAY`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "provider_status" => {
          "fields" => [
            {
              "name" => "ErrorCodes",
              "title" => "Error Codes",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "Items",
              "title" => "Items",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "ResultCode",
              "title" => "Result Code",
              "type" => "`$INTEGER`",
              "req" => true,
              "format" => "int32",
            },
          ],
          "name" => "provider_status",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/api/V1/GetProviderStatus",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "GetProviderStatus",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "GetProviderStatus",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "provider_code",
                        "orig" => "providerCodes",
                        "type" => "`$ARRAY`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "region" => {
          "fields" => [
            {
              "name" => "ErrorCodes",
              "title" => "Error Codes",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "Items",
              "title" => "Items",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "The list of regions that the system uses.",
            },
            {
              "name" => "ResultCode",
              "title" => "Result Code",
              "type" => "`$INTEGER`",
              "req" => true,
              "format" => "int32",
            },
          ],
          "name" => "region",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/api/V1/GetRegions",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "GetRegions",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "GetRegions",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "country_iso",
                        "orig" => "countryIsos",
                        "type" => "`$ARRAY`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "send_transfer" => {
          "fields" => [
            {
              "name" => "AccountNumber",
              "title" => "Account Number",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "The account number to target",
            },
            {
              "name" => "BillRef",
              "title" => "Bill Ref",
              "type" => "`$STRING`",
              "short" => "Bill reference.",
            },
            {
              "name" => "DistributorRef",
              "title" => "Distributor Ref",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "Unique identifier in the distributor system to be associated with the transfer",
            },
            {
              "name" => "ErrorCodes",
              "title" => "Error Codes",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "ResultCode",
              "title" => "Result Code",
              "type" => "`$INTEGER`",
              "req" => true,
              "format" => "int32",
            },
            {
              "name" => "SendCurrencyIso",
              "title" => "Send Currency Iso",
              "type" => "`$STRING`",
              "short" => "The currency of the `SendValue`.",
            },
            {
              "name" => "SendValue",
              "title" => "Send Value",
              "type" => "`$NUMBER`",
              "req" => true,
              "short" => "The transfer value to be sent.",
              "format" => "decimal",
            },
            {
              "name" => "Settings",
              "title" => "Settings",
              "type" => "`$ARRAY`",
              "short" => "Product specific name/value pairs to be associated with the transfer request",
            },
            {
              "name" => "SkuCode",
              "title" => "Sku Code",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "Code provided by GetProducts API",
            },
            {
              "name" => "TransferRecord",
              "title" => "Transfer Record",
              "type" => "`$OBJECT`",
              "req" => true,
            },
            {
              "name" => "ValidateOnly",
              "title" => "Validate Only",
              "type" => "`$BOOLEAN`",
              "req" => true,
              "short" => "Validate the request with the provider without doing a transfer",
            },
          ],
          "name" => "send_transfer",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/api/V1/SendTransfer",
                  "segments" => [
                    {
                      "lit" => "api",
                    },
                    {
                      "lit" => "V1",
                    },
                    {
                      "lit" => "SendTransfer",
                    },
                  ],
                  "parts" => [
                    "api",
                    "V1",
                    "SendTransfer",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_correlation_id",
                        "orig" => "X-Correlation-Id",
                        "type" => "`$STRING`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "alternatives" => [
                      {
                        "kind" => "json",
                        "media" => "text/json",
                      },
                    ],
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
      },
    }
  end


  def self.make_feature(name)
    require_relative 'features'
    DingconnectFeatures.make_feature(name)
  end
end
