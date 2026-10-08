

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


describe('AccountLookupEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when DINGCONNECT_TEST_LIVE=TRUE.
  afterEach(liveDelay('DINGCONNECT_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = DingconnectSDK.test()
    const ent = testsdk.AccountLookup()
    assert(null != ent)
  })


  class FailHook extends BaseFeature {
    name = 'failhook'
    version = '0.0.1'
    active = true
    unexpected = 0
    init() { }
    PreSpec() { throw new Error('account_lookup hook failed') }
    PreUnexpected() { this.unexpected++ }
  }

  test('stream-error', async () => {
    const offline = { net: { offline: true } }
    await assert.rejects(async () => {
      for await (const _item of DingconnectSDK.test(offline).AccountLookup().stream('list')) { }
    }, /offline/)

    for await (const _item of DingconnectSDK.test(offline).AccountLookup()
      .stream('list', undefined, { ctrl: { throw: false } })) { }

    if (null != (config as any).feature?.rbac) {
      const denied = DingconnectSDK.test(undefined, { feature: { rbac: { active: true, deny: true } } })
      await assert.rejects(async () => {
        for await (const _item of denied.AccountLookup().stream('list')) { }
      }, (err: any) => 'rbac_denied' === err.code)
    }
  })

  test('stream-ctrl', async () => {
    const explain: any = {}
    const ctrl: any = { explain }
    for await (const _item of DingconnectSDK.test().AccountLookup().stream('list', undefined, { ctrl })) { }
    assert.deepStrictEqual(Object.keys(ctrl), ['explain'])
    assert(explain === ctrl.explain && 0 < Object.keys(explain).length)
  })

  test('unexpected', async () => {
    const hook = new FailHook()
    const client = new DingconnectSDK({ feature: { test: { active: true } }, extend: [hook] })
    await assert.rejects(client.AccountLookup().list(), /hook failed/)
    assert(0 < hook.unexpected)

    const fired = hook.unexpected
    assert.strictEqual(await client.AccountLookup().list(undefined, { throw: false }), undefined)
    assert(fired < hook.unexpected)
  })

  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = DingconnectSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.AccountLookup().list({"account_number":1} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.DINGCONNECT_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'account_lookup.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"AccountNumberNormalized":{"a":true,"h":"Account Number Normalized","n":"AccountNumberNormalized","r":false,"sh":"We attempt to normalize phone numbers following the public telecommunication numbering plan <a href=\"https://en.wikipedia.org/wiki/E.164\" target=\"_blank\">E.164</a>, if we succeed the normalized number will be returned in this field formatt…","t":"`$STRING`","key$":"AccountNumberNormalized","index$":0},"CountryIso":{"a":true,"h":"Country Iso","n":"CountryIso","r":false,"sh":"The country of the account number","t":"`$STRING`","key$":"CountryIso","index$":1},"ErrorCodes":{"a":true,"h":"Error Codes","n":"ErrorCodes","r":true,"t":"`$ARRAY`","key$":"ErrorCodes","index$":2},"Items":{"a":true,"h":"Items","n":"Items","r":true,"sh":"This will contain provider information associated to the account number.","t":"`$ARRAY`","key$":"Items","index$":3},"ResultCode":{"a":true,"fo":"int32","h":"Result Code","n":"ResultCode","r":true,"t":"`$INTEGER`","key$":"ResultCode","index$":4}},"name":"account_lookup","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /api/V1/GetAccountLookup","source":"swagger2","version":2},"g":{"header":[{"a":true,"k":"header","n":"x_correlation_id","or":"X-Correlation-Id","r":false,"t":"`$STRING`","index$":0}],"query":[{"a":true,"ex":"","k":"query","n":"account_number","or":"accountNumber","r":false,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/api/V1/GetAccountLookup","q":{},"r":{},"rs":{"alternatives":[{"kind":"json","media":"text/json"}],"kind":"json","media":"application/json"},"s":[{"lit":"api"},{"lit":"V1"},{"lit":"GetAccountLookup"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[]},"key$":"account_lookup","name__orig":"account_lookup","Name":"AccountLookup","name_":"account_lookup","name-":"account-lookup","NAME":"ACCOUNT_LOOKUP","index$":0}, {"active":true,"entity":"account_lookup","key$":"BasicAccountLookupFlow","kind":"basic","name":"BasicAccountLookupFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"account_lookup_ref01"}}],"index$":0}]}, 'AccountLookup', {"GET /api/V1/GetAccountLookup":{"protocol":"http","parameters":[{"in":"query","name":"accountNumber","description":"For phone number based products, the account number should be in international phone number format.","type":"string","default":"","index$":0},{"in":"header","name":"X-Correlation-Id","description":"Correlates HTTP requests between a client and server","type":"String","index$":1}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let account_lookup_ref01_data = Object.values(setup.data.existing.account_lookup)[0] as any

    // LIST
    const account_lookup_ref01_ent = client.AccountLookup()
    const account_lookup_ref01_match: any = {}

    const account_lookup_ref01_list = (await account_lookup_ref01_ent.list(account_lookup_ref01_match)).map((e: any) => e.data())


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
      '../../../../.sdk/test/entity/account_lookup/AccountLookupTestData.json')

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
    ['account_lookup01','account_lookup02','account_lookup03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'DINGCONNECT_TEST_ACCOUNT_LOOKUP_ENTID': idmap,
    'DINGCONNECT_TEST_LIVE': 'FALSE',
    'DINGCONNECT_TEST_EXPLAIN': 'FALSE',
    'DINGCONNECT_APIKEY': '',
  })

  idmap = env['DINGCONNECT_TEST_ACCOUNT_LOOKUP_ENTID']

  const live = 'TRUE' === env.DINGCONNECT_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['DINGCONNECT_TEST_ACCOUNT_LOOKUP_ENTID']
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
  
