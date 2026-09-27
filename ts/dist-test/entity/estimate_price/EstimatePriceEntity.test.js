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
(0, node_test_1.describe)('EstimatePriceEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when DINGCONNECT_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('DINGCONNECT_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.DingconnectSDK.test();
        const ent = testsdk.EstimatePrice();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.DINGCONNECT_TEST_LIVE;
        for (const op of ['create']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'estimate_price.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "ErrorCodes": { "a": true, "h": "Error Codes", "n": "ErrorCodes", "r": true, "t": "`$ARRAY`", "key$": "ErrorCodes", "index$": 0 }, "Items": { "a": true, "h": "Items", "n": "Items", "r": true, "t": "`$ARRAY`", "key$": "Items", "index$": 1 }, "ResultCode": { "a": true, "fo": "int32", "h": "Result Code", "n": "ResultCode", "r": true, "t": "`$INTEGER`", "key$": "ResultCode", "index$": 2 } }, "name": "estimate_price", "op": { "create": { "input": "data", "name": "create", "points": [{ "a": true, "co": { "id": "POST /api/V1/EstimatePrices", "source": "swagger2", "version": 2 }, "g": { "header": [{ "a": true, "k": "header", "n": "x_correlation_id", "or": "x_correlation_id", "r": false, "t": "`$STRING`", "index$": 0 }], "query": [{ "a": true, "k": "query", "n": "requested_estimation", "or": "requested_estimation", "r": true, "t": "`$ARRAY`", "index$": 0 }] }, "k": "http", "m": "POST", "o": "/api/V1/EstimatePrices", "q": { "exist": ["requested_estimation", "x_correlation_id"] }, "r": {}, "s": [{ "lit": "api" }, { "lit": "V1" }, { "lit": "EstimatePrices" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "create" } }, "relations": { "ancestors": [] }, "key$": "estimate_price", "name__orig": "estimate_price", "Name": "EstimatePrice", "name_": "estimate_price", "name-": "estimate-price", "NAME": "ESTIMATE_PRICE", "index$": 6 }, { "active": true, "entity": "estimate_price", "key$": "BasicEstimatePriceFlow", "kind": "basic", "name": "BasicEstimatePriceFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "estimate_price_ref01" }, "m": {}, "o": "create", "s": [], "v": [], "index$": 0 }] }, 'EstimatePrice', { "POST /api/V1/EstimatePrices": { "protocol": "http", "parameters": [{ "in": "header", "name": "X-Correlation-Id", "description": "Correlates HTTP requests between a client and server", "type": "String", "index$": 0 }, { "in": "body", "name": "requestedEstimations", "required": true, "schema": { "type": "array", "items": { "required": ["BatchItemRef", "SkuCode"], "type": "object", "properties": { "SendValue": { "format": "decimal", "description": "Desired send value", "type": "number" }, "SendCurrencyIso": { "description": "The currency of `SendValue`", "type": "string" }, "ReceiveValue": { "format": "decimal", "description": "Desired receive value", "type": "number" }, "SkuCode": { "description": "Product to estimate a price for", "type": "string" }, "BatchItemRef": { "description": "A unique reference for an item in a batched request.", "type": "string" } }, "additionalProperties": false, "x-ref": "#/definitions/EstimationRequest" } }, "index$": 1 }] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const estimate_price_ref01_ent = client.EstimatePrice();
        let estimate_price_ref01_data = setup.data.new.estimate_price['estimate_price_ref01'];
        estimate_price_ref01_data = (await estimate_price_ref01_ent.create(estimate_price_ref01_data)).data();
        (0, node_assert_1.default)(null != estimate_price_ref01_data);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/estimate_price/EstimatePriceTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.DingconnectSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['estimate_price01', 'estimate_price02', 'estimate_price03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'DINGCONNECT_TEST_ESTIMATE_PRICE_ENTID': idmap,
        'DINGCONNECT_TEST_LIVE': 'FALSE',
        'DINGCONNECT_TEST_EXPLAIN': 'FALSE',
        'DINGCONNECT_APIKEY': '',
    });
    idmap = env['DINGCONNECT_TEST_ESTIMATE_PRICE_ENTID'];
    const live = 'TRUE' === env.DINGCONNECT_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['DINGCONNECT_TEST_ESTIMATE_PRICE_ENTID'];
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
//# sourceMappingURL=EstimatePriceEntity.test.js.map