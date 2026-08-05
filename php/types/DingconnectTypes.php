<?php
declare(strict_types=1);

// Typed models for the Dingconnect SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** AccountLookup entity data model. */
class AccountLookup
{
    public ?string $account_number_normalized = null;
    public ?string $country_iso = null;
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Request payload for AccountLookup#list. */
class AccountLookupListMatch
{
    public ?string $account_number_normalized = null;
    public ?string $country_iso = null;
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Balance entity data model. */
class Balance
{
    public ?string $code = null;
    public ?string $context = null;
}

/** Request payload for Balance#list. */
class BalanceListMatch
{
    public ?string $code = null;
    public ?string $context = null;
}

/** CancelResult entity data model. */
class CancelResult
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Request payload for CancelResult#create. */
class CancelResultCreateData
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Country entity data model. */
class Country
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Request payload for Country#list. */
class CountryListMatch
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Currency entity data model. */
class Currency
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Request payload for Currency#list. */
class CurrencyListMatch
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** ErrorCodeDescription entity data model. */
class ErrorCodeDescription
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Request payload for ErrorCodeDescription#list. */
class ErrorCodeDescriptionListMatch
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Estimate entity data model. */
class Estimate
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Request payload for Estimate#create. */
class EstimateCreateData
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Product entity data model. */
class Product
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Request payload for Product#list. */
class ProductListMatch
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** ProductDescription entity data model. */
class ProductDescription
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Request payload for ProductDescription#list. */
class ProductDescriptionListMatch
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Promotion entity data model. */
class Promotion
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Request payload for Promotion#list. */
class PromotionListMatch
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** PromotionDescription entity data model. */
class PromotionDescription
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Request payload for PromotionDescription#list. */
class PromotionDescriptionListMatch
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Provider entity data model. */
class Provider
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Request payload for Provider#list. */
class ProviderListMatch
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** ProviderStatus entity data model. */
class ProviderStatus
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Request payload for ProviderStatus#list. */
class ProviderStatusListMatch
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Region entity data model. */
class Region
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** Request payload for Region#list. */
class RegionListMatch
{
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
}

/** SendTransfer entity data model. */
class SendTransfer
{
    public string $account_number;
    public string $distributor_ref;
    public ?array $error_code = null;
    public ?int $result_code = null;
    public ?string $send_currency_iso = null;
    public float $send_value;
    public ?array $setting = null;
    public string $sku_code;
    public ?array $transfer_record = null;
    public ?bool $validate_only = null;
}

/** Request payload for SendTransfer#create. */
class SendTransferCreateData
{
    public string $account_number;
    public string $distributor_ref;
    public ?array $error_code = null;
    public ?int $result_code = null;
    public ?string $send_currency_iso = null;
    public float $send_value;
    public ?array $setting = null;
    public string $sku_code;
    public ?array $transfer_record = null;
    public ?bool $validate_only = null;
}

/** TransferRecord entity data model. */
class TransferRecord
{
    public ?string $account_number = null;
    public ?array $distributor_ref = null;
    public ?string $ended_at_utc = null;
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
    public int $skip;
    public ?array $sku_code = null;
    public ?string $started_at_utc = null;
    public int $take;
    public ?bool $there_are_more_item = null;
    public ?array $transfer_ref = null;
}

/** Request payload for TransferRecord#create. */
class TransferRecordCreateData
{
    public ?string $account_number = null;
    public ?array $distributor_ref = null;
    public ?string $ended_at_utc = null;
    public ?array $error_code = null;
    public ?array $item = null;
    public ?int $result_code = null;
    public int $skip;
    public ?array $sku_code = null;
    public ?string $started_at_utc = null;
    public int $take;
    public ?bool $there_are_more_item = null;
    public ?array $transfer_ref = null;
}

