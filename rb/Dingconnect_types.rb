# frozen_string_literal: true

# Typed models for the Dingconnect SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# AccountLookup entity data model.
#
# @!attribute [rw] AccountNumberNormalized
#   @return [String, nil]
#
# @!attribute [rw] CountryIso
#   @return [String, nil]
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
AccountLookup = Struct.new(
  :AccountNumberNormalized,
  :CountryIso,
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# Request payload for AccountLookup#list.
#
# @!attribute [rw] account_number
#   @return [String, nil]
AccountLookupListMatch = Struct.new(
  :account_number,
  keyword_init: true
)

# Balance entity data model.
#
# @!attribute [rw] Code
#   @return [String]
#
# @!attribute [rw] Context
#   @return [String, nil]
Balance = Struct.new(
  :Code,
  :Context,
  keyword_init: true
)

# Request payload for Balance#list.
#
# @!attribute [rw] Code
#   @return [String, nil]
#
# @!attribute [rw] Context
#   @return [String, nil]
BalanceListMatch = Struct.new(
  :Code,
  :Context,
  keyword_init: true
)

# CancelTransfer entity data model.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
#
# @!attribute [rw] cancellations
#   @return [Array, nil]
CancelTransfer = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  :cancellations,
  keyword_init: true
)

# Request payload for CancelTransfer#create.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
#
# @!attribute [rw] cancellations
#   @return [Array, nil]
CancelTransferCreateData = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  :cancellations,
  keyword_init: true
)

# Country entity data model.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
Country = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# Request payload for Country#list.
#
# @!attribute [rw] ErrorCodes
#   @return [Array, nil]
#
# @!attribute [rw] Items
#   @return [Array, nil]
#
# @!attribute [rw] ResultCode
#   @return [Integer, nil]
CountryListMatch = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# Currency entity data model.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
Currency = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# Request payload for Currency#list.
#
# @!attribute [rw] ErrorCodes
#   @return [Array, nil]
#
# @!attribute [rw] Items
#   @return [Array, nil]
#
# @!attribute [rw] ResultCode
#   @return [Integer, nil]
CurrencyListMatch = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# ErrorCodeDescription entity data model.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
ErrorCodeDescription = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# Request payload for ErrorCodeDescription#list.
#
# @!attribute [rw] ErrorCodes
#   @return [Array, nil]
#
# @!attribute [rw] Items
#   @return [Array, nil]
#
# @!attribute [rw] ResultCode
#   @return [Integer, nil]
ErrorCodeDescriptionListMatch = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# EstimatePrice entity data model.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
#
# @!attribute [rw] estimations
#   @return [Array, nil]
EstimatePrice = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  :estimations,
  keyword_init: true
)

# Request payload for EstimatePrice#create.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
#
# @!attribute [rw] estimations
#   @return [Array, nil]
EstimatePriceCreateData = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  :estimations,
  keyword_init: true
)

# ListTransferRecord entity data model.
#
# @!attribute [rw] AccountNumber
#   @return [String, nil]
#
# @!attribute [rw] DistributorRef
#   @return [String, nil]
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
#
# @!attribute [rw] Skip
#   @return [Integer, nil]
#
# @!attribute [rw] Take
#   @return [Integer]
#
# @!attribute [rw] ThereAreMoreItems
#   @return [Boolean]
#
# @!attribute [rw] TransferRef
#   @return [String, nil]
ListTransferRecord = Struct.new(
  :AccountNumber,
  :DistributorRef,
  :ErrorCodes,
  :Items,
  :ResultCode,
  :Skip,
  :Take,
  :ThereAreMoreItems,
  :TransferRef,
  keyword_init: true
)

# Request payload for ListTransferRecord#create.
#
# @!attribute [rw] AccountNumber
#   @return [String, nil]
#
# @!attribute [rw] DistributorRef
#   @return [String, nil]
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
#
# @!attribute [rw] Skip
#   @return [Integer, nil]
#
# @!attribute [rw] Take
#   @return [Integer]
#
# @!attribute [rw] ThereAreMoreItems
#   @return [Boolean]
#
# @!attribute [rw] TransferRef
#   @return [String, nil]
ListTransferRecordCreateData = Struct.new(
  :AccountNumber,
  :DistributorRef,
  :ErrorCodes,
  :Items,
  :ResultCode,
  :Skip,
  :Take,
  :ThereAreMoreItems,
  :TransferRef,
  keyword_init: true
)

# LookupBill entity data model.
#
# @!attribute [rw] AccountNumber
#   @return [String]
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
#
# @!attribute [rw] Settings
#   @return [Array, nil]
#
# @!attribute [rw] SkuCode
#   @return [String]
LookupBill = Struct.new(
  :AccountNumber,
  :ErrorCodes,
  :Items,
  :ResultCode,
  :Settings,
  :SkuCode,
  keyword_init: true
)

