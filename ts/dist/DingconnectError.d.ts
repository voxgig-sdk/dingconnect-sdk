import { Context } from './Context';
declare class DingconnectError extends Error {
    isDingconnectError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { DingconnectError };
