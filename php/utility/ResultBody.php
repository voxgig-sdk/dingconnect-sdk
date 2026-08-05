<?php
declare(strict_types=1);

// Dingconnect SDK utility: result_body

class DingconnectResultBody
{
    public static function call(DingconnectContext $ctx): ?DingconnectResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
