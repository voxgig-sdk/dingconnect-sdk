import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { SendTransfer, SendTransferCreateData } from '../DingconnectTypes';
declare class SendTransferEntity extends DingconnectEntityBase<SendTransfer> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: SendTransferEntity): SendTransferEntity;
    create(this: any, reqdata?: SendTransferCreateData, ctrl?: Control): Promise<SendTransferEntity>;
}
export { SendTransferEntity };
