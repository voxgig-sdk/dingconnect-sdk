

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { DingconnectSDK, BaseFeature, config, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


loadEnvLocal(__dirname + '/../../../.env.local')


describe('EstimatePriceEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when DINGCONNECT_TEST_LIVE=TRUE.
  afterEach(liveDelay('DINGCONNECT_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = DingconnectSDK.test()
    const ent = testsdk.EstimatePrice()
    assert(null != ent)
  })


  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = DingconnectSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.EstimatePrice().create({"ErrorCodes":"x","Items":"x","ResultCode":"x"} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.DINGCONNECT_TEST_LIVE
    for (const op of ['create']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'estimate_price.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"ErrorCodes":{"a":true,"h":"Error Codes","n":"ErrorCodes","r":true,"t":"`$ARRAY`","key$":"ErrorCodes","index$":0},"Items":{"a":true,"h":"Items","n":"Items","r":true,"t":"`$ARRAY`","key$":"Items","index$":1},"ResultCode":{"a":true,"fo":"int32","h":"Result Code","n":"ResultCode","r":true,"t":"`$INTEGER`","key$":"ResultCode","index$":2},"estimations":{"a":true,"h":"Estimations","n":"estimations","r":false,"t":"`$ARRAY`","key$":"estimations","index$":3}},"name":"estimate_price","op":{"create":{"input":"data","name":"create","points":[{"a":true,"bf":["estimations"],"co":{"id":"POST /api/V1/EstimatePrices","source":"swagger2","version":2},"g":{"header":[{"a":true,"k":"header","n":"x_correlation_id","or":"X-Correlation-Id","r":false,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/api/V1/EstimatePrices","q":{"exist":["estimations"]},"r":{},"rs":{"alternatives":[{"kind":"json","media":"text/json"}],"kind":"json","media":"application/json"},"s":[{"lit":"api"},{"lit":"V1"},{"lit":"EstimatePrices"}],"t":{"req":"`reqdata.estimations`","res":"`body`"},"index$":0}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"estimate_price","name__orig":"estimate_price","Name":"EstimatePrice","name_":"estimate_price","name-":"estimate-price","NAME":"ESTIMATE_PRICE","index$":6}, {"active":true,"entity":"estimate_price","key$":"BasicEstimatePriceFlow","kind":"basic","name":"BasicEstimatePriceFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"estimate_price_ref01"},"m":{},"o":"create","s":[],"v":[],"index$":0}]}, 'EstimatePrice', {"POST /api/V1/EstimatePrices":{"protocol":"http","parameters":[{"in":"header","name":"X-Correlation-Id","description":"Correlates HTTP requests between a client and server","type":"String","index$":0},{"in":"body","name":"requestedEstimations","required":true,"schema":{"type":"array","items":{"required":["BatchItemRef","SkuCode"],"type":"object","properties":{"SendValue":{"format":"decimal","description":"Desired send value","type":"number"},"SendCurrencyIso":{"description":"The currency of `SendValue`","type":"string"},"ReceiveValue":{"format":"decimal","description":"Desired receive value","type":"number"},"SkuCode":{"description":"Product to estimate a price for","type":"string"},"BatchItemRef":{"description":"A unique reference for an item in a batched request.","type":"string"}},"additionalProperties":false,"x-ref":"#/definitions/EstimationRequest"},"index$":1},"index$":1}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const estimate_price_ref01_ent = client.EstimatePrice()
    let estimate_price_ref01_data = setup.data.new.estimate_price['estimate_price_ref01']

    estimate_price_ref01_data = (await estimate_price_ref01_ent.create(estimate_price_ref01_data)).data()
    assert(null != estimate_price_ref01_data)


  })
})



// main.kit.test.live.strict is true (the default is true): a live
// request that fails, or a live test missing an input it needs,
// fails the test.
// An account with no record for a test to read skips it either way.
const LIVE_STRICT = true

function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/estimate_price/EstimatePriceTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = DingconnectSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['estimate_price01','estimate_price02','estimate_price03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'DINGCONNECT_TEST_ESTIMATE_PRICE_ENTID': idmap,
    'DINGCONNECT_TEST_LIVE': 'FALSE',
    'DINGCONNECT_TEST_EXPLAIN': 'FALSE',
    'DINGCONNECT_APIKEY': '',
  })

  idmap = env['DINGCONNECT_TEST_ESTIMATE_PRICE_ENTID']

  const live = 'TRUE' === env.DINGCONNECT_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['DINGCONNECT_TEST_ESTIMATE_PRICE_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new DingconnectSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
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
    ]))
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
  }

  return setup
}
  
