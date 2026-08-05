<?php
declare(strict_types=1);

// Dingconnect SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class DingconnectMakeContext
{
    public static function call(array $ctxmap, ?DingconnectContext $basectx): DingconnectContext
    {
        return new DingconnectContext($ctxmap, $basectx);
    }
}
