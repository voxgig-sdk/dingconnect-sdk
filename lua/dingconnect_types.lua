-- Typed models for the Dingconnect SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class AccountLookup
---@field account_number_normalized? string
---@field country_iso? string
---@field error_code? table
---@field item? table
---@field result_code? number

---@class AccountLookupListMatch
---@field account_number_normalized? string
---@field country_iso? string
---@field error_code? table
---@field item? table
---@field result_code? number

---@class Balance
---@field code? string
---@field context? string

---@class BalanceListMatch
---@field code? string
---@field context? string

---@class CancelResult
---@field error_code? table
---@field item? table
---@field result_code? number

---@class CancelResultCreateData
---@field error_code? table
---@field item? table
---@field result_code? number

---@class Country
---@field error_code? table
---@field item? table
---@field result_code? number

---@class CountryListMatch
---@field error_code? table
---@field item? table
---@field result_code? number

---@class Currency
---@field error_code? table
---@field item? table
---@field result_code? number

---@class CurrencyListMatch
---@field error_code? table
---@field item? table
---@field result_code? number

---@class ErrorCodeDescription
---@field error_code? table
---@field item? table
---@field result_code? number

---@class ErrorCodeDescriptionListMatch
---@field error_code? table
---@field item? table
---@field result_code? number

---@class Estimate
---@field error_code? table
---@field item? table
---@field result_code? number

---@class EstimateCreateData
---@field error_code? table
---@field item? table
---@field result_code? number

---@class Product
---@field error_code? table
---@field item? table
---@field result_code? number

---@class ProductListMatch
---@field error_code? table
---@field item? table
---@field result_code? number

---@class ProductDescription
---@field error_code? table
---@field item? table
---@field result_code? number

---@class ProductDescriptionListMatch
---@field error_code? table
---@field item? table
---@field result_code? number

---@class Promotion
---@field error_code? table
---@field item? table
---@field result_code? number

---@class PromotionListMatch
---@field error_code? table
---@field item? table
---@field result_code? number

---@class PromotionDescription
---@field error_code? table
---@field item? table
---@field result_code? number

---@class PromotionDescriptionListMatch
---@field error_code? table
---@field item? table
---@field result_code? number

---@class Provider
---@field error_code? table
---@field item? table
---@field result_code? number

---@class ProviderListMatch
---@field error_code? table
---@field item? table
---@field result_code? number

---@class ProviderStatus
---@field error_code? table
---@field item? table
---@field result_code? number

---@class ProviderStatusListMatch
---@field error_code? table
---@field item? table
---@field result_code? number

---@class Region
---@field error_code? table
---@field item? table
---@field result_code? number

---@class RegionListMatch
---@field error_code? table
---@field item? table
---@field result_code? number

---@class SendTransfer
---@field account_number string
---@field distributor_ref string
---@field error_code? table
---@field result_code? number
---@field send_currency_iso? string
---@field send_value number
---@field setting? table
---@field sku_code string
---@field transfer_record? table
---@field validate_only? boolean

---@class SendTransferCreateData
---@field account_number string
---@field distributor_ref string
---@field error_code? table
---@field result_code? number
---@field send_currency_iso? string
---@field send_value number
---@field setting? table
---@field sku_code string
---@field transfer_record? table
---@field validate_only? boolean

---@class TransferRecord
---@field account_number? string
---@field distributor_ref? table
---@field ended_at_utc? string
---@field error_code? table
---@field item? table
---@field result_code? number
---@field skip number
---@field sku_code? table
---@field started_at_utc? string
---@field take number
---@field there_are_more_item? boolean
---@field transfer_ref? table

---@class TransferRecordCreateData
---@field account_number? string
---@field distributor_ref? table
---@field ended_at_utc? string
---@field error_code? table
---@field item? table
---@field result_code? number
---@field skip number
---@field sku_code? table
---@field started_at_utc? string
---@field take number
---@field there_are_more_item? boolean
---@field transfer_ref? table

local M = {}

return M
