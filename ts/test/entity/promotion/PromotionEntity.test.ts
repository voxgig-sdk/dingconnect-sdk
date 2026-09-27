

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { DingconnectSDK, BaseFeature, stdutil } from '../../..'

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


describe('PromotionEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when DINGCONNECT_TEST_LIVE=TRUE.
  afterEach(liveDelay('DINGCONNECT_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = DingconnectSDK.test()
    const ent = testsdk.Promotion()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.DINGCONNECT_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'promotion.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"ErrorCodes":{"a":true,"h":"Error Codes","n":"ErrorCodes","r":true,"t":"`$ARRAY`","key$":"ErrorCodes","index$":0},"Items":{"a":true,"h":"Items","n":"Items","r":true,"sh":"List of available promotions","t":"`$ARRAY`","key$":"Items","index$":1},"ResultCode":{"a":true,"fo":"int32","h":"Result Code","n":"ResultCode","r":true,"t":"`$INTEGER`","key$":"ResultCode","index$":2}},"name":"promotion","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /api/V1/GetPromotions","source":"swagger2","version":2},"g":{"header":[{"a":true,"k":"header","n":"x_correlation_id","or":"x_correlation_id","r":false,"t":"`$STRING`","index$":0}],"query":[{"a":true,"k":"query","n":"account_number","or":"account_number","r":false,"t":"`$INTEGER`","index$":0},{"a":true,"k":"query","n":"country_iso","or":"country_iso","r":false,"t":"`$ANY`","index$":1},{"a":true,"k":"query","n":"provider_code","or":"provider_code","r":false,"t":"`$ANY`","index$":2}]},"k":"http","m":"GET","o":"/api/V1/GetPromotions","q":{"exist":["account_number","country_iso","provider_code","x_correlation_id"]},"r":{},"s":[{"lit":"api"},{"lit":"V1"},{"lit":"GetPromotions"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[]},"key$":"promotion","name__orig":"promotion","Name":"Promotion","name_":"promotion","name-":"promotion","NAME":"PROMOTION","index$":11}, {"active":true,"entity":"promotion","key$":"BasicPromotionFlow","kind":"basic","name":"BasicPromotionFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"promotion_ref01"}}],"index$":0}]}, 'Promotion', {"GET /api/V1/GetPromotions":{"protocol":"http","parameters":[{"in":"query","name":"countryIsos","description":"Filter the list to promotions for countries with the given ISOs.","type":"array","items":{"type":"string"},"collectionFormat":"multi","index$":0},{"in":"query","name":"providerCodes","description":"Filter the list to promotions on products supplied by providers with the submitted provider codes.","type":"array","items":{"type":"string"},"collectionFormat":"multi","index$":1},{"in":"query","name":"accountNumber","description":"Filter the list to promotions on products that are valid for the submitted account number. For phone number based products, the account number should be in international phone number format.","type":"string","default":"","index$":2},{"in":"header","name":"X-Correlation-Id","description":"Correlates HTTP requests between a client and server","type":"String","index$":3}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let promotion_ref01_data = Object.values(setup.data.existing.promotion)[0] as any

    // LIST
    const promotion_ref01_ent = client.Promotion()
    const promotion_ref01_match: any = {}

    const promotion_ref01_list = (await promotion_ref01_ent.list(promotion_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/promotion/PromotionTestData.json')

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
    ['promotion01','promotion02','promotion03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'DINGCONNECT_TEST_PROMOTION_ENTID': idmap,
    'DINGCONNECT_TEST_LIVE': 'FALSE',
    'DINGCONNECT_TEST_EXPLAIN': 'FALSE',
    'DINGCONNECT_APIKEY': '',
  })

  idmap = env['DINGCONNECT_TEST_PROMOTION_ENTID']

  const live = 'TRUE' === env.DINGCONNECT_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['DINGCONNECT_TEST_PROMOTION_ENTID']
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
  
