package core

import (
	"fmt"
	"strings"

	vs "github.com/voxgig-sdk/dingconnect-sdk/go/utility/struct"
)

type DingconnectSDK struct {
	Mode     string
	options  map[string]any
	utility  *Utility
	Features []Feature
	rootctx  *Context
}

func NewDingconnectSDK(options map[string]any) *DingconnectSDK {
	sdk := &DingconnectSDK{
		Mode:     "live",
		Features: []Feature{},
	}

	sdk.utility = NewUtility()

	config := MakeConfig()

	sdk.rootctx = sdk.utility.MakeContext(map[string]any{
		"client":  sdk,
		"utility": sdk.utility,
		"config":  config,
		"options": options,
		"shared":  map[string]any{},
	}, nil)

	sdk.options = sdk.utility.MakeOptions(sdk.rootctx)

	if vs.GetPath([]any{"feature", "test", "active"}, sdk.options) == true {
		sdk.Mode = "test"
	}

	sdk.rootctx.Options = sdk.options

	// Add features in the resolved order (MakeOptions puts an explicit array
	// order first, else defaults to test-first). Ordering matters: the `test`
	// feature installs the base mock transport and the transport features
	// (retry/cache/netsim/proxy/ratelimit) wrap whatever is current, so `test`
	// must be added before them to sit at the base of the chain.
	featureOpts := ToMapAny(vs.GetProp(sdk.options, "feature"))
	if featureOpts != nil {
		if fo, ok := vs.GetPath([]any{"__derived__", "featureorder"}, sdk.options).([]any); ok {
			for _, n := range fo {
				fname, _ := n.(string)
				fopts := ToMapAny(featureOpts[fname])
				if fopts != nil {
					if active, ok := fopts["active"]; ok {
						if ab, ok := active.(bool); ok && ab {
							sdk.utility.FeatureAdd(sdk.rootctx, makeFeature(fname))
						}
					}
				}
			}
		}
	}

	// Add extension features.
	if extend := vs.GetProp(sdk.options, "extend"); extend != nil {
		if extList, ok := extend.([]any); ok {
			for _, f := range extList {
				if feat, ok := f.(Feature); ok {
					sdk.utility.FeatureAdd(sdk.rootctx, feat)
				}
			}
		}
	}

	// Initialize features.
	for _, f := range sdk.Features {
		sdk.utility.FeatureInit(sdk.rootctx, f)
	}

	sdk.utility.FeatureHook(sdk.rootctx, "PostConstruct")

	return sdk
}

func (sdk *DingconnectSDK) OptionsMap() map[string]any {
	out := vs.Clone(sdk.options)
	if om, ok := out.(map[string]any); ok {
		return om
	}
	return map[string]any{}
}

func (sdk *DingconnectSDK) GetUtility() *Utility {
	return CopyUtility(sdk.utility)
}

func (sdk *DingconnectSDK) GetRootCtx() *Context {
	return sdk.rootctx
}

func (sdk *DingconnectSDK) Prepare(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "prepare",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	options := sdk.options

	path, _ := vs.GetProp(fetchargs, "path").(string)
	method, _ := vs.GetProp(fetchargs, "method").(string)
	if method == "" {
		method = "GET"
	}

	params := ToMapAny(vs.GetProp(fetchargs, "params"))
	if params == nil {
		params = map[string]any{}
	}
	query := ToMapAny(vs.GetProp(fetchargs, "query"))
	if query == nil {
		query = map[string]any{}
	}

	headers := utility.PrepareHeaders(ctx)

	base, _ := vs.GetProp(options, "base").(string)
	prefix, _ := vs.GetProp(options, "prefix").(string)
	suffix, _ := vs.GetProp(options, "suffix").(string)

	ctx.Spec = NewSpec(map[string]any{
		"base":    base,
		"prefix":  prefix,
		"suffix":  suffix,
		"path":    path,
		"method":  method,
		"params":  params,
		"query":   query,
		"headers": headers,
		"body":    vs.GetProp(fetchargs, "body"),
		"step":    "start",
	})

	// Merge user-provided headers.
	if uh := vs.GetProp(fetchargs, "headers"); uh != nil {
		if uhm, ok := uh.(map[string]any); ok {
			for k, v := range uhm {
				ctx.Spec.Headers[k] = v
			}
		}
	}

	_, err := utility.PrepareAuth(ctx)
	if err != nil {
		return nil, err
	}

	return utility.MakeFetchDef(ctx)
}

// Raw endpoint access is operator-controllable, like every entity op.
// Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
// either one reaches the same endpoint.
func (sdk *DingconnectSDK) Direct(fetchargs map[string]any) (map[string]any, error) {
	if !sdk.opAllowed("direct") {
		return sdk.opDenied("direct"), nil
	}

	return sdk.rawRequest(fetchargs)
}

