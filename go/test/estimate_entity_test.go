package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/dingconnect-sdk/go"
	"github.com/voxgig-sdk/dingconnect-sdk/go/core"

	vs "github.com/voxgig-sdk/dingconnect-sdk/go/utility/struct"
)

func TestEstimateEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Estimate(nil)
		if ent == nil {
			t.Fatal("expected non-nil EstimateEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := estimateBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "estimate." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set DINGCONNECT_TEST_ESTIMATE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		estimateRef01Ent := client.Estimate(nil)
		estimateRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "estimate"}, setup.data), "estimate_ref01"))

		estimateRef01DataResult, err := estimateRef01Ent.Create(estimateRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		estimateRef01Data = core.ToMapAny(estimateRef01DataResult)
		if estimateRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

	})
}

func estimateBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "estimate", "EstimateTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read estimate test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse estimate test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"estimate01", "estimate02", "estimate03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("DINGCONNECT_TEST_ESTIMATE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"DINGCONNECT_TEST_ESTIMATE_ENTID": idmap,
		"DINGCONNECT_TEST_LIVE":      "FALSE",
		"DINGCONNECT_TEST_EXPLAIN":   "FALSE",
		"DINGCONNECT_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["DINGCONNECT_TEST_ESTIMATE_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["DINGCONNECT_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["DINGCONNECT_APIKEY"],
			},
			extra,
		})
		client = sdk.NewDingconnectSDK(core.ToMapAny(mergedOpts))
	}

	live := env["DINGCONNECT_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["DINGCONNECT_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
