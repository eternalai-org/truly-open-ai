/** Agent token */

import { AgentInfo } from "./agent";

export interface IAgentToken extends AgentInfo {
  meme: IToken;
  token_network_id: number;
  token_network_name: string;
  active_latest_time?: string;
  base_token_symbol?: string;
  thumbnail?: string;
}

export interface IToken {
  id: number;
  created_at: Date;
  updated_at: Date;
  owner_address: string;
  owner: null;
  token_address: string;
  name: string;
  description: string;
  ticker: string;
  image: string;
  twitter: string;
  telegram: string;
  website: string;
  tx_hash: string;
  status: string;
  reply_count: number;
  last_reply: null;
  pool: string;
  uniswap_pool: string;
  supply: string;
  price: string;
  price_usd: string;
  price_last24h: string;
  volume_last24h: string;
  total_volume: string;
  base_token_symbol: string;
  percent: number;
  decimals: number;
  pool_fee: number;
  market_cap: string;
  total_balance: string;
  system_prompt: string;
  holders: number;
  shared: number;
  agent_info: null;
  latest_twitter_post: null;
  trade_url: string;
  network_id: string;
}

/** Get Agent tokens params */

export enum OrderOption {
  Desc = "0",
  Asc = "1",
}

export enum SortOption {
  MarketCap = "meme_market_cap",
  Percent = "meme_percent",
  LastReply = "reply_latest_time",
  Price = "meme_price",
  Volume24h = "meme_volume_last24h",
  CreatedAt = "created_at",
}

export interface IGetAgentTokensParams {
  limit?: number;
  page?: number;
  sort_type?: SortOption;
  sort_col?: OrderOption;
  search?: string;
  chain?: number;
}
