# Dingconnect SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

DingconnectUtility.registrar = ->(u) {
  u.clean = DingconnectUtilities::Clean
  u.done = DingconnectUtilities::Done
  u.make_error = DingconnectUtilities::MakeError
  u.feature_add = DingconnectUtilities::FeatureAdd
  u.feature_hook = DingconnectUtilities::FeatureHook
  u.feature_init = DingconnectUtilities::FeatureInit
  u.fetcher = DingconnectUtilities::Fetcher
  u.make_fetch_def = DingconnectUtilities::MakeFetchDef
  u.make_context = DingconnectUtilities::MakeContext
  u.make_options = DingconnectUtilities::MakeOptions
  u.make_request = DingconnectUtilities::MakeRequest
  u.make_response = DingconnectUtilities::MakeResponse
  u.make_result = DingconnectUtilities::MakeResult
  u.make_point = DingconnectUtilities::MakePoint
  u.make_spec = DingconnectUtilities::MakeSpec
  u.make_url = DingconnectUtilities::MakeUrl
  u.param = DingconnectUtilities::Param
  u.prepare_auth = DingconnectUtilities::PrepareAuth
  u.prepare_body = DingconnectUtilities::PrepareBody
  u.prepare_headers = DingconnectUtilities::PrepareHeaders
  u.prepare_method = DingconnectUtilities::PrepareMethod
  u.prepare_params = DingconnectUtilities::PrepareParams
  u.prepare_path = DingconnectUtilities::PreparePath
  u.prepare_query = DingconnectUtilities::PrepareQuery
  u.result_basic = DingconnectUtilities::ResultBasic
  u.result_body = DingconnectUtilities::ResultBody
  u.result_headers = DingconnectUtilities::ResultHeaders
  u.transform_request = DingconnectUtilities::TransformRequest
  u.transform_response = DingconnectUtilities::TransformResponse
}
