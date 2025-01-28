import { LegoComponentIcon } from '../components';
import { StudioCategoryType } from '../enums/category';
import { StudioCategory } from '../types';

const AGENT_CATEGORY: StudioCategory = {
  idx: 'agent',
  title: 'Agent',
  required: true,
  icon: LegoComponentIcon,
  options: [
    {
      idx: 'agent_new',
      title: 'New Agent',
      tooltip: '',
      data: {
        agentName: {
          type: 'text',
          label: 'Agent Name',
          placeholder: 'Agent Name',
          defaultValue: '',
        },
      },
    },
  ],
};

const PERSONALITY_CATEGORY: StudioCategory = {
  idx: 'personality',
  title: 'Personality',
  required: true,
  tooltip:
    'Create an agent for your NFT, Ordinals, token, —or start fresh with a new idea. This section defines your agent’s lore and backstory.',
  icon: LegoComponentIcon,
  color: '#12DAC2',
  options: [
    {
      idx: 'personality_customize',
      title: 'New personality',
      tooltip: '',
      icon: 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_nft.svg',
      type: StudioCategoryType.LINK,
    },
    {
      idx: 'personality_nft',
      title: 'Import from NFT',
      icon: 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_ordinals.svg',
      type: StudioCategoryType.LINK,
    },
    {
      idx: 'personality_ordinals',
      title: 'Import from Ordinals',
      icon: 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_token.svg',
    },
    {
      idx: 'personality_token',
      title: 'Import from Token',
      icon: 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_custom.svg',
      customizeRenderOnBoard: () => {
        return (
          <div
            style={{
              width: 400,
              height: 400,
            }}
          >
            custom
          </div>
        );
      },
    },
  ],
};

const AI_FRAMEWORK_CATEGORY: StudioCategory = {
  idx: 'ai_framework',
  title: 'AI Framework',
  required: true,
  tooltip:
    'Pick the blockchain where your agent lives. Each option has unique deployment fees, performance, and ongoing costs. Choose what best fits your needs.',
  icon: LegoComponentIcon,
  color: '#12DAC2',
  options: [
    {
      idx: 'ai_framework_eternal_ai',
      title: 'Eternal AI',
      tooltip: '',
      icon: 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_nft.svg',
    },
    {
      idx: 'ai_framework_eliza',
      title: 'Eliza',
      disabled: true,
    },
    {
      idx: 'ai_framework_zerepy',
      title: 'ZerePy',
      disabled: true,
    },
  ],
};

const BLOCKCHAIN_CATEGORY: StudioCategory = {
  idx: 'blockchain',
  title: 'Blockchains',
  required: true,
  icon: LegoComponentIcon,
  color: '#368cdc',
  tooltip: '',
  options: [
    {
      idx: 'blockchain_base',
      title: 'Base',
      tooltip: '',
      icon: '/icons/blockchains/ic_base.svg',
    },
    {
      idx: 'blockchain_arbitrum',
      title: 'Arbitrum',
      icon: '/icons/blockchains/ic_arbitrum.svg',
    },
    {
      idx: 'blockchain_bnb',
      title: 'BNB',
      icon: '/icons/blockchains/ic-bsc.png',
    },
    {
      idx: 'blockchain_bitcoin',
      title: 'Bitcoin',
      icon: '/icons/blockchains/ic_bitcoin.svg',
    },
    {
      idx: 'blockchain_bittensor',
      title: 'Bittensor',
      icon: '/icons/blockchains/ic-tao.svg',
    },
    {
      idx: 'blockchain_ape_chain',
      title: 'Ape Chain',
      icon: '/icons/blockchains/ic-ape.png',
    },
    {
      idx: 'blockchain_solana',
      title: 'Solana',
      icon: '/icons/blockchains/ic_solana.svg',
    },
    {
      idx: 'blockchain_polygon',
      title: 'Polygon',
      icon: '/icons/blockchains/ic_polygon.svg',
    },
    {
      idx: 'blockchain_zksync_era',
      title: 'ZKsync Era',
      icon: '/icons/blockchains/ic_zksync.svg',
    },
    {
      idx: 'blockchain_avalanche',
      title: 'Avalanche C-Chain',
      icon: '/icons/blockchains/ic_avax.svg',
    },
    {
      idx: 'blockchain_abstract_testnet',
      title: 'Abstract Testnet',
      icon: '/icons/blockchains/ic-abs.png',
    },
    {
      idx: 'blockchain_duck_chain',
      title: 'DuckChain',
      icon: '/icons/blockchains/ic-duck.svg',
    },
    {
      idx: 'blockchain_symbiosis',
      title: 'Symbiosis',
      icon: '/icons/blockchains/ic_nbs.svg',
    },
  ],
};

