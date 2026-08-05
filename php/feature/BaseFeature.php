<?php
declare(strict_types=1);

// Dingconnect SDK base feature

class DingconnectBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(DingconnectContext $ctx, array $options): void {}
    public function PostConstruct(DingconnectContext $ctx): void {}
    public function PostConstructEntity(DingconnectContext $ctx): void {}
    public function SetData(DingconnectContext $ctx): void {}
    public function GetData(DingconnectContext $ctx): void {}
    public function GetMatch(DingconnectContext $ctx): void {}
    public function SetMatch(DingconnectContext $ctx): void {}
    public function PrePoint(DingconnectContext $ctx): void {}
    public function PreSpec(DingconnectContext $ctx): void {}
    public function PreRequest(DingconnectContext $ctx): void {}
    public function PreResponse(DingconnectContext $ctx): void {}
    public function PreResult(DingconnectContext $ctx): void {}
    public function PreDone(DingconnectContext $ctx): void {}
    public function PreUnexpected(DingconnectContext $ctx): void {}
}
