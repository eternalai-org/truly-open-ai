import { FarcasterAgentClient } from "@eternal-dagent/plugin-farcaster";
import { BaseDagent, InitAgent } from "@eternal-dagent/core";

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

  init = async () => {
    const accessToken = await this.configAccessToken();
    if (!this.accessToken) {
      throw new Error("Access token is not loaded.");
    }
    this.api.setAuthToken(accessToken || "");
    return accessToken || "";
  };

  setupMissions = async (agentId: string) => {
    await this.api.setupMissions({
      agentId: agentId,
      missions: this.dagentCharacter?.farcasterMissions || [],
    });
  };
}

export default DagentFarcaster;