const DECENTRALIZED_CATEGORY: StudioCategory = {
  idx: 'decentralized',
  title: 'Decentralize Inference',
  required: true,
  tooltip:
    'Create an agent for your NFT, Ordinals, token, —or start fresh with a new idea. This section defines your agent’s lore and backstory.',
  icon: LegoComponentIcon,
  color: '#15C888',
  options: [
    {
      idx: 'decentralize_inference_hermes_3_70b',
      title: 'Hermes 3 70B',
      tooltip: '',
    },
    {
      idx: 'decentralize_inference_intellect_1_10b',
      title: 'INTELLECT 1 10B',
      tooltip: '',
    },
    {
      idx: 'decentralize_inference_llama_3_1_405b',
      title: 'Llama 3.1 405B',
      tooltip: '',
    },
    {
      idx: 'decentralize_inference_llama_3_3_470b',
      title: 'Llama 3.3 470B',
      tooltip: '',
      disabled: true,
    },
  ],
};

const TOKEN_CATEGORY: StudioCategory = {
  idx: 'token',
  title: 'Tokens',
  required: false,
  tooltip:
    'Select the blockchain to issue your agent’s token. Consider factors like accessibility, liquidity, and trading volume.',
  icon: LegoComponentIcon,
  color: '#A041FF',
  options: [
    {
      idx: 'token_base',
      title: 'Base',
      tooltip: '',
      icon: '/icons/blockchains/ic_base.svg',
      data: {
        value: {
          type: 'hidden',
          label: 'Token Name',
          placeholder: 'Token Name',
          defaultValue: '8453',
        },
      },
    },
    {
      idx: 'token_solana',
      title: 'Solana',
      tooltip: '',
      icon: '/icons/blockchains/ic_solana.svg',
      data: {
        value: {
          type: 'hidden',
          label: 'Token Name',
          placeholder: 'Token Name',
          defaultValue: '8454',
        },
      },
    },
    {
      idx: 'token_arbitrum',
      title: 'Arbitrum',
      tooltip: '',
      icon: '/icons/blockchains/ic_arbitrum.svg',
      data: {
        value: {
          type: 'hidden',
          label: 'Token Name',
          placeholder: 'Token Name',
          defaultValue: '8455',
        },
      },
    },
    {
      idx: 'token_bnb',
      title: 'BNB',
      tooltip: '',
      icon: '/icons/blockchains/ic-bsc.png',
      data: {
        value: {
          type: 'hidden',
          label: 'Token Name',
          placeholder: 'Token Name',
          defaultValue: '8456',
        },
      },
    },
  ],
};

const MISSION_ON_X_CATEGORY: StudioCategory = {
  idx: 'mission_on_x',
  title: 'X',
  required: false,
  tooltip: 'Set your agent’s activities on X. Choose what it does (e.g., posts, replies) and how often it does them.',
  icon: LegoComponentIcon,
  color: '#6B39F9',
  options: [
    {
      idx: 'mission_on_x_post',
      title: 'Post on X',
      tooltip: '',
      icon: 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_post.svg',
      type: StudioCategoryType.LINK,
    },
    {
      idx: 'mission_on_x_reply',
      title: 'Reply on X',
      tooltip: '',
      icon: 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_reply.svg',
      type: StudioCategoryType.LINK,
    },
    {
      idx: 'mission_on_x_engage',
      title: 'Engage on X',
      tooltip: '',
      icon: 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_engage.svg',
      type: StudioCategoryType.LINK,
    },
    {
      idx: 'mission_on_x_follow',
      title: 'Follow on X',
      tooltip: '',
      icon: 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_follow.svg',
      type: StudioCategoryType.LINK,
    },
  ],
};

const MISSION_ON_FARCASTER_CATEGORY: StudioCategory = {
  idx: 'mission_on_farcaster',
  title: 'Farcaster',
  required: false,
  tooltip: 'Set your agent’s activities on X. Choose what it does (e.g., posts, replies) and how often it does them.',
  icon: LegoComponentIcon,
  color: '#6B39F9',
  options: [
    {
      idx: 'mission_on_farcaster_post',
      title: 'Post on Farcaster',
      tooltip: '',
      icon: 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_post.svg',
      type: StudioCategoryType.LINK,
    },
    {
      idx: 'mission_on_farcaster_reply',
      title: 'Reply on Farcaster',
      tooltip: '',
      icon: 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_reply.svg',
      type: StudioCategoryType.LINK,
    },
  ],
};

const AGENT_MODEL_CATEGORIES: StudioCategory[] = [
  AGENT_CATEGORY,
  PERSONALITY_CATEGORY,
  AI_FRAMEWORK_CATEGORY,
  BLOCKCHAIN_CATEGORY,
  DECENTRALIZED_CATEGORY,
  TOKEN_CATEGORY,
  MISSION_ON_X_CATEGORY,
  MISSION_ON_FARCASTER_CATEGORY,
];

export default AGENT_MODEL_CATEGORIES;
