export interface AccountLookup {
    AccountNumberNormalized?: string;
    CountryIso?: string;
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface AccountLookupListMatch {
    account_number?: string;
}
export interface Balance {
    Code: string;
    Context?: string;
}
export interface BalanceListMatch {
    Code?: string;
    Context?: string;
}
export interface CancelTransfer {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
    cancellations?: any[];
}
export interface CancelTransferCreateData {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
    cancellations?: any[];
}
export interface Country {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface CountryListMatch {
    ErrorCodes?: any[];
    Items?: any[];
    ResultCode?: number;
}
export interface Currency {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface CurrencyListMatch {
    ErrorCodes?: any[];
    Items?: any[];
    ResultCode?: number;
}
export interface ErrorCodeDescription {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface ErrorCodeDescriptionListMatch {
    ErrorCodes?: any[];
    Items?: any[];
    ResultCode?: number;
}
export interface EstimatePrice {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
    estimations?: any[];
}
export interface EstimatePriceCreateData {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
    estimations?: any[];
}
export interface ListTransferRecord {
    AccountNumber?: string;
    DistributorRef?: string;
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
    Skip?: number;
    Take: number;
    ThereAreMoreItems: boolean;
    TransferRef?: string;
}
export interface ListTransferRecordCreateData {
    AccountNumber?: string;
    DistributorRef?: string;
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
    Skip?: number;
    Take: number;
    ThereAreMoreItems: boolean;
    TransferRef?: string;
}
export interface LookupBill {
    AccountNumber: string;
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
    Settings?: any[];
    SkuCode: string;
}
export interface LookupBillCreateData {
    AccountNumber: string;
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
    Settings?: any[];
    SkuCode: string;
}
export interface Product {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface ProductListMatch {
    account_number?: string;
    benefit?: any[];
    country_iso?: any[];
    provider_code?: any[];
    region_code?: any[];
    sku_code?: any[];
}
export interface ProductDescription {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface ProductDescriptionListMatch {
    language_code?: any[];
    sku_code?: any[];
}
export interface Promotion {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface PromotionListMatch {
    account_number?: string;
    country_iso?: any[];
    provider_code?: any[];
}
export interface PromotionDescription {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface PromotionDescriptionListMatch {
    language_code?: any[];
}
export interface Provider {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface ProviderListMatch {
    account_number?: string;
    country_iso?: any[];
    provider_code?: any[];
    region_code?: any[];
}
export interface ProviderStatus {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface ProviderStatusListMatch {
    provider_code?: any[];
}
export interface Region {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface RegionListMatch {
    country_iso?: any[];
}
export interface SendTransfer {
    AccountNumber: string;
    BillRef?: string;
    DistributorRef: string;
    ErrorCodes: any[];
    ResultCode: number;
    SendCurrencyIso?: string;
    SendValue: number;
    Settings?: any[];
    SkuCode: string;
    TransferRecord: Record<string, any>;
    ValidateOnly: boolean;
}
export interface SendTransferCreateData {
    AccountNumber: string;
    BillRef?: string;
    DistributorRef: string;
    ErrorCodes: any[];
    ResultCode: number;
    SendCurrencyIso?: string;
    SendValue: number;
    Settings?: any[];
    SkuCode: string;
    TransferRecord: Record<string, any>;
    ValidateOnly: boolean;
}
