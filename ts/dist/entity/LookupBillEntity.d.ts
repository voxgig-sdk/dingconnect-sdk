import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { LookupBill, LookupBillCreateData } from '../DingconnectTypes';
declare class LookupBillEntity extends DingconnectEntityBase<LookupBill> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: LookupBillEntity): LookupBillEntity;
    create(this: any, reqdata?: LookupBillCreateData, ctrl?: Control): Promise<LookupBill>;
}
export { LookupBillEntity };
