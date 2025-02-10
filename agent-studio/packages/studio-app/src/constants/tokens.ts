import { ChainId } from "./networks";

type IToken = {
  id: string;
  title: string;
  required_amount: number;
  icon: string;
  chainId: any;
  desc?: string;
  deploy_token_amount: number;
  deploy_agent_amount: number;
  hide?: boolean;
  api_key_label?: string;
};

export const tokens: IToken[] = [
  {
    id: "base",
    title: "Base",
    required_amount: 1,
    icon: "ic_base.svg",
    chainId: ChainId.Base,
    deploy_token_amount: 50,
    deploy_agent_amount: 0,
    api_key_label:
      "API Key for other networks (Base, Arbitrum, Solana, ZKsync, Bittensor, etc.)",
  },
  {
    id: "arbitrum",
    title: "Arbitrum",
    required_amount: 1,
    icon: "ic_arbitrum.svg",
    chainId: ChainId.Arbitrum,
    deploy_token_amount: 50,
    deploy_agent_amount: 0,
    api_key_label:
      "API Key for other networks (Base, Arbitrum, Solana, ZKsync, Bittensor, etc.)",
  },
  {
    id: "bsc",
    title: "BNB",
    required_amount: 1,
    icon: "ic-bsc.png",
    chainId: ChainId.BSC,
    deploy_token_amount: 50,
    deploy_agent_amount: 0,
    api_key_label:
      "API Key for other networks (Base, Arbitrum, Solana, ZKsync, Bittensor, etc.)",
  },
  {
    id: "ethereum",
    title: "Ethereum",
    required_amount: 500,
    icon: "ic-ether.png",
    chainId: ChainId.ETH,
    deploy_token_amount: 0,
    deploy_agent_amount: 100,
    api_key_label:
      "API Key for other networks (Base, Arbitrum, Solana, ZKsync, Bittensor, etc.)",
  },
  {
    id: "bitcoin",
    title: "Bitcoin",
    required_amount: 600,
    icon: "ic_bitcoin_gold.svg",
    chainId: ChainId.ShardAI,
    deploy_token_amount: 50,
    deploy_agent_amount: 200,
    api_key_label:
      "API Key for other networks (Base, Arbitrum, Solana, ZKsync, Bittensor, etc.)",
  },
  {
    id: "tao",
    title: "Bittensor",
    // required_amount: 0.1,
    required_amount: 100,
    icon: "ic-tao.svg",
    chainId: ChainId.Bittensor,
    deploy_token_amount: 0,
    deploy_agent_amount: 20,
    api_key_label:
      "API Key for other networks (Base, Arbitrum, Solana, ZKsync, Bittensor, etc.)",
  },
  {
    id: "ape",
    title: "ApeChain",
    // required_amount: 0.1,
    required_amount: 1,
    icon: "ic-ape.png",
    chainId: ChainId.Ape,
    deploy_token_amount: 50,
    deploy_agent_amount: 0,
    api_key_label:
      "API Key for other networks (Base, Arbitrum, Solana, ZKsync, Bittensor, etc.)",
  },
  {
    id: "solana",
    title: "Solana",
    required_amount: 50,
    icon: "ic_solana.svg",
    chainId: ChainId.Solana,
    deploy_token_amount: 50,
    deploy_agent_amount: 125,
    api_key_label:
      "API Key for other networks (Base, Arbitrum, Solana, ZKsync, Bittensor, etc.)",
  },
  {
    id: "polygon",
    title: "Polygon",
    required_amount: 1,
    icon: "ic_polygon.svg",
    chainId: ChainId.Polygon,
    deploy_token_amount: 50,
    deploy_agent_amount: 0,
    api_key_label:
      "API Key for other networks (Base, Arbitrum, Solana, ZKsync, Bittensor, etc.)",
  },
  {
    id: "zksync",
    title: "ZKsync Era",
    // required_amount: 0.1,
    required_amount: 1,
    icon: "ic_zksync.svg",
    chainId: ChainId.ZkSync,
    deploy_token_amount: 50,
    deploy_agent_amount: 0,
    api_key_label:
      "API Key for other networks (Base, Arbitrum, Solana, ZKsync, Bittensor, etc.)",
  },
  {
    id: "avax",
    title: "Avalanche C-Chain",
    // required_amount: 0.1,
    required_amount: 1,
    icon: "ic_avax.svg",
    chainId: ChainId.Avax,
    deploy_token_amount: 50,
    deploy_agent_amount: 0,
    api_key_label:
      "API Key for other networks (Base, Arbitrum, Solana, ZKsync, Bittensor, etc.)",
  },
  {
    id: "tron",
    title: "Tron",
    // required_amount: 0.1,
    required_amount: 50,
    icon: "ic-tron.svg",
    chainId: ChainId.Tron,
    deploy_token_amount: 50,
    deploy_agent_amount: 200,
    api_key_label:
      "API Key for other networks (Base, Arbitrum, Solana, ZKsync, Bittensor, etc.)",
  },
  {
    id: "abs",
    title: "Abstract Testnet",
    // required_amount: 0.1,
    required_amount: 1,
    icon: "ic-abs.png",
    chainId: ChainId.Abstract,
    deploy_token_amount: 50,
    deploy_agent_amount: 0,
    api_key_label:
      "API Key for other networks (Base, Arbitrum, Solana, ZKsync, Bittensor, etc.)",
  },
  {
    id: "duck",
    title: "DuckChain",
    // required_amount: 0.1,
    required_amount: 1,
    icon: "ic-duck.svg",
    chainId: ChainId.Duck,
    deploy_token_amount: 50,
    deploy_agent_amount: 0,
    api_key_label:
      "API Key for other networks (Base, Arbitrum, Solana, ZKsync, Bittensor, etc.)",
  },
  {
    id: "symbiosis",
    title: "Symbiosis",
    required_amount: 1,
    icon: "ic_nbs.svg",
    chainId: ChainId.HERMES,
    desc: "Symbiosis is ZK rollup on Bitcoin optimized for onchain AI agents and onchain AI models. It combines high performance and low cost, while leveraging the security and decentralization of Bitcoin.",
    deploy_token_amount: 50,
    deploy_agent_amount: 0,
    api_key_label: "API Key for Symbiosis",
  },
  {
    id: "runes",
    title: "Runes",
    required_amount: -1,
    icon: "ic-rune.svg",
    chainId: ChainId.Runes,
    deploy_token_amount: 50,
    deploy_agent_amount: 0,
    hide: true,
  },
  // {
  //   id: 'no-token',
  //   title: 'No Token',
  //   required_amount: -1,
  //   icon: '',
  //   chainId: '-1',
  //   deploy_amount: 0,
  // },
];
