import { DingconnectEntityBase } from '../DingconnectEntityBase';
import type { DingconnectSDK } from '../DingconnectSDK';
import type { Control } from '../types';
import type { Country, CountryListMatch } from '../DingconnectTypes';
declare class CountryEntity extends DingconnectEntityBase<Country> {
    constructor(client: DingconnectSDK, entopts: any);
    make(this: CountryEntity): CountryEntity;
    list(this: any, reqmatch?: CountryListMatch, ctrl?: Control): Promise<CountryEntity[]>;
}
export { CountryEntity };
