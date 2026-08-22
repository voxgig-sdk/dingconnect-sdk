import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { ListTransferRecord, ListTransferRecordCreateData } from '../DingconnectTypes';
declare class ListTransferRecordEntity extends DingconnectEntityBase<ListTransferRecord> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: ListTransferRecordEntity): ListTransferRecordEntity;
    create(this: any, reqdata?: ListTransferRecordCreateData, ctrl?: Control): Promise<ListTransferRecordEntity>;
}
export { ListTransferRecordEntity };
