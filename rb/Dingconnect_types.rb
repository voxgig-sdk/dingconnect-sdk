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
# @!attribute [rw] account_number_normalized
#   @return [String, nil]
#
# @!attribute [rw] country_iso
#   @return [String, nil]
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
AccountLookup = Struct.new(
  :account_number_normalized,
  :country_iso,
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Request payload for AccountLookup#list.
#
# @!attribute [rw] account_number_normalized
#   @return [String, nil]
#
# @!attribute [rw] country_iso
#   @return [String, nil]
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
AccountLookupListMatch = Struct.new(
  :account_number_normalized,
  :country_iso,
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Balance entity data model.
#
# @!attribute [rw] code
#   @return [String, nil]
#
# @!attribute [rw] context
#   @return [String, nil]
Balance = Struct.new(
  :code,
  :context,
  keyword_init: true
)

# Request payload for Balance#list.
#
# @!attribute [rw] code
#   @return [String, nil]
#
# @!attribute [rw] context
#   @return [String, nil]
BalanceListMatch = Struct.new(
  :code,
  :context,
  keyword_init: true
)

# CancelResult entity data model.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
CancelResult = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Request payload for CancelResult#create.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
CancelResultCreateData = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Country entity data model.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
Country = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Request payload for Country#list.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
CountryListMatch = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Currency entity data model.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
Currency = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Request payload for Currency#list.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
CurrencyListMatch = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# ErrorCodeDescription entity data model.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
ErrorCodeDescription = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Request payload for ErrorCodeDescription#list.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
ErrorCodeDescriptionListMatch = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Estimate entity data model.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
Estimate = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Request payload for Estimate#create.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
EstimateCreateData = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Product entity data model.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
Product = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Request payload for Product#list.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
ProductListMatch = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# ProductDescription entity data model.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
ProductDescription = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Request payload for ProductDescription#list.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
ProductDescriptionListMatch = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Promotion entity data model.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
Promotion = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Request payload for Promotion#list.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
PromotionListMatch = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# PromotionDescription entity data model.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
PromotionDescription = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Request payload for PromotionDescription#list.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
PromotionDescriptionListMatch = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Provider entity data model.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
Provider = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Request payload for Provider#list.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
ProviderListMatch = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# ProviderStatus entity data model.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
ProviderStatus = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Request payload for ProviderStatus#list.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
ProviderStatusListMatch = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Region entity data model.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
Region = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# Request payload for Region#list.
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
RegionListMatch = Struct.new(
  :error_code,
  :item,
  :result_code,
  keyword_init: true
)

# SendTransfer entity data model.
#
# @!attribute [rw] account_number
#   @return [String]
#
# @!attribute [rw] distributor_ref
#   @return [String]
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
#
# @!attribute [rw] send_currency_iso
#   @return [String, nil]
#
# @!attribute [rw] send_value
#   @return [Float]
#
# @!attribute [rw] setting
#   @return [Array, nil]
#
# @!attribute [rw] sku_code
#   @return [String]
#
# @!attribute [rw] transfer_record
#   @return [Hash, nil]
#
# @!attribute [rw] validate_only
#   @return [Boolean, nil]
SendTransfer = Struct.new(
  :account_number,
  :distributor_ref,
  :error_code,
  :result_code,
  :send_currency_iso,
  :send_value,
  :setting,
  :sku_code,
  :transfer_record,
  :validate_only,
  keyword_init: true
)

# Request payload for SendTransfer#create.
#
# @!attribute [rw] account_number
#   @return [String]
#
# @!attribute [rw] distributor_ref
#   @return [String]
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
#
# @!attribute [rw] send_currency_iso
#   @return [String, nil]
#
# @!attribute [rw] send_value
#   @return [Float]
#
# @!attribute [rw] setting
#   @return [Array, nil]
#
# @!attribute [rw] sku_code
#   @return [String]
#
# @!attribute [rw] transfer_record
#   @return [Hash, nil]
#
# @!attribute [rw] validate_only
#   @return [Boolean, nil]
SendTransferCreateData = Struct.new(
  :account_number,
  :distributor_ref,
  :error_code,
  :result_code,
  :send_currency_iso,
  :send_value,
  :setting,
  :sku_code,
  :transfer_record,
  :validate_only,
  keyword_init: true
)

# TransferRecord entity data model.
#
# @!attribute [rw] account_number
#   @return [String, nil]
#
# @!attribute [rw] distributor_ref
#   @return [Array, nil]
#
# @!attribute [rw] ended_at_utc
#   @return [String, nil]
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
#
# @!attribute [rw] skip
#   @return [Integer]
#
# @!attribute [rw] sku_code
#   @return [Array, nil]
#
# @!attribute [rw] started_at_utc
#   @return [String, nil]
#
# @!attribute [rw] take
#   @return [Integer]
#
# @!attribute [rw] there_are_more_item
#   @return [Boolean, nil]
#
# @!attribute [rw] transfer_ref
#   @return [Array, nil]
TransferRecord = Struct.new(
  :account_number,
  :distributor_ref,
  :ended_at_utc,
  :error_code,
  :item,
  :result_code,
  :skip,
  :sku_code,
  :started_at_utc,
  :take,
  :there_are_more_item,
  :transfer_ref,
  keyword_init: true
)

# Request payload for TransferRecord#create.
#
# @!attribute [rw] account_number
#   @return [String, nil]
#
# @!attribute [rw] distributor_ref
#   @return [Array, nil]
#
# @!attribute [rw] ended_at_utc
#   @return [String, nil]
#
# @!attribute [rw] error_code
#   @return [Array, nil]
#
# @!attribute [rw] item
#   @return [Array, nil]
#
# @!attribute [rw] result_code
#   @return [Integer, nil]
#
# @!attribute [rw] skip
#   @return [Integer]
#
# @!attribute [rw] sku_code
#   @return [Array, nil]
#
# @!attribute [rw] started_at_utc
#   @return [String, nil]
#
# @!attribute [rw] take
#   @return [Integer]
#
# @!attribute [rw] there_are_more_item
#   @return [Boolean, nil]
#
# @!attribute [rw] transfer_ref
#   @return [Array, nil]
TransferRecordCreateData = Struct.new(
  :account_number,
  :distributor_ref,
  :ended_at_utc,
  :error_code,
  :item,
  :result_code,
  :skip,
  :sku_code,
  :started_at_utc,
  :take,
  :there_are_more_item,
  :transfer_ref,
  keyword_init: true
)

