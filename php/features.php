<?php
declare(strict_types=1);

// Dingconnect SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class DingconnectFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new DingconnectBaseFeature();
            case "test":
                return new DingconnectTestFeature();
            default:
                return new DingconnectBaseFeature();
        }
    }
}
