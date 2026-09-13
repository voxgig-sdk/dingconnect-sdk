package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "Dingconnect",
			"slug": "dingconnect",
			"version": "0.1.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"transport": "base",
			},
		},
		"options": map[string]any{
			"base": "https://api.dingconnect.com",
			"auth": map[string]any{
				"prefix": "",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"account_lookup": map[string]any{},
				"balance": map[string]any{},
				"cancel_transfer": map[string]any{},
				"country": map[string]any{},
				"currency": map[string]any{},
				"error_code_description": map[string]any{},
				"estimate_price": map[string]any{},
				"list_transfer_record": map[string]any{},
				"lookup_bill": map[string]any{},
				"product": map[string]any{},
				"product_description": map[string]any{},
				"promotion": map[string]any{},
				"promotion_description": map[string]any{},
				"provider": map[string]any{},
				"provider_status": map[string]any{},
				"region": map[string]any{},
				"send_transfer": map[string]any{},
			},
		},
		"entity": map[string]any{
			"account_lookup": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "AccountNumberNormalized",
						"short": "We attempt to normalize phone numbers following the public telecommunication numbering plan <a href=\"https://en.wikipedia.org/wiki/E.164\" target=\"_blank\">E.164</a>, if we succeed the normalized number will be returned in this field formatt…",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "CountryIso",
						"short": "The country of the account number",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "Items",
						"req": true,
						"short": "This will contain provider information associated to the account number.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "int32",
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
				},
				"name": "account_lookup",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "account_number",
											"orig": "account_number",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/V1/GetAccountLookup",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "GetAccountLookup",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_number",
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"api",
									"V1",
									"GetAccountLookup",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"balance": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "Code",
						"req": true,
						"short": "The code that can be used to lookup the explanatory message associated with the error",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "Context",
						"short": "API specific context as to the reason for the specific code",
						"type": "`$STRING`",
					},
				},
				"name": "balance",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/V1/GetBalance",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "GetBalance",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.ErrorCodes`",
								},
								"parts": []any{
									"api",
									"V1",
									"GetBalance",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"cancel_transfer": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "int32",
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
				},
				"name": "cancel_transfer",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "cancellation_request",
											"orig": "cancellation_request",
											"reqd": true,
											"type": "`$ARRAY`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/V1/CancelTransfers",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "CancelTransfers",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"cancellation_request",
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"api",
									"V1",
									"CancelTransfers",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"country": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "Items",
						"req": true,
						"short": "The list of countries that our system is aware of.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "int32",
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
				},
				"name": "country",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/V1/GetCountries",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "GetCountries",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"api",
									"V1",
									"GetCountries",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"currency": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "int32",
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
				},
				"name": "currency",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/V1/GetCurrencies",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "GetCurrencies",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"api",
									"V1",
									"GetCurrencies",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"error_code_description": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "Items",
						"req": true,
						"short": "A list of ErrorCodes and their localized descriptions",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "int32",
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
				},
				"name": "error_code_description",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/V1/GetErrorCodeDescriptions",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "GetErrorCodeDescriptions",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"api",
									"V1",
									"GetErrorCodeDescriptions",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"estimate_price": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "int32",
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
				},
				"name": "estimate_price",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "requested_estimation",
											"orig": "requested_estimation",
											"reqd": true,
											"type": "`$ARRAY`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/V1/EstimatePrices",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "EstimatePrices",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"requested_estimation",
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"api",
									"V1",
									"EstimatePrices",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"list_transfer_record": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "Items",
						"req": true,
						"short": "The list of items satisfying the transfer query.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "int32",
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ThereAreMoreItems",
						"req": true,
						"short": "Indicates if the caller should execute the query again.",
						"type": "`$BOOLEAN`",
					},
				},
				"name": "list_transfer_record",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "request",
											"orig": "request",
											"reqd": true,
											"type": "`$OBJECT`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/V1/ListTransferRecords",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "ListTransferRecords",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"request",
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"api",
									"V1",
									"ListTransferRecords",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"lookup_bill": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "int32",
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
				},
				"name": "lookup_bill",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "request",
											"orig": "request",
											"reqd": true,
											"type": "`$OBJECT`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/V1/LookupBills",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "LookupBills",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"request",
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"api",
									"V1",
									"LookupBills",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"product": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "Items",
						"req": true,
						"short": "A list of products that fulfil the submitted criteria.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "int32",
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
				},
				"name": "product",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "account_number",
											"orig": "account_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "benefit",
											"orig": "benefit",
											"type": "`$ANY`",
										},
										map[string]any{
											"kind": "query",
											"name": "country_iso",
											"orig": "country_iso",
											"type": "`$ANY`",
										},
										map[string]any{
											"kind": "query",
											"name": "provider_code",
											"orig": "provider_code",
											"type": "`$ANY`",
										},
										map[string]any{
											"kind": "query",
											"name": "region_code",
											"orig": "region_code",
											"type": "`$ANY`",
										},
										map[string]any{
											"kind": "query",
											"name": "sku_code",
											"orig": "sku_code",
											"type": "`$ANY`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/V1/GetProducts",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "GetProducts",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_number",
										"benefit",
										"country_iso",
										"provider_code",
										"region_code",
										"sku_code",
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"api",
									"V1",
									"GetProducts",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"product_description": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "Items",
						"req": true,
						"short": "A localized list of product descriptions.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "int32",
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
				},
				"name": "product_description",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "language_code",
											"orig": "language_code",
											"type": "`$ANY`",
										},
										map[string]any{
											"kind": "query",
											"name": "sku_code",
											"orig": "sku_code",
											"type": "`$ANY`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/V1/GetProductDescriptions",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "GetProductDescriptions",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"language_code",
										"sku_code",
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"api",
									"V1",
									"GetProductDescriptions",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"promotion": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "Items",
						"req": true,
						"short": "List of available promotions",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "int32",
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
				},
				"name": "promotion",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "account_number",
											"orig": "account_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "country_iso",
											"orig": "country_iso",
											"type": "`$ANY`",
										},
										map[string]any{
											"kind": "query",
											"name": "provider_code",
											"orig": "provider_code",
											"type": "`$ANY`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/V1/GetPromotions",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "GetPromotions",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_number",
										"country_iso",
										"provider_code",
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"api",
									"V1",
									"GetPromotions",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"promotion_description": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "Items",
						"req": true,
						"short": "A localized list of promotions.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "int32",
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
				},
				"name": "promotion_description",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "language_code",
											"orig": "language_code",
											"type": "`$ANY`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/V1/GetPromotionDescriptions",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "GetPromotionDescriptions",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"language_code",
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"api",
									"V1",
									"GetPromotionDescriptions",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"provider": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "Items",
						"req": true,
						"short": "A list of providers that the distributor has Products for.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "int32",
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
				},
				"name": "provider",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "account_number",
											"orig": "account_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "country_iso",
											"orig": "country_iso",
											"type": "`$ANY`",
										},
										map[string]any{
											"kind": "query",
											"name": "provider_code",
											"orig": "provider_code",
											"type": "`$ANY`",
										},
										map[string]any{
											"kind": "query",
											"name": "region_code",
											"orig": "region_code",
											"type": "`$ANY`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/V1/GetProviders",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "GetProviders",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_number",
										"country_iso",
										"provider_code",
										"region_code",
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"api",
									"V1",
									"GetProviders",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"provider_status": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "int32",
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
				},
				"name": "provider_status",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "provider_code",
											"orig": "provider_code",
											"type": "`$ANY`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/V1/GetProviderStatus",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "GetProviderStatus",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"provider_code",
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"api",
									"V1",
									"GetProviderStatus",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"region": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "Items",
						"req": true,
						"short": "The list of regions that the system uses.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "int32",
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
				},
				"name": "region",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "country_iso",
											"orig": "country_iso",
											"type": "`$ANY`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/V1/GetRegions",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "GetRegions",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"country_iso",
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"api",
									"V1",
									"GetRegions",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"send_transfer": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "int32",
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "TransferRecord",
						"req": true,
						"type": "`$OBJECT`",
					},
				},
				"name": "send_transfer",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "request",
											"orig": "request",
											"reqd": true,
											"type": "`$OBJECT`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/V1/SendTransfer",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "V1",
									},
									map[string]any{
										"lit": "SendTransfer",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"request",
										"x_correlation_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"api",
									"V1",
									"SendTransfer",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