// Is this raw-access op permitted by the SDK's allow.op option?
func (sdk *DingconnectSDK) opAllowed(op string) bool {
	allowOp, _ := vs.GetPath([]any{"allow", "op"}, sdk.options).(string)
	return strings.Contains(allowOp, op)
}

func (sdk *DingconnectSDK) opDenied(op string) map[string]any {
	allowOp, _ := vs.GetPath([]any{"allow", "op"}, sdk.options).(string)
	return map[string]any{
		"ok": false,
		"err": fmt.Errorf("DingconnectSDK: %s: operation not allowed by"+
			" SDK option allow.op value: \"%s\"", op, allowOp),
	}
}

// Ungated request path shared by Direct and Graphql, each of which checks
// its own allow.op token first. Unexported, rather than a flag on fetchargs:
// a caller-supplied marker would let anyone opt straight back out of the
// gate by passing it.
func (sdk *DingconnectSDK) rawRequest(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	fetchdef, err := sdk.Prepare(fetchargs)
	if err != nil {
		return map[string]any{"ok": false, "err": err}, nil
	}

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "direct",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	url, _ := fetchdef["url"].(string)
	fetched, fetchErr := utility.Fetcher(ctx, url, fetchdef)

	if fetchErr != nil {
		return map[string]any{"ok": false, "err": fetchErr}, nil
	}

	if fetched == nil {
		return map[string]any{
			"ok":  false,
			"err": ctx.MakeError("direct_no_response", "response: undefined"),
		}, nil
	}

	if fm, ok := fetched.(map[string]any); ok {
		status := ToInt(vs.GetProp(fm, "status"))
		headers := vs.GetProp(fm, "headers")

		// No-body responses (204, 304) and explicit zero content-length
		// must skip JSON parsing — calling json() on an empty body errors.
		var contentLength string
		if hm, ok := headers.(map[string]any); ok {
			if cl, ok := hm["content-length"]; ok {
				contentLength = fmt.Sprintf("%v", cl)
			}
		}
		noBody := status == 204 || status == 304 || contentLength == "0"

		var jsonData any
		if !noBody {
			if jf := vs.GetProp(fm, "json"); jf != nil {
				if f, ok := jf.(func() any); ok {
					// f() returns nil on parse error in our fetcher.
					jsonData = f()
				}
			}
		}

		return map[string]any{
			"ok":      status >= 200 && status < 300,
			"status":  status,
			"headers": headers,
			"data":    jsonData,
		}, nil
	}

	return map[string]any{"ok": false, "err": ctx.MakeError("direct_invalid", "invalid response type")}, nil
}

// Raw GraphQL access: the pressure valve that makes the generated surface's
// deliberate omissions (per-call selection sets, typed filter builders,
// batching, subscriptions) livable — the whole schema stays reachable.
//
// Thin wrapper over the same prepare/fetch path Direct uses, with the one
// thing raw Direct cannot do for GraphQL: a GraphQL failure rides HTTP 200
// as a top-level `errors` array, so status alone would report a failed query
// as ok.
//
// NOTE: like Direct, this bypasses the feature pipeline — no retry,
// ratelimit or paging features apply.
func (sdk *DingconnectSDK) Graphql(
	query string, variables map[string]any, ctrl map[string]any,
) (map[string]any, error) {
	if !sdk.opAllowed("graphql") {
		return sdk.opDenied("graphql"), nil
	}

	if variables == nil {
		variables = map[string]any{}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	res, err := sdk.rawRequest(map[string]any{
		"method":  "POST",
		"headers": map[string]any{"content-type": "application/json"},
		"body":    map[string]any{"query": query, "variables": variables},
		"ctrl":    ctrl,
	})

	if err != nil {
		return res, err
	}

	// Errors are read BEFORE any status check: a GraphQL parse or validation
	// failure comes back as HTTP 400 carrying the standard { errors: [...] }
	// body, and the raw path represents a non-2xx as ok:false with no err —
	// so returning early on status would discard the server's own
	// diagnostics, which are the only useful part of that response.
	errors, _ := vs.GetPath([]any{"data", "errors"}, res).([]any)

	if 0 < len(errors) {
		msg, _ := vs.GetProp(errors[0], "message").(string)
		if msg == "" {
			msg = "graphql error"
		}
		res["ok"] = false
		res["err"] = fmt.Errorf("DingconnectSDK: graphql: %s", msg)
		res["graphql"] = errors
	}

	return res, nil
}


// AccountLookup returns a AccountLookup entity bound to this client.
// Idiomatic usage: client.AccountLookup(nil).List(nil, nil) or
// client.AccountLookup(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) AccountLookup(data map[string]any) DingconnectEntity {
	return NewAccountLookupEntityFunc(sdk, data)
}


// Balance returns a Balance entity bound to this client.
// Idiomatic usage: client.Balance(nil).List(nil, nil) or
// client.Balance(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) Balance(data map[string]any) DingconnectEntity {
	return NewBalanceEntityFunc(sdk, data)
}


