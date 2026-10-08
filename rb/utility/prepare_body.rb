# Dingconnect SDK utility: prepare_body
require_relative 'media'
module DingconnectUtilities
  PrepareBody = ->(ctx) {
    return nil unless ctx.op.input == "data"
    return DingconnectUtilities.raw_body(ctx.reqdata) if DingconnectUtilities.raw_request?(ctx.point)
    ctx.utility.transform_request.call(ctx)
  }
end
