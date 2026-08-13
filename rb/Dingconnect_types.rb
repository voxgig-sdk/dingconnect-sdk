# frozen_string_literal: true

# Typed models for the Dingconnect SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
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
# @!attribute [rw] AccountNumberNormalized
#   @return [String, nil]
#
# @!attribute [rw] CountryIso
#   @return [String, nil]
#
# @!attribute [rw] ErrorCodes
#   @return [Array, nil]
#
# @!attribute [rw] Items
#   @return [Array, nil]
#
# @!attribute [rw] ResultCode
#   @return [Integer, nil]
AccountLookupListMatch = Struct.new(
  :AccountNumberNormalized,
  :CountryIso,
  :ErrorCodes,
  :Items,
  :ResultCode,
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
CancelTransfer = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
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
CancelTransferCreateData = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
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
EstimatePrice = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
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
EstimatePriceCreateData = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# ListTransferRecord entity data model.
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
# @!attribute [rw] ThereAreMoreItems
#   @return [Boolean]
ListTransferRecord = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  :ThereAreMoreItems,
  keyword_init: true
)

# Request payload for ListTransferRecord#create.
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
# @!attribute [rw] ThereAreMoreItems
#   @return [Boolean]
ListTransferRecordCreateData = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  :ThereAreMoreItems,
  keyword_init: true
)

# LookupBill entity data model.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
LookupBill = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# Request payload for LookupBill#create.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] Items
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
LookupBillCreateData = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
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
# @!attribute [rw] ErrorCodes
#   @return [Array, nil]
#
# @!attribute [rw] Items
#   @return [Array, nil]
#
# @!attribute [rw] ResultCode
#   @return [Integer, nil]
ProductListMatch = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
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
# @!attribute [rw] ErrorCodes
#   @return [Array, nil]
#
# @!attribute [rw] Items
#   @return [Array, nil]
#
# @!attribute [rw] ResultCode
#   @return [Integer, nil]
ProductDescriptionListMatch = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
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
# @!attribute [rw] ErrorCodes
#   @return [Array, nil]
#
# @!attribute [rw] Items
#   @return [Array, nil]
#
# @!attribute [rw] ResultCode
#   @return [Integer, nil]
PromotionListMatch = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
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
# @!attribute [rw] ErrorCodes
#   @return [Array, nil]
#
# @!attribute [rw] Items
#   @return [Array, nil]
#
# @!attribute [rw] ResultCode
#   @return [Integer, nil]
PromotionDescriptionListMatch = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
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
# @!attribute [rw] ErrorCodes
#   @return [Array, nil]
#
# @!attribute [rw] Items
#   @return [Array, nil]
#
# @!attribute [rw] ResultCode
#   @return [Integer, nil]
ProviderListMatch = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
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
# @!attribute [rw] ErrorCodes
#   @return [Array, nil]
#
# @!attribute [rw] Items
#   @return [Array, nil]
#
# @!attribute [rw] ResultCode
#   @return [Integer, nil]
ProviderStatusListMatch = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
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
# @!attribute [rw] ErrorCodes
#   @return [Array, nil]
#
# @!attribute [rw] Items
#   @return [Array, nil]
#
# @!attribute [rw] ResultCode
#   @return [Integer, nil]
RegionListMatch = Struct.new(
  :ErrorCodes,
  :Items,
  :ResultCode,
  keyword_init: true
)

# SendTransfer entity data model.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
#
# @!attribute [rw] TransferRecord
#   @return [Hash]
SendTransfer = Struct.new(
  :ErrorCodes,
  :ResultCode,
  :TransferRecord,
  keyword_init: true
)

# Request payload for SendTransfer#create.
#
# @!attribute [rw] ErrorCodes
#   @return [Array]
#
# @!attribute [rw] ResultCode
#   @return [Integer]
#
# @!attribute [rw] TransferRecord
#   @return [Hash]
SendTransferCreateData = Struct.new(
  :ErrorCodes,
  :ResultCode,
  :TransferRecord,
  keyword_init: true
)

