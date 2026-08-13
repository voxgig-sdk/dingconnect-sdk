<?php
declare(strict_types=1);

// LookupBill entity test

require_once __DIR__ . '/../dingconnect_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class LookupBillEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = DingconnectSDK::test(null, null);
        $ent = $testsdk->LookupBill(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = lookup_bill_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "lookup_bill." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set DINGCONNECT_TEST_LOOKUP_BILL_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $lookup_bill_ref01_ent = $client->LookupBill(null);
        $lookup_bill_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.lookup_bill"), "lookup_bill_ref01"));

        $lookup_bill_ref01_data_result = $lookup_bill_ref01_ent->create($lookup_bill_ref01_data, null);
        $lookup_bill_ref01_data = Helpers::to_map(is_object($lookup_bill_ref01_data_result) && method_exists($lookup_bill_ref01_data_result, 'data_get') ? $lookup_bill_ref01_data_result->data_get() : $lookup_bill_ref01_data_result);
        $this->assertNotNull($lookup_bill_ref01_data);

    }
}

function lookup_bill_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/lookup_bill/LookupBillTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = DingconnectSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["lookup_bill01", "lookup_bill02", "lookup_bill03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("DINGCONNECT_TEST_LOOKUP_BILL_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "DINGCONNECT_TEST_LOOKUP_BILL_ENTID" => $idmap,
        "DINGCONNECT_TEST_LIVE" => "FALSE",
        "DINGCONNECT_TEST_EXPLAIN" => "FALSE",
        "DINGCONNECT_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["DINGCONNECT_TEST_LOOKUP_BILL_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["DINGCONNECT_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
                "apikey" => $env["DINGCONNECT_APIKEY"],
            ],
            $extra ?? [],
        ]);
        $client = new DingconnectSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["DINGCONNECT_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["DINGCONNECT_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
