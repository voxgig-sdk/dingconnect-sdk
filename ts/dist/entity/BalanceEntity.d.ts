import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { Balance, BalanceListMatch } from '../DingconnectTypes';
declare class BalanceEntity extends DingconnectEntityBase<Balance> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: BalanceEntity): BalanceEntity;
    list(this: any, reqmatch?: BalanceListMatch, ctrl?: Control): Promise<BalanceEntity[]>;
}
export { BalanceEntity };
