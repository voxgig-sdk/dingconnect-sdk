import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { ErrorCodeDescription, ErrorCodeDescriptionListMatch } from '../DingconnectTypes';
declare class ErrorCodeDescriptionEntity extends DingconnectEntityBase<ErrorCodeDescription> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: ErrorCodeDescriptionEntity): ErrorCodeDescriptionEntity;
    list(this: any, reqmatch?: ErrorCodeDescriptionListMatch, ctrl?: Control): Promise<ErrorCodeDescriptionEntity[]>;
}
export { ErrorCodeDescriptionEntity };
