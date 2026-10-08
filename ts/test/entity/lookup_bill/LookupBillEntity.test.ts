

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


describe('LookupBillEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when DINGCONNECT_TEST_LIVE=TRUE.
  afterEach(liveDelay('DINGCONNECT_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = DingconnectSDK.test()
    const ent = testsdk.LookupBill()
    assert(null != ent)
  })


  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = DingconnectSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.LookupBill().create({"AccountNumber":1,"ErrorCodes":"x","Items":"x","ResultCode":1,"SkuCode":"x"} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.DINGCONNECT_TEST_LIVE
    for (const op of ['create']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'lookup_bill.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"AccountNumber":{"a":true,"h":"Account Number","n":"AccountNumber","r":true,"sh":"The account number to target","t":"`$STRING`","key$":"AccountNumber","index$":0},"ErrorCodes":{"a":true,"h":"Error Codes","n":"ErrorCodes","r":true,"t":"`$ARRAY`","key$":"ErrorCodes","index$":1},"Items":{"a":true,"h":"Items","n":"Items","r":true,"t":"`$ARRAY`","key$":"Items","index$":2},"ResultCode":{"a":true,"fo":"int32","h":"Result Code","n":"ResultCode","r":true,"t":"`$INTEGER`","key$":"ResultCode","index$":3},"Settings":{"a":true,"h":"Settings","n":"Settings","r":false,"sh":"Product specific name/value pairs to be associated with the lookup bills request","t":"`$ARRAY`","key$":"Settings","index$":4},"SkuCode":{"a":true,"h":"Sku Code","n":"SkuCode","r":true,"sh":"Code provided by GetProducts API","t":"`$STRING`","key$":"SkuCode","index$":5}},"name":"lookup_bill","op":{"create":{"input":"data","name":"create","points":[{"a":true,"bf":["AccountNumber","Settings","SkuCode"],"co":{"id":"POST /api/V1/LookupBills","source":"swagger2","version":2},"g":{"header":[{"a":true,"k":"header","n":"x_correlation_id","or":"X-Correlation-Id","r":false,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/api/V1/LookupBills","q":{},"r":{},"rs":{"alternatives":[{"kind":"json","media":"text/json"}],"kind":"json","media":"application/json"},"s":[{"lit":"api"},{"lit":"V1"},{"lit":"LookupBills"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"lookup_bill","name__orig":"lookup_bill","Name":"LookupBill","name_":"lookup_bill","name-":"lookup-bill","NAME":"LOOKUP_BILL","index$":8}, {"active":true,"entity":"lookup_bill","key$":"BasicLookupBillFlow","kind":"basic","name":"BasicLookupBillFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"lookup_bill_ref01"},"m":{},"o":"create","s":[],"v":[],"index$":0}]}, 'LookupBill', {"POST /api/V1/LookupBills":{"protocol":"http","parameters":[{"in":"header","name":"X-Correlation-Id","description":"Correlates HTTP requests between a client and server","type":"String","index$":0},{"in":"body","name":"request","required":true,"schema":{"required":["AccountNumber","SkuCode"],"type":"object","properties":{"SkuCode":{"description":"Code provided by GetProducts API","type":"string","key$":"SkuCode"},"AccountNumber":{"description":"The account number to target","type":"string","key$":"AccountNumber"},"Settings":{"description":"Product specific name/value pairs to be associated with the lookup bills request","type":"array","items":{"description":"A simple name/value pair","required":["Name","Value"],"type":"object","properties":{"Name":{"description":"The name of the setting as defined in the SettingDefinition","type":"string"},"Value":{"description":"The transfer specific value to associate with the Name","type":"string"}},"additionalProperties":false,"x-ref":"#/definitions/Setting"},"key$":"Settings"}},"additionalProperties":false,"x-ref":"#/definitions/LookupBillsRequest","index$":1},"index$":1}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const lookup_bill_ref01_ent = client.LookupBill()
    let lookup_bill_ref01_data = setup.data.new.lookup_bill['lookup_bill_ref01']

    lookup_bill_ref01_data = (await lookup_bill_ref01_ent.create(lookup_bill_ref01_data)).data()
    assert(null != lookup_bill_ref01_data)


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
      '../../../../.sdk/test/entity/lookup_bill/LookupBillTestData.json')

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
    ['lookup_bill01','lookup_bill02','lookup_bill03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'DINGCONNECT_TEST_LOOKUP_BILL_ENTID': idmap,
    'DINGCONNECT_TEST_LIVE': 'FALSE',
    'DINGCONNECT_TEST_EXPLAIN': 'FALSE',
    'DINGCONNECT_APIKEY': '',
  })

  idmap = env['DINGCONNECT_TEST_LOOKUP_BILL_ENTID']

  const live = 'TRUE' === env.DINGCONNECT_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['DINGCONNECT_TEST_LOOKUP_BILL_ENTID']
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
  
