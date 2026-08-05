<?php
declare(strict_types=1);

// Dingconnect SDK utility: prepare_body

class DingconnectPrepareBody
{
    public static function call(DingconnectContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
