# ListTransferRecord entity test

require "minitest/autorun"
require "json"
require_relative "../Dingconnect_sdk"
require_relative "runner"

class ListTransferRecordEntityTest < Minitest::Test
  def test_create_instance
    testsdk = DingconnectSDK.test(nil, nil)
    ent = testsdk.ListTransferRecord(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = list_transfer_record_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "list_transfer_record." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set DINGCONNECT_TEST_LIST_TRANSFER_RECORD_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    list_transfer_record_ref01_ent = client.ListTransferRecord(nil)
    list_transfer_record_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.list_transfer_record"), "list_transfer_record_ref01"))

    list_transfer_record_ref01_data_result = list_transfer_record_ref01_ent.create(list_transfer_record_ref01_data, nil)
    list_transfer_record_ref01_data = Helpers.to_map(list_transfer_record_ref01_data_result.respond_to?(:data_get) ? list_transfer_record_ref01_data_result.data_get : list_transfer_record_ref01_data_result)
    assert !list_transfer_record_ref01_data.nil?

  end
end

def list_transfer_record_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "list_transfer_record", "ListTransferRecordTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = DingconnectSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["list_transfer_record01", "list_transfer_record02", "list_transfer_record03"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Detect ENTID env override before envOverride consumes it. When live
  # mode is on without a real override, the basic test runs against synthetic
  # IDs from the fixture and 4xx's. Surface this so the test can skip.
  entid_env_raw = ENV["DINGCONNECT_TEST_LIST_TRANSFER_RECORD_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "DINGCONNECT_TEST_LIST_TRANSFER_RECORD_ENTID" => idmap,
    "DINGCONNECT_TEST_LIVE" => "FALSE",
    "DINGCONNECT_TEST_EXPLAIN" => "FALSE",
    "DINGCONNECT_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["DINGCONNECT_TEST_LIST_TRANSFER_RECORD_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["DINGCONNECT_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      {
        "apikey" => env["DINGCONNECT_APIKEY"],
      },
      extra || {},
    ])
    client = DingconnectSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["DINGCONNECT_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["DINGCONNECT_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
