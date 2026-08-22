import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { Promotion, PromotionListMatch } from '../DingconnectTypes';
declare class PromotionEntity extends DingconnectEntityBase<Promotion> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: PromotionEntity): PromotionEntity;
    list(this: any, reqmatch?: PromotionListMatch, ctrl?: Control): Promise<PromotionEntity[]>;
}
export { PromotionEntity };
