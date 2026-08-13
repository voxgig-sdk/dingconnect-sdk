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
    AccountNumberNormalized: str
    CountryIso: str
    ErrorCodes: list
    Items: list
    ResultCode: int


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
    ErrorCodes: list
    Items: list
    ResultCode: int


class ListTransferRecord(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int
    ThereAreMoreItems: bool


class ListTransferRecordCreateData(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int
    ThereAreMoreItems: bool


class LookupBill(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class LookupBillCreateData(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class Product(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class ProductListMatch(TypedDict, total=False):
    ErrorCodes: list
    Items: list
    ResultCode: int


class ProductDescription(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class ProductDescriptionListMatch(TypedDict, total=False):
    ErrorCodes: list
    Items: list
    ResultCode: int


class Promotion(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class PromotionListMatch(TypedDict, total=False):
    ErrorCodes: list
    Items: list
    ResultCode: int


class PromotionDescription(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class PromotionDescriptionListMatch(TypedDict, total=False):
    ErrorCodes: list
    Items: list
    ResultCode: int


class Provider(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class ProviderListMatch(TypedDict, total=False):
    ErrorCodes: list
    Items: list
    ResultCode: int


class ProviderStatus(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class ProviderStatusListMatch(TypedDict, total=False):
    ErrorCodes: list
    Items: list
    ResultCode: int


class Region(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class RegionListMatch(TypedDict, total=False):
    ErrorCodes: list
    Items: list
    ResultCode: int


class SendTransfer(TypedDict):
    ErrorCodes: list
    ResultCode: int
    TransferRecord: dict


class SendTransferCreateData(TypedDict):
    ErrorCodes: list
    ResultCode: int
    TransferRecord: dict
