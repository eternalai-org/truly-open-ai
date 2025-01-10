/** Chain id of Agent */
export enum AgentChainId {
  Bitcoin = "222672",
  Arbitrum = "42161",
  Base = "8453",
  Solana = "1111",
  BSC = "56",
  Symbiosis = "45762",
}

/** Chain id of Agent token */
export enum AgentTokenChainId {
  Arbitrum = "42161",
  Base = "8453",
  Solana = "1111",
  BSC = "56",
}

/**  Chain connected */
export type IChainConnected = {
  id: string;
  created_at: string;
  updated_at: string;
  chain_id: string;
  rpc: string;
  name: string;
  explorer: string;
  eai_erc20: string;
  nft_address: string;
  paymaster_address: string;
  paymaster_fee_zero: boolean;
  paymaster_token: string;
  workerhub_address: string;
  zk_sync: boolean;
  eai_native: boolean;
  formatBalance?: string;
  balance?: string;
  thumbnail?: string;
  tag?: string;
  model_ids?: string[];
  model_details?: any[];
  support_model_names?: {
    [key: string]: string;
  };
};
