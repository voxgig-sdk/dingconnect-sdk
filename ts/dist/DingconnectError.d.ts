import { Context } from './Context';
declare class DingconnectError extends Error {
    isDingconnectError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    constructor(code: string, msg: string, ctx: Context);
}
export { DingconnectError };
