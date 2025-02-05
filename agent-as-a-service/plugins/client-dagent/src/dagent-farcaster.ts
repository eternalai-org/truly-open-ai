import { FarcasterAgentClient } from "@eternalai-dagent/plugin-farcaster";
import {BaseDagent, IGetAccessTokenParams, InitAgent} from "@eternalai-dagent/core";

interface IDagentFarcaster {
  init: () => Promise<string>;
  setupMissions: (agentId: string) => Promise<void>;
}

class DagentFarcaster extends BaseDagent implements IDagentFarcaster {
  public api: FarcasterAgentClient;

  constructor(params: InitAgent) {
    super(params);
    this.api = new FarcasterAgentClient({
      endpoint: this.env.ETERNAL_AI_URL,
    });
  }

  init = async (params?: IGetAccessTokenParams) => {
    const _accessToken = await this.configAccessToken(params);
    if (!this.accessToken) {
      throw new Error("Access token is not loaded.");
    }
    this.api.setAuthToken(_accessToken || "");
    return _accessToken || "";
  };

  setupMissions = async (agentId: string) => {
    await this.api.setupMissions({
      agentId: agentId,
      missions: this.dagentCharacter?.farcasterMissions || [],
    });
  };
}

export default DagentFarcaster;