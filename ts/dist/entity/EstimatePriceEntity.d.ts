import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { EstimatePrice, EstimatePriceCreateData } from '../DingconnectTypes';
declare class EstimatePriceEntity extends DingconnectEntityBase<EstimatePrice> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: EstimatePriceEntity): EstimatePriceEntity;
    create(this: any, reqdata?: EstimatePriceCreateData, ctrl?: Control): Promise<EstimatePriceEntity>;
}
export { EstimatePriceEntity };
