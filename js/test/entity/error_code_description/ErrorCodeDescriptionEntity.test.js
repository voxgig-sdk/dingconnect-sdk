
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

const Path = require('node:path')
const Fs = require('node:fs')

const { test, describe } = require('node:test')
const assert = require('node:assert')


const { DingconnectSDK, BaseFeature, stdutil, config } = require('../../..')

const {
  envOverride,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
} = require('../../utility')


describe('ErrorCodeDescriptionEntity', async () => {

  test('instance', async () => {
    const testsdk = DingconnectSDK.test()
    const ent = testsdk.ErrorCodeDescription()
    assert(null != ent)
  })


  test('basic', async () => {

    const setup = basicSetup()
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let error_code_description_ref01_data = Object.values(setup.data.existing.error_code_description)[0]

    // LIST
    const error_code_description_ref01_ent = client.ErrorCodeDescription()
    const error_code_description_ref01_match = {}

    const error_code_description_ref01_list = await error_code_description_ref01_ent.list(error_code_description_ref01_match)


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname,
      '../../../../.sdk/test/entity/error_code_description/ErrorCodeDescriptionTestData.json')

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
    ['error_code_description01','error_code_description02','error_code_description03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'DINGCONNECT_TEST_ERROR_CODE_DESCRIPTION_ENTID': idmap,
    'DINGCONNECT_TEST_LIVE': 'FALSE',
    'DINGCONNECT_TEST_EXPLAIN': 'FALSE',
    'DINGCONNECT_APIKEY': 'NONE',
  })

  idmap = env['DINGCONNECT_TEST_ERROR_CODE_DESCRIPTION_ENTID']

  if ('TRUE' === env.DINGCONNECT_TEST_LIVE) {
    client = new DingconnectSDK(merge([
      {
        apikey: env.DINGCONNECT_APIKEY,
      },
      extra
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
    now: Date.now(),
  }

  return setup
}
  
