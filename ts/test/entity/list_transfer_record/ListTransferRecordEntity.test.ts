

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


describe('ListTransferRecordEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when DINGCONNECT_TEST_LIVE=TRUE.
  afterEach(liveDelay('DINGCONNECT_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = DingconnectSDK.test()
    const ent = testsdk.ListTransferRecord()
    assert(null != ent)
  })


  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = DingconnectSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.ListTransferRecord().create({"AccountNumber":1,"ErrorCodes":"x","Items":"x","ResultCode":1,"Take":1,"ThereAreMoreItems":true} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.DINGCONNECT_TEST_LIVE
    for (const op of ['create']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'list_transfer_record.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"AccountNumber":{"a":true,"h":"Account Number","n":"AccountNumber","r":false,"sh":"Filter transfers by AccountNumber","t":"`$STRING`","key$":"AccountNumber","index$":0},"DistributorRef":{"a":true,"h":"Distributor Ref","n":"DistributorRef","r":false,"sh":"Filter transfers by DistributorRef.","t":"`$STRING`","key$":"DistributorRef","index$":1},"ErrorCodes":{"a":true,"h":"Error Codes","n":"ErrorCodes","r":true,"t":"`$ARRAY`","key$":"ErrorCodes","index$":2},"Items":{"a":true,"h":"Items","n":"Items","r":true,"sh":"The list of items satisfying the transfer query.","t":"`$ARRAY`","key$":"Items","index$":3},"ResultCode":{"a":true,"fo":"int32","h":"Result Code","n":"ResultCode","r":true,"t":"`$INTEGER`","key$":"ResultCode","index$":4},"Skip":{"a":true,"fo":"int32","h":"Skip","n":"Skip","r":false,"sh":"The amount of records to by-pass before returning the remaining records","t":"`$INTEGER`","key$":"Skip","index$":5},"Take":{"a":true,"fo":"int32","h":"Take","n":"Take","r":true,"sh":"The amount of records to return","t":"`$INTEGER`","key$":"Take","index$":6},"ThereAreMoreItems":{"a":true,"h":"There Are More Items","n":"ThereAreMoreItems","r":true,"sh":"Indicates if the caller should execute the query again.","t":"`$BOOLEAN`","key$":"ThereAreMoreItems","index$":7},"TransferRef":{"a":true,"h":"Transfer Ref","n":"TransferRef","r":false,"sh":"Filter by Ding TransferRef","t":"`$STRING`","key$":"TransferRef","index$":8}},"name":"list_transfer_record","op":{"create":{"input":"data","name":"create","points":[{"a":true,"bf":["AccountNumber","DistributorRef","Skip","Take","TransferRef"],"co":{"id":"POST /api/V1/ListTransferRecords","source":"swagger2","version":2},"g":{"header":[{"a":true,"k":"header","n":"x_correlation_id","or":"X-Correlation-Id","r":false,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/api/V1/ListTransferRecords","q":{},"r":{},"rs":{"alternatives":[{"kind":"json","media":"text/json"}],"kind":"json","media":"application/json"},"s":[{"lit":"api"},{"lit":"V1"},{"lit":"ListTransferRecords"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"list_transfer_record","name__orig":"list_transfer_record","Name":"ListTransferRecord","name_":"list_transfer_record","name-":"list-transfer-record","NAME":"LIST_TRANSFER_RECORD","index$":7}, {"active":true,"entity":"list_transfer_record","key$":"BasicListTransferRecordFlow","kind":"basic","name":"BasicListTransferRecordFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"list_transfer_record_ref01"},"m":{},"o":"create","s":[],"v":[],"index$":0}]}, 'ListTransferRecord', {"POST /api/V1/ListTransferRecords":{"protocol":"http","parameters":[{"in":"header","name":"X-Correlation-Id","description":"Correlates HTTP requests between a client and server","type":"String","index$":0},{"in":"body","name":"request","required":true,"schema":{"required":["Take"],"type":"object","properties":{"TransferRef":{"description":"Filter by Ding TransferRef","type":"string","key$":"TransferRef"},"DistributorRef":{"description":"Filter transfers by DistributorRef.","type":"string","key$":"DistributorRef"},"AccountNumber":{"description":"Filter transfers by AccountNumber","type":"string","key$":"AccountNumber"},"Skip":{"format":"int32","description":"The amount of records to by-pass before returning the remaining records","type":"integer","key$":"Skip"},"Take":{"format":"int32","description":"The amount of records to return","type":"integer","key$":"Take"}},"additionalProperties":false,"x-ref":"#/definitions/ListTransferRecordsRequest","index$":1},"index$":1}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const list_transfer_record_ref01_ent = client.ListTransferRecord()
    let list_transfer_record_ref01_data = setup.data.new.list_transfer_record['list_transfer_record_ref01']

    list_transfer_record_ref01_data = (await list_transfer_record_ref01_ent.create(list_transfer_record_ref01_data)).data()
    assert(null != list_transfer_record_ref01_data)


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
      '../../../../.sdk/test/entity/list_transfer_record/ListTransferRecordTestData.json')

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
    ['list_transfer_record01','list_transfer_record02','list_transfer_record03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'DINGCONNECT_TEST_LIST_TRANSFER_RECORD_ENTID': idmap,
    'DINGCONNECT_TEST_LIVE': 'FALSE',
    'DINGCONNECT_TEST_EXPLAIN': 'FALSE',
    'DINGCONNECT_APIKEY': '',
  })

  idmap = env['DINGCONNECT_TEST_LIST_TRANSFER_RECORD_ENTID']

  const live = 'TRUE' === env.DINGCONNECT_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['DINGCONNECT_TEST_LIST_TRANSFER_RECORD_ENTID']
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
  
