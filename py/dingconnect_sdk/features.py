# Dingconnect SDK feature factory

from dingconnect_sdk.feature.base_feature import DingconnectBaseFeature
from dingconnect_sdk.feature.debug_feature import DingconnectDebugFeature
from dingconnect_sdk.feature.idempotency_feature import DingconnectIdempotencyFeature
from dingconnect_sdk.feature.metrics_feature import DingconnectMetricsFeature
from dingconnect_sdk.feature.paging_feature import DingconnectPagingFeature
from dingconnect_sdk.feature.ratelimit_feature import DingconnectRatelimitFeature
from dingconnect_sdk.feature.retry_feature import DingconnectRetryFeature
from dingconnect_sdk.feature.test_feature import DingconnectTestFeature
from dingconnect_sdk.feature.timeout_feature import DingconnectTimeoutFeature


_FEATURES = {
    "base": lambda: DingconnectBaseFeature(),
    "debug": lambda: DingconnectDebugFeature(),
    "idempotency": lambda: DingconnectIdempotencyFeature(),
    "metrics": lambda: DingconnectMetricsFeature(),
    "paging": lambda: DingconnectPagingFeature(),
    "ratelimit": lambda: DingconnectRatelimitFeature(),
    "retry": lambda: DingconnectRetryFeature(),
    "test": lambda: DingconnectTestFeature(),
    "timeout": lambda: DingconnectTimeoutFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
