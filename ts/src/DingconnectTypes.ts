// Typed models for the Dingconnect SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface AccountLookup {
  AccountNumberNormalized?: string
  CountryIso?: string
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface AccountLookupListMatch {
  account_number?: number
}

export interface Balance {
  Code: string
  Context?: string
}

export interface BalanceListMatch {
  Code?: string
  Context?: string
}

export interface CancelTransfer {
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface CancelTransferCreateData {
  cancellation_request: any[]
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface Country {
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface CountryListMatch {
  ErrorCodes?: any[]
  Items?: any[]
  ResultCode?: number
}

export interface Currency {
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface CurrencyListMatch {
  ErrorCodes?: any[]
  Items?: any[]
  ResultCode?: number
}

export interface ErrorCodeDescription {
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface ErrorCodeDescriptionListMatch {
  ErrorCodes?: any[]
  Items?: any[]
  ResultCode?: number
}

export interface EstimatePrice {
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface EstimatePriceCreateData {
  requested_estimation: any[]
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface ListTransferRecord {
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
  ThereAreMoreItems: boolean
}

export interface ListTransferRecordCreateData {
  request: Record<string, any>
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
  ThereAreMoreItems: boolean
}

export interface LookupBill {
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface LookupBillCreateData {
  request: Record<string, any>
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface Product {
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface ProductListMatch {
  account_number?: number
  benefit?: any
  country_iso?: any
  provider_code?: any
  region_code?: any
  sku_code?: any
}

export interface ProductDescription {
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface ProductDescriptionListMatch {
  language_code?: any
  sku_code?: any
}

export interface Promotion {
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface PromotionListMatch {
  account_number?: number
  country_iso?: any
  provider_code?: any
}

export interface PromotionDescription {
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface PromotionDescriptionListMatch {
  language_code?: any
}

export interface Provider {
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface ProviderListMatch {
  account_number?: number
  country_iso?: any
  provider_code?: any
  region_code?: any
}

export interface ProviderStatus {
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface ProviderStatusListMatch {
  provider_code?: any
}

export interface Region {
  ErrorCodes: any[]
  Items: any[]
  ResultCode: number
}

export interface RegionListMatch {
  country_iso?: any
}

export interface SendTransfer {
  ErrorCodes: any[]
  ResultCode: number
  TransferRecord: Record<string, any>
}

export interface SendTransferCreateData {
  request: Record<string, any>
  ErrorCodes: any[]
  ResultCode: number
  TransferRecord: Record<string, any>
}

