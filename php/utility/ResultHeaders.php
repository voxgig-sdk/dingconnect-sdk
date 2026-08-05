<?php
declare(strict_types=1);

// Dingconnect SDK utility: result_headers

class DingconnectResultHeaders
{
    public static function call(DingconnectContext $ctx): ?DingconnectResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
