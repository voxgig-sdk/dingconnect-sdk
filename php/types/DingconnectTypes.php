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
    public ?string $AccountNumberNormalized = null;
    public ?string $CountryIso = null;
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** Request payload for AccountLookup#list. */
class AccountLookupListMatch
{
    public ?string $AccountNumberNormalized = null;
    public ?string $CountryIso = null;
    public ?array $ErrorCodes = null;
    public ?array $Items = null;
    public ?int $ResultCode = null;
}

/** Balance entity data model. */
class Balance
{
    public string $Code;
    public ?string $Context = null;
}

/** Request payload for Balance#list. */
class BalanceListMatch
{
    public ?string $Code = null;
    public ?string $Context = null;
}

/** CancelTransfer entity data model. */
class CancelTransfer
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** Request payload for CancelTransfer#create. */
class CancelTransferCreateData
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** Country entity data model. */
class Country
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** Request payload for Country#list. */
class CountryListMatch
{
    public ?array $ErrorCodes = null;
    public ?array $Items = null;
    public ?int $ResultCode = null;
}

/** Currency entity data model. */
class Currency
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** Request payload for Currency#list. */
class CurrencyListMatch
{
    public ?array $ErrorCodes = null;
    public ?array $Items = null;
    public ?int $ResultCode = null;
}

/** ErrorCodeDescription entity data model. */
class ErrorCodeDescription
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** Request payload for ErrorCodeDescription#list. */
class ErrorCodeDescriptionListMatch
{
    public ?array $ErrorCodes = null;
    public ?array $Items = null;
    public ?int $ResultCode = null;
}

/** EstimatePrice entity data model. */
class EstimatePrice
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** Request payload for EstimatePrice#create. */
class EstimatePriceCreateData
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** ListTransferRecord entity data model. */
class ListTransferRecord
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
    public bool $ThereAreMoreItems;
}

/** Request payload for ListTransferRecord#create. */
class ListTransferRecordCreateData
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
    public bool $ThereAreMoreItems;
}

/** LookupBill entity data model. */
class LookupBill
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** Request payload for LookupBill#create. */
class LookupBillCreateData
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** Product entity data model. */
class Product
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** Request payload for Product#list. */
class ProductListMatch
{
    public ?array $ErrorCodes = null;
    public ?array $Items = null;
    public ?int $ResultCode = null;
}

/** ProductDescription entity data model. */
class ProductDescription
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** Request payload for ProductDescription#list. */
class ProductDescriptionListMatch
{
    public ?array $ErrorCodes = null;
    public ?array $Items = null;
    public ?int $ResultCode = null;
}

/** Promotion entity data model. */
class Promotion
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** Request payload for Promotion#list. */
class PromotionListMatch
{
    public ?array $ErrorCodes = null;
    public ?array $Items = null;
    public ?int $ResultCode = null;
}

/** PromotionDescription entity data model. */
class PromotionDescription
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** Request payload for PromotionDescription#list. */
class PromotionDescriptionListMatch
{
    public ?array $ErrorCodes = null;
    public ?array $Items = null;
    public ?int $ResultCode = null;
}

/** Provider entity data model. */
class Provider
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** Request payload for Provider#list. */
class ProviderListMatch
{
    public ?array $ErrorCodes = null;
    public ?array $Items = null;
    public ?int $ResultCode = null;
}

/** ProviderStatus entity data model. */
class ProviderStatus
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** Request payload for ProviderStatus#list. */
class ProviderStatusListMatch
{
    public ?array $ErrorCodes = null;
    public ?array $Items = null;
    public ?int $ResultCode = null;
}

/** Region entity data model. */
class Region
{
    public array $ErrorCodes;
    public array $Items;
    public int $ResultCode;
}

/** Request payload for Region#list. */
class RegionListMatch
{
    public ?array $ErrorCodes = null;
    public ?array $Items = null;
    public ?int $ResultCode = null;
}

/** SendTransfer entity data model. */
class SendTransfer
{
    public array $ErrorCodes;
    public int $ResultCode;
    public array $TransferRecord;
}

/** Request payload for SendTransfer#create. */
class SendTransferCreateData
{
    public array $ErrorCodes;
    public int $ResultCode;
    public array $TransferRecord;
}

