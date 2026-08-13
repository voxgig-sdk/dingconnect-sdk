

class DingconnectError extends Error {

  isDingconnectError = true

  sdk = 'Dingconnect'

  constructor(code, msg, ctx) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

module.exports = {
  DingconnectError
}

