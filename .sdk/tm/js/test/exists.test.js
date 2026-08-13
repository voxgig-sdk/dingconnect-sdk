
const { test, describe } = require('node:test')
const { equal } = require('node:assert')


const { DingconnectSDK } = require('..')


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await DingconnectSDK.test()
    equal(null !== testsdk, true)
  })

})
