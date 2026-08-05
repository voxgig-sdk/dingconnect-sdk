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


class AccountLookup(TypedDict, total=False):
    account_number_normalized: str
    country_iso: str
    error_code: list
    item: list
    result_code: int


class AccountLookupListMatch(TypedDict, total=False):
    account_number_normalized: str
    country_iso: str
    error_code: list
    item: list
    result_code: int


class Balance(TypedDict, total=False):
    code: str
    context: str


class BalanceListMatch(TypedDict, total=False):
    code: str
    context: str


class CancelResult(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class CancelResultCreateData(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class Country(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class CountryListMatch(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class Currency(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class CurrencyListMatch(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class ErrorCodeDescription(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class ErrorCodeDescriptionListMatch(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class Estimate(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class EstimateCreateData(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class Product(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class ProductListMatch(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class ProductDescription(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class ProductDescriptionListMatch(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class Promotion(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class PromotionListMatch(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class PromotionDescription(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class PromotionDescriptionListMatch(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class Provider(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class ProviderListMatch(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class ProviderStatus(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class ProviderStatusListMatch(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class Region(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class RegionListMatch(TypedDict, total=False):
    error_code: list
    item: list
    result_code: int


class SendTransferRequired(TypedDict):
    account_number: str
    distributor_ref: str
    send_value: float
    sku_code: str


class SendTransfer(SendTransferRequired, total=False):
    error_code: list
    result_code: int
    send_currency_iso: str
    setting: list
    transfer_record: dict
    validate_only: bool


class SendTransferCreateDataRequired(TypedDict):
    account_number: str
    distributor_ref: str
    send_value: float
    sku_code: str


class SendTransferCreateData(SendTransferCreateDataRequired, total=False):
    error_code: list
    result_code: int
    send_currency_iso: str
    setting: list
    transfer_record: dict
    validate_only: bool


class TransferRecordRequired(TypedDict):
    skip: int
    take: int


class TransferRecord(TransferRecordRequired, total=False):
    account_number: str
    distributor_ref: list
    ended_at_utc: str
    error_code: list
    item: list
    result_code: int
    sku_code: list
    started_at_utc: str
    there_are_more_item: bool
    transfer_ref: list


class TransferRecordCreateDataRequired(TypedDict):
    skip: int
    take: int


class TransferRecordCreateData(TransferRecordCreateDataRequired, total=False):
    account_number: str
    distributor_ref: list
    ended_at_utc: str
    error_code: list
    item: list
    result_code: int
    sku_code: list
    started_at_utc: str
    there_are_more_item: bool
    transfer_ref: list
