# Dingconnect SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/debug_feature'
require_relative 'feature/idempotency_feature'
require_relative 'feature/metrics_feature'
require_relative 'feature/paging_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module DingconnectFeatures
  def self.make_feature(name)
    case name
    when "base"
      DingconnectBaseFeature.new
    when "debug"
      DingconnectDebugFeature.new
    when "idempotency"
      DingconnectIdempotencyFeature.new
    when "metrics"
      DingconnectMetricsFeature.new
    when "paging"
      DingconnectPagingFeature.new
    when "ratelimit"
      DingconnectRatelimitFeature.new
    when "retry"
      DingconnectRetryFeature.new
    when "test"
      DingconnectTestFeature.new
    when "timeout"
      DingconnectTimeoutFeature.new
    else
      DingconnectBaseFeature.new
    end
  end
end
