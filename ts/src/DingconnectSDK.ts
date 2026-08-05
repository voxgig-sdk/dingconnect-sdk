// Dingconnect Ts SDK

import { AccountLookupEntity } from './entity/AccountLookupEntity'
import { BalanceEntity } from './entity/BalanceEntity'
import { CancelResultEntity } from './entity/CancelResultEntity'
import { CountryEntity } from './entity/CountryEntity'
import { CurrencyEntity } from './entity/CurrencyEntity'
import { ErrorCodeDescriptionEntity } from './entity/ErrorCodeDescriptionEntity'
import { EstimateEntity } from './entity/EstimateEntity'
import { ProductEntity } from './entity/ProductEntity'
import { ProductDescriptionEntity } from './entity/ProductDescriptionEntity'
import { PromotionEntity } from './entity/PromotionEntity'
import { PromotionDescriptionEntity } from './entity/PromotionDescriptionEntity'
import { ProviderEntity } from './entity/ProviderEntity'
import { ProviderStatusEntity } from './entity/ProviderStatusEntity'
import { RegionEntity } from './entity/RegionEntity'
import { SendTransferEntity } from './entity/SendTransferEntity'
import { TransferRecordEntity } from './entity/TransferRecordEntity'

export type * from './DingconnectTypes'


import { inspect } from 'node:util'

import type { Context, Feature } from './types'

import { config } from './Config'
import { DingconnectEntityBase } from './DingconnectEntityBase'
import { Utility } from './utility/Utility'


import { BaseFeature } from './feature/base/BaseFeature'


const stdutil = new Utility()


class DingconnectSDK {
  _mode: string = 'live'
  _options: any
  _utility = new Utility()
  _features: Feature[]
  _rootctx: Context

