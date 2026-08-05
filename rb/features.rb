# Dingconnect SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module DingconnectFeatures
  def self.make_feature(name)
    case name
    when "base"
      DingconnectBaseFeature.new
    when "test"
      DingconnectTestFeature.new
    else
      DingconnectBaseFeature.new
    end
  end
end
