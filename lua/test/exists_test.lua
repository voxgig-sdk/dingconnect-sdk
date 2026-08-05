-- Dingconnect SDK exists test

local sdk = require("dingconnect_sdk")

describe("DingconnectSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
