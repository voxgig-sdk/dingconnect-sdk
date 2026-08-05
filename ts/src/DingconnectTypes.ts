// Typed models for the Dingconnect SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface AccountLookup {
  account_number_normalized?: string
  country_iso?: string
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface AccountLookupListMatch {
  account_number_normalized?: string
  country_iso?: string
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface Balance {
  code?: string
  context?: string
}

export interface BalanceListMatch {
  code?: string
  context?: string
}

export interface CancelResult {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface CancelResultCreateData {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface Country {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface CountryListMatch {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface Currency {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface CurrencyListMatch {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface ErrorCodeDescription {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface ErrorCodeDescriptionListMatch {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface Estimate {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface EstimateCreateData {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface Product {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface ProductListMatch {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface ProductDescription {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface ProductDescriptionListMatch {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface Promotion {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface PromotionListMatch {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface PromotionDescription {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface PromotionDescriptionListMatch {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface Provider {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface ProviderListMatch {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface ProviderStatus {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface ProviderStatusListMatch {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface Region {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface RegionListMatch {
  error_code?: any[]
  item?: any[]
  result_code?: number
}

export interface SendTransfer {
  account_number: string
  distributor_ref: string
  error_code?: any[]
  result_code?: number
  send_currency_iso?: string
  send_value: number
  setting?: any[]
  sku_code: string
  transfer_record?: Record<string, any>
  validate_only?: boolean
}

export interface SendTransferCreateData {
  account_number: string
  distributor_ref: string
  error_code?: any[]
  result_code?: number
  send_currency_iso?: string
  send_value: number
  setting?: any[]
  sku_code: string
  transfer_record?: Record<string, any>
  validate_only?: boolean
}

export interface TransferRecord {
  account_number?: string
  distributor_ref?: any[]
  ended_at_utc?: string
  error_code?: any[]
  item?: any[]
  result_code?: number
  skip: number
  sku_code?: any[]
  started_at_utc?: string
  take: number
  there_are_more_item?: boolean
  transfer_ref?: any[]
}

export interface TransferRecordCreateData {
  account_number?: string
  distributor_ref?: any[]
  ended_at_utc?: string
  error_code?: any[]
  item?: any[]
  result_code?: number
  skip: number
  sku_code?: any[]
  started_at_utc?: string
  take: number
  there_are_more_item?: boolean
  transfer_ref?: any[]
}

