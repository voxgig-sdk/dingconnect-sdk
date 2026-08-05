# Dingconnect SDK exists test

require "minitest/autorun"
require_relative "../Dingconnect_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = DingconnectSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
