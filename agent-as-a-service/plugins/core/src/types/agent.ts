import {IAgentCharacter, ICharacter} from "./character";

export interface IAgent extends ICharacter {
  id: string;
  created_at: string;
  updated_at: string;
  agent_group_id: string;
  infer_fee: string;
  infer_fee_number: number;
  creator: string;
  meta_data: string;
  social_info?: Array<{ account_name: string; fee: number }>;
  contract_agent_id: string;
  agent_contract_address: string;
  token_address: string;

  open_ai_assistant_id: string;
  open_ai_assistant_model: string;
  open_ai_assistant_created_at: number;
  open_ai_assistant_name: string;
  open_ai_assistant_description: string;
  open_ai_assistant_instructions: string;
  open_ai_tools: any[];

  is_file_uploaded: boolean;
  type: number;
  minimum_required: number;

  is_enable: boolean;
  thumbnail: string;
  is_added: boolean;
  agent_balance: string;
  status: string;
  user_prompt: string;
  wakeup_interval: string;
  pump_fun_address?: string;
  create_token_mode: TokenSetupMode;
  token_image?: string;
  token_image_url?: string;
  token_chain_id?: string;

  neynar_signers?: INeynarSigners;

  agent_info: AgentInfo;
}

/** Agent token info */
export interface AgentInfo {
  id: number;
  created_at: string;
  twitter_info_id: number;
  dex_url: string;
  twitter_info?: TwitterInfo;
  agent_id: string;
  agent_contract_id: string;
  agent_contract_address: string;
  agent_name: string;
  agent_snapshot_mission: AgentSnapshotMission[];
  network_id: number;
  network_name: string;
  eth_address: string;
  tip_amount: string;
  wallet_balance: string;
  creator: string;
  token_symbol: string;
  price_usd: number;
  usd_market_cap: number;
  personality: string;
  tmp_twitter_info?: TwitterInfo;
  sol_address: string;
  mentions: number;
  x_followers: number;
  tip_eth_address: string;
  tip_btc_address: string;
  tip_sol_address: string;
  is_faucet: boolean;
  user_prompt: string;
  token_name: string;
  token_address: string;
  token_image_url: string;
  token_mode: string;
  total_supply: number;
  latest_twitter_post: null;
  is_claimed: boolean;
}

/** Twitter info */
export interface TwitterInfo {
  twitter_id: string;
  twitter_avatar: string;
  twitter_username: string;
  twitter_name: string;
  description: string;
  re_link: boolean;
}

export enum ETwitterMissionToolSet {
  DEFAULT = "default",
  REPLY_MENTIONS = "reply_mentions",
  REPLY_NON_MENTIONS = "reply_non_mentions",
  FOLLOW = "follow",
  POST = "post",
  CREATE_TOKEN = "create_token",
}

export enum EFarcasterMissionToolSet {
  DEFAULT = "farcaster_default",
  REPLY_NON_MENTIONS = "farcaster_reply_non_mentions",
  POST = "farcaster_post",
}


export enum MissionTypeEnum {
  CHAT = 0, // Standard LLM
  CHAIN_OF_THOUGHT = 1, // Advanced reasoning with Chain of Thought
}

export interface AgentSnapshotMission {
  user_prompt: string;
  interval: number;
  agent_type?: MissionTypeEnum;
  tool_set?: ETwitterMissionToolSet | EFarcasterMissionToolSet;
}

export interface TwitterAgentSnapshotMission extends AgentSnapshotMission {
  tool_set: ETwitterMissionToolSet;
}

export interface FarcasterAgentSnapshotMission extends AgentSnapshotMission {
  tool_set: EFarcasterMissionToolSet;
}

/** Farcaster info */
export interface INeynarSigners {
  id: string;
  uuid: string;
  status: string;
  fid: number;
  user_address: string;
  assistant_id: string;
}

/** Mode token */

export enum TokenSetupMode {
  CREATE_TOKEN = "auto_create",
  NO_TOKEN = "no_token",
  LINK_EXISTING = "link_existing",
  AUTO_CREATE_RUNE = "auto_create_rune",
}

export type IENV = {
  PRIVATE_KEY: string;
  ETERNAL_AI_URL: string;
  TWITTER: {
    CLIENT_ID: string;
  }
  FARCASTER: {
    CLIENT_ID: string;
  }
}

export type InitAgent = {
  dagentCharacter: IAgentCharacter;
  environment: IENV;
}

export type AgentAppMission = {
  agent_store_id: number;
  id: number;
  created_at: string;
  name: string;
  description: string;
  user_prompt: string;
  price: string;
  tool_list: string;
  icon: string;
}

export type AgentApps = {
  id: number;
  created_at: string;
  name: string;
  description: string;
  authen_url: string;
  icon: string;
  agent_store_missions: AgentAppMission[];
}

export interface AgentSnapshotMissionVer2 {
  user_prompt: string,
  interval: number,
  tool_set: string,
  agent_base_model: string,
  agent_store_mission_id?: number
}