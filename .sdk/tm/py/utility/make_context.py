# Dingconnect SDK utility: make_context

from core.context import DingconnectContext


def make_context_util(ctxmap, basectx):
    return DingconnectContext(ctxmap, basectx)
