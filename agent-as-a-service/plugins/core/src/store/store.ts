import BaseAPI from "../client/base";
import {Store} from "../types/store";

export class StoreAgent extends BaseAPI {
    createApp = async (params: Store) => {
        const appId = await this.api.post(`/agent-store/save`, {
            ...params.info
        })

        for (let i = 0; i < params.missions.length; i++) {
            const mission = params.missions[i];
            await this.api.post(`/agent-store/${appId}/mission`, {
                ...mission
            })
        }
        return appId;
    }

}