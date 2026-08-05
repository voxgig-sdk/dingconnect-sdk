<?php
declare(strict_types=1);

// Dingconnect SDK configuration

class DingconnectConfig
{
    public static function make_config(): array
    {
        return [
            "main" => [
                "name" => "Dingconnect",
            ],
            "feature" => [
                "test" => [
          'options' => [
            'active' => false,
          ],
        ],
            ],
            "options" => [
                "base" => "https://api.dingconnect.com/api/V1",
                "auth" => [
                    "prefix" => "",
                ],
                "headers" => [
          'content-type' => 'application/json',
        ],
                "entity" => [
                    "account_lookup" => [],
                    "balance" => [],
                    "cancel_result" => [],
                    "country" => [],
                    "currency" => [],
                    "error_code_description" => [],
                    "estimate" => [],
                    "product" => [],
                    "product_description" => [],
                    "promotion" => [],
                    "promotion_description" => [],
                    "provider" => [],
                    "provider_status" => [],
                    "region" => [],
                    "send_transfer" => [],
                    "transfer_record" => [],
                ],
            ],
            "entity" => [
        'account_lookup' => [
          'fields' => [
            [
              'active' => true,
              'name' => 'account_number_normalized',
              'req' => false,
              'type' => '`$STRING`',
              'index$' => 0,
            ],
            [
              'active' => true,
              'name' => 'country_iso',
              'req' => false,
              'type' => '`$STRING`',
              'index$' => 1,
            ],
            [
              'active' => true,
              'name' => 'error_code',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 2,
            ],
            [
              'active' => true,
              'name' => 'item',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 3,
            ],
            [
              'active' => true,
              'name' => 'result_code',
              'req' => false,
              'type' => '`$INTEGER`',
              'index$' => 4,
            ],
          ],
          'name' => 'account_lookup',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'active' => true,
                  'args' => [
                    'query' => [
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'account_number',
                        'orig' => 'account_number',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/GetAccountLookup',
                  'parts' => [
                    'GetAccountLookup',
                  ],
                  'select' => [
                    'exist' => [
                      'account_number',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'index$' => 0,
                ],
              ],
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'balance' => [
          'fields' => [
            [
              'active' => true,
              'name' => 'code',
              'req' => false,
              'type' => '`$STRING`',
              'index$' => 0,
            ],
            [
              'active' => true,
              'name' => 'context',
              'req' => false,
              'type' => '`$STRING`',
              'index$' => 1,
            ],
          ],
          'name' => 'balance',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'active' => true,
                  'args' => [],
                  'method' => 'GET',
                  'orig' => '/GetBalance',
                  'parts' => [
                    'GetBalance',
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'index$' => 0,
                ],
              ],
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'cancel_result' => [
          'fields' => [
            [
              'active' => true,
              'name' => 'error_code',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 0,
            ],
            [
              'active' => true,
              'name' => 'item',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 1,
            ],
            [
              'active' => true,
              'name' => 'result_code',
              'req' => false,
              'type' => '`$INTEGER`',
              'index$' => 2,
            ],
          ],
          'name' => 'cancel_result',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'active' => true,
                  'args' => [],
                  'method' => 'POST',
                  'orig' => '/CancelTransfers',
                  'parts' => [
                    'CancelTransfers',
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'index$' => 0,
                ],
              ],
              'key$' => 'create',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'country' => [
          'fields' => [
            [
              'active' => true,
              'name' => 'error_code',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 0,
            ],
            [
              'active' => true,
              'name' => 'item',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 1,
            ],
            [
              'active' => true,
              'name' => 'result_code',
              'req' => false,
              'type' => '`$INTEGER`',
              'index$' => 2,
            ],
          ],
          'name' => 'country',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'active' => true,
                  'args' => [],
                  'method' => 'GET',
                  'orig' => '/GetCountries',
                  'parts' => [
                    'GetCountries',
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'index$' => 0,
                ],
              ],
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'currency' => [
          'fields' => [
            [
              'active' => true,
              'name' => 'error_code',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 0,
            ],
            [
              'active' => true,
              'name' => 'item',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 1,
            ],
            [
              'active' => true,
              'name' => 'result_code',
              'req' => false,
              'type' => '`$INTEGER`',
              'index$' => 2,
            ],
          ],
          'name' => 'currency',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'active' => true,
                  'args' => [],
                  'method' => 'GET',
                  'orig' => '/GetCurrencies',
                  'parts' => [
                    'GetCurrencies',
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'index$' => 0,
                ],
              ],
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'error_code_description' => [
          'fields' => [
            [
              'active' => true,
              'name' => 'error_code',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 0,
            ],
            [
              'active' => true,
              'name' => 'item',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 1,
            ],
            [
              'active' => true,
              'name' => 'result_code',
              'req' => false,
              'type' => '`$INTEGER`',
              'index$' => 2,
            ],
          ],
          'name' => 'error_code_description',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'active' => true,
                  'args' => [],
                  'method' => 'GET',
                  'orig' => '/GetErrorCodeDescriptions',
                  'parts' => [
                    'GetErrorCodeDescriptions',
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'index$' => 0,
                ],
              ],
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'estimate' => [
          'fields' => [
            [
              'active' => true,
              'name' => 'error_code',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 0,
            ],
            [
              'active' => true,
              'name' => 'item',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 1,
            ],
            [
              'active' => true,
              'name' => 'result_code',
              'req' => false,
              'type' => '`$INTEGER`',
              'index$' => 2,
            ],
          ],
          'name' => 'estimate',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'active' => true,
                  'args' => [],
                  'method' => 'POST',
                  'orig' => '/EstimatePrices',
                  'parts' => [
                    'EstimatePrices',
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'index$' => 0,
                ],
              ],
              'key$' => 'create',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'product' => [
          'fields' => [
            [
              'active' => true,
              'name' => 'error_code',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 0,
            ],
            [
              'active' => true,
              'name' => 'item',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 1,
            ],
            [
              'active' => true,
              'name' => 'result_code',
              'req' => false,
              'type' => '`$INTEGER`',
              'index$' => 2,
            ],
          ],
          'name' => 'product',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'active' => true,
                  'args' => [
                    'query' => [
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'account_number',
                        'orig' => 'account_number',
                        'reqd' => false,
                        'type' => '`$STRING`',
                      ],
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'benefit',
                        'orig' => 'benefit',
                        'reqd' => false,
                        'type' => '`$ARRAY`',
                      ],
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'country_iso',
                        'orig' => 'country_iso',
                        'reqd' => false,
                        'type' => '`$ARRAY`',
                      ],
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'provider_code',
                        'orig' => 'provider_code',
                        'reqd' => false,
                        'type' => '`$ARRAY`',
                      ],
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'region_code',
                        'orig' => 'region_code',
                        'reqd' => false,
                        'type' => '`$ARRAY`',
                      ],
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'sku_code',
                        'orig' => 'sku_code',
                        'reqd' => false,
                        'type' => '`$ARRAY`',
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/GetProducts',
                  'parts' => [
                    'GetProducts',
                  ],
                  'select' => [
                    'exist' => [
                      'account_number',
                      'benefit',
                      'country_iso',
                      'provider_code',
                      'region_code',
                      'sku_code',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'index$' => 0,
                ],
              ],
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'product_description' => [
          'fields' => [
            [
              'active' => true,
              'name' => 'error_code',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 0,
            ],
            [
              'active' => true,
              'name' => 'item',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 1,
            ],
            [
              'active' => true,
              'name' => 'result_code',
              'req' => false,
              'type' => '`$INTEGER`',
              'index$' => 2,
            ],
          ],
          'name' => 'product_description',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'active' => true,
                  'args' => [
                    'query' => [
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'language_code',
                        'orig' => 'language_code',
                        'reqd' => false,
                        'type' => '`$ARRAY`',
                      ],
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'sku_code',
                        'orig' => 'sku_code',
                        'reqd' => false,
                        'type' => '`$ARRAY`',
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/GetProductDescriptions',
                  'parts' => [
                    'GetProductDescriptions',
                  ],
                  'select' => [
                    'exist' => [
                      'language_code',
                      'sku_code',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'index$' => 0,
                ],
              ],
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'promotion' => [
          'fields' => [
            [
              'active' => true,
              'name' => 'error_code',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 0,
            ],
            [
              'active' => true,
              'name' => 'item',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 1,
            ],
            [
              'active' => true,
              'name' => 'result_code',
              'req' => false,
              'type' => '`$INTEGER`',
              'index$' => 2,
            ],
          ],
          'name' => 'promotion',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'active' => true,
                  'args' => [
                    'query' => [
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'account_number',
                        'orig' => 'account_number',
                        'reqd' => false,
                        'type' => '`$STRING`',
                      ],
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'country_iso',
                        'orig' => 'country_iso',
                        'reqd' => false,
                        'type' => '`$ARRAY`',
                      ],
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'provider_code',
                        'orig' => 'provider_code',
                        'reqd' => false,
                        'type' => '`$ARRAY`',
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/GetPromotions',
                  'parts' => [
                    'GetPromotions',
                  ],
                  'select' => [
                    'exist' => [
                      'account_number',
                      'country_iso',
                      'provider_code',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'index$' => 0,
                ],
              ],
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'promotion_description' => [
          'fields' => [
            [
              'active' => true,
              'name' => 'error_code',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 0,
            ],
            [
              'active' => true,
              'name' => 'item',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 1,
            ],
            [
              'active' => true,
              'name' => 'result_code',
              'req' => false,
              'type' => '`$INTEGER`',
              'index$' => 2,
            ],
          ],
          'name' => 'promotion_description',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'active' => true,
                  'args' => [
                    'query' => [
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'language_code',
                        'orig' => 'language_code',
                        'reqd' => false,
                        'type' => '`$ARRAY`',
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/GetPromotionDescriptions',
                  'parts' => [
                    'GetPromotionDescriptions',
                  ],
                  'select' => [
                    'exist' => [
                      'language_code',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'index$' => 0,
                ],
              ],
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'provider' => [
          'fields' => [
            [
              'active' => true,
              'name' => 'error_code',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 0,
            ],
            [
              'active' => true,
              'name' => 'item',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 1,
            ],
            [
              'active' => true,
              'name' => 'result_code',
              'req' => false,
              'type' => '`$INTEGER`',
              'index$' => 2,
            ],
          ],
          'name' => 'provider',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'active' => true,
                  'args' => [
                    'query' => [
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'account_number',
                        'orig' => 'account_number',
                        'reqd' => false,
                        'type' => '`$STRING`',
                      ],
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'country_iso',
                        'orig' => 'country_iso',
                        'reqd' => false,
                        'type' => '`$ARRAY`',
                      ],
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'provider_code',
                        'orig' => 'provider_code',
                        'reqd' => false,
                        'type' => '`$ARRAY`',
                      ],
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'region_code',
                        'orig' => 'region_code',
                        'reqd' => false,
                        'type' => '`$ARRAY`',
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/GetProviders',
                  'parts' => [
                    'GetProviders',
                  ],
                  'select' => [
                    'exist' => [
                      'account_number',
                      'country_iso',
                      'provider_code',
                      'region_code',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'index$' => 0,
                ],
              ],
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'provider_status' => [
          'fields' => [
            [
              'active' => true,
              'name' => 'error_code',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 0,
            ],
            [
              'active' => true,
              'name' => 'item',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 1,
            ],
            [
              'active' => true,
              'name' => 'result_code',
              'req' => false,
              'type' => '`$INTEGER`',
              'index$' => 2,
            ],
          ],
          'name' => 'provider_status',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'active' => true,
                  'args' => [
                    'query' => [
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'provider_code',
                        'orig' => 'provider_code',
                        'reqd' => false,
                        'type' => '`$ARRAY`',
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/GetProviderStatus',
                  'parts' => [
                    'GetProviderStatus',
                  ],
                  'select' => [
                    'exist' => [
                      'provider_code',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'index$' => 0,
                ],
              ],
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'region' => [
          'fields' => [
            [
              'active' => true,
              'name' => 'error_code',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 0,
            ],
            [
              'active' => true,
              'name' => 'item',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 1,
            ],
            [
              'active' => true,
              'name' => 'result_code',
              'req' => false,
              'type' => '`$INTEGER`',
              'index$' => 2,
            ],
          ],
          'name' => 'region',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'active' => true,
                  'args' => [
                    'query' => [
                      [
                        'active' => true,
                        'kind' => 'query',
                        'name' => 'country_iso',
                        'orig' => 'country_iso',
                        'reqd' => false,
                        'type' => '`$ARRAY`',
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/GetRegions',
                  'parts' => [
                    'GetRegions',
                  ],
                  'select' => [
                    'exist' => [
                      'country_iso',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'index$' => 0,
                ],
              ],
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'send_transfer' => [
          'fields' => [
            [
              'active' => true,
              'name' => 'account_number',
              'req' => true,
              'type' => '`$STRING`',
              'index$' => 0,
            ],
            [
              'active' => true,
              'name' => 'distributor_ref',
              'req' => true,
              'type' => '`$STRING`',
              'index$' => 1,
            ],
            [
              'active' => true,
              'name' => 'error_code',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 2,
            ],
            [
              'active' => true,
              'name' => 'result_code',
              'req' => false,
              'type' => '`$INTEGER`',
              'index$' => 3,
            ],
            [
              'active' => true,
              'name' => 'send_currency_iso',
              'req' => false,
              'type' => '`$STRING`',
              'index$' => 4,
            ],
            [
              'active' => true,
              'name' => 'send_value',
              'req' => true,
              'type' => '`$NUMBER`',
              'index$' => 5,
            ],
            [
              'active' => true,
              'name' => 'setting',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 6,
            ],
            [
              'active' => true,
              'name' => 'sku_code',
              'req' => true,
              'type' => '`$STRING`',
              'index$' => 7,
            ],
            [
              'active' => true,
              'name' => 'transfer_record',
              'req' => false,
              'type' => '`$OBJECT`',
              'index$' => 8,
            ],
            [
              'active' => true,
              'name' => 'validate_only',
              'req' => false,
              'type' => '`$BOOLEAN`',
              'index$' => 9,
            ],
          ],
          'name' => 'send_transfer',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'active' => true,
                  'args' => [],
                  'method' => 'POST',
                  'orig' => '/SendTransfer',
                  'parts' => [
                    'SendTransfer',
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'index$' => 0,
                ],
              ],
              'key$' => 'create',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'transfer_record' => [
          'fields' => [
            [
              'active' => true,
              'name' => 'account_number',
              'req' => false,
              'type' => '`$STRING`',
              'index$' => 0,
            ],
            [
              'active' => true,
              'name' => 'distributor_ref',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 1,
            ],
            [
              'active' => true,
              'name' => 'ended_at_utc',
              'req' => false,
              'type' => '`$STRING`',
              'index$' => 2,
            ],
            [
              'active' => true,
              'name' => 'error_code',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 3,
            ],
            [
              'active' => true,
              'name' => 'item',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 4,
            ],
            [
              'active' => true,
              'name' => 'result_code',
              'req' => false,
              'type' => '`$INTEGER`',
              'index$' => 5,
            ],
            [
              'active' => true,
              'name' => 'skip',
              'req' => true,
              'type' => '`$INTEGER`',
              'index$' => 6,
            ],
            [
              'active' => true,
              'name' => 'sku_code',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 7,
            ],
            [
              'active' => true,
              'name' => 'started_at_utc',
              'req' => false,
              'type' => '`$STRING`',
              'index$' => 8,
            ],
            [
              'active' => true,
              'name' => 'take',
              'req' => true,
              'type' => '`$INTEGER`',
              'index$' => 9,
            ],
            [
              'active' => true,
              'name' => 'there_are_more_item',
              'req' => false,
              'type' => '`$BOOLEAN`',
              'index$' => 10,
            ],
            [
              'active' => true,
              'name' => 'transfer_ref',
              'req' => false,
              'type' => '`$ARRAY`',
              'index$' => 11,
            ],
          ],
          'name' => 'transfer_record',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'active' => true,
                  'args' => [],
                  'method' => 'POST',
                  'orig' => '/ListTransferRecords',
                  'parts' => [
                    'ListTransferRecords',
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'index$' => 0,
                ],
              ],
              'key$' => 'create',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
      ],
        ];
    }


    public static function make_feature(string $name)
    {
        require_once __DIR__ . '/features.php';
        return DingconnectFeatures::make_feature($name);
    }
}
