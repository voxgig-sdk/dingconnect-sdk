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
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "CountryIso",
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
						"type": "`$ARRAY`",
					},
					map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetAccountLookup",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "Context",
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
								"parts": []any{
									"api",
									"V1",
									"GetBalance",
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
								"parts": []any{
									"api",
									"V1",
									"CancelTransfers",
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
						"type": "`$ARRAY`",
					},
					map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetCountries",
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
								"parts": []any{
									"api",
									"V1",
									"GetCurrencies",
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
						"type": "`$ARRAY`",
					},
					map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetErrorCodeDescriptions",
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
								"parts": []any{
									"api",
									"V1",
									"EstimatePrices",
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
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ThereAreMoreItems",
						"req": true,
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
								"parts": []any{
									"api",
									"V1",
									"ListTransferRecords",
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
								"parts": []any{
									"api",
									"V1",
									"LookupBills",
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
						"type": "`$ARRAY`",
					},
					map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetProducts",
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
						"type": "`$ARRAY`",
					},
					map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetProductDescriptions",
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
						"type": "`$ARRAY`",
					},
					map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetPromotions",
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
						"type": "`$ARRAY`",
					},
					map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetPromotionDescriptions",
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
						"type": "`$ARRAY`",
					},
					map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetProviders",
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
								"parts": []any{
									"api",
									"V1",
									"GetProviderStatus",
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
						"type": "`$ARRAY`",
					},
					map[string]any{
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
								"parts": []any{
									"api",
									"V1",
									"GetRegions",
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
								"parts": []any{
									"api",
									"V1",
									"SendTransfer",
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
