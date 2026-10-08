package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	sdk "github.com/voxgig-sdk/dingconnect-sdk/go"
)

// ListArgs is what an agent sends to dingconnect_list.
type ListArgs struct {
	Entity string         `json:"entity" jsonschema:"one of: account_lookup | balance | country | currency | error_code_description | product | product_description | promotion | promotion_description | provider | provider_status | region"`
	Query  map[string]any `json:"query,omitempty" jsonschema:"optional filter map; omit it for the first page"`
}

func registerTools(server *mcp.Server, client *sdk.DingconnectSDK) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "dingconnect_list",
		Description: "List records from Dingconnect. Args: entity, query (optional filter map; omit it for the first page). Returns the first page of records as JSON.",
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true},
		InputSchema: entitySchema[ListArgs]("account_lookup", "balance", "country", "currency", "error_code_description", "product", "product_description", "promotion", "promotion_description", "provider", "provider_status", "region"),
	}, func(ctx context.Context, req *mcp.CallToolRequest, args ListArgs) (*mcp.CallToolResult, any, error) {
		return runOp(ctx, client, "list", args.Entity, args.Query)
	})
}

// entitySchema is the schema inferred from In, its entity limited to the
// entities the tool serves.
func entitySchema[In any](names ...string) *jsonschema.Schema {
	schema, err := jsonschema.For[In](nil)
	if err != nil {
		panic(err)
	}
	enum := make([]any, len(names))
	for i, name := range names {
		enum[i] = name
	}
	schema.Properties["entity"].Enum = enum
	return schema
}

func runOp(_ context.Context, client *sdk.DingconnectSDK, op string, entity string, input map[string]any) (*mcp.CallToolResult, any, error) {
	ent, err := entityFor(client, entity)
	if err != nil {
		return toolError(err.Error())
	}

	var result any
	switch op {
	case "list":
		result, err = ent.List(input, nil)
	case "load":
		result, err = ent.Load(input, nil)
	case "create":
		result, err = ent.Create(input, nil)
	case "update":
		result, err = ent.Update(input, nil)
	case "patch":
		result, err = ent.Patch(input, nil)
	case "remove":
		result, err = ent.Remove(input, nil)
	default:
		return toolError(fmt.Sprintf("unknown op %q", op))
	}
	if err != nil {
		return toolError(err.Error())
	}

	// SDK returns *Entity wrappers; unwrap each via .Data() to get a
	// plain map[string]any (or []any of maps for list) suitable for
	// JSON marshalling.
	data := extractData(result)
	body, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return toolError(fmt.Sprintf("marshal: %v", err))
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(body)},
		},
	}, data, nil
}

// entityFor dispatches on the lowercase entity name. The generator
// emits one `case "<name>":` per entity defined in the SDK model.
func entityFor(client *sdk.DingconnectSDK, name string) (sdk.DingconnectEntity, error) {
	switch strings.ToLower(name) {
	case "account_lookup":
		return client.AccountLookup(nil), nil
	case "balance":
		return client.Balance(nil), nil
	case "cancel_transfer":
		return client.CancelTransfer(nil), nil
	case "country":
		return client.Country(nil), nil
	case "currency":
		return client.Currency(nil), nil
	case "error_code_description":
		return client.ErrorCodeDescription(nil), nil
	case "estimate_price":
		return client.EstimatePrice(nil), nil
	case "list_transfer_record":
		return client.ListTransferRecord(nil), nil
	case "lookup_bill":
		return client.LookupBill(nil), nil
	case "product":
		return client.Product(nil), nil
	case "product_description":
		return client.ProductDescription(nil), nil
	case "promotion":
		return client.Promotion(nil), nil
	case "promotion_description":
		return client.PromotionDescription(nil), nil
	case "provider":
		return client.Provider(nil), nil
	case "provider_status":
		return client.ProviderStatus(nil), nil
	case "region":
		return client.Region(nil), nil
	case "send_transfer":
		return client.SendTransfer(nil), nil
	}
	return nil, fmt.Errorf("unknown entity %q", name)
}

func extractData(x any) any {
	switch v := x.(type) {
	case sdk.Entity:
		return extractData(v.Data())
	case []any:
		out := make([]any, len(v))
		for i, e := range v {
			out[i] = extractData(e)
		}
		return out
	case map[string]any:
		out := make(map[string]any, len(v))
		for k, vv := range v {
			out[k] = extractData(vv)
		}
		return out
	}
	return x
}

func toolError(msg string) (*mcp.CallToolResult, any, error) {
	return &mcp.CallToolResult{
		IsError: true,
		Content: []mcp.Content{
			&mcp.TextContent{Text: msg},
		},
	}, nil, nil
}

// hint is an MCP annotation that defaults to true unless stated.
func hint(b bool) *bool {
	return &b
}
