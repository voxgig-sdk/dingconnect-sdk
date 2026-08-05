// Typed models for the Dingconnect SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// AccountLookup is the typed data model for the account_lookup entity.
type AccountLookup struct {
	AccountNumberNormalized *string `json:"account_number_normalized,omitempty"`
	CountryIso *string `json:"country_iso,omitempty"`
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// AccountLookupListMatch is the typed request payload for AccountLookup.ListTyped.
type AccountLookupListMatch struct {
	AccountNumberNormalized *string `json:"account_number_normalized,omitempty"`
	CountryIso *string `json:"country_iso,omitempty"`
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// Balance is the typed data model for the balance entity.
type Balance struct {
	Code *string `json:"code,omitempty"`
	Context *string `json:"context,omitempty"`
}

// BalanceListMatch is the typed request payload for Balance.ListTyped.
type BalanceListMatch struct {
	Code *string `json:"code,omitempty"`
	Context *string `json:"context,omitempty"`
}

// CancelResult is the typed data model for the cancel_result entity.
type CancelResult struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// CancelResultCreateData is the typed request payload for CancelResult.CreateTyped.
type CancelResultCreateData struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// Country is the typed data model for the country entity.
type Country struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// CountryListMatch is the typed request payload for Country.ListTyped.
type CountryListMatch struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// Currency is the typed data model for the currency entity.
type Currency struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// CurrencyListMatch is the typed request payload for Currency.ListTyped.
type CurrencyListMatch struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// ErrorCodeDescription is the typed data model for the error_code_description entity.
type ErrorCodeDescription struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// ErrorCodeDescriptionListMatch is the typed request payload for ErrorCodeDescription.ListTyped.
type ErrorCodeDescriptionListMatch struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// Estimate is the typed data model for the estimate entity.
type Estimate struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// EstimateCreateData is the typed request payload for Estimate.CreateTyped.
type EstimateCreateData struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// Product is the typed data model for the product entity.
type Product struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// ProductListMatch is the typed request payload for Product.ListTyped.
type ProductListMatch struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// ProductDescription is the typed data model for the product_description entity.
type ProductDescription struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// ProductDescriptionListMatch is the typed request payload for ProductDescription.ListTyped.
type ProductDescriptionListMatch struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// Promotion is the typed data model for the promotion entity.
type Promotion struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// PromotionListMatch is the typed request payload for Promotion.ListTyped.
type PromotionListMatch struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// PromotionDescription is the typed data model for the promotion_description entity.
type PromotionDescription struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// PromotionDescriptionListMatch is the typed request payload for PromotionDescription.ListTyped.
type PromotionDescriptionListMatch struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// Provider is the typed data model for the provider entity.
type Provider struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// ProviderListMatch is the typed request payload for Provider.ListTyped.
type ProviderListMatch struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// ProviderStatus is the typed data model for the provider_status entity.
type ProviderStatus struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// ProviderStatusListMatch is the typed request payload for ProviderStatus.ListTyped.
type ProviderStatusListMatch struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// Region is the typed data model for the region entity.
type Region struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// RegionListMatch is the typed request payload for Region.ListTyped.
type RegionListMatch struct {
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
}

// SendTransfer is the typed data model for the send_transfer entity.
type SendTransfer struct {
	AccountNumber string `json:"account_number"`
	DistributorRef string `json:"distributor_ref"`
	ErrorCode *[]any `json:"error_code,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
	SendCurrencyIso *string `json:"send_currency_iso,omitempty"`
	SendValue float64 `json:"send_value"`
	Setting *[]any `json:"setting,omitempty"`
	SkuCode string `json:"sku_code"`
	TransferRecord *map[string]any `json:"transfer_record,omitempty"`
	ValidateOnly *bool `json:"validate_only,omitempty"`
}

// SendTransferCreateData is the typed request payload for SendTransfer.CreateTyped.
type SendTransferCreateData struct {
	AccountNumber string `json:"account_number"`
	DistributorRef string `json:"distributor_ref"`
	ErrorCode *[]any `json:"error_code,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
	SendCurrencyIso *string `json:"send_currency_iso,omitempty"`
	SendValue float64 `json:"send_value"`
	Setting *[]any `json:"setting,omitempty"`
	SkuCode string `json:"sku_code"`
	TransferRecord *map[string]any `json:"transfer_record,omitempty"`
	ValidateOnly *bool `json:"validate_only,omitempty"`
}

// TransferRecord is the typed data model for the transfer_record entity.
type TransferRecord struct {
	AccountNumber *string `json:"account_number,omitempty"`
	DistributorRef *[]any `json:"distributor_ref,omitempty"`
	EndedAtUtc *string `json:"ended_at_utc,omitempty"`
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
	Skip int `json:"skip"`
	SkuCode *[]any `json:"sku_code,omitempty"`
	StartedAtUtc *string `json:"started_at_utc,omitempty"`
	Take int `json:"take"`
	ThereAreMoreItem *bool `json:"there_are_more_item,omitempty"`
	TransferRef *[]any `json:"transfer_ref,omitempty"`
}

// TransferRecordCreateData is the typed request payload for TransferRecord.CreateTyped.
type TransferRecordCreateData struct {
	AccountNumber *string `json:"account_number,omitempty"`
	DistributorRef *[]any `json:"distributor_ref,omitempty"`
	EndedAtUtc *string `json:"ended_at_utc,omitempty"`
	ErrorCode *[]any `json:"error_code,omitempty"`
	Item *[]any `json:"item,omitempty"`
	ResultCode *int `json:"result_code,omitempty"`
	Skip int `json:"skip"`
	SkuCode *[]any `json:"sku_code,omitempty"`
	StartedAtUtc *string `json:"started_at_utc,omitempty"`
	Take int `json:"take"`
	ThereAreMoreItem *bool `json:"there_are_more_item,omitempty"`
	TransferRef *[]any `json:"transfer_ref,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
