# Dingconnect SDK

from dingconnect_sdk.utility.voxgig_struct import voxgig_struct as vs
from dingconnect_sdk.core.utility_type import DingconnectUtility
from dingconnect_sdk.core.spec import DingconnectSpec
from dingconnect_sdk.core import helpers

# Load utility registration (populates Utility._registrar)
from dingconnect_sdk.utility import register

# Load features
from dingconnect_sdk.feature.base_feature import DingconnectBaseFeature
from dingconnect_sdk.features import _make_feature


class DingconnectSDK:

    def __init__(self, options=None):
        self.mode = "live"
        self.features = []
        self.options = None

        utility = DingconnectUtility()
        self._utility = utility

        from dingconnect_sdk.config import shared_config
        config = shared_config()

        self._rootctx = utility.make_context({
            "client": self,
            "utility": utility,
            "config": config,
            "options": options if options is not None else {},
            "shared": {},
        }, None)

        self.options = utility.make_options(self._rootctx)

        if vs.getpath(self.options, "feature.test.active") is True:
            self.mode = "test"

        self._rootctx.options = self.options

        # Add features in the resolved order (make_options puts an explicit
        # list order first, else defaults to test-first). Ordering matters: the
        # `test` feature installs the base mock transport and the transport
        # features (retry/cache/netsim/proxy/ratelimit) wrap whatever is
        # current, so `test` must be added before them to sit at the base.
        feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
        if feature_opts is not None:
            featureorder = vs.getpath(self.options, "__derived__.featureorder")
            if isinstance(featureorder, list):
                for fname in featureorder:
                    fopts = helpers.to_map(feature_opts.get(fname))
                    if fopts is not None and fopts.get("active") is True:
                        utility.feature_add(self._rootctx, _make_feature(fname))

        # Add extension features.
        extend = vs.getprop(self.options, "extend")
        if isinstance(extend, list):
            for f in extend:
                if isinstance(f, dict) or (hasattr(f, "get_name") and callable(f.get_name)):
                    utility.feature_add(self._rootctx, f)

        # Initialize features.
        for f in self.features:
            utility.feature_init(self._rootctx, f)

        utility.feature_hook(self._rootctx, "PostConstruct")

        # #BuildFeatures

    def options_map(self):
        out = vs.clone(self.options)
        if isinstance(out, dict):
            return out
        return {}

    def get_utility(self):
        return DingconnectUtility.copy(self._utility)

    def get_root_ctx(self):
        return self._rootctx

    def prepare(self, fetchargs=None):
        utility = self._utility

        if fetchargs is None:
            fetchargs = {}

        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "prepare",
            "ctrl": ctrl,
        }, self._rootctx)

        options = self.options

        path = vs.getprop(fetchargs, "path") or ""
        if not isinstance(path, str):
            path = ""

        method = vs.getprop(fetchargs, "method") or "GET"
        if not isinstance(method, str):
            method = "GET"

        params = helpers.to_map(vs.getprop(fetchargs, "params"))
        if params is None:
            params = {}
        query = helpers.to_map(vs.getprop(fetchargs, "query"))
        if query is None:
            query = {}

        headers = utility.prepare_headers(ctx)

        base = vs.getprop(options, "base") or ""
        if not isinstance(base, str):
            base = ""
        prefix = vs.getprop(options, "prefix") or ""
        if not isinstance(prefix, str):
            prefix = ""
        suffix = vs.getprop(options, "suffix") or ""
        if not isinstance(suffix, str):
            suffix = ""

        ctx.spec = DingconnectSpec({
            "base": base,
            "prefix": prefix,
            "suffix": suffix,
            "path": path,
            "method": method,
            "params": params,
            "query": query,
            "headers": headers,
            "body": vs.getprop(fetchargs, "body"),
            "step": "start",
        })

        # Merge user-provided headers.
        uh = vs.getprop(fetchargs, "headers")
        if isinstance(uh, dict):
            for k, v in uh.items():
                ctx.spec.headers[k] = v

        _, err = utility.prepare_auth(ctx)
        if err is not None:
            raise err

        fetchdef, err = utility.make_fetch_def(ctx)
        if err is not None:
            raise err

        return fetchdef

    # Raw endpoint access is operator-controllable, like every entity op.
    # Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
    # either one reaches the same endpoint.
    def direct(self, fetchargs=None):
        if not self._op_allowed("direct"):
            return self._op_denied("direct")

        return self._raw_request(fetchargs)

    # Is this raw-access op permitted by the SDK's allow.op option?
    def _op_allowed(self, op):
        allow_op = vs.getpath(self.options, "allow.op")
        return isinstance(allow_op, str) and op in allow_op

    def _op_denied(self, op):
        allow_op = vs.getpath(self.options, "allow.op")
        return {
            "ok": False,
            "err": Exception(
                "DingconnectSDK: " + op + ": operation not allowed by"
                ' SDK option allow.op value: "' + str(allow_op) + '"'),
        }

    # Ungated request path shared by direct and graphql, each of which checks
    # its own allow.op token first. Private, rather than a flag on fetchargs:
    # a caller-supplied marker would let anyone opt straight back out of the
    # gate by passing it.
    def _raw_request(self, fetchargs=None):
        utility = self._utility

        try:
            fetchdef = self.prepare(fetchargs)
        except Exception as err:
            # direct() is the raw-HTTP escape hatch: it never raises, it
            # returns a result object callers branch on via result["ok"].
            return {"ok": False, "err": err}

        if fetchargs is None:
            fetchargs = {}
        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "direct",
            "ctrl": ctrl,
        }, self._rootctx)

        url = fetchdef.get("url", "")
        fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

        if fetch_err is not None:
            return {"ok": False, "err": fetch_err}

        if fetched is None:
            return {
                "ok": False,
                "err": ctx.make_error("direct_no_response", "response: undefined"),
            }

        if isinstance(fetched, dict):
            status = helpers.to_int(vs.getprop(fetched, "status"))
            headers = vs.getprop(fetched, "headers") or {}

            # No-body responses (204, 304) and explicit zero content-length
            # must skip JSON parsing — calling json() on an empty body raises.
            content_length = None
            if isinstance(headers, dict):
                content_length = headers.get("content-length")
            no_body = status in (204, 304) or str(content_length) == "0"

            json_data = None
            if not no_body:
                jf = vs.getprop(fetched, "json")
                if callable(jf):
                    try:
                        json_data = jf()
                    except Exception:
                        # Non-JSON body (e.g. text/plain, text/html). Surface
                        # status + headers but leave data as None.
                        json_data = None

            return {
                "ok": status >= 200 and status < 300,
                "status": status,
                "headers": headers,
                "data": json_data,
            }

        return {
            "ok": False,
            "err": ctx.make_error("direct_invalid", "invalid response type"),
        }

    # Raw GraphQL access: the pressure valve that makes the generated
    # surface's deliberate omissions (per-call selection sets, typed filter
    # builders, batching, subscriptions) livable — the whole schema stays
    # reachable.
    #
    # Thin wrapper over the same prepare/fetch path direct uses, with the one
    # thing raw direct cannot do for GraphQL: a GraphQL failure rides HTTP 200
    # as a top-level `errors` array, so status alone would report a failed
    # query as ok.
    #
    # NOTE: like direct, this bypasses the feature pipeline — no retry,
    # ratelimit or paging features apply.
    def graphql(self, query, variables=None, ctrl=None):
        if not self._op_allowed("graphql"):
            return self._op_denied("graphql")

        res = self._raw_request({
            "method": "POST",
            "headers": {"content-type": "application/json"},
            "body": {"query": query, "variables": variables or {}},
            "ctrl": ctrl or {},
        })

        # Errors are read BEFORE any status check: a GraphQL parse or
        # validation failure comes back as HTTP 400 carrying the standard
        # { errors: [...] } body, and the raw path represents a non-2xx as
        # ok:False with no err — so returning early on status would discard
        # the server's own diagnostics, which are the only useful part of
        # that response.
        errors = vs.getpath(res, "data.errors")

        if isinstance(errors, list) and 0 < len(errors):
            first = errors[0] if isinstance(errors[0], dict) else {}
            msg = first.get("message") or "graphql error"
            res["ok"] = False
            res["err"] = Exception("DingconnectSDK: graphql: " + str(msg))
            res["graphql"] = errors

        return res


    def AccountLookup(self, data=None) -> "AccountLookupEntity":
        """Entity factory: client.AccountLookup().list() / client.AccountLookup().load({"id": ...})."""
        from dingconnect_sdk.entity.account_lookup_entity import AccountLookupEntity
        return AccountLookupEntity(self, data)


    def Balance(self, data=None) -> "BalanceEntity":
        """Entity factory: client.Balance().list() / client.Balance().load({"id": ...})."""
        from dingconnect_sdk.entity.balance_entity import BalanceEntity
        return BalanceEntity(self, data)


    def CancelTransfer(self, data=None) -> "CancelTransferEntity":
        """Entity factory: client.CancelTransfer().list() / client.CancelTransfer().load({"id": ...})."""
        from dingconnect_sdk.entity.cancel_transfer_entity import CancelTransferEntity
        return CancelTransferEntity(self, data)


    def Country(self, data=None) -> "CountryEntity":
        """Entity factory: client.Country().list() / client.Country().load({"id": ...})."""
        from dingconnect_sdk.entity.country_entity import CountryEntity
        return CountryEntity(self, data)


    def Currency(self, data=None) -> "CurrencyEntity":
        """Entity factory: client.Currency().list() / client.Currency().load({"id": ...})."""
        from dingconnect_sdk.entity.currency_entity import CurrencyEntity
        return CurrencyEntity(self, data)


    def ErrorCodeDescription(self, data=None) -> "ErrorCodeDescriptionEntity":
        """Entity factory: client.ErrorCodeDescription().list() / client.ErrorCodeDescription().load({"id": ...})."""
        from dingconnect_sdk.entity.error_code_description_entity import ErrorCodeDescriptionEntity
        return ErrorCodeDescriptionEntity(self, data)


    def EstimatePrice(self, data=None) -> "EstimatePriceEntity":
        """Entity factory: client.EstimatePrice().list() / client.EstimatePrice().load({"id": ...})."""
        from dingconnect_sdk.entity.estimate_price_entity import EstimatePriceEntity
        return EstimatePriceEntity(self, data)


    def ListTransferRecord(self, data=None) -> "ListTransferRecordEntity":
        """Entity factory: client.ListTransferRecord().list() / client.ListTransferRecord().load({"id": ...})."""
        from dingconnect_sdk.entity.list_transfer_record_entity import ListTransferRecordEntity
        return ListTransferRecordEntity(self, data)


    def LookupBill(self, data=None) -> "LookupBillEntity":
        """Entity factory: client.LookupBill().list() / client.LookupBill().load({"id": ...})."""
        from dingconnect_sdk.entity.lookup_bill_entity import LookupBillEntity
        return LookupBillEntity(self, data)


    def Product(self, data=None) -> "ProductEntity":
        """Entity factory: client.Product().list() / client.Product().load({"id": ...})."""
        from dingconnect_sdk.entity.product_entity import ProductEntity
        return ProductEntity(self, data)


    def ProductDescription(self, data=None) -> "ProductDescriptionEntity":
        """Entity factory: client.ProductDescription().list() / client.ProductDescription().load({"id": ...})."""
        from dingconnect_sdk.entity.product_description_entity import ProductDescriptionEntity
        return ProductDescriptionEntity(self, data)


    def Promotion(self, data=None) -> "PromotionEntity":
        """Entity factory: client.Promotion().list() / client.Promotion().load({"id": ...})."""
        from dingconnect_sdk.entity.promotion_entity import PromotionEntity
        return PromotionEntity(self, data)


    def PromotionDescription(self, data=None) -> "PromotionDescriptionEntity":
        """Entity factory: client.PromotionDescription().list() / client.PromotionDescription().load({"id": ...})."""
        from dingconnect_sdk.entity.promotion_description_entity import PromotionDescriptionEntity
        return PromotionDescriptionEntity(self, data)


    def Provider(self, data=None) -> "ProviderEntity":
        """Entity factory: client.Provider().list() / client.Provider().load({"id": ...})."""
        from dingconnect_sdk.entity.provider_entity import ProviderEntity
        return ProviderEntity(self, data)


    def ProviderStatus(self, data=None) -> "ProviderStatusEntity":
        """Entity factory: client.ProviderStatus().list() / client.ProviderStatus().load({"id": ...})."""
        from dingconnect_sdk.entity.provider_status_entity import ProviderStatusEntity
        return ProviderStatusEntity(self, data)


    def Region(self, data=None) -> "RegionEntity":
        """Entity factory: client.Region().list() / client.Region().load({"id": ...})."""
        from dingconnect_sdk.entity.region_entity import RegionEntity
        return RegionEntity(self, data)


    def SendTransfer(self, data=None) -> "SendTransferEntity":
        """Entity factory: client.SendTransfer().list() / client.SendTransfer().load({"id": ...})."""
        from dingconnect_sdk.entity.send_transfer_entity import SendTransferEntity
        return SendTransferEntity(self, data)



    @classmethod
    def test(cls, testopts=None, sdkopts=None) -> "DingconnectSDK":
        if sdkopts is None:
            sdkopts = {}
        sdkopts = vs.clone(sdkopts)
        if not isinstance(sdkopts, dict):
            sdkopts = {}

        if testopts is None:
            testopts = {}
        testopts = vs.clone(testopts)
        if not isinstance(testopts, dict):
            testopts = {}
        testopts["active"] = True

        vs.setpath(sdkopts, "feature.test", testopts)

        sdk = cls(sdkopts)
        sdk.mode = "test"

        return sdk


