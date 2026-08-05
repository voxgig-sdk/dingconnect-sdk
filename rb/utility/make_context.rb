# Dingconnect SDK utility: make_context
require_relative '../core/context'
module DingconnectUtilities
  MakeContext = ->(ctxmap, basectx) {
    DingconnectContext.new(ctxmap, basectx)
  }
end
