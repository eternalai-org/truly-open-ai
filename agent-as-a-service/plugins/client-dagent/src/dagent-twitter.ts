import {
  dagentLogger,
  BaseDagent,
  InitAgent,
} from "@eternal-dagent/core";
import { TwitterAgentClient } from "@eternal-dagent/plugin-twitter";
import { ILinkTwitterParams } from "./types";

class DagentTwitter extends BaseDagent {
  public api: TwitterAgentClient;

  constructor(params: InitAgent) {
    super(params);
    this.api = new TwitterAgentClient({
      endpoint: this.env.ETERNAL_AI_URL,
    });
  }

  init = async () => {
    const accessToken = await this.configAccessToken();
    if (!this.accessToken) {
      throw new Error("Access token is not loaded.");
    }
    this.api.setAuthToken(accessToken || "");
  };

  linkTwitter = async (params: ILinkTwitterParams) => {
    const url = this.api.getTwitterOauthUrl({
      agent_id: params.agent_id,
      callback_url: params.callback_url,
      twitter_client_id: params.twitter_client_id,
      twitter_oauth_url: params.twitter_oauth_url,
      wallet_address: this.signer.address,
    });
    dagentLogger.warn(`Twitter OAuth URL: ${url}`);
    return url;
  };

  setupMissions = async (agentId: string) => {
    await this.api.setupMissions({
      agentId: agentId,
      missions: this.dagentCharacter?.twitterMissions || [],
    });
  };
}

export default DagentTwitter;