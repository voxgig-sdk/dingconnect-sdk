// Typed models for the Dingconnect SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/dingconnect-sdk/go/core"
)

// AccountLookup is the typed data model for the account_lookup entity.
type AccountLookup struct {
}

// AccountLookupListMatch is the typed request payload for AccountLookup.ListTyped.
type AccountLookupListMatch struct {
	AccountNumber *string `json:"account_number,omitempty"`
}

// Balance is the typed data model for the balance entity.
type Balance struct {
}

// BalanceListMatch is the typed request payload for Balance.ListTyped.
type BalanceListMatch struct {
	Code *string `json:"Code,omitempty"`
	Context *string `json:"Context,omitempty"`
}

// CancelTransfer is the typed data model for the cancel_transfer entity.
type CancelTransfer struct {
}

// CancelTransferCreateData is the typed request payload for CancelTransfer.CreateTyped.
type CancelTransferCreateData struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
	Cancellations *[]any `json:"cancellations,omitempty"`
}

// Country is the typed data model for the country entity.
type Country struct {
}

// CountryListMatch is the typed request payload for Country.ListTyped.
type CountryListMatch struct {
	ErrorCodes *[]any `json:"ErrorCodes,omitempty"`
	Items *[]any `json:"Items,omitempty"`
	ResultCode *int `json:"ResultCode,omitempty"`
}

// Currency is the typed data model for the currency entity.
type Currency struct {
}

// CurrencyListMatch is the typed request payload for Currency.ListTyped.
type CurrencyListMatch struct {
	ErrorCodes *[]any `json:"ErrorCodes,omitempty"`
	Items *[]any `json:"Items,omitempty"`
	ResultCode *int `json:"ResultCode,omitempty"`
}

// ErrorCodeDescription is the typed data model for the error_code_description entity.
type ErrorCodeDescription struct {
}

// ErrorCodeDescriptionListMatch is the typed request payload for ErrorCodeDescription.ListTyped.
type ErrorCodeDescriptionListMatch struct {
	ErrorCodes *[]any `json:"ErrorCodes,omitempty"`
	Items *[]any `json:"Items,omitempty"`
	ResultCode *int `json:"ResultCode,omitempty"`
}

// EstimatePrice is the typed data model for the estimate_price entity.
type EstimatePrice struct {
}

// EstimatePriceCreateData is the typed request payload for EstimatePrice.CreateTyped.
type EstimatePriceCreateData struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
	Estimations *[]any `json:"estimations,omitempty"`
}

// ListTransferRecord is the typed data model for the list_transfer_record entity.
type ListTransferRecord struct {
}

// ListTransferRecordCreateData is the typed request payload for ListTransferRecord.CreateTyped.
type ListTransferRecordCreateData struct {
	AccountNumber *string `json:"AccountNumber,omitempty"`
	DistributorRef *string `json:"DistributorRef,omitempty"`
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
	Skip *int `json:"Skip,omitempty"`
	Take int `json:"Take"`
	ThereAreMoreItems bool `json:"ThereAreMoreItems"`
	TransferRef *string `json:"TransferRef,omitempty"`
}

// LookupBill is the typed data model for the lookup_bill entity.
type LookupBill struct {
}

// LookupBillCreateData is the typed request payload for LookupBill.CreateTyped.
type LookupBillCreateData struct {
	AccountNumber string `json:"AccountNumber"`
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
	Settings *[]any `json:"Settings,omitempty"`
	SkuCode string `json:"SkuCode"`
}

// Product is the typed data model for the product entity.
type Product struct {
}

// ProductListMatch is the typed request payload for Product.ListTyped.
type ProductListMatch struct {
	AccountNumber *string `json:"account_number,omitempty"`
	Benefit *[]any `json:"benefit,omitempty"`
	CountryIso *[]any `json:"country_iso,omitempty"`
	ProviderCode *[]any `json:"provider_code,omitempty"`
	RegionCode *[]any `json:"region_code,omitempty"`
	SkuCode *[]any `json:"sku_code,omitempty"`
}

// ProductDescription is the typed data model for the product_description entity.
type ProductDescription struct {
}

// ProductDescriptionListMatch is the typed request payload for ProductDescription.ListTyped.
type ProductDescriptionListMatch struct {
	LanguageCode *[]any `json:"language_code,omitempty"`
	SkuCode *[]any `json:"sku_code,omitempty"`
}

// Promotion is the typed data model for the promotion entity.
type Promotion struct {
}

// PromotionListMatch is the typed request payload for Promotion.ListTyped.
type PromotionListMatch struct {
	AccountNumber *string `json:"account_number,omitempty"`
	CountryIso *[]any `json:"country_iso,omitempty"`
	ProviderCode *[]any `json:"provider_code,omitempty"`
}

// PromotionDescription is the typed data model for the promotion_description entity.
type PromotionDescription struct {
}

// PromotionDescriptionListMatch is the typed request payload for PromotionDescription.ListTyped.
type PromotionDescriptionListMatch struct {
	LanguageCode *[]any `json:"language_code,omitempty"`
}

// Provider is the typed data model for the provider entity.
type Provider struct {
}

// ProviderListMatch is the typed request payload for Provider.ListTyped.
type ProviderListMatch struct {
	AccountNumber *string `json:"account_number,omitempty"`
	CountryIso *[]any `json:"country_iso,omitempty"`
	ProviderCode *[]any `json:"provider_code,omitempty"`
	RegionCode *[]any `json:"region_code,omitempty"`
}

// ProviderStatus is the typed data model for the provider_status entity.
type ProviderStatus struct {
}

// ProviderStatusListMatch is the typed request payload for ProviderStatus.ListTyped.
type ProviderStatusListMatch struct {
	ProviderCode *[]any `json:"provider_code,omitempty"`
}

// Region is the typed data model for the region entity.
type Region struct {
}

// RegionListMatch is the typed request payload for Region.ListTyped.
type RegionListMatch struct {
	CountryIso *[]any `json:"country_iso,omitempty"`
}

// SendTransfer is the typed data model for the send_transfer entity.
type SendTransfer struct {
}

// SendTransferCreateData is the typed request payload for SendTransfer.CreateTyped.
type SendTransferCreateData struct {
	AccountNumber string `json:"AccountNumber"`
	BillRef *string `json:"BillRef,omitempty"`
	DistributorRef string `json:"DistributorRef"`
	ErrorCodes []any `json:"ErrorCodes"`
	ResultCode int `json:"ResultCode"`
	SendCurrencyIso *string `json:"SendCurrencyIso,omitempty"`
	SendValue float64 `json:"SendValue"`
	Settings *[]any `json:"Settings,omitempty"`
	SkuCode string `json:"SkuCode"`
	TransferRecord map[string]any `json:"TransferRecord"`
	ValidateOnly bool `json:"ValidateOnly"`
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

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
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

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
