
import { Context } from './Context'


class DingconnectError extends Error {

  isDingconnectError = true

  sdk = 'Dingconnect'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  DingconnectError
}

