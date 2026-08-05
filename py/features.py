# Dingconnect SDK feature factory

from feature.base_feature import DingconnectBaseFeature
from feature.test_feature import DingconnectTestFeature


def _make_feature(name):
    features = {
        "base": lambda: DingconnectBaseFeature(),
        "test": lambda: DingconnectTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
