"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const node_test_1 = require("node:test");
const __1 = require("..");
const definition_runner_1 = require("./definition-runner");
const utility_1 = require("./utility");
// Generated from the API definition, not from the model this SDK was built
// from: the route, the declared query parameters, the credential the security
// scheme names, and the definition's own response example.
const PLAN = [
    {
        "entity": "account_lookup",
        "accessor": "AccountLookup",
        "op": "list",
        "method": "GET",
        "path": "/api/V1/GetAccountLookup",
        "args": [],
        "select": {
            "account_number": ""
        },
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [
            "accountNumber"
        ],
        "queryArgs": [
            {
                "name": "account_number",
                "wire": "accountNumber"
            }
        ],
        "auth": [],
        "status": 200,
        "sample": {
            "CountryIso": "x",
            "AccountNumberNormalized": "x",
            "Items": [
                {
                    "ProviderCode": "x",
                    "RegionCode": "x"
                }
            ],
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ]
        },
        "idField": "id"
    },
    {
        "entity": "balance",
        "accessor": "Balance",
        "op": "list",
        "method": "GET",
        "path": "/api/V1/GetBalance",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": [],
        "status": 200,
        "sample": {
            "Balance": 1,
            "CurrencyIso": "x",
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ]
        },
        "idField": "id"
    },
    {
        "entity": "cancel_transfer",
        "accessor": "CancelTransfer",
        "op": "create",
        "method": "POST",
        "path": "/api/V1/CancelTransfers",
        "args": [],
        "select": {
            "cancellations": "v1"
        },
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": [],
        "status": 200,
        "sample": {
            "Items": [
                {
                    "TransferId": {
                        "TransferRef": "x",
                        "DistributorRef": "x"
                    },
                    "ProcessingState": "x",
                    "BatchItemRef": "x",
                    "ResultCode": 1,
                    "ErrorCodes": [
                        {
                            "Code": "x",
                            "Context": "x"
                        }
                    ]
                }
            ],
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ]
        },
        "idField": "id"
    },
    {
        "entity": "country",
        "accessor": "Country",
        "op": "list",
        "method": "GET",
        "path": "/api/V1/GetCountries",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": [],
        "status": 200,
        "sample": {
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ],
            "Items": [
                {
                    "CountryIso": "x",
                    "CountryName": "x",
                    "InternationalDialingInformation": [
                        {
                            "Prefix": "x",
                            "MinimumLength": 1,
                            "MaximumLength": 1
                        }
                    ],
                    "RegionCodes": [
                        "x"
                    ]
                }
            ]
        },
        "idField": "id"
    },
    {
        "entity": "currency",
        "accessor": "Currency",
        "op": "list",
        "method": "GET",
        "path": "/api/V1/GetCurrencies",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": [],
        "status": 200,
        "sample": {
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ],
            "Items": [
                {
                    "CurrencyIso": "x",
                    "CurrencyName": "x"
                }
            ]
        },
        "idField": "id"
    },
    {
        "entity": "error_code_description",
        "accessor": "ErrorCodeDescription",
        "op": "list",
        "method": "GET",
        "path": "/api/V1/GetErrorCodeDescriptions",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": [],
        "status": 200,
        "sample": {
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ],
            "Items": [
                {
                    "Message": "x",
                    "Code": "x"
                }
            ]
        },
        "idField": "id"
    },
    {
        "entity": "estimate_price",
        "accessor": "EstimatePrice",
        "op": "create",
        "method": "POST",
        "path": "/api/V1/EstimatePrices",
        "args": [],
        "select": {
            "estimations": "v1"
        },
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": [],
        "status": 200,
        "sample": {
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ],
            "Items": [
                {
                    "Price": {
                        "CustomerFee": 1,
                        "DistributorFee": 1,
                        "ReceiveValue": 1,
                        "ReceiveCurrencyIso": "x",
                        "ReceiveValueExcludingTax": 1,
                        "TaxRate": 1,
                        "TaxName": "x",
                        "TaxCalculation": "x",
                        "SendValue": 1,
                        "SendCurrencyIso": "x"
                    },
                    "SkuCode": "x",
                    "BatchItemRef": "x",
                    "ResultCode": 1,
                    "ErrorCodes": [
                        {
                            "Code": "x",
                            "Context": "x"
                        }
                    ]
                }
            ]
        },
        "idField": "id"
    },
    {
        "entity": "list_transfer_record",
        "accessor": "ListTransferRecord",
        "op": "create",
        "method": "POST",
        "path": "/api/V1/ListTransferRecords",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": [],
        "status": 200,
        "sample": {
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ],
            "Items": [
                {
                    "TransferRecord": {
                        "TransferId": {
                            "TransferRef": "x",
                            "DistributorRef": "x"
                        },
                        "SkuCode": "x",
                        "Price": {
                            "CustomerFee": 1,
                            "DistributorFee": 1,
                            "ReceiveValue": 1,
                            "ReceiveCurrencyIso": "x",
                            "ReceiveValueExcludingTax": 1,
                            "TaxRate": 1,
                            "TaxName": "x",
                            "TaxCalculation": "x",
                            "SendValue": 1,
                            "SendCurrencyIso": "x"
                        },
                        "CommissionApplied": 1,
                        "StartedUtc": "2026-01-01T00:00:00Z",
                        "CompletedUtc": "2026-01-01T00:00:00Z",
                        "ProcessingState": "x",
                        "ReceiptText": "x",
                        "ReceiptParams": {},
                        "AccountNumber": "x"
                    },
                    "ResultCode": 1,
                    "ErrorCodes": [
                        {
                            "Code": "x",
                            "Context": "x"
                        }
                    ]
                }
            ],
            "ThereAreMoreItems": true
        },
        "idField": "id"
    },
    {
        "entity": "lookup_bill",
        "accessor": "LookupBill",
        "op": "create",
        "method": "POST",
        "path": "/api/V1/LookupBills",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": [],
        "status": 200,
        "sample": {
            "Items": [
                {
                    "Price": {
                        "CustomerFee": 1,
                        "DistributorFee": 1,
                        "ReceiveValue": 1,
                        "ReceiveCurrencyIso": "x",
                        "ReceiveValueExcludingTax": 1,
                        "TaxRate": 1,
                        "TaxName": "x",
                        "TaxCalculation": "x",
                        "SendValue": 1,
                        "SendCurrencyIso": "x"
                    },
                    "BillRef": "x",
                    "AdditionalInfo": {}
                }
            ],
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ]
        },
        "idField": "id"
    },
    {
        "entity": "product",
        "accessor": "Product",
        "op": "list",
        "method": "GET",
        "path": "/api/V1/GetProducts",
        "args": [],
        "select": {
            "account_number": "",
            "benefit": "v1",
            "country_iso": "v1",
            "provider_code": "v1",
            "region_code": "v1",
            "sku_code": "v1"
        },
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [
            "countryIsos",
            "providerCodes",
            "skuCodes",
            "benefits",
            "regionCodes",
            "accountNumber"
        ],
        "queryArgs": [
            {
                "name": "account_number",
                "wire": "accountNumber"
            },
            {
                "name": "benefit",
                "wire": "benefits"
            },
            {
                "name": "country_iso",
                "wire": "countryIsos"
            },
            {
                "name": "provider_code",
                "wire": "providerCodes"
            },
            {
                "name": "region_code",
                "wire": "regionCodes"
            },
            {
                "name": "sku_code",
                "wire": "skuCodes"
            }
        ],
        "auth": [],
        "status": 200,
        "sample": {
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ],
            "Items": [
                {
                    "ProviderCode": "x",
                    "SkuCode": "x",
                    "LocalizationKey": "x",
                    "SettingDefinitions": [
                        {
                            "Name": "x",
                            "Description": "x",
                            "IsMandatory": true
                        }
                    ],
                    "Maximum": {
                        "CustomerFee": 1,
                        "DistributorFee": 1,
                        "ReceiveValue": 1,
                        "ReceiveCurrencyIso": "x",
                        "ReceiveValueExcludingTax": 1,
                        "TaxRate": 1,
                        "TaxName": "x",
                        "TaxCalculation": "x",
                        "SendValue": 1,
                        "SendCurrencyIso": "x"
                    },
                    "Minimum": {
                        "CustomerFee": 1,
                        "DistributorFee": 1,
                        "ReceiveValue": 1,
                        "ReceiveCurrencyIso": "x",
                        "ReceiveValueExcludingTax": 1,
                        "TaxRate": 1,
                        "TaxName": "x",
                        "TaxCalculation": "x",
                        "SendValue": 1,
                        "SendCurrencyIso": "x"
                    },
                    "CommissionRate": 1,
                    "ProcessingMode": "x",
                    "RedemptionMechanism": "x",
                    "Benefits": [
                        "x"
                    ],
                    "ValidityPeriodIso": "x",
                    "UatNumber": "x",
                    "AdditionalInformation": "x",
                    "DefaultDisplayText": "x",
                    "RegionCode": "x",
                    "PaymentTypes": [
                        "x"
                    ],
                    "LookupBillsRequired": true
                }
            ]
        },
        "idField": "id"
    },
    {
        "entity": "product_description",
        "accessor": "ProductDescription",
        "op": "list",
        "method": "GET",
        "path": "/api/V1/GetProductDescriptions",
        "args": [],
        "select": {
            "language_code": "v1",
            "sku_code": "v1"
        },
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [
            "languageCodes",
            "skuCodes"
        ],
        "queryArgs": [
            {
                "name": "language_code",
                "wire": "languageCodes"
            },
            {
                "name": "sku_code",
                "wire": "skuCodes"
            }
        ],
        "auth": [],
        "status": 200,
        "sample": {
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ],
            "Items": [
                {
                    "DisplayText": "x",
                    "DescriptionMarkdown": "x",
                    "ReadMoreMarkdown": "x",
                    "LocalizationKey": "x",
                    "LanguageCode": "x"
                }
            ]
        },
        "idField": "id"
    },
    {
        "entity": "promotion",
        "accessor": "Promotion",
        "op": "list",
        "method": "GET",
        "path": "/api/V1/GetPromotions",
        "args": [],
        "select": {
            "account_number": "",
            "country_iso": "v1",
            "provider_code": "v1"
        },
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [
            "countryIsos",
            "providerCodes",
            "accountNumber"
        ],
        "queryArgs": [
            {
                "name": "account_number",
                "wire": "accountNumber"
            },
            {
                "name": "country_iso",
                "wire": "countryIsos"
            },
            {
                "name": "provider_code",
                "wire": "providerCodes"
            }
        ],
        "auth": [],
        "status": 200,
        "sample": {
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ],
            "Items": [
                {
                    "ProviderCode": "x",
                    "StartUtc": "2026-01-01T00:00:00Z",
                    "EndUtc": "2026-01-01T00:00:00Z",
                    "CurrencyIso": "x",
                    "ValidityPeriodIso": "x",
                    "MinimumSendAmount": 1,
                    "LocalizationKey": "x"
                }
            ]
        },
        "idField": "id"
    },
    {
        "entity": "promotion_description",
        "accessor": "PromotionDescription",
        "op": "list",
        "method": "GET",
        "path": "/api/V1/GetPromotionDescriptions",
        "args": [],
        "select": {
            "language_code": "v1"
        },
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [
            "languageCodes"
        ],
        "queryArgs": [
            {
                "name": "language_code",
                "wire": "languageCodes"
            }
        ],
        "auth": [],
        "status": 200,
        "sample": {
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ],
            "Items": [
                {
                    "Dates": "x",
                    "Headline": "x",
                    "TermsAndConditionsMarkDown": "x",
                    "BonusValidity": "x",
                    "PromotionType": "x",
                    "LocalizationKey": "x",
                    "LanguageCode": "x",
                    "SendAmounts": [
                        {
                            "Minimum": 1,
                            "Maximum": 1,
                            "Bonuses": [
                                {}
                            ]
                        }
                    ]
                }
            ]
        },
        "idField": "id"
    },
    {
        "entity": "provider",
        "accessor": "Provider",
        "op": "list",
        "method": "GET",
        "path": "/api/V1/GetProviders",
        "args": [],
        "select": {
            "account_number": "",
            "country_iso": "v1",
            "provider_code": "v1",
            "region_code": "v1"
        },
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [
            "providerCodes",
            "countryIsos",
            "regionCodes",
            "accountNumber"
        ],
        "queryArgs": [
            {
                "name": "account_number",
                "wire": "accountNumber"
            },
            {
                "name": "country_iso",
                "wire": "countryIsos"
            },
            {
                "name": "provider_code",
                "wire": "providerCodes"
            },
            {
                "name": "region_code",
                "wire": "regionCodes"
            }
        ],
        "auth": [],
        "status": 200,
        "sample": {
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ],
            "Items": [
                {
                    "ProviderCode": "x",
                    "CountryIso": "x",
                    "Name": "x",
                    "ShortName": "x",
                    "ValidationRegex": "x",
                    "CustomerCareNumber": "x",
                    "RegionCodes": [
                        "x"
                    ],
                    "PaymentTypes": [
                        "x"
                    ],
                    "LogoUrl": "x"
                }
            ]
        },
        "idField": "id"
    },
    {
        "entity": "provider_status",
        "accessor": "ProviderStatus",
        "op": "list",
        "method": "GET",
        "path": "/api/V1/GetProviderStatus",
        "args": [],
        "select": {
            "provider_code": "v1"
        },
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [
            "providerCodes"
        ],
        "queryArgs": [
            {
                "name": "provider_code",
                "wire": "providerCodes"
            }
        ],
        "auth": [],
        "status": 200,
        "sample": {
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ],
            "Items": [
                {
                    "ProviderCode": "x",
                    "IsProcessingTransfers": true,
                    "Message": "x"
                }
            ]
        },
        "idField": "id"
    },
    {
        "entity": "region",
        "accessor": "Region",
        "op": "list",
        "method": "GET",
        "path": "/api/V1/GetRegions",
        "args": [],
        "select": {
            "country_iso": "v1"
        },
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [
            "countryIsos"
        ],
        "queryArgs": [
            {
                "name": "country_iso",
                "wire": "countryIsos"
            }
        ],
        "auth": [],
        "status": 200,
        "sample": {
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ],
            "Items": [
                {
                    "RegionCode": "x",
                    "RegionName": "x",
                    "CountryIso": "x"
                }
            ]
        },
        "idField": "id"
    },
    {
        "entity": "send_transfer",
        "accessor": "SendTransfer",
        "op": "create",
        "method": "POST",
        "path": "/api/V1/SendTransfer",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_correlation_id",
                "wire": "X-Correlation-Id",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": [],
        "status": 200,
        "sample": {
            "TransferRecord": {
                "TransferId": {
                    "TransferRef": "x",
                    "DistributorRef": "x"
                },
                "SkuCode": "x",
                "Price": {
                    "CustomerFee": 1,
                    "DistributorFee": 1,
                    "ReceiveValue": 1,
                    "ReceiveCurrencyIso": "x",
                    "ReceiveValueExcludingTax": 1,
                    "TaxRate": 1,
                    "TaxName": "x",
                    "TaxCalculation": "x",
                    "SendValue": 1,
                    "SendCurrencyIso": "x"
                },
                "CommissionApplied": 1,
                "StartedUtc": "2026-01-01T00:00:00Z",
                "CompletedUtc": "2026-01-01T00:00:00Z",
                "ProcessingState": "x",
                "ReceiptText": "x",
                "ReceiptParams": {},
                "AccountNumber": "x"
            },
            "ResultCode": 1,
            "ErrorCodes": [
                {
                    "Code": "x",
                    "Context": "x"
                }
            ]
        },
        "idField": "id"
    }
];
(0, node_test_1.describe)('definition', () => {
    for (const point of PLAN) {
        (0, node_test_1.test)(point.entity + '.' + point.op + ' ' + point.method + ' ' + point.path, async (t) => {
            const control = (0, utility_1.isControlSkipped)('entityOp', point.entity + '.' + point.op, 'definition');
            if (control.skip) {
                t.skip(control.reason || 'skipped via sdk-test-control.json');
                return;
            }
            await (0, definition_runner_1.runDefinitionPoint)(__1.SDK, point);
        });
    }
});
//# sourceMappingURL=definition.test.js.map