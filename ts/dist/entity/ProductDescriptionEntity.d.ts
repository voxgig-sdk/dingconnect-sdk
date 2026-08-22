import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { ProductDescription, ProductDescriptionListMatch } from '../DingconnectTypes';
declare class ProductDescriptionEntity extends DingconnectEntityBase<ProductDescription> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: ProductDescriptionEntity): ProductDescriptionEntity;
    list(this: any, reqmatch?: ProductDescriptionListMatch, ctrl?: Control): Promise<ProductDescriptionEntity[]>;
}
export { ProductDescriptionEntity };
