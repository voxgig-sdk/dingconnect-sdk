"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('ListTransferRecordEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when DINGCONNECT_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('DINGCONNECT_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.DingconnectSDK.test();
        const ent = testsdk.ListTransferRecord();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('validate', async (t) => {
        if (null == __1.config.feature?.validate) {
            t.skip('feature not present in this SDK: validate');
            return;
        }
        const client = __1.DingconnectSDK.test(undefined, { feature: { validate: { active: true } } });
        await node_assert_1.default.rejects(client.ListTransferRecord().create({ "AccountNumber": 1, "ErrorCodes": "x", "Items": "x", "ResultCode": 1, "Take": 1, "ThereAreMoreItems": true }), (err) => 'validate_failed' === err.code);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.DINGCONNECT_TEST_LIVE;
        for (const op of ['create']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'list_transfer_record.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "AccountNumber": { "a": true, "h": "Account Number", "n": "AccountNumber", "r": false, "sh": "Filter transfers by AccountNumber", "t": "`$STRING`", "key$": "AccountNumber", "index$": 0 }, "DistributorRef": { "a": true, "h": "Distributor Ref", "n": "DistributorRef", "r": false, "sh": "Filter transfers by DistributorRef.", "t": "`$STRING`", "key$": "DistributorRef", "index$": 1 }, "ErrorCodes": { "a": true, "h": "Error Codes", "n": "ErrorCodes", "r": true, "t": "`$ARRAY`", "key$": "ErrorCodes", "index$": 2 }, "Items": { "a": true, "h": "Items", "n": "Items", "r": true, "sh": "The list of items satisfying the transfer query.", "t": "`$ARRAY`", "key$": "Items", "index$": 3 }, "ResultCode": { "a": true, "fo": "int32", "h": "Result Code", "n": "ResultCode", "r": true, "t": "`$INTEGER`", "key$": "ResultCode", "index$": 4 }, "Skip": { "a": true, "fo": "int32", "h": "Skip", "n": "Skip", "r": false, "sh": "The amount of records to by-pass before returning the remaining records", "t": "`$INTEGER`", "key$": "Skip", "index$": 5 }, "Take": { "a": true, "fo": "int32", "h": "Take", "n": "Take", "r": true, "sh": "The amount of records to return", "t": "`$INTEGER`", "key$": "Take", "index$": 6 }, "ThereAreMoreItems": { "a": true, "h": "There Are More Items", "n": "ThereAreMoreItems", "r": true, "sh": "Indicates if the caller should execute the query again.", "t": "`$BOOLEAN`", "key$": "ThereAreMoreItems", "index$": 7 }, "TransferRef": { "a": true, "h": "Transfer Ref", "n": "TransferRef", "r": false, "sh": "Filter by Ding TransferRef", "t": "`$STRING`", "key$": "TransferRef", "index$": 8 } }, "name": "list_transfer_record", "op": { "create": { "input": "data", "name": "create", "points": [{ "a": true, "bf": ["AccountNumber", "DistributorRef", "Skip", "Take", "TransferRef"], "co": { "id": "POST /api/V1/ListTransferRecords", "source": "swagger2", "version": 2 }, "g": { "header": [{ "a": true, "k": "header", "n": "x_correlation_id", "or": "X-Correlation-Id", "r": false, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "POST", "o": "/api/V1/ListTransferRecords", "q": {}, "r": {}, "rs": { "alternatives": [{ "kind": "json", "media": "text/json" }], "kind": "json", "media": "application/json" }, "s": [{ "lit": "api" }, { "lit": "V1" }, { "lit": "ListTransferRecords" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "create" } }, "relations": { "ancestors": [] }, "key$": "list_transfer_record", "name__orig": "list_transfer_record", "Name": "ListTransferRecord", "name_": "list_transfer_record", "name-": "list-transfer-record", "NAME": "LIST_TRANSFER_RECORD", "index$": 7 }, { "active": true, "entity": "list_transfer_record", "key$": "BasicListTransferRecordFlow", "kind": "basic", "name": "BasicListTransferRecordFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "list_transfer_record_ref01" }, "m": {}, "o": "create", "s": [], "v": [], "index$": 0 }] }, 'ListTransferRecord', { "POST /api/V1/ListTransferRecords": { "protocol": "http", "parameters": [{ "in": "header", "name": "X-Correlation-Id", "description": "Correlates HTTP requests between a client and server", "type": "String", "index$": 0 }, { "in": "body", "name": "request", "required": true, "schema": { "required": ["Take"], "type": "object", "properties": { "TransferRef": { "description": "Filter by Ding TransferRef", "type": "string", "key$": "TransferRef" }, "DistributorRef": { "description": "Filter transfers by DistributorRef.", "type": "string", "key$": "DistributorRef" }, "AccountNumber": { "description": "Filter transfers by AccountNumber", "type": "string", "key$": "AccountNumber" }, "Skip": { "format": "int32", "description": "The amount of records to by-pass before returning the remaining records", "type": "integer", "key$": "Skip" }, "Take": { "format": "int32", "description": "The amount of records to return", "type": "integer", "key$": "Take" } }, "additionalProperties": false, "x-ref": "#/definitions/ListTransferRecordsRequest", "index$": 1 }, "index$": 1 }] } }, { strict: LIVE_STRICT, t });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const list_transfer_record_ref01_ent = client.ListTransferRecord();
        let list_transfer_record_ref01_data = setup.data.new.list_transfer_record['list_transfer_record_ref01'];
        list_transfer_record_ref01_data = (await list_transfer_record_ref01_ent.create(list_transfer_record_ref01_data)).data();
        (0, node_assert_1.default)(null != list_transfer_record_ref01_data);
    });
});
// main.kit.test.live.strict is true (the default is true): a live
// request that fails, or a live test missing an input it needs,
// fails the test.
// An account with no record for a test to read skips it either way.
const LIVE_STRICT = true;
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/list_transfer_record/ListTransferRecordTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.DingconnectSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['list_transfer_record01', 'list_transfer_record02', 'list_transfer_record03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'DINGCONNECT_TEST_LIST_TRANSFER_RECORD_ENTID': idmap,
        'DINGCONNECT_TEST_LIVE': 'FALSE',
        'DINGCONNECT_TEST_EXPLAIN': 'FALSE',
        'DINGCONNECT_APIKEY': '',
    });
    idmap = env['DINGCONNECT_TEST_LIST_TRANSFER_RECORD_ENTID'];
    const live = 'TRUE' === env.DINGCONNECT_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['DINGCONNECT_TEST_LIST_TRANSFER_RECORD_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.DingconnectSDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
            {
                apikey: env.DINGCONNECT_APIKEY,
            },
            // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
            // last entry is undefined, and basicSetup is normally called with no
            // argument at all - so a bare 'extra' silently discarded the apikey
            // and server values above and handed the SDK undefined. Harmless
            // while there was nothing in that object; not harmless now.
            extra || {},
            { system: { fetch: transport.fetch } }
        ]));
    }
    const setup = {
        idmap,
        env,
        options,
        client,
        struct,
        data: entityData,
        explain: 'TRUE' === env.DINGCONNECT_TEST_EXPLAIN,
        live,
        transport,
        now: Date.now(),
    };
    return setup;
}
//# sourceMappingURL=ListTransferRecordEntity.test.js.map