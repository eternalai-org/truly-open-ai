import { IChainConnected } from "@eternalai-dagent/core";
import { eternalAPIClient } from "../clients";

const ChainAPI = {
  getChainList: async (): Promise<IChainConnected[]> => {
    const res = await eternalAPIClient.get(`/api/chain-config/list`);
    return res as any;
  },
};

export default ChainAPI;
