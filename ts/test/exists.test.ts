
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { DingconnectSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await DingconnectSDK.test()
    equal(null !== testsdk, true)
  })

})
