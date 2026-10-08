# Typed models for the Dingconnect SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Field/param types come from the
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
    account_number: str


class BalanceRequired(TypedDict):
    Code: str


class Balance(BalanceRequired, total=False):
    Context: str


class BalanceListMatch(TypedDict, total=False):
    Code: str
    Context: str


class CancelTransferRequired(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class CancelTransfer(CancelTransferRequired, total=False):
    cancellations: list


class CancelTransferCreateDataRequired(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class CancelTransferCreateData(CancelTransferCreateDataRequired, total=False):
    cancellations: list


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


class EstimatePriceRequired(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class EstimatePrice(EstimatePriceRequired, total=False):
    estimations: list


class EstimatePriceCreateDataRequired(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class EstimatePriceCreateData(EstimatePriceCreateDataRequired, total=False):
    estimations: list


class ListTransferRecordRequired(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int
    Take: int
    ThereAreMoreItems: bool


class ListTransferRecord(ListTransferRecordRequired, total=False):
    AccountNumber: str
    DistributorRef: str
    Skip: int
    TransferRef: str


class ListTransferRecordCreateDataRequired(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int
    Take: int
    ThereAreMoreItems: bool


class ListTransferRecordCreateData(ListTransferRecordCreateDataRequired, total=False):
    AccountNumber: str
    DistributorRef: str
    Skip: int
    TransferRef: str


class LookupBillRequired(TypedDict):
    AccountNumber: str
    ErrorCodes: list
    Items: list
    ResultCode: int
    SkuCode: str


class LookupBill(LookupBillRequired, total=False):
    Settings: list


class LookupBillCreateDataRequired(TypedDict):
    AccountNumber: str
    ErrorCodes: list
    Items: list
    ResultCode: int
    SkuCode: str


class LookupBillCreateData(LookupBillCreateDataRequired, total=False):
    Settings: list


class Product(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class ProductListMatch(TypedDict, total=False):
    account_number: str
    benefit: list
    country_iso: list
    provider_code: list
    region_code: list
    sku_code: list


class ProductDescription(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class ProductDescriptionListMatch(TypedDict, total=False):
    language_code: list
    sku_code: list


class Promotion(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class PromotionListMatch(TypedDict, total=False):
    account_number: str
    country_iso: list
    provider_code: list


class PromotionDescription(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class PromotionDescriptionListMatch(TypedDict, total=False):
    language_code: list


class Provider(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class ProviderListMatch(TypedDict, total=False):
    account_number: str
    country_iso: list
    provider_code: list
    region_code: list


class ProviderStatus(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class ProviderStatusListMatch(TypedDict, total=False):
    provider_code: list


class Region(TypedDict):
    ErrorCodes: list
    Items: list
    ResultCode: int


class RegionListMatch(TypedDict, total=False):
    country_iso: list


class SendTransferRequired(TypedDict):
    AccountNumber: str
    DistributorRef: str
    ErrorCodes: list
    ResultCode: int
    SendValue: float
    SkuCode: str
    TransferRecord: dict
    ValidateOnly: bool


class SendTransfer(SendTransferRequired, total=False):
    BillRef: str
    SendCurrencyIso: str
    Settings: list


class SendTransferCreateDataRequired(TypedDict):
    AccountNumber: str
    DistributorRef: str
    ErrorCodes: list
    ResultCode: int
    SendValue: float
    SkuCode: str
    TransferRecord: dict
    ValidateOnly: bool


class SendTransferCreateData(SendTransferCreateDataRequired, total=False):
    BillRef: str
    SendCurrencyIso: str
    Settings: list
