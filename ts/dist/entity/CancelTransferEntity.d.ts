import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { CancelTransfer, CancelTransferCreateData } from '../DingconnectTypes';
declare class CancelTransferEntity extends DingconnectEntityBase<CancelTransfer> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: CancelTransferEntity): CancelTransferEntity;
    create(this: any, reqdata?: CancelTransferCreateData, ctrl?: Control): Promise<CancelTransfer>;
}
export { CancelTransferEntity };
