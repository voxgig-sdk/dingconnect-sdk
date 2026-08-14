-- Dingconnect SDK configuration

-- Build a fresh, fully materialised config table. Every call rebuilds the
-- whole structure, so prefer require("config_shared") unless you need a
-- private copy you intend to mutate.
local function make_config()
  return {
    main = {
      name = "Dingconnect",
    },
    feature = {
      ["test"] = {
        ["options"] = {
          ["active"] = false,
        },
      },
    },
    options = {
      base = "https://api.dingconnect.com",
      auth = {
        prefix = "",
      },
      headers = {
        ["content-type"] = "application/json",
      },
      entity = {
        ["account_lookup"] = {},
        ["balance"] = {},
        ["cancel_transfer"] = {},
        ["country"] = {},
        ["currency"] = {},
        ["error_code_description"] = {},
        ["estimate_price"] = {},
        ["list_transfer_record"] = {},
        ["lookup_bill"] = {},
        ["product"] = {},
        ["product_description"] = {},
        ["promotion"] = {},
        ["promotion_description"] = {},
        ["provider"] = {},
        ["provider_status"] = {},
        ["region"] = {},
        ["send_transfer"] = {},
      },
    },
    entity = {
      ["account_lookup"] = {
        ["fields"] = {
          {
            ["name"] = "AccountNumberNormalized",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "CountryIso",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ErrorCodes",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "Items",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ResultCode",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "account_lookup",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "account_number",
                      ["orig"] = "account_number",
                      ["type"] = "`$INTEGER`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/V1/GetAccountLookup",
                ["parts"] = {
                  "api",
                  "V1",
                  "GetAccountLookup",
                },
                ["select"] = {
                  ["exist"] = {
                    "account_number",
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["balance"] = {
        ["fields"] = {
          {
            ["name"] = "Code",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "Context",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "balance",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/V1/GetBalance",
                ["parts"] = {
                  "api",
                  "V1",
                  "GetBalance",
                },
                ["select"] = {
                  ["exist"] = {
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.ErrorCodes`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["cancel_transfer"] = {
        ["fields"] = {
          {
            ["name"] = "ErrorCodes",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "Items",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ResultCode",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "cancel_transfer",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "cancellation_request",
                      ["orig"] = "cancellation_request",
                      ["reqd"] = true,
                      ["type"] = "`$ARRAY`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/V1/CancelTransfers",
                ["parts"] = {
                  "api",
                  "V1",
                  "CancelTransfers",
                },
                ["select"] = {
                  ["exist"] = {
                    "cancellation_request",
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["country"] = {
        ["fields"] = {
          {
            ["name"] = "ErrorCodes",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "Items",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ResultCode",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "country",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/V1/GetCountries",
                ["parts"] = {
                  "api",
                  "V1",
                  "GetCountries",
                },
                ["select"] = {
                  ["exist"] = {
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["currency"] = {
        ["fields"] = {
          {
            ["name"] = "ErrorCodes",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "Items",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ResultCode",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "currency",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/V1/GetCurrencies",
                ["parts"] = {
                  "api",
                  "V1",
                  "GetCurrencies",
                },
                ["select"] = {
                  ["exist"] = {
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["error_code_description"] = {
        ["fields"] = {
          {
            ["name"] = "ErrorCodes",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "Items",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ResultCode",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "error_code_description",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/V1/GetErrorCodeDescriptions",
                ["parts"] = {
                  "api",
                  "V1",
                  "GetErrorCodeDescriptions",
                },
                ["select"] = {
                  ["exist"] = {
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["estimate_price"] = {
        ["fields"] = {
          {
            ["name"] = "ErrorCodes",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "Items",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ResultCode",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "estimate_price",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "requested_estimation",
                      ["orig"] = "requested_estimation",
                      ["reqd"] = true,
                      ["type"] = "`$ARRAY`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/V1/EstimatePrices",
                ["parts"] = {
                  "api",
                  "V1",
                  "EstimatePrices",
                },
                ["select"] = {
                  ["exist"] = {
                    "requested_estimation",
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["list_transfer_record"] = {
        ["fields"] = {
          {
            ["name"] = "ErrorCodes",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "Items",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ResultCode",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ThereAreMoreItems",
            ["req"] = true,
            ["type"] = "`$BOOLEAN`",
          },
        },
        ["name"] = "list_transfer_record",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "request",
                      ["orig"] = "request",
                      ["reqd"] = true,
                      ["type"] = "`$OBJECT`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/V1/ListTransferRecords",
                ["parts"] = {
                  "api",
                  "V1",
                  "ListTransferRecords",
                },
                ["select"] = {
                  ["exist"] = {
                    "request",
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["lookup_bill"] = {
        ["fields"] = {
          {
            ["name"] = "ErrorCodes",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "Items",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ResultCode",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "lookup_bill",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "request",
                      ["orig"] = "request",
                      ["reqd"] = true,
                      ["type"] = "`$OBJECT`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/V1/LookupBills",
                ["parts"] = {
                  "api",
                  "V1",
                  "LookupBills",
                },
                ["select"] = {
                  ["exist"] = {
                    "request",
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["product"] = {
        ["fields"] = {
          {
            ["name"] = "ErrorCodes",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "Items",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ResultCode",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "product",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "account_number",
                      ["orig"] = "account_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "benefit",
                      ["orig"] = "benefit",
                      ["type"] = "`$ANY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "country_iso",
                      ["orig"] = "country_iso",
                      ["type"] = "`$ANY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "provider_code",
                      ["orig"] = "provider_code",
                      ["type"] = "`$ANY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "region_code",
                      ["orig"] = "region_code",
                      ["type"] = "`$ANY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sku_code",
                      ["orig"] = "sku_code",
                      ["type"] = "`$ANY`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/V1/GetProducts",
                ["parts"] = {
                  "api",
                  "V1",
                  "GetProducts",
                },
                ["select"] = {
                  ["exist"] = {
                    "account_number",
                    "benefit",
                    "country_iso",
                    "provider_code",
                    "region_code",
                    "sku_code",
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["product_description"] = {
        ["fields"] = {
          {
            ["name"] = "ErrorCodes",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "Items",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ResultCode",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "product_description",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "language_code",
                      ["orig"] = "language_code",
                      ["type"] = "`$ANY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sku_code",
                      ["orig"] = "sku_code",
                      ["type"] = "`$ANY`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/V1/GetProductDescriptions",
                ["parts"] = {
                  "api",
                  "V1",
                  "GetProductDescriptions",
                },
                ["select"] = {
                  ["exist"] = {
                    "language_code",
                    "sku_code",
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["promotion"] = {
        ["fields"] = {
          {
            ["name"] = "ErrorCodes",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "Items",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ResultCode",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "promotion",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "account_number",
                      ["orig"] = "account_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "country_iso",
                      ["orig"] = "country_iso",
                      ["type"] = "`$ANY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "provider_code",
                      ["orig"] = "provider_code",
                      ["type"] = "`$ANY`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/V1/GetPromotions",
                ["parts"] = {
                  "api",
                  "V1",
                  "GetPromotions",
                },
                ["select"] = {
                  ["exist"] = {
                    "account_number",
                    "country_iso",
                    "provider_code",
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["promotion_description"] = {
        ["fields"] = {
          {
            ["name"] = "ErrorCodes",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "Items",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ResultCode",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "promotion_description",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "language_code",
                      ["orig"] = "language_code",
                      ["type"] = "`$ANY`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/V1/GetPromotionDescriptions",
                ["parts"] = {
                  "api",
                  "V1",
                  "GetPromotionDescriptions",
                },
                ["select"] = {
                  ["exist"] = {
                    "language_code",
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["provider"] = {
        ["fields"] = {
          {
            ["name"] = "ErrorCodes",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "Items",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ResultCode",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "provider",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "account_number",
                      ["orig"] = "account_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "country_iso",
                      ["orig"] = "country_iso",
                      ["type"] = "`$ANY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "provider_code",
                      ["orig"] = "provider_code",
                      ["type"] = "`$ANY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "region_code",
                      ["orig"] = "region_code",
                      ["type"] = "`$ANY`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/V1/GetProviders",
                ["parts"] = {
                  "api",
                  "V1",
                  "GetProviders",
                },
                ["select"] = {
                  ["exist"] = {
                    "account_number",
                    "country_iso",
                    "provider_code",
                    "region_code",
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["provider_status"] = {
        ["fields"] = {
          {
            ["name"] = "ErrorCodes",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "Items",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ResultCode",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "provider_status",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "provider_code",
                      ["orig"] = "provider_code",
                      ["type"] = "`$ANY`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/V1/GetProviderStatus",
                ["parts"] = {
                  "api",
                  "V1",
                  "GetProviderStatus",
                },
                ["select"] = {
                  ["exist"] = {
                    "provider_code",
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["region"] = {
        ["fields"] = {
          {
            ["name"] = "ErrorCodes",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "Items",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ResultCode",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "region",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "country_iso",
                      ["orig"] = "country_iso",
                      ["type"] = "`$ANY`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/V1/GetRegions",
                ["parts"] = {
                  "api",
                  "V1",
                  "GetRegions",
                },
                ["select"] = {
                  ["exist"] = {
                    "country_iso",
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["send_transfer"] = {
        ["fields"] = {
          {
            ["name"] = "ErrorCodes",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ResultCode",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "TransferRecord",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
        },
        ["name"] = "send_transfer",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["header"] = {
                    {
                      ["kind"] = "header",
                      ["name"] = "x_correlation_id",
                      ["orig"] = "x_correlation_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "request",
                      ["orig"] = "request",
                      ["reqd"] = true,
                      ["type"] = "`$OBJECT`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/V1/SendTransfer",
                ["parts"] = {
                  "api",
                  "V1",
                  "SendTransfer",
                },
                ["select"] = {
                  ["exist"] = {
                    "request",
                    "x_correlation_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
    },
  }
end


local function make_feature(name)
  local features = require("features")
  local factory = features[name]
  if factory ~= nil then
    return factory()
  end
  return features.base()
end


-- Attach make_feature to the SDK class
local function setup_sdk(SDK)
  SDK._make_feature = make_feature
end


return make_config
