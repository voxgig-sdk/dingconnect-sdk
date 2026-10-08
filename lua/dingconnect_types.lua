-- Typed models for the Dingconnect SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
-- params (op.<name>.points[].g.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class AccountLookup
---@field AccountNumberNormalized? string
---@field CountryIso? string
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class AccountLookupListMatch
---@field account_number? string

---@class Balance
---@field Code string
---@field Context? string

---@class BalanceListMatch
---@field Code? string
---@field Context? string

---@class CancelTransfer
---@field ErrorCodes table
---@field Items table
---@field ResultCode number
---@field cancellations? table

---@class CancelTransferCreateData
---@field ErrorCodes table
---@field Items table
---@field ResultCode number
---@field cancellations? table

---@class Country
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class CountryListMatch
---@field ErrorCodes? table
---@field Items? table
---@field ResultCode? number

---@class Currency
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class CurrencyListMatch
---@field ErrorCodes? table
---@field Items? table
---@field ResultCode? number

---@class ErrorCodeDescription
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class ErrorCodeDescriptionListMatch
---@field ErrorCodes? table
---@field Items? table
---@field ResultCode? number

---@class EstimatePrice
---@field ErrorCodes table
---@field Items table
---@field ResultCode number
---@field estimations? table

---@class EstimatePriceCreateData
---@field ErrorCodes table
---@field Items table
---@field ResultCode number
---@field estimations? table

---@class ListTransferRecord
---@field AccountNumber? string
---@field DistributorRef? string
---@field ErrorCodes table
---@field Items table
---@field ResultCode number
---@field Skip? number
---@field Take number
---@field ThereAreMoreItems boolean
---@field TransferRef? string

---@class ListTransferRecordCreateData
---@field AccountNumber? string
---@field DistributorRef? string
---@field ErrorCodes table
---@field Items table
---@field ResultCode number
---@field Skip? number
---@field Take number
---@field ThereAreMoreItems boolean
---@field TransferRef? string

---@class LookupBill
---@field AccountNumber string
---@field ErrorCodes table
---@field Items table
---@field ResultCode number
---@field Settings? table
---@field SkuCode string

---@class LookupBillCreateData
---@field AccountNumber string
---@field ErrorCodes table
---@field Items table
---@field ResultCode number
---@field Settings? table
---@field SkuCode string

---@class Product
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class ProductListMatch
---@field account_number? string
---@field benefit? table
---@field country_iso? table
---@field provider_code? table
---@field region_code? table
---@field sku_code? table

---@class ProductDescription
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class ProductDescriptionListMatch
---@field language_code? table
---@field sku_code? table

---@class Promotion
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class PromotionListMatch
---@field account_number? string
---@field country_iso? table
---@field provider_code? table

---@class PromotionDescription
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class PromotionDescriptionListMatch
---@field language_code? table

---@class Provider
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class ProviderListMatch
---@field account_number? string
---@field country_iso? table
---@field provider_code? table
---@field region_code? table

---@class ProviderStatus
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class ProviderStatusListMatch
---@field provider_code? table

---@class Region
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class RegionListMatch
---@field country_iso? table

---@class SendTransfer
---@field AccountNumber string
---@field BillRef? string
---@field DistributorRef string
---@field ErrorCodes table
---@field ResultCode number
---@field SendCurrencyIso? string
---@field SendValue number
---@field Settings? table
---@field SkuCode string
---@field TransferRecord table
---@field ValidateOnly boolean

---@class SendTransferCreateData
---@field AccountNumber string
---@field BillRef? string
---@field DistributorRef string
---@field ErrorCodes table
---@field ResultCode number
---@field SendCurrencyIso? string
---@field SendValue number
---@field Settings? table
---@field SkuCode string
---@field TransferRecord table
---@field ValidateOnly boolean

local M = {}

return M