# Request payload for LookupBill#create.
#
# @!attribute [rw] AccountNumber
#   @return [String]
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
#
# @!attribute [rw] Settings
#   @return [Array, nil]
#
# @!attribute [rw] SkuCode
#   @return [String]
LookupBillCreateData = Struct.new(
  :AccountNumber,
  :ErrorCodes,
  :Items,
  :ResultCode,
  :Settings,
  :SkuCode,
  keyword_init: true
)

# Product entity data model.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
Product = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# Request payload for Product#list.
#
# @!attribute [rw] account_number
#   @return [String, nil]
#
# @!attribute [rw] benefit
#   @return [Array, nil]
#
# @!attribute [rw] country_iso
#   @return [Array, nil]
#
# @!attribute [rw] provider_code
#   @return [Array, nil]
#
# @!attribute [rw] region_code
#   @return [Array, nil]
#
# @!attribute [rw] sku_code
#   @return [Array, nil]
ProductListMatch = Struct.new(
  :account_number,
  :benefit,
  :country_iso,
  :provider_code,
  :region_code,
  :sku_code,
  keyword_init: true
)

# ProductDescription entity data model.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
ProductDescription = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# Request payload for ProductDescription#list.
#
# @!attribute [rw] language_code
#   @return [Array, nil]
#
# @!attribute [rw] sku_code
#   @return [Array, nil]
ProductDescriptionListMatch = Struct.new(
  :language_code,
  :sku_code,
  keyword_init: true
)

# Promotion entity data model.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
Promotion = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# Request payload for Promotion#list.
#
# @!attribute [rw] account_number
#   @return [String, nil]
#
# @!attribute [rw] country_iso
#   @return [Array, nil]
#
# @!attribute [rw] provider_code
#   @return [Array, nil]
PromotionListMatch = Struct.new(
  :account_number,
  :country_iso,
  :provider_code,
  keyword_init: true
)

# PromotionDescription entity data model.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
PromotionDescription = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# Request payload for PromotionDescription#list.
#
# @!attribute [rw] language_code
#   @return [Array, nil]
PromotionDescriptionListMatch = Struct.new(
  :language_code,
  keyword_init: true
)

# Provider entity data model.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
Provider = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# Request payload for Provider#list.
#
# @!attribute [rw] account_number
#   @return [String, nil]
#
# @!attribute [rw] country_iso
#   @return [Array, nil]
#
# @!attribute [rw] provider_code
#   @return [Array, nil]
#
# @!attribute [rw] region_code
#   @return [Array, nil]
ProviderListMatch = Struct.new(
  :account_number,
  :country_iso,
  :provider_code,
  :region_code,
  keyword_init: true
)

# ProviderStatus entity data model.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
ProviderStatus = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# Request payload for ProviderStatus#list.
#
# @!attribute [rw] provider_code
#   @return [Array, nil]
ProviderStatusListMatch = Struct.new(
  :provider_code,
  keyword_init: true
)

# Region entity data model.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
Region = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# Request payload for Region#list.
#
# @!attribute [rw] country_iso
#   @return [Array, nil]
RegionListMatch = Struct.new(
  :country_iso,
  keyword_init: true
)

# SendTransfer entity data model.
#
# @!attribute [rw] AccountNumber
#   @return [String]
#
# @!attribute [rw] BillRef
#   @return [String, nil]
#
# @!attribute [rw] DistributorRef
#   @return [String]
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
#
# @!attribute [rw] SendCurrencyIso
#   @return [String, nil]
#
# @!attribute [rw] SendValue
#   @return [Float]
#
# @!attribute [rw] Settings
#   @return [Array, nil]
#
# @!attribute [rw] SkuCode
#   @return [String]
#
# @!attribute [rw] TransferRecord
#   @return [Hash]
#
# @!attribute [rw] ValidateOnly
#   @return [Boolean]
SendTransfer = Struct.new(
  :AccountNumber,
  :BillRef,
  :DistributorRef,
  :ErrorCodes,
  :ResultCode,
  :SendCurrencyIso,
  :SendValue,
  :Settings,
  :SkuCode,
  :TransferRecord,
  :ValidateOnly,
  keyword_init: true
)

# Request payload for SendTransfer#create.
#
# @!attribute [rw] AccountNumber
#   @return [String]
#
# @!attribute [rw] BillRef
#   @return [String, nil]
#
# @!attribute [rw] DistributorRef
#   @return [String]
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
#
# @!attribute [rw] SendCurrencyIso
#   @return [String, nil]
#
# @!attribute [rw] SendValue
#   @return [Float]
#
# @!attribute [rw] Settings
#   @return [Array, nil]
#
# @!attribute [rw] SkuCode
#   @return [String]
#
# @!attribute [rw] TransferRecord
#   @return [Hash]
#
# @!attribute [rw] ValidateOnly
#   @return [Boolean]
SendTransferCreateData = Struct.new(
  :AccountNumber,
  :BillRef,
  :DistributorRef,
  :ErrorCodes,
  :ResultCode,
  :SendCurrencyIso,
  :SendValue,
  :Settings,
  :SkuCode,
  :TransferRecord,
  :ValidateOnly,
  keyword_init: true
)

