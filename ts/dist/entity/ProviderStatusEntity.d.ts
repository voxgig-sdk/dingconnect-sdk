import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { ProviderStatus, ProviderStatusListMatch } from '../DingconnectTypes';
declare class ProviderStatusEntity extends DingconnectEntityBase<ProviderStatus> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: ProviderStatusEntity): ProviderStatusEntity;
    list(this: any, reqmatch?: ProviderStatusListMatch, ctrl?: Control): Promise<ProviderStatus[]>;
}
export { ProviderStatusEntity };
