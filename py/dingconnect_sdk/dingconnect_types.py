# Typed models for the Dingconnect SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class AccountLookupRequired(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class AccountLookup(AccountLookupRequired, total=False):
    AccountNumberNormalized: str
    CountryIso: str


class AccountLookupListMatch(TypedDict, total=False):
    account_number: int


class BalanceRequired(TypedDict):
    Code: str


class Balance(BalanceRequired, total=False):
    Context: str


class BalanceListMatch(TypedDict, total=False):
    Code: str
    Context: str


class CancelTransfer(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class CancelTransferCreateData(TypedDict):
    cancellation_request: list
    ErrorCodes: list
    Items: list
    ResultCode: int


class Country(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class CountryListMatch(TypedDict, total=False):
    ErrorCodes: list
    Items: list
    ResultCode: int


class Currency(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class CurrencyListMatch(TypedDict, total=False):
    ErrorCodes: list
    Items: list
    ResultCode: int


class ErrorCodeDescription(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class ErrorCodeDescriptionListMatch(TypedDict, total=False):
    ErrorCodes: list
    Items: list
    ResultCode: int


class EstimatePrice(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class EstimatePriceCreateData(TypedDict):
    requested_estimation: list
    ErrorCodes: list
    Items: list
    ResultCode: int


class ListTransferRecord(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int
    ThereAreMoreItems: bool


class ListTransferRecordCreateData(TypedDict):
    request: dict
    ErrorCodes: list
    Items: list
    ResultCode: int
    ThereAreMoreItems: bool


class LookupBill(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class LookupBillCreateData(TypedDict):
    request: dict
    ErrorCodes: list
    Items: list
    ResultCode: int


class Product(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class ProductListMatch(TypedDict, total=False):
    account_number: int
    benefit: Any
    country_iso: Any
    provider_code: Any
    region_code: Any
    sku_code: Any


class ProductDescription(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class ProductDescriptionListMatch(TypedDict, total=False):
    language_code: Any
    sku_code: Any


class Promotion(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class PromotionListMatch(TypedDict, total=False):
    account_number: int
    country_iso: Any
    provider_code: Any


class PromotionDescription(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class PromotionDescriptionListMatch(TypedDict, total=False):
    language_code: Any


class Provider(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class ProviderListMatch(TypedDict, total=False):
    account_number: int
    country_iso: Any
    provider_code: Any
    region_code: Any


class ProviderStatus(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class ProviderStatusListMatch(TypedDict, total=False):
    provider_code: Any


class Region(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class RegionListMatch(TypedDict, total=False):
    country_iso: Any


class SendTransfer(TypedDict):
    ErrorCodes: list
    ResultCode: int
    TransferRecord: dict


class SendTransferCreateData(TypedDict):
    request: dict
    ErrorCodes: list
    ResultCode: int
    TransferRecord: dict
