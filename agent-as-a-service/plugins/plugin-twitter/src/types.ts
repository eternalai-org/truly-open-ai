import { AgentSnapshotMission } from "@eternalai-dagent/core";

/** Twitter params */
export interface IConnectTwitterParams {
  agent_id: string;
  twitter_client_id: string;
  twitter_client_secret: string;
}

/** Get Twitter Oauth Url params */
export interface IGetTwitterOauthParams {
  agent_id: string;
  callback_url: string;
  twitter_oauth_url: string;
  twitter_client_id: string;
  wallet_address: string;
}

/** Twitter missions */

export interface ITwitterConfigMissionParams {
  agent_id: string;
  missions: AgentSnapshotMission[];
}
