import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { Product, ProductListMatch } from '../DingconnectTypes';
declare class ProductEntity extends DingconnectEntityBase<Product> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: ProductEntity): ProductEntity;
    list(this: any, reqmatch?: ProductListMatch, ctrl?: Control): Promise<ProductEntity[]>;
}
export { ProductEntity };
