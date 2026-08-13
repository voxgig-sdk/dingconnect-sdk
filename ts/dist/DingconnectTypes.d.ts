export interface AccountLookup {
    AccountNumberNormalized?: string;
    CountryIso?: string;
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface AccountLookupListMatch {
    AccountNumberNormalized?: string;
    CountryIso?: string;
    ErrorCodes?: any[];
    Items?: any[];
    ResultCode?: number;
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
}
export interface CancelTransferCreateData {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
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
}
export interface EstimatePriceCreateData {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface ListTransferRecord {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
    ThereAreMoreItems: boolean;
}
export interface ListTransferRecordCreateData {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
    ThereAreMoreItems: boolean;
}
export interface LookupBill {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface LookupBillCreateData {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface Product {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface ProductListMatch {
    ErrorCodes?: any[];
    Items?: any[];
    ResultCode?: number;
}
export interface ProductDescription {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface ProductDescriptionListMatch {
    ErrorCodes?: any[];
    Items?: any[];
    ResultCode?: number;
}
export interface Promotion {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface PromotionListMatch {
    ErrorCodes?: any[];
    Items?: any[];
    ResultCode?: number;
}
export interface PromotionDescription {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface PromotionDescriptionListMatch {
    ErrorCodes?: any[];
    Items?: any[];
    ResultCode?: number;
}
export interface Provider {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface ProviderListMatch {
    ErrorCodes?: any[];
    Items?: any[];
    ResultCode?: number;
}
export interface ProviderStatus {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface ProviderStatusListMatch {
    ErrorCodes?: any[];
    Items?: any[];
    ResultCode?: number;
}
export interface Region {
    ErrorCodes: any[];
    Items: any[];
    ResultCode: number;
}
export interface RegionListMatch {
    ErrorCodes?: any[];
    Items?: any[];
    ResultCode?: number;
}
export interface SendTransfer {
    ErrorCodes: any[];
    ResultCode: number;
    TransferRecord: Record<string, any>;
}
export interface SendTransferCreateData {
    ErrorCodes: any[];
    ResultCode: number;
    TransferRecord: Record<string, any>;
}
