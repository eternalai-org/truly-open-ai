import {
  dagentLogger,
  BaseDagent,
  InitAgent, IGetAccessTokenParams,
} from "@eternal-ai/core";
import { TwitterAgentClient } from "@eternal-ai/plugin-twitter";
import { ILinkTwitterParams } from "./types";

class DagentTwitter extends BaseDagent {
  public api: TwitterAgentClient;

  constructor(params: InitAgent) {
    super(params);
    this.api = new TwitterAgentClient({
      endpoint: this.env.ETERNAL_AI_URL,
    });
  }

  init = async (params?: IGetAccessTokenParams) => {
    const accessToken = await this.configAccessToken(params);
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
      wallet_address: this.getSignerAddress(),
    });
    dagentLogger.warn(`Please Link Twitter OAuth via URL: ${url}\n`);
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