

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


describe('SendTransferEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when DINGCONNECT_TEST_LIVE=TRUE.
  afterEach(liveDelay('DINGCONNECT_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = DingconnectSDK.test()
    const ent = testsdk.SendTransfer()
    assert(null != ent)
  })


  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = DingconnectSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.SendTransfer().create({"AccountNumber":1,"DistributorRef":"x","ErrorCodes":"x","ResultCode":1,"SendValue":1,"SkuCode":"x","TransferRecord":"x","ValidateOnly":true} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.DINGCONNECT_TEST_LIVE
    for (const op of ['create']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'send_transfer.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"AccountNumber":{"a":true,"h":"Account Number","n":"AccountNumber","r":true,"sh":"The account number to target","t":"`$STRING`","key$":"AccountNumber","index$":0},"BillRef":{"a":true,"h":"Bill Ref","n":"BillRef","r":false,"sh":"Bill reference.","t":"`$STRING`","key$":"BillRef","index$":1},"DistributorRef":{"a":true,"h":"Distributor Ref","n":"DistributorRef","r":true,"sh":"Unique identifier in the distributor system to be associated with the transfer","t":"`$STRING`","key$":"DistributorRef","index$":2},"ErrorCodes":{"a":true,"h":"Error Codes","n":"ErrorCodes","r":true,"t":"`$ARRAY`","key$":"ErrorCodes","index$":3},"ResultCode":{"a":true,"fo":"int32","h":"Result Code","n":"ResultCode","r":true,"t":"`$INTEGER`","key$":"ResultCode","index$":4},"SendCurrencyIso":{"a":true,"h":"Send Currency Iso","n":"SendCurrencyIso","r":false,"sh":"The currency of the `SendValue`.","t":"`$STRING`","key$":"SendCurrencyIso","index$":5},"SendValue":{"a":true,"fo":"decimal","h":"Send Value","n":"SendValue","r":true,"sh":"The transfer value to be sent.","t":"`$NUMBER`","key$":"SendValue","index$":6},"Settings":{"a":true,"h":"Settings","n":"Settings","r":false,"sh":"Product specific name/value pairs to be associated with the transfer request","t":"`$ARRAY`","key$":"Settings","index$":7},"SkuCode":{"a":true,"h":"Sku Code","n":"SkuCode","r":true,"sh":"Code provided by GetProducts API","t":"`$STRING`","key$":"SkuCode","index$":8},"TransferRecord":{"a":true,"h":"Transfer Record","n":"TransferRecord","r":true,"t":"`$OBJECT`","key$":"TransferRecord","index$":9},"ValidateOnly":{"a":true,"h":"Validate Only","n":"ValidateOnly","r":true,"sh":"Validate the request with the provider without doing a transfer","t":"`$BOOLEAN`","key$":"ValidateOnly","index$":10}},"name":"send_transfer","op":{"create":{"input":"data","name":"create","points":[{"a":true,"bf":["AccountNumber","BillRef","DistributorRef","SendCurrencyIso","SendValue","Settings","SkuCode","ValidateOnly"],"co":{"id":"POST /api/V1/SendTransfer","source":"swagger2","version":2},"g":{"header":[{"a":true,"k":"header","n":"x_correlation_id","or":"X-Correlation-Id","r":false,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/api/V1/SendTransfer","q":{},"r":{},"rs":{"alternatives":[{"kind":"json","media":"text/json"}],"kind":"json","media":"application/json"},"s":[{"lit":"api"},{"lit":"V1"},{"lit":"SendTransfer"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"send_transfer","name__orig":"send_transfer","Name":"SendTransfer","name_":"send_transfer","name-":"send-transfer","NAME":"SEND_TRANSFER","index$":16}, {"active":true,"entity":"send_transfer","key$":"BasicSendTransferFlow","kind":"basic","name":"BasicSendTransferFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"send_transfer_ref01"},"m":{},"o":"create","s":[],"v":[],"index$":0}]}, 'SendTransfer', {"POST /api/V1/SendTransfer":{"protocol":"http","parameters":[{"in":"header","name":"X-Correlation-Id","description":"Correlates HTTP requests between a client and server","type":"String","index$":0},{"in":"body","name":"request","required":true,"schema":{"required":["AccountNumber","DistributorRef","SendValue","SkuCode","ValidateOnly"],"type":"object","properties":{"SkuCode":{"description":"Code provided by GetProducts API","type":"string","key$":"SkuCode"},"SendValue":{"format":"decimal","description":"The transfer value to be sent. Specified to two decimal places of accuracy of the major currency unit, e.g. 3.17 USD.","type":"number","key$":"SendValue"},"SendCurrencyIso":{"description":"The currency of the `SendValue`. If this is null or empty, we will assume distributor currency.","type":"string","key$":"SendCurrencyIso"},"AccountNumber":{"description":"The account number to target","type":"string","key$":"AccountNumber"},"DistributorRef":{"description":"Unique identifier in the distributor system to be associated with the transfer","type":"string","key$":"DistributorRef"},"Settings":{"description":"Product specific name/value pairs to be associated with the transfer request","type":"array","items":{"description":"A simple name/value pair","required":["Name","Value"],"type":"object","properties":{"Name":{"description":"The name of the setting as defined in the SettingDefinition","type":"string"},"Value":{"description":"The transfer specific value to associate with the Name","type":"string"}},"additionalProperties":false,"x-ref":"#/definitions/Setting"},"key$":"Settings"},"ValidateOnly":{"description":"Validate the request with the provider without doing a transfer","type":"boolean","key$":"ValidateOnly"},"BillRef":{"description":"Bill reference. Required when product has \"LookupBillsRequired\" set to true.","type":"string","key$":"BillRef"}},"additionalProperties":false,"x-ref":"#/definitions/SendTransferRequest","index$":1},"index$":1}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const send_transfer_ref01_ent = client.SendTransfer()
    let send_transfer_ref01_data = setup.data.new.send_transfer['send_transfer_ref01']

    send_transfer_ref01_data = (await send_transfer_ref01_ent.create(send_transfer_ref01_data)).data()
    assert(null != send_transfer_ref01_data)


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
      '../../../../.sdk/test/entity/send_transfer/SendTransferTestData.json')

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
    ['send_transfer01','send_transfer02','send_transfer03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'DINGCONNECT_TEST_SEND_TRANSFER_ENTID': idmap,
    'DINGCONNECT_TEST_LIVE': 'FALSE',
    'DINGCONNECT_TEST_EXPLAIN': 'FALSE',
    'DINGCONNECT_APIKEY': '',
  })

  idmap = env['DINGCONNECT_TEST_SEND_TRANSFER_ENTID']

  const live = 'TRUE' === env.DINGCONNECT_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['DINGCONNECT_TEST_SEND_TRANSFER_ENTID']
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
  
