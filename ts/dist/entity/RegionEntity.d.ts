import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { Region, RegionListMatch } from '../DingconnectTypes';
declare class RegionEntity extends DingconnectEntityBase<Region> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: RegionEntity): RegionEntity;
    list(this: any, reqmatch?: RegionListMatch, ctrl?: Control): Promise<Region[]>;
}
export { RegionEntity };
