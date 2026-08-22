import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { Currency, CurrencyListMatch } from '../DingconnectTypes';
declare class CurrencyEntity extends DingconnectEntityBase<Currency> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: CurrencyEntity): CurrencyEntity;
    list(this: any, reqmatch?: CurrencyListMatch, ctrl?: Control): Promise<CurrencyEntity[]>;
}
export { CurrencyEntity };
