import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { AccountLookup, AccountLookupListMatch } from '../DingconnectTypes';
declare class AccountLookupEntity extends DingconnectEntityBase<AccountLookup> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: AccountLookupEntity): AccountLookupEntity;
    list(this: any, reqmatch?: AccountLookupListMatch, ctrl?: Control): Promise<AccountLookup[]>;
}
export { AccountLookupEntity };
