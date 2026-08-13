// Typed models for the Dingconnect SDK (JSDoc typedefs).
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
// edit by hand.

/**
 * @typedef {Object} AccountLookup
 * @property {string} [AccountNumberNormalized]
 * @property {string} [CountryIso]
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} AccountLookupListMatch
 * @property {string} [AccountNumberNormalized]
 * @property {string} [CountryIso]
 * @property {Array} [ErrorCodes]
 * @property {Array} [Items]
 * @property {number} [ResultCode]
 */

/**
 * @typedef {Object} Balance
 * @property {string} Code
 * @property {string} [Context]
 */

/**
 * @typedef {Object} BalanceListMatch
 * @property {string} [Code]
 * @property {string} [Context]
 */

/**
 * @typedef {Object} CancelTransfer
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} CancelTransferCreateData
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} Country
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} CountryListMatch
 * @property {Array} [ErrorCodes]
 * @property {Array} [Items]
 * @property {number} [ResultCode]
 */

/**
 * @typedef {Object} Currency
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} CurrencyListMatch
 * @property {Array} [ErrorCodes]
 * @property {Array} [Items]
 * @property {number} [ResultCode]
 */

/**
 * @typedef {Object} ErrorCodeDescription
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} ErrorCodeDescriptionListMatch
 * @property {Array} [ErrorCodes]
 * @property {Array} [Items]
 * @property {number} [ResultCode]
 */

/**
 * @typedef {Object} EstimatePrice
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} EstimatePriceCreateData
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} ListTransferRecord
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 * @property {boolean} ThereAreMoreItems
 */

/**
 * @typedef {Object} ListTransferRecordCreateData
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 * @property {boolean} ThereAreMoreItems
 */

/**
 * @typedef {Object} LookupBill
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} LookupBillCreateData
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} Product
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} ProductListMatch
 * @property {Array} [ErrorCodes]
 * @property {Array} [Items]
 * @property {number} [ResultCode]
 */

/**
 * @typedef {Object} ProductDescription
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} ProductDescriptionListMatch
 * @property {Array} [ErrorCodes]
 * @property {Array} [Items]
 * @property {number} [ResultCode]
 */

/**
 * @typedef {Object} Promotion
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} PromotionListMatch
 * @property {Array} [ErrorCodes]
 * @property {Array} [Items]
 * @property {number} [ResultCode]
 */

/**
 * @typedef {Object} PromotionDescription
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} PromotionDescriptionListMatch
 * @property {Array} [ErrorCodes]
 * @property {Array} [Items]
 * @property {number} [ResultCode]
 */

/**
 * @typedef {Object} Provider
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} ProviderListMatch
 * @property {Array} [ErrorCodes]
 * @property {Array} [Items]
 * @property {number} [ResultCode]
 */

/**
 * @typedef {Object} ProviderStatus
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} ProviderStatusListMatch
 * @property {Array} [ErrorCodes]
 * @property {Array} [Items]
 * @property {number} [ResultCode]
 */

/**
 * @typedef {Object} Region
 * @property {Array} ErrorCodes
 * @property {Array} Items
 * @property {number} ResultCode
 */

/**
 * @typedef {Object} RegionListMatch
 * @property {Array} [ErrorCodes]
 * @property {Array} [Items]
 * @property {number} [ResultCode]
 */

/**
 * @typedef {Object} SendTransfer
 * @property {Array} ErrorCodes
 * @property {number} ResultCode
 * @property {Object} TransferRecord
 */

/**
 * @typedef {Object} SendTransferCreateData
 * @property {Array} ErrorCodes
 * @property {number} ResultCode
 * @property {Object} TransferRecord
 */

