import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { Provider, ProviderListMatch } from '../DingconnectTypes';
declare class ProviderEntity extends DingconnectEntityBase<Provider> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: ProviderEntity): ProviderEntity;
    list(this: any, reqmatch?: ProviderListMatch, ctrl?: Control): Promise<ProviderEntity[]>;
}
export { ProviderEntity };
