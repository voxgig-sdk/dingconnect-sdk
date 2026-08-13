-- Typed models for the Dingconnect SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
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
---@field AccountNumberNormalized? string
---@field CountryIso? string
---@field ErrorCodes? table
---@field Items? table
---@field ResultCode? number

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

---@class CancelTransferCreateData
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

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

---@class EstimatePriceCreateData
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class ListTransferRecord
---@field ErrorCodes table
---@field Items table
---@field ResultCode number
---@field ThereAreMoreItems boolean

---@class ListTransferRecordCreateData
---@field ErrorCodes table
---@field Items table
---@field ResultCode number
---@field ThereAreMoreItems boolean

---@class LookupBill
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class LookupBillCreateData
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class Product
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class ProductListMatch
---@field ErrorCodes? table
---@field Items? table
---@field ResultCode? number

---@class ProductDescription
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class ProductDescriptionListMatch
---@field ErrorCodes? table
---@field Items? table
---@field ResultCode? number

---@class Promotion
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class PromotionListMatch
---@field ErrorCodes? table
---@field Items? table
---@field ResultCode? number

---@class PromotionDescription
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class PromotionDescriptionListMatch
---@field ErrorCodes? table
---@field Items? table
---@field ResultCode? number

---@class Provider
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class ProviderListMatch
---@field ErrorCodes? table
---@field Items? table
---@field ResultCode? number

---@class ProviderStatus
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class ProviderStatusListMatch
---@field ErrorCodes? table
---@field Items? table
---@field ResultCode? number

---@class Region
---@field ErrorCodes table
---@field Items table
---@field ResultCode number

---@class RegionListMatch
---@field ErrorCodes? table
---@field Items? table
---@field ResultCode? number

---@class SendTransfer
---@field ErrorCodes table
---@field ResultCode number
---@field TransferRecord table

---@class SendTransferCreateData
---@field ErrorCodes table
---@field ResultCode number
---@field TransferRecord table

local M = {}

return M