// CancelTransfer returns a CancelTransfer entity bound to this client.
// Idiomatic usage: client.CancelTransfer(nil).List(nil, nil) or
// client.CancelTransfer(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) CancelTransfer(data map[string]any) DingconnectEntity {
	return NewCancelTransferEntityFunc(sdk, data)
}


// Country returns a Country entity bound to this client.
// Idiomatic usage: client.Country(nil).List(nil, nil) or
// client.Country(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) Country(data map[string]any) DingconnectEntity {
	return NewCountryEntityFunc(sdk, data)
}


// Currency returns a Currency entity bound to this client.
// Idiomatic usage: client.Currency(nil).List(nil, nil) or
// client.Currency(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) Currency(data map[string]any) DingconnectEntity {
	return NewCurrencyEntityFunc(sdk, data)
}


// ErrorCodeDescription returns a ErrorCodeDescription entity bound to this client.
// Idiomatic usage: client.ErrorCodeDescription(nil).List(nil, nil) or
// client.ErrorCodeDescription(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) ErrorCodeDescription(data map[string]any) DingconnectEntity {
	return NewErrorCodeDescriptionEntityFunc(sdk, data)
}


// EstimatePrice returns a EstimatePrice entity bound to this client.
// Idiomatic usage: client.EstimatePrice(nil).List(nil, nil) or
// client.EstimatePrice(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) EstimatePrice(data map[string]any) DingconnectEntity {
	return NewEstimatePriceEntityFunc(sdk, data)
}


// ListTransferRecord returns a ListTransferRecord entity bound to this client.
// Idiomatic usage: client.ListTransferRecord(nil).List(nil, nil) or
// client.ListTransferRecord(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) ListTransferRecord(data map[string]any) DingconnectEntity {
	return NewListTransferRecordEntityFunc(sdk, data)
}


// LookupBill returns a LookupBill entity bound to this client.
// Idiomatic usage: client.LookupBill(nil).List(nil, nil) or
// client.LookupBill(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) LookupBill(data map[string]any) DingconnectEntity {
	return NewLookupBillEntityFunc(sdk, data)
}


// Product returns a Product entity bound to this client.
// Idiomatic usage: client.Product(nil).List(nil, nil) or
// client.Product(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) Product(data map[string]any) DingconnectEntity {
	return NewProductEntityFunc(sdk, data)
}


// ProductDescription returns a ProductDescription entity bound to this client.
// Idiomatic usage: client.ProductDescription(nil).List(nil, nil) or
// client.ProductDescription(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) ProductDescription(data map[string]any) DingconnectEntity {
	return NewProductDescriptionEntityFunc(sdk, data)
}


// Promotion returns a Promotion entity bound to this client.
// Idiomatic usage: client.Promotion(nil).List(nil, nil) or
// client.Promotion(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) Promotion(data map[string]any) DingconnectEntity {
	return NewPromotionEntityFunc(sdk, data)
}


// PromotionDescription returns a PromotionDescription entity bound to this client.
// Idiomatic usage: client.PromotionDescription(nil).List(nil, nil) or
// client.PromotionDescription(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) PromotionDescription(data map[string]any) DingconnectEntity {
	return NewPromotionDescriptionEntityFunc(sdk, data)
}


// Provider returns a Provider entity bound to this client.
// Idiomatic usage: client.Provider(nil).List(nil, nil) or
// client.Provider(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) Provider(data map[string]any) DingconnectEntity {
	return NewProviderEntityFunc(sdk, data)
}


// ProviderStatus returns a ProviderStatus entity bound to this client.
// Idiomatic usage: client.ProviderStatus(nil).List(nil, nil) or
// client.ProviderStatus(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) ProviderStatus(data map[string]any) DingconnectEntity {
	return NewProviderStatusEntityFunc(sdk, data)
}


// Region returns a Region entity bound to this client.
// Idiomatic usage: client.Region(nil).List(nil, nil) or
// client.Region(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) Region(data map[string]any) DingconnectEntity {
	return NewRegionEntityFunc(sdk, data)
}


// SendTransfer returns a SendTransfer entity bound to this client.
// Idiomatic usage: client.SendTransfer(nil).List(nil, nil) or
// client.SendTransfer(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *DingconnectSDK) SendTransfer(data map[string]any) DingconnectEntity {
	return NewSendTransferEntityFunc(sdk, data)
}



func TestSDK(testopts map[string]any, sdkopts map[string]any) *DingconnectSDK {
	if sdkopts == nil {
		sdkopts = map[string]any{}
	}
	sdkopts = vs.Clone(sdkopts).(map[string]any)

	if testopts == nil {
		testopts = map[string]any{}
	}
	testopts = vs.Clone(testopts).(map[string]any)
	testopts["active"] = true

	vs.SetPath(sdkopts, []any{"feature", "test"}, testopts)

	sdk := NewDingconnectSDK(sdkopts)
	sdk.Mode = "test"

	return sdk
}
