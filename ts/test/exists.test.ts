
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { DingconnectSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = DingconnectSDK.test()
    equal(testsdk instanceof DingconnectSDK, true,
      'DingconnectSDK.test() must return a client synchronously')
  })

})