from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from dingconnect_sdk.entity.account_lookup_entity import AccountLookupEntity
    from dingconnect_sdk.entity.balance_entity import BalanceEntity
    from dingconnect_sdk.entity.cancel_transfer_entity import CancelTransferEntity
    from dingconnect_sdk.entity.country_entity import CountryEntity
    from dingconnect_sdk.entity.currency_entity import CurrencyEntity
    from dingconnect_sdk.entity.error_code_description_entity import ErrorCodeDescriptionEntity
    from dingconnect_sdk.entity.estimate_price_entity import EstimatePriceEntity
    from dingconnect_sdk.entity.list_transfer_record_entity import ListTransferRecordEntity
    from dingconnect_sdk.entity.lookup_bill_entity import LookupBillEntity
    from dingconnect_sdk.entity.product_entity import ProductEntity
    from dingconnect_sdk.entity.product_description_entity import ProductDescriptionEntity
    from dingconnect_sdk.entity.promotion_entity import PromotionEntity
    from dingconnect_sdk.entity.promotion_description_entity import PromotionDescriptionEntity
    from dingconnect_sdk.entity.provider_entity import ProviderEntity
    from dingconnect_sdk.entity.provider_status_entity import ProviderStatusEntity
    from dingconnect_sdk.entity.region_entity import RegionEntity
    from dingconnect_sdk.entity.send_transfer_entity import SendTransferEntity
