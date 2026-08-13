package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewAccountLookupEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

var NewBalanceEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

var NewCancelTransferEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

var NewCountryEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

var NewCurrencyEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

var NewErrorCodeDescriptionEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

var NewEstimatePriceEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

var NewListTransferRecordEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

var NewLookupBillEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

var NewProductEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

var NewProductDescriptionEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

var NewPromotionEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

var NewPromotionDescriptionEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

var NewProviderEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

var NewProviderStatusEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

var NewRegionEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

var NewSendTransferEntityFunc func(client *DingconnectSDK, entopts map[string]any) DingconnectEntity

