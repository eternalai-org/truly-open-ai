import { FarcasterAgentSnapshotMission, TokenSetupMode, TwitterAgentSnapshotMission } from "./agent";
import { AgentChainId, AgentTokenChainId } from "./chain";
import { ICharacter } from "./character";

/** Authentic params */
export interface IAuthenticParams {
  endpoint: string;
  accessToken?: string;
}

/** Deploy token params */
export interface IDeployAgentTokenParams {
  agent_id: string;
  create_token_mode?: TokenSetupMode;
  chain_id: AgentChainId;
  token_chain_id: AgentTokenChainId;
  token_address?: string;
  token_name?: string;
  ticker?: string;
}

/** Get agents params */
export interface IGetAgentsParams {
  limit?: number;
  page?: number;
  chain_id?: number;
  creator?: string;
  contract_agent_id?: string;
  joined_group?: boolean;
  status?: string;
  group_id?: string;
  linked_twitter?: boolean;
}

/** Update agent params */
export interface IUpdateAgentParams extends ICharacter {
  id: string;
  user_prompt?: string;
  wakeup_interval?: number;
  wakeup_interval_unit?: string;
  social_info?: {
    account_name: string;
    fee: number;
  }[];
  thumbnail?: string;
}

export interface IGetChallengeParams {
  address: string;
  referrer?: string;
}

export interface IGetAccessTokenParams {
  signature: string;
  address: string;
  message: string;
}
/** ChatCompletionType */

export type ChatCompletionType = {
  role: string;
  content: string;
};

export type ITwitterSetupMission = {
  agentId: string;
  missions: TwitterAgentSnapshotMission[];
}

export type IFarcasterSetupMission = {
  agentId: string;
  missions: FarcasterAgentSnapshotMission[];
}