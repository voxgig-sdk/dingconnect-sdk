# Dingconnect SDK exists test

import pytest
from dingconnect_sdk import DingconnectSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = DingconnectSDK.test(None, None)
        assert testsdk is not None
