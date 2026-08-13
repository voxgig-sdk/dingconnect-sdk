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

func TestListTransferRecordEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.ListTransferRecord(nil)
		if ent == nil {
			t.Fatal("expected non-nil ListTransferRecordEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := list_transfer_recordBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "list_transfer_record." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set DINGCONNECT_TEST_LIST_TRANSFER_RECORD_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		listTransferRecordRef01Ent := client.ListTransferRecord(nil)
		listTransferRecordRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "list_transfer_record"}, setup.data), "list_transfer_record_ref01"))

		listTransferRecordRef01DataResult, err := listTransferRecordRef01Ent.Create(listTransferRecordRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		listTransferRecordRef01Data = core.ToMapAny(entityData(listTransferRecordRef01DataResult))
		if listTransferRecordRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

	})
}

func list_transfer_recordBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "list_transfer_record", "ListTransferRecordTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read list_transfer_record test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse list_transfer_record test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"list_transfer_record01", "list_transfer_record02", "list_transfer_record03"},
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
	entidEnvRaw := os.Getenv("DINGCONNECT_TEST_LIST_TRANSFER_RECORD_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"DINGCONNECT_TEST_LIST_TRANSFER_RECORD_ENTID": idmap,
		"DINGCONNECT_TEST_LIVE":      "FALSE",
		"DINGCONNECT_TEST_EXPLAIN":   "FALSE",
		"DINGCONNECT_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["DINGCONNECT_TEST_LIST_TRANSFER_RECORD_ENTID"])
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
