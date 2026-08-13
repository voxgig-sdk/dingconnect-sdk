import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { PromotionDescription, PromotionDescriptionListMatch } from '../DingconnectTypes';
declare class PromotionDescriptionEntity extends DingconnectEntityBase<PromotionDescription> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: PromotionDescriptionEntity): PromotionDescriptionEntity;
    list(this: any, reqmatch?: PromotionDescriptionListMatch, ctrl?: Control): Promise<PromotionDescription[]>;
}
export { PromotionDescriptionEntity };
