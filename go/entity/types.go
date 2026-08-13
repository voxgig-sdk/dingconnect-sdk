// Typed models for the Dingconnect SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/dingconnect-sdk/go/core"
)

// AccountLookup is the typed data model for the account_lookup entity.
type AccountLookup struct {
	AccountNumberNormalized *string `json:"AccountNumberNormalized,omitempty"`
	CountryIso *string `json:"CountryIso,omitempty"`
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// AccountLookupListMatch is the typed request payload for AccountLookup.ListTyped.
type AccountLookupListMatch struct {
	AccountNumberNormalized *string `json:"AccountNumberNormalized,omitempty"`
	CountryIso *string `json:"CountryIso,omitempty"`
	ErrorCodes *[]any `json:"ErrorCodes,omitempty"`
	Items *[]any `json:"Items,omitempty"`
	ResultCode *int `json:"ResultCode,omitempty"`
}

// Balance is the typed data model for the balance entity.
type Balance struct {
	Code string `json:"Code"`
	Context *string `json:"Context,omitempty"`
}

// BalanceListMatch is the typed request payload for Balance.ListTyped.
type BalanceListMatch struct {
	Code *string `json:"Code,omitempty"`
	Context *string `json:"Context,omitempty"`
}

// CancelTransfer is the typed data model for the cancel_transfer entity.
type CancelTransfer struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// CancelTransferCreateData is the typed request payload for CancelTransfer.CreateTyped.
type CancelTransferCreateData struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// Country is the typed data model for the country entity.
type Country struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// CountryListMatch is the typed request payload for Country.ListTyped.
type CountryListMatch struct {
	ErrorCodes *[]any `json:"ErrorCodes,omitempty"`
	Items *[]any `json:"Items,omitempty"`
	ResultCode *int `json:"ResultCode,omitempty"`
}

// Currency is the typed data model for the currency entity.
type Currency struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// CurrencyListMatch is the typed request payload for Currency.ListTyped.
type CurrencyListMatch struct {
	ErrorCodes *[]any `json:"ErrorCodes,omitempty"`
	Items *[]any `json:"Items,omitempty"`
	ResultCode *int `json:"ResultCode,omitempty"`
}

// ErrorCodeDescription is the typed data model for the error_code_description entity.
type ErrorCodeDescription struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// ErrorCodeDescriptionListMatch is the typed request payload for ErrorCodeDescription.ListTyped.
type ErrorCodeDescriptionListMatch struct {
	ErrorCodes *[]any `json:"ErrorCodes,omitempty"`
	Items *[]any `json:"Items,omitempty"`
	ResultCode *int `json:"ResultCode,omitempty"`
}

// EstimatePrice is the typed data model for the estimate_price entity.
type EstimatePrice struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// EstimatePriceCreateData is the typed request payload for EstimatePrice.CreateTyped.
type EstimatePriceCreateData struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// ListTransferRecord is the typed data model for the list_transfer_record entity.
type ListTransferRecord struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
	ThereAreMoreItems bool `json:"ThereAreMoreItems"`
}

// ListTransferRecordCreateData is the typed request payload for ListTransferRecord.CreateTyped.
type ListTransferRecordCreateData struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
	ThereAreMoreItems bool `json:"ThereAreMoreItems"`
}

// LookupBill is the typed data model for the lookup_bill entity.
type LookupBill struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// LookupBillCreateData is the typed request payload for LookupBill.CreateTyped.
type LookupBillCreateData struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// Product is the typed data model for the product entity.
type Product struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// ProductListMatch is the typed request payload for Product.ListTyped.
type ProductListMatch struct {
	ErrorCodes *[]any `json:"ErrorCodes,omitempty"`
	Items *[]any `json:"Items,omitempty"`
	ResultCode *int `json:"ResultCode,omitempty"`
}

// ProductDescription is the typed data model for the product_description entity.
type ProductDescription struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// ProductDescriptionListMatch is the typed request payload for ProductDescription.ListTyped.
type ProductDescriptionListMatch struct {
	ErrorCodes *[]any `json:"ErrorCodes,omitempty"`
	Items *[]any `json:"Items,omitempty"`
	ResultCode *int `json:"ResultCode,omitempty"`
}

// Promotion is the typed data model for the promotion entity.
type Promotion struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// PromotionListMatch is the typed request payload for Promotion.ListTyped.
type PromotionListMatch struct {
	ErrorCodes *[]any `json:"ErrorCodes,omitempty"`
	Items *[]any `json:"Items,omitempty"`
	ResultCode *int `json:"ResultCode,omitempty"`
}

// PromotionDescription is the typed data model for the promotion_description entity.
type PromotionDescription struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// PromotionDescriptionListMatch is the typed request payload for PromotionDescription.ListTyped.
type PromotionDescriptionListMatch struct {
	ErrorCodes *[]any `json:"ErrorCodes,omitempty"`
	Items *[]any `json:"Items,omitempty"`
	ResultCode *int `json:"ResultCode,omitempty"`
}

// Provider is the typed data model for the provider entity.
type Provider struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// ProviderListMatch is the typed request payload for Provider.ListTyped.
type ProviderListMatch struct {
	ErrorCodes *[]any `json:"ErrorCodes,omitempty"`
	Items *[]any `json:"Items,omitempty"`
	ResultCode *int `json:"ResultCode,omitempty"`
}

// ProviderStatus is the typed data model for the provider_status entity.
type ProviderStatus struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// ProviderStatusListMatch is the typed request payload for ProviderStatus.ListTyped.
type ProviderStatusListMatch struct {
	ErrorCodes *[]any `json:"ErrorCodes,omitempty"`
	Items *[]any `json:"Items,omitempty"`
	ResultCode *int `json:"ResultCode,omitempty"`
}

// Region is the typed data model for the region entity.
type Region struct {
	ErrorCodes []any `json:"ErrorCodes"`
	Items []any `json:"Items"`
	ResultCode int `json:"ResultCode"`
}

// RegionListMatch is the typed request payload for Region.ListTyped.
type RegionListMatch struct {
	ErrorCodes *[]any `json:"ErrorCodes,omitempty"`
	Items *[]any `json:"Items,omitempty"`
	ResultCode *int `json:"ResultCode,omitempty"`
}

// SendTransfer is the typed data model for the send_transfer entity.
type SendTransfer struct {
	ErrorCodes []any `json:"ErrorCodes"`
	ResultCode int `json:"ResultCode"`
	TransferRecord map[string]any `json:"TransferRecord"`
}

// SendTransferCreateData is the typed request payload for SendTransfer.CreateTyped.
type SendTransferCreateData struct {
	ErrorCodes []any `json:"ErrorCodes"`
	ResultCode int `json:"ResultCode"`
	TransferRecord map[string]any `json:"TransferRecord"`
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