  constructor(options?: any) {

    this._rootctx = this._utility.makeContext({
      client: this,
      utility: this._utility,
      config,
      options,
      shared: new WeakMap()
    })

    this._options = this._utility.makeOptions(this._rootctx)

    const struct = this._utility.struct
    const getpath = struct.getpath

    if (true === getpath(this._options.feature, 'test.active')) {
      this._mode = 'test'
    }

    this._rootctx.options = this._options

    this._features = []

    const featureAdd = this._utility.featureAdd
    const featureInit = this._utility.featureInit

    // Add features in the resolved order (makeOptions puts an explicit
    // array order first, else defaults to test-first). Ordering matters:
    // the `test` feature installs the base mock transport and the transport
    // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is current,
    // so `test` must be added before them to sit at the base of the chain.
    const featureorder = getpath(this._options, '__derived__.featureorder') || []
    for (const fname of featureorder) {
      const fopts = this._options.feature[fname] || {}
      if (fopts.active) {
        featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname))
      }
    }

    if (null != this._options.extend) {
      for (let f of this._options.extend) {
        featureAdd(this._rootctx, f)
      }
    }

    for (let f of this._features) {
      featureInit(this._rootctx, f)
    }

    const featureHook = this._utility.featureHook
    featureHook(this._rootctx, 'PostConstruct')
  }


  options() {
    return this._utility.struct.clone(this._options)
  }


  utility() {
    return this._utility.struct.clone(this._utility)
  }


  async prepare(fetchargs?: any) {
    const utility = this._utility
    const struct = utility.struct
    const clone = struct.clone

    const {
      makeContext,
      makeFetchDef,
      prepareHeaders,
      prepareAuth,
    } = utility

    fetchargs = fetchargs || {}

    let ctx: Context = makeContext({
      opname: 'prepare',
      ctrl: fetchargs.ctrl || {},
    }, this._rootctx)

    const options = this._options

    // Build spec directly from SDK options + user-provided fetch args.
    const spec: any = {
      base: options.base,
      prefix: options.prefix,
      suffix: options.suffix,
      path: fetchargs.path || '',
      method: fetchargs.method || 'GET',
      params: fetchargs.params || {},
      query: fetchargs.query || {},
      headers: prepareHeaders(ctx),
      body: fetchargs.body,
      step: 'start',
    }

    ctx.spec = spec

    // Merge user-provided headers over SDK defaults.
    if (fetchargs.headers) {
      const uheaders = fetchargs.headers
      for (let key in uheaders) {
        spec.headers[key] = uheaders[key]
      }
    }

    // Apply SDK auth (apikey, auth prefix, etc.)
    const authResult = prepareAuth(ctx)
    if (authResult instanceof Error) {
      return authResult
    }

    return makeFetchDef(ctx)
  }


  async direct(fetchargs?: any) {
    const utility = this._utility
    const fetcher = utility.fetcher
    const makeContext = utility.makeContext

    const fetchdef = await this.prepare(fetchargs)
    if (fetchdef instanceof Error) {
      return fetchdef
    }

    let ctx: Context = makeContext({
      opname: 'direct',
      ctrl: (fetchargs || {}).ctrl || {},
    }, this._rootctx)

    try {
      const fetched = await fetcher(ctx, fetchdef.url, fetchdef)

      if (null == fetched) {
        return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') }
      }
      else if (fetched instanceof Error) {
        return { ok: false, err: fetched }
      }

      const status = fetched.status

      // No body responses (204 No Content, 304 Not Modified) and explicit
      // zero content-length must skip JSON parsing — fetched.json() would
      // throw `Unexpected end of JSON input` on an empty body.
      const headers = fetched.headers
      const contentLength = headers && 'function' === typeof headers.get
        ? headers.get('content-length')
        : (headers || {})['content-length']
      const noBody = 204 === status || 304 === status || '0' === String(contentLength)

      let json: any = undefined
      if (!noBody) {
        try {
          json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json
        }
        catch (parseErr) {
          // Body wasn't valid JSON — surface the raw response rather than
          // throwing. data stays undefined; callers can inspect status/headers.
          json = undefined
        }
      }

      return {
        ok: status >= 200 && status < 300,
        status,
        headers: fetched.headers,
        data: json,
      }
    }
    catch (err: any) {
      return { ok: false, err }
    }
  }



  // Entity access: `client.AccountLookup().list()` / `client.AccountLookup().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  AccountLookup(entopts?: Record<string, any>) {
    const self = this
    return new AccountLookupEntity(self, entopts)
  }


  // Entity access: `client.Balance().list()` / `client.Balance().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Balance(entopts?: Record<string, any>) {
    const self = this
    return new BalanceEntity(self, entopts)
  }


  // Entity access: `client.CancelResult().list()` / `client.CancelResult().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CancelResult(entopts?: Record<string, any>) {
    const self = this
    return new CancelResultEntity(self, entopts)
  }


  // Entity access: `client.Country().list()` / `client.Country().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Country(entopts?: Record<string, any>) {
    const self = this
    return new CountryEntity(self, entopts)
  }


  // Entity access: `client.Currency().list()` / `client.Currency().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Currency(entopts?: Record<string, any>) {
    const self = this
    return new CurrencyEntity(self, entopts)
  }


  // Entity access: `client.ErrorCodeDescription().list()` / `client.ErrorCodeDescription().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ErrorCodeDescription(entopts?: Record<string, any>) {
    const self = this
    return new ErrorCodeDescriptionEntity(self, entopts)
  }


  // Entity access: `client.Estimate().list()` / `client.Estimate().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Estimate(entopts?: Record<string, any>) {
    const self = this
    return new EstimateEntity(self, entopts)
  }


  // Entity access: `client.Product().list()` / `client.Product().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Product(entopts?: Record<string, any>) {
    const self = this
    return new ProductEntity(self, entopts)
  }


  // Entity access: `client.ProductDescription().list()` / `client.ProductDescription().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ProductDescription(entopts?: Record<string, any>) {
    const self = this
    return new ProductDescriptionEntity(self, entopts)
  }


  // Entity access: `client.Promotion().list()` / `client.Promotion().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Promotion(entopts?: Record<string, any>) {
    const self = this
    return new PromotionEntity(self, entopts)
  }


  // Entity access: `client.PromotionDescription().list()` / `client.PromotionDescription().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PromotionDescription(entopts?: Record<string, any>) {
    const self = this
    return new PromotionDescriptionEntity(self, entopts)
  }


  // Entity access: `client.Provider().list()` / `client.Provider().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Provider(entopts?: Record<string, any>) {
    const self = this
    return new ProviderEntity(self, entopts)
  }


  // Entity access: `client.ProviderStatus().list()` / `client.ProviderStatus().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ProviderStatus(entopts?: Record<string, any>) {
    const self = this
    return new ProviderStatusEntity(self, entopts)
  }


  // Entity access: `client.Region().list()` / `client.Region().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Region(entopts?: Record<string, any>) {
    const self = this
    return new RegionEntity(self, entopts)
  }


  // Entity access: `client.SendTransfer().list()` / `client.SendTransfer().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SendTransfer(entopts?: Record<string, any>) {
    const self = this
    return new SendTransferEntity(self, entopts)
  }


  // Entity access: `client.TransferRecord().list()` / `client.TransferRecord().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  TransferRecord(entopts?: Record<string, any>) {
    const self = this
    return new TransferRecordEntity(self, entopts)
  }




  static test(testoptsarg?: any, sdkoptsarg?: any) {
    const struct = stdutil.struct
    const setpath = struct.setpath
    const getdef = struct.getdef
    const clone = struct.clone
    const setprop = struct.setprop

    const sdkopts = getdef(clone(sdkoptsarg), {})
    const testopts = getdef(clone(testoptsarg), {})
    setprop(testopts, 'active', true)
    setpath(sdkopts, 'feature.test', testopts)

    const testsdk = new DingconnectSDK(sdkopts)
    testsdk._mode = 'test'

    return testsdk
  }


  tester(testopts?: any, sdkopts?: any) {
    return DingconnectSDK.test(testopts, sdkopts)
  }


  toJSON() {
    return { name: 'Dingconnect' }
  }

  toString() {
    return 'Dingconnect ' + this._utility.struct.jsonify(this.toJSON())
  }

  [inspect.custom]() {
    return this.toString()
  }

}




const SDK = DingconnectSDK


export {
  stdutil,
  config,

  BaseFeature,
  DingconnectEntityBase,

  DingconnectSDK,
  SDK,
}


