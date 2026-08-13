package core

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
						"active": true,
						"name": "AccountNumberNormalized",
						"req": false,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "CountryIso",
						"req": false,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 4,
					},
				},
				"name": "account_lookup",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "account_number",
											"orig": "account_number",
											"reqd": false,
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
								"index$": 0,
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
						"active": true,
						"name": "Code",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "Context",
						"req": false,
						"type": "`$STRING`",
						"index$": 1,
					},
				},
				"name": "balance",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
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
								"index$": 0,
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
						"active": true,
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 2,
					},
				},
				"name": "cancel_transfer",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
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
								"index$": 0,
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
						"active": true,
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 2,
					},
				},
				"name": "country",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
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
								"index$": 0,
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
						"active": true,
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 2,
					},
				},
				"name": "currency",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
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
								"index$": 0,
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
						"active": true,
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 2,
					},
				},
				"name": "error_code_description",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
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
								"index$": 0,
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
						"active": true,
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 2,
					},
				},
				"name": "estimate_price",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
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
								"index$": 0,
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
						"active": true,
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "ThereAreMoreItems",
						"req": true,
						"type": "`$BOOLEAN`",
						"index$": 3,
					},
				},
				"name": "list_transfer_record",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
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
								"index$": 0,
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
						"active": true,
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 2,
					},
				},
				"name": "lookup_bill",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
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
								"index$": 0,
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
						"active": true,
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 2,
					},
				},
				"name": "product",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "account_number",
											"orig": "account_number",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "benefit",
											"orig": "benefit",
											"reqd": false,
											"type": "`$ANY`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "country_iso",
											"orig": "country_iso",
											"reqd": false,
											"type": "`$ANY`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "provider_code",
											"orig": "provider_code",
											"reqd": false,
											"type": "`$ANY`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "region_code",
											"orig": "region_code",
											"reqd": false,
											"type": "`$ANY`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "sku_code",
											"orig": "sku_code",
											"reqd": false,
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
								"index$": 0,
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
						"active": true,
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 2,
					},
				},
				"name": "product_description",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "language_code",
											"orig": "language_code",
											"reqd": false,
											"type": "`$ANY`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "sku_code",
											"orig": "sku_code",
											"reqd": false,
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
								"index$": 0,
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
						"active": true,
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 2,
					},
				},
				"name": "promotion",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "account_number",
											"orig": "account_number",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "country_iso",
											"orig": "country_iso",
											"reqd": false,
											"type": "`$ANY`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "provider_code",
											"orig": "provider_code",
											"reqd": false,
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
								"index$": 0,
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
						"active": true,
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 2,
					},
				},
				"name": "promotion_description",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "language_code",
											"orig": "language_code",
											"reqd": false,
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
								"index$": 0,
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
						"active": true,
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 2,
					},
				},
				"name": "provider",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "account_number",
											"orig": "account_number",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "country_iso",
											"orig": "country_iso",
											"reqd": false,
											"type": "`$ANY`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "provider_code",
											"orig": "provider_code",
											"reqd": false,
											"type": "`$ANY`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "region_code",
											"orig": "region_code",
											"reqd": false,
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
								"index$": 0,
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
						"active": true,
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 2,
					},
				},
				"name": "provider_status",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "provider_code",
											"orig": "provider_code",
											"reqd": false,
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
								"index$": 0,
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
						"active": true,
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "Items",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 2,
					},
				},
				"name": "region",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "country_iso",
											"orig": "country_iso",
											"reqd": false,
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
								"index$": 0,
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
						"active": true,
						"name": "ErrorCodes",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "ResultCode",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "TransferRecord",
						"req": true,
						"type": "`$OBJECT`",
						"index$": 2,
					},
				},
				"name": "send_transfer",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "x_correlation_id",
											"orig": "x_correlation_id",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
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
								"index$": 0,
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
