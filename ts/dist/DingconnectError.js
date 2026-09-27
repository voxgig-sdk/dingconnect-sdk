"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.DingconnectError = void 0;
class DingconnectError extends Error {
    isDingconnectError = true;
    sdk = 'Dingconnect';
    code;
    ctx;
    status = -1;
    // `err.notFound` rather than a magic number at every call site.
    get notFound() { return 404 === this.status; }
    constructor(code, msg, ctx) {
        super(msg);
        this.code = code;
        this.ctx = ctx;
    }
}
exports.DingconnectError = DingconnectError;
//# sourceMappingURL=DingconnectError.js.map