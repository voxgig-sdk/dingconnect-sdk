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
			"debug": map[string]any{
				"options": map[string]any{
					"active": false,
					"max": 100,
					"redact": []any{
						"authorization",
						"cookie",
						"set-cookie",
						"api-key",
						"apikey",
						"x-api-key",
						"idempotency-key",
					},
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"onEntry": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"idempotency": map[string]any{
				"options": map[string]any{
					"active": false,
					"header": "Idempotency-Key",
					"methods": []any{
						"POST",
						"PUT",
						"PATCH",
						"DELETE",
					},
					"ops": []any{
						"create",
						"update",
						"remove",
					},
				},
				"optspec": map[string]any{
					"keygen": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"metrics": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"paging": map[string]any{
				"options": map[string]any{
					"active": false,
					"afterVar": "after",
					"cursorParam": "cursor",
					"firstVar": "first",
					"limitParam": "limit",
					"pageParam": "page",
					"startPage": 1,
				},
				"optspec": map[string]any{
					"limit": "`$NUMBER`",
					"ops": "`$LIST`",
				},
				"strict": false,
				"transport": "none",
			},
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
		},
		"options": map[string]any{
			"base": "https://api.dingconnect.com",
			"auth": map[string]any{
				"prefix": "",
				"name": "api_key",
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
						"title": "Account Number Normalized",
						"type": "`$STRING`",
						"short": "We attempt to normalize phone numbers following the public telecommunication numbering plan <a href=\"https://en.wikipedia.org/wiki/E.164\" target=\"_blank\">E.164</a>, if we succeed the normalized number will be returned in this field formatt…",
					},
					map[string]any{
						"name": "CountryIso",
						"title": "Country Iso",
						"type": "`$STRING`",
						"short": "The country of the account number",
					},
					map[string]any{
						"name": "ErrorCodes",
						"title": "Error Codes",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "Items",
						"title": "Items",
						"type": "`$ARRAY`",
						"req": true,
						"short": "This will contain provider information associated to the account number.",
					},
					map[string]any{
						"name": "ResultCode",
						"title": "Result Code",
						"type": "`$INTEGER`",
						"req": true,
						"format": "int32",
					},
				},
				"name": "account_lookup",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetAccountLookup",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "account_number",
											"orig": "account_number",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_number",
										"x_correlation_id",
									},
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
						"title": "Code",
						"type": "`$STRING`",
						"req": true,
						"short": "The code that can be used to lookup the explanatory message associated with the error",
					},
					map[string]any{
						"name": "Context",
						"title": "Context",
						"type": "`$STRING`",
						"short": "API specific context as to the reason for the specific code",
					},
				},
				"name": "balance",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetBalance",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.ErrorCodes`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_correlation_id",
									},
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
						"title": "Error Codes",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "Items",
						"title": "Items",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "ResultCode",
						"title": "Result Code",
						"type": "`$INTEGER`",
						"req": true,
						"format": "int32",
					},
				},
				"name": "cancel_transfer",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"CancelTransfers",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "cancellation_request",
											"orig": "cancellation_request",
											"type": "`$ARRAY`",
											"kind": "query",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"cancellation_request",
										"x_correlation_id",
									},
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
						"title": "Error Codes",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "Items",
						"title": "Items",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The list of countries that our system is aware of.",
					},
					map[string]any{
						"name": "ResultCode",
						"title": "Result Code",
						"type": "`$INTEGER`",
						"req": true,
						"format": "int32",
					},
				},
				"name": "country",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetCountries",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_correlation_id",
									},
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
						"title": "Error Codes",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "Items",
						"title": "Items",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "ResultCode",
						"title": "Result Code",
						"type": "`$INTEGER`",
						"req": true,
						"format": "int32",
					},
				},
				"name": "currency",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetCurrencies",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_correlation_id",
									},
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
						"title": "Error Codes",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "Items",
						"title": "Items",
						"type": "`$ARRAY`",
						"req": true,
						"short": "A list of ErrorCodes and their localized descriptions",
					},
					map[string]any{
						"name": "ResultCode",
						"title": "Result Code",
						"type": "`$INTEGER`",
						"req": true,
						"format": "int32",
					},
				},
				"name": "error_code_description",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetErrorCodeDescriptions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_correlation_id",
									},
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
						"title": "Error Codes",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "Items",
						"title": "Items",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "ResultCode",
						"title": "Result Code",
						"type": "`$INTEGER`",
						"req": true,
						"format": "int32",
					},
				},
				"name": "estimate_price",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"EstimatePrices",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "requested_estimation",
											"orig": "requested_estimation",
											"type": "`$ARRAY`",
											"kind": "query",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"requested_estimation",
										"x_correlation_id",
									},
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
						"title": "Error Codes",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "Items",
						"title": "Items",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The list of items satisfying the transfer query.",
					},
					map[string]any{
						"name": "ResultCode",
						"title": "Result Code",
						"type": "`$INTEGER`",
						"req": true,
						"format": "int32",
					},
					map[string]any{
						"name": "ThereAreMoreItems",
						"title": "There Are More Items",
						"type": "`$BOOLEAN`",
						"req": true,
						"short": "Indicates if the caller should execute the query again.",
					},
				},
				"name": "list_transfer_record",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"ListTransferRecords",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "request",
											"orig": "request",
											"type": "`$OBJECT`",
											"kind": "query",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"request",
										"x_correlation_id",
									},
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
						"title": "Error Codes",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "Items",
						"title": "Items",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "ResultCode",
						"title": "Result Code",
						"type": "`$INTEGER`",
						"req": true,
						"format": "int32",
					},
				},
				"name": "lookup_bill",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"LookupBills",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "request",
											"orig": "request",
											"type": "`$OBJECT`",
											"kind": "query",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"request",
										"x_correlation_id",
									},
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
						"title": "Error Codes",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "Items",
						"title": "Items",
						"type": "`$ARRAY`",
						"req": true,
						"short": "A list of products that fulfil the submitted criteria.",
					},
					map[string]any{
						"name": "ResultCode",
						"title": "Result Code",
						"type": "`$INTEGER`",
						"req": true,
						"format": "int32",
					},
				},
				"name": "product",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetProducts",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "account_number",
											"orig": "account_number",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "benefit",
											"orig": "benefit",
											"type": "`$ANY`",
											"kind": "query",
										},
										map[string]any{
											"name": "country_iso",
											"orig": "country_iso",
											"type": "`$ANY`",
											"kind": "query",
										},
										map[string]any{
											"name": "provider_code",
											"orig": "provider_code",
											"type": "`$ANY`",
											"kind": "query",
										},
										map[string]any{
											"name": "region_code",
											"orig": "region_code",
											"type": "`$ANY`",
											"kind": "query",
										},
										map[string]any{
											"name": "sku_code",
											"orig": "sku_code",
											"type": "`$ANY`",
											"kind": "query",
										},
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
						"title": "Error Codes",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "Items",
						"title": "Items",
						"type": "`$ARRAY`",
						"req": true,
						"short": "A localized list of product descriptions.",
					},
					map[string]any{
						"name": "ResultCode",
						"title": "Result Code",
						"type": "`$INTEGER`",
						"req": true,
						"format": "int32",
					},
				},
				"name": "product_description",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetProductDescriptions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "language_code",
											"orig": "language_code",
											"type": "`$ANY`",
											"kind": "query",
										},
										map[string]any{
											"name": "sku_code",
											"orig": "sku_code",
											"type": "`$ANY`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"language_code",
										"sku_code",
										"x_correlation_id",
									},
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
						"title": "Error Codes",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "Items",
						"title": "Items",
						"type": "`$ARRAY`",
						"req": true,
						"short": "List of available promotions",
					},
					map[string]any{
						"name": "ResultCode",
						"title": "Result Code",
						"type": "`$INTEGER`",
						"req": true,
						"format": "int32",
					},
				},
				"name": "promotion",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetPromotions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "account_number",
											"orig": "account_number",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "country_iso",
											"orig": "country_iso",
											"type": "`$ANY`",
											"kind": "query",
										},
										map[string]any{
											"name": "provider_code",
											"orig": "provider_code",
											"type": "`$ANY`",
											"kind": "query",
										},
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
						"title": "Error Codes",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "Items",
						"title": "Items",
						"type": "`$ARRAY`",
						"req": true,
						"short": "A localized list of promotions.",
					},
					map[string]any{
						"name": "ResultCode",
						"title": "Result Code",
						"type": "`$INTEGER`",
						"req": true,
						"format": "int32",
					},
				},
				"name": "promotion_description",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetPromotionDescriptions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "language_code",
											"orig": "language_code",
											"type": "`$ANY`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"language_code",
										"x_correlation_id",
									},
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
						"title": "Error Codes",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "Items",
						"title": "Items",
						"type": "`$ARRAY`",
						"req": true,
						"short": "A list of providers that the distributor has Products for.",
					},
					map[string]any{
						"name": "ResultCode",
						"title": "Result Code",
						"type": "`$INTEGER`",
						"req": true,
						"format": "int32",
					},
				},
				"name": "provider",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetProviders",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "account_number",
											"orig": "account_number",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "country_iso",
											"orig": "country_iso",
											"type": "`$ANY`",
											"kind": "query",
										},
										map[string]any{
											"name": "provider_code",
											"orig": "provider_code",
											"type": "`$ANY`",
											"kind": "query",
										},
										map[string]any{
											"name": "region_code",
											"orig": "region_code",
											"type": "`$ANY`",
											"kind": "query",
										},
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
						"title": "Error Codes",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "Items",
						"title": "Items",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "ResultCode",
						"title": "Result Code",
						"type": "`$INTEGER`",
						"req": true,
						"format": "int32",
					},
				},
				"name": "provider_status",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetProviderStatus",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "provider_code",
											"orig": "provider_code",
											"type": "`$ANY`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"provider_code",
										"x_correlation_id",
									},
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
						"title": "Error Codes",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "Items",
						"title": "Items",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The list of regions that the system uses.",
					},
					map[string]any{
						"name": "ResultCode",
						"title": "Result Code",
						"type": "`$INTEGER`",
						"req": true,
						"format": "int32",
					},
				},
				"name": "region",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetRegions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "country_iso",
											"orig": "country_iso",
											"type": "`$ANY`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"country_iso",
										"x_correlation_id",
									},
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
						"title": "Error Codes",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "ResultCode",
						"title": "Result Code",
						"type": "`$INTEGER`",
						"req": true,
						"format": "int32",
					},
					map[string]any{
						"name": "TransferRecord",
						"title": "Transfer Record",
						"type": "`$OBJECT`",
						"req": true,
					},
				},
				"name": "send_transfer",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"SendTransfer",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "request",
											"orig": "request",
											"type": "`$OBJECT`",
											"kind": "query",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"request",
										"x_correlation_id",
									},
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
	case "debug":
		if NewDebugFeatureFunc != nil {
			return NewDebugFeatureFunc()
		}
	case "idempotency":
		if NewIdempotencyFeatureFunc != nil {
			return NewIdempotencyFeatureFunc()
		}
	case "metrics":
		if NewMetricsFeatureFunc != nil {
			return NewMetricsFeatureFunc()
		}
	case "paging":
		if NewPagingFeatureFunc != nil {
			return NewPagingFeatureFunc()
		}
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
