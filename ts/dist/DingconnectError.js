"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.DingconnectError = void 0;
class DingconnectError extends Error {
    isDingconnectError = true;
    sdk = 'Dingconnect';
    code;
    ctx;
    constructor(code, msg, ctx) {
        super(msg);
        this.code = code;
        this.ctx = ctx;
    }
}
exports.DingconnectError = DingconnectError;
//# sourceMappingURL=DingconnectError.js.map