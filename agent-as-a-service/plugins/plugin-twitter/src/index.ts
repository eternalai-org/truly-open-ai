import { AgentClient, IAgent, IAgentClient } from "@eternal-ai/core";
import {
  IConnectTwitterParams,
  IGetTwitterOauthParams,
} from "./types";

export * from "./types";
export interface ITwitterAgentClient extends IAgentClient {
  /** Link Twitter */
  linkTwitter: (params: IConnectTwitterParams) => Promise<IAgent>;
  /** Get Twitter Oauth url */
  getTwitterOauthUrl: (params: IGetTwitterOauthParams) => string;
  /** Unlink Twitter */
  unlinkTwitter: (agentId: string) => Promise<IAgent>;
}
export class TwitterAgentClient
  extends AgentClient
  implements ITwitterAgentClient
{
  linkTwitter = async (params: IConnectTwitterParams): Promise<IAgent> => {
    try {
      const res: IAgent = await this.api.post(
        `/agent/twitter-info/${params.agent_id}`,
        {
          twitter_client_id: params.twitter_client_id,
          twitter_client_secret: params.twitter_client_secret,
        }
      );
      return res;
    } catch (e) {
      throw e;
    }
  };

  unlinkTwitter = async (agentId: string): Promise<IAgent> => {
    try {
      const res: IAgent = await this.api.post(`/agent/unlink/${agentId}`, {});
      return res;
    } catch (e) {
      throw e;
    }
  };

  getTwitterOauthUrl = (params: IGetTwitterOauthParams) => {
    const rootUrl = "https://twitter.com/i/oauth2/authorize";
    const URL = `${params?.callback_url}&address=${params?.wallet_address}&agent_id=${params.agent_id}&client_id=${params.twitter_client_id}`;
    const options = {
      redirect_uri: `${params?.twitter_oauth_url}?callback=${URL}`,
      client_id: params.twitter_client_id,
      state: "state",
      response_type: "code",
      code_challenge: "challenge",
      code_challenge_method: "plain",
      scope: [
        "offline.access",
        "tweet.read",
        "tweet.write",
        "users.read",
        "tweet.moderate.write",
        "follows.write",
        "like.write",
        "list.write",
        "block.write",
        "bookmark.write",
        "block.read",
        "follows.read",
        "bookmark.read",
        "list.read",
        "space.read",
        "like.read",
        "users.read",
        "mute.read",
      ].join(" "),
    };
    const qs = new URLSearchParams(options).toString();
    return `${rootUrl}?${qs}`;
  };
}
