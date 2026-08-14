package voxgigdingconnectsdk

import (
	"github.com/voxgig-sdk/dingconnect-sdk/go/core"
	"github.com/voxgig-sdk/dingconnect-sdk/go/entity"
	"github.com/voxgig-sdk/dingconnect-sdk/go/feature"
	_ "github.com/voxgig-sdk/dingconnect-sdk/go/utility"
)

// Type aliases preserve external API.
type DingconnectSDK = core.DingconnectSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type DingconnectEntity = core.DingconnectEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type DingconnectError = core.DingconnectError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewAccountLookupEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewAccountLookupEntity(client, entopts)
	}
	core.NewBalanceEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewBalanceEntity(client, entopts)
	}
	core.NewCancelTransferEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewCancelTransferEntity(client, entopts)
	}
	core.NewCountryEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewCountryEntity(client, entopts)
	}
	core.NewCurrencyEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewCurrencyEntity(client, entopts)
	}
	core.NewErrorCodeDescriptionEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewErrorCodeDescriptionEntity(client, entopts)
	}
	core.NewEstimatePriceEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewEstimatePriceEntity(client, entopts)
	}
	core.NewListTransferRecordEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewListTransferRecordEntity(client, entopts)
	}
	core.NewLookupBillEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewLookupBillEntity(client, entopts)
	}
	core.NewProductEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewProductEntity(client, entopts)
	}
	core.NewProductDescriptionEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewProductDescriptionEntity(client, entopts)
	}
	core.NewPromotionEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewPromotionEntity(client, entopts)
	}
	core.NewPromotionDescriptionEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewPromotionDescriptionEntity(client, entopts)
	}
	core.NewProviderEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewProviderEntity(client, entopts)
	}
	core.NewProviderStatusEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewProviderStatusEntity(client, entopts)
	}
	core.NewRegionEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewRegionEntity(client, entopts)
	}
	core.NewSendTransferEntityFunc = func(client *core.DingconnectSDK, entopts map[string]any) core.DingconnectEntity {
		return entity.NewSendTransferEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewDingconnectSDK = core.NewDingconnectSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var SharedConfig = core.SharedConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewDingconnectSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *DingconnectSDK  { return NewDingconnectSDK(nil) }
func Test() *DingconnectSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
