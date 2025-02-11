import BaseAPI from "../client/base";
import { Store } from "../types";
import logger from "../logger";

export class StoreAgent extends BaseAPI {
    createApp = async (params: Store) => {
        const app = (await this.api.post(`/agent-store/save`, {
            ...params.info
        })) as any;

        const appId = app?.id;

        for (let i = 0; i < params.missions.length; i++) {
            const mission = params.missions[i];
            const _params = {
                ...mission,
                tool_list: `[${mission.tool_list.map(tool => `${JSON.stringify(tool)}`)}]`
            }
            await this.api.post(`/agent-store/${appId}/mission`, {
                ..._params
            })
        }

        logger.info(`App created with id: ${appId}`);
        return appId;
    }

}