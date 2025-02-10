import type { Meta, StoryObj } from '@storybook/react';
import { useRef, useState } from 'react';

import { AGENT_DATA_SOURCE } from './__mocks__/agent-data-source';
import { Studio, StudioProps, StudioRef } from './Studio';
import { GraphData } from './types';

type Story = StoryObj<typeof Studio>;

const args = {
  categories: [
    {
      'idx': 'agent',
      'title': 'Agent',
      'required': true,
      'multipleOption': false,
      'disabled': false,
      'tooltip': 'New Agent',
      'options': [
        {
          'idx': 'agent_new',
          'title': 'New Agent',
          'tooltip': 'New Agent',
          'zIndex': 100,
          'data': {
            'agentName': {
              'type': 'text',
              'label': 'Agent Name',
              'placeholder': 'Agent Name',
              'defaultValue': '',
              'disabled': false,
            },
          },
        },
      ],
    },
    {
      'idx': 'personality',
      'title': 'Personality',
      'required': true,
      'tooltip':
        'Create an agent for your NFT, Ordinals, token, —or start fresh with a new idea. This section defines your agent’s lore and backstory.',
      'color': '#12DAC2',
      'disabled': false,
      'options': [
        {
          'idx': 'personality_customize',
          'title': 'New personality',
          'tooltip': 'New Agent',
          'icon': 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_custom.svg',
          customizeRenderOnBoard: () => {
            return (
              <div
                style={{
                  width: 400,
                  height: 28,
                }}
              >
                Customize
              </div>
            );
          },
        },
        {
          'idx': 'personality_nft',
          'title': 'Import from NFT',
          'icon': 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_nft.svg',
          customizeRenderOnBoard: () => {
            return (
              <div
                style={{
                  width: 400,
                  height: 400,
                }}
              >
                Customize
              </div>
            );
          },
        },
        {
          'idx': 'personality_ordinals',
          'title': 'Import from Ordinals',
          'icon': 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_ordinals.svg',
          customizeRenderOnBoard: () => {
            return (
              <div
                style={{
                  width: 400,
                  height: 400,
                }}
              >
                Customize
              </div>
            );
          },
        },
        {
          'idx': 'personality_token',
          'title': 'Import from Token',
          'icon': 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_token.svg',
          customizeRenderOnBoard: () => {
            return (
              <div
                style={{
                  width: 400,
                  height: 400,
                }}
              >
                Customize
              </div>
            );
          },
        },
      ],
    },
    {
      'idx': 'ai_framework',
      'title': 'AI Framework',
      'required': true,
      'tooltip':
        'Pick the blockchain where your agent lives. Each option has unique deployment fees, performance, and ongoing costs. Choose what best fits your needs.',
      'color': '#12DAC2',
      'disabled': false,
      'options': [
        {
          'idx': 'ai_framework_eternal_ai',
          'title': 'Eternal AI',
          'tooltip': '',
          'icon': 'https://storage.googleapis.com/eternal-ai/agent-studio/ic_eternal_ai.svg',
          'data': {
            'aiFrameworkId': {
              'type': 'hidden',
              'label': 'AI Framework Id',
              'placeholder': 'AI Framework Id',
              'defaultValue': 'ai_framework_eternal_ai',
            },
          },
        },
        {
          'idx': 'ai_framework_eliza',
          'title': 'Eliza',
          'disabled': true,
        },
        {
          'idx': 'ai_framework_zerepy',
          'title': 'ZerePy',
          'disabled': true,
        },
      ],
    },
    {
      'idx': 'blockchain',
      'title': 'Blockchains',
      'required': true,
      'color': '#368cdc',
      'tooltip': '',
      'disabled': false,
      'options': [
        {
          'idx': 'blockchain_base',
          'title': 'Base',
          'tooltip': '',
          'icon': '/icons/blockchains/ic_base.svg',
          'data': {
            'chainId': {
              'type': 'hidden',
              'label': 'Chain Id',
              'placeholder': 'Chain Id',
              'defaultValue': 'base',
            },
          },
        },
        {
          'idx': 'blockchain_arbitrum',
          'title': 'Arbitrum',
          'icon': '/icons/blockchains/ic_arbitrum.svg',
          'data': {
            'chainId': {
              'type': 'hidden',
              'label': 'Chain Id',
              'placeholder': 'Chain Id',
              'defaultValue': 'arbitrum',
            },
          },
        },
        {
          'idx': 'blockchain_bnb',
          'title': 'BNB',
          'icon': '/icons/blockchains/ic-bsc.png',
          'data': {
            'chainId': {
              'type': 'hidden',
              'label': 'Chain Id',
              'placeholder': 'Chain Id',
              'defaultValue': 'bsc',
            },
          },
        },
        {
          'idx': 'blockchain_bitcoin',
          'title': 'Bitcoin',
          'icon': '/icons/blockchains/ic_bitcoin.svg',
          'data': {
            'chainId': {
              'type': 'hidden',
              'label': 'Chain Id',
              'placeholder': 'Chain Id',
              'defaultValue': 'bitcoin',
            },
          },
        },
        {
          'idx': 'blockchain_bittensor',
          'title': 'Bittensor',
          'icon': '/icons/blockchains/ic-tao.svg',
          'data': {
            'chainId': {
              'type': 'hidden',
              'label': 'Chain Id',
              'placeholder': 'Chain Id',
              'defaultValue': 'tao',
            },
          },
        },
        {
          'idx': 'blockchain_ape_chain',
          'title': 'Ape Chain',
          'icon': '/icons/blockchains/ic-ape.png',
          'data': {
            'chainId': {
              'type': 'hidden',
              'label': 'Chain Id',
              'placeholder': 'Chain Id',
              'defaultValue': 'ape',
            },
          },
        },
        {
          'idx': 'blockchain_solana',
          'title': 'Solana',
          'icon': '/icons/blockchains/ic_solana.svg',
          'data': {
            'chainId': {
              'type': 'hidden',
              'label': 'Chain Id',
              'placeholder': 'Chain Id',
              'defaultValue': 'solana',
            },
          },
        },
        {
          'idx': 'blockchain_polygon',
          'title': 'Polygon',
          'icon': '/icons/blockchains/ic_polygon.svg',
          'data': {
            'chainId': {
              'type': 'hidden',
              'label': 'Chain Id',
              'placeholder': 'Chain Id',
              'defaultValue': 'polygon',
            },
          },
        },
        {
          'idx': 'blockchain_zksync_era',
          'title': 'ZKsync Era',
          'icon': '/icons/blockchains/ic_zksync.svg',
          'data': {
            'chainId': {
              'type': 'hidden',
              'label': 'Chain Id',
              'placeholder': 'Chain Id',
              'defaultValue': 'zksync',
            },
          },
        },
        {
          'idx': 'blockchain_avalanche',
          'title': 'Avalanche C-Chain',
          'icon': '/icons/blockchains/ic_avax.svg',
          'data': {
            'chainId': {
              'type': 'hidden',
              'label': 'Chain Id',
              'placeholder': 'Chain Id',
              'defaultValue': 'avax',
            },
          },
        },
        {
          'idx': 'blockchain_abstract_testnet',
          'title': 'Abstract Testnet',
          'icon': '/icons/blockchains/ic-abs.png',
          'data': {
            'chainId': {
              'type': 'hidden',
              'label': 'Chain Id',
              'placeholder': 'Chain Id',
              'defaultValue': 'abs',
            },
          },
        },
        {
          'idx': 'blockchain_duck_chain',
          'title': 'DuckChain',
          'icon': '/icons/blockchains/ic-duck.svg',
          'data': {
            'chainId': {
              'type': 'hidden',
              'label': 'Chain Id',
              'placeholder': 'Chain Id',
              'defaultValue': 'duck',
            },
          },
        },
        {
          'idx': 'blockchain_symbiosis',
          'title': 'Symbiosis',
          'icon': '/icons/blockchains/ic_nbs.svg',
          'data': {
            'chainId': {
              'type': 'hidden',
              'label': 'Chain Id',
              'placeholder': 'Chain Id',
              'defaultValue': 'symbiosis',
            },
          },
        },
      ],
    },
    {
      'idx': 'decentralized_inference',
      'title': 'Decentralize Inference',
      'required': true,
      'tooltip':
        'Create an agent for your NFT, Ordinals, token, —or start fresh with a new idea. This section defines your agent’s lore and backstory.',
      'color': '#15C888',
      'disabled': false,
      'options': [
        {
          'idx': 'decentralize_inference_hermes_3_70b',
          'title': 'Hermes 3 70B',
          'tooltip': '',
          'data': {
            'decentralizeId': {
              'type': 'hidden',
              'label': 'Decentralize Id',
              'placeholder': 'Decentralize Id',
              'defaultValue': 'NousResearch/Hermes-3-Llama-3.1-70B-FP8',
            },
          },
        },
        {
          'idx': 'decentralize_inference_intellect_1_10b',
          'title': 'INTELLECT 1 10B',
          'tooltip': '',
          'data': {
            'decentralizeId': {
              'type': 'hidden',
              'label': 'Decentralize Id',
              'placeholder': 'Decentralize Id',
              'defaultValue': 'PrimeIntellect/INTELLECT-1-Instruct',
            },
          },
        },
        {
          'idx': 'decentralize_inference_llama_3_1_405b',
          'title': 'Llama 3.1 405B',
          'tooltip': '',
          'data': {
            'decentralizeId': {
              'type': 'hidden',
              'label': 'Decentralize Id',
              'placeholder': 'Decentralize Id',
              'defaultValue': 'neuralmagic/Meta-Llama-3.1-405B-Instruct-quantized.w4a16',
            },
          },
        },
        {
          'idx': 'decentralize_inference_llama_3_3_470b',
          'title': 'Llama 3.3 470B',
          'tooltip': '',
          'disabled': true,
          'data': {
            'decentralizeId': {
              'type': 'hidden',
              'label': 'Decentralize Id',
              'placeholder': 'Decentralize Id',
              'defaultValue': 'unsloth/Llama-3.3-70B-Instruct-bnb-4bit',
            },
          },
        },
      ],
    },
    {
      'idx': 'token',
      'title': 'Tokens',
      'required': false,
      'tooltip':
        'Select the blockchain to issue your agent’s token. Consider factors like accessibility, liquidity, and trading volume.',
      'color': '#A041FF',
      'disabled': false,
      'options': [
        {
          'idx': 'token_base',
          'title': 'Token on Base',
          'tooltip': '',
          'icon': '/icons/blockchains/ic_base.svg',
          'data': {
            'tokenId': {
              'type': 'hidden',
              'label': 'Token Name',
              'placeholder': 'Token Name',
              'defaultValue': '8453',
            },
          },
        },
        {
          'idx': 'token_solana',
          'title': 'Token on Solana',
          'tooltip': '',
          'icon': '/icons/blockchains/ic_solana.svg',
          'data': {
            'tokenId': {
              'type': 'hidden',
              'label': 'Token Name',
              'placeholder': 'Token Name',
              'defaultValue': '1111',
            },
          },
        },
        {
          'idx': 'token_arbitrum',
          'title': 'Token on Arbitrum',
          'tooltip': '',
          'icon': '/icons/blockchains/ic_arbitrum.svg',
          'data': {
            'tokenId': {
              'type': 'hidden',
              'label': 'Token Name',
              'placeholder': 'Token Name',
              'defaultValue': '42161',
            },
          },
        },
        {
          'idx': 'token_bnb',
          'title': 'Token on BNB',
          'tooltip': '',
          'icon': '/icons/blockchains/ic-bsc.png',
          'data': {
            'tokenId': {
              'type': 'hidden',
              'label': 'Token Name',
              'placeholder': 'Token Name',
              'defaultValue': '56',
            },
          },
        },
        {
          'idx': 'token_ape',
          'title': 'Token on ApeChain',
          'tooltip': '',
          'icon': '/icons/blockchains/ic-ape.png',
          'data': {
            'tokenId': {
              'type': 'hidden',
              'label': 'Token Name',
              'placeholder': 'Token Name',
              'defaultValue': '33139',
            },
          },
        },
        {
          'idx': 'token_avax',
          'title': 'Token on Avalance C-Chain',
          'tooltip': '',
          'icon': '/icons/blockchains/ic_avax.svg',
          'data': {
            'tokenId': {
              'type': 'hidden',
              'label': 'Token Name',
              'placeholder': 'Token Name',
              'defaultValue': '43114',
            },
          },
        },
      ],
    },
    {
      'idx': 'mission_on_x',
      'title': 'X',
      'required': false,
      'tooltip':
        'Set your agent’s activities on X. Choose what it does (e.g., posts, replies) and how often it does them.',
      'color': '#6B39F9',
      'options': [
        {
          'idx': 'mission_on_x_post',
          'title': 'Post on X',
          'tooltip': '',
          'icon': 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_post.svg',
          'type': 'link',
        },
        {
          'idx': 'mission_on_x_reply',
          'title': 'Reply on X',
          'tooltip': '',
          'icon': 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_reply.svg',
          'type': 'link',
        },
        {
          'idx': 'mission_on_x_engage',
          'title': 'Engage on X',
          'tooltip': '',
          'icon': 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_engage.svg',
          'type': 'link',
        },
        {
          'idx': 'mission_on_x_follow',
          'title': 'Follow on X',
          'tooltip': '',
          'icon': 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_follow.svg',
          'type': 'link',
        },
      ],
    },
    {
      'idx': 'mission_on_farcaster',
      'title': 'Farcaster',
      'required': false,
      'tooltip':
        'Set your agent’s activities on X. Choose what it does (e.g., posts, replies) and how often it does them.',
      'color': '#6B39F9',
      'options': [
        {
          'idx': 'mission_on_farcaster_post',
          'title': 'Post on Farcaster',
          'tooltip': '',
          'icon': 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_post.svg',
          'type': 'link',
        },
        {
          'idx': 'mission_on_farcaster_reply',
          'title': 'Reply on Farcaster',
          'tooltip': '',
          'icon': 'https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_reply.svg',
          'type': 'link',
        },
      ],
    },
  ],
  dataSource: AGENT_DATA_SOURCE,
  graphData: {
    'data': [
      {
        'id': '721c2948-ccab-4425-9779-936172fb01f2',
        'idx': 'agent_new',
        'title': 'New Agent',
        'children': [
          {
            'id': '12e9dea7-ecab-40f8-8dd1-c8f1da61e80a',
            'idx': 'personality_customize',
            'title': 'New personality',
            'children': [],
            'data': {},
            'rect': {
              'position': {
                'x': 130,
                'y': 99.5,
              },
            },
            'categoryIdx': 'personality',
          },
          {
            'id': '307e60ef-cf1e-4584-b343-830cac55035f',
            'idx': 'decentralize_inference_hermes_3_70b',
            'title': 'Hermes 3 70B',
            'children': [],
            'data': {
              'decentralizeId': 'NousResearch/Hermes-3-Llama-3.1-70B-FP8',
            },
            'rect': {
              'position': {
                'x': -356.325830278392,
                'y': 579.7615898194902,
              },
            },
            'categoryIdx': 'decentralized_inference',
          },
          {
            'id': '0989eac7-90e5-4d41-86d2-4e3006c0d5fa',
            'idx': 'blockchain_base',
            'title': 'Base',
            'children': [],
            'data': {
              'chainId': 'base',
            },
            'rect': {
              'position': {
                'x': -344.171875,
                'y': 561,
              },
            },
            'categoryIdx': 'blockchain',
          },
          {
            'id': '56398488-5df3-48f5-b660-a58c0ceab703',
            'idx': 'ai_framework_eternal_ai',
            'title': 'Eternal AI',
            'children': [],
            'data': {
              'aiFrameworkId': 'ai_framework_eternal_ai',
            },
            'rect': {
              'position': {
                'x': -333.171875,
                'y': 560,
              },
            },
            'categoryIdx': 'ai_framework',
          },
          {
            'id': 'f6f0ba36-b89a-4334-abe1-11b125a5d4c2',
            'idx': 'decentralize_inference_hermes_3_70b',
            'title': 'Hermes 3 70B',
            'children': [],
            'data': {
              'decentralizeId': 'NousResearch/Hermes-3-Llama-3.1-70B-FP8',
            },
            'rect': {
              'position': {
                'x': -409.171875,
                'y': 705,
              },
            },
            'categoryIdx': 'decentralized_inference',
          },
          {
            'id': 'e141c450-4033-4fd1-b1e0-ca7eaa441d2b',
            'idx': 'token_base',
            'title': 'Token on Base',
            'children': [],
            'data': {
              'tokenId': '8453',
            },
            'rect': {
              'position': {
                'x': -456.171875,
                'y': 724,
              },
            },
            'categoryIdx': 'token',
          },
          {
            'id': '0b2a65a8-7954-42d8-b775-07028aeb996f',
            'idx': 'blockchain_symbiosis',
            'title': 'Symbiosis',
            'children': [],
            'data': {
              'chainId': 'symbiosis',
            },
            'rect': {
              'position': {
                'x': -437.171875,
                'y': 565,
              },
            },
            'categoryIdx': 'blockchain',
          },
          {
            'id': '6760e0ac-5f11-42dd-b892-41e9bdf237cc',
            'idx': 'ai_framework_eternal_ai',
            'title': 'Eternal AI',
            'children': [],
            'data': {
              'aiFrameworkId': 'ai_framework_eternal_ai',
            },
            'rect': {
              'position': {
                'x': -436.171875,
                'y': 555,
              },
            },
            'categoryIdx': 'ai_framework',
          },
          {
            'id': '90685e57-5c1a-45d5-9ac8-303776a82ee5',
            'idx': 'blockchain_bnb',
            'title': 'BNB',
            'children': [],
            'data': {
              'chainId': 'bsc',
            },
            'rect': {
              'position': {
                'x': -429.171875,
                'y': 577,
              },
            },
            'categoryIdx': 'blockchain',
          },
          {
            'id': '2bc99401-1bac-4951-a031-fc2636417035',
            'idx': 'decentralize_inference_intellect_1_10b',
            'title': 'INTELLECT 1 10B',
            'children': [],
            'data': {
              'decentralizeId': 'PrimeIntellect/INTELLECT-1-Instruct',
            },
            'rect': {
              'position': {
                'x': -489.171875,
                'y': 685,
              },
            },
            'categoryIdx': 'decentralized_inference',
          },
          {
            'id': '507e468d-e189-476c-b55d-36c9063a438b',
            'idx': 'mission_on_x_post',
            'title': 'Post on X',
            'children': [],
            'data': {},
            'rect': {
              'position': {
                'x': -629.171875,
                'y': 913,
              },
            },
            'categoryIdx': 'mission_on_x',
          },
        ],
        'data': {
          'agentName': '',
        },
        'rect': {
          'position': {
            'x': -445,
            'y': 310.5,
          },
        },
        'categoryIdx': 'agent',
      },
      {
        'id': '8a0487c6-ce57-49ab-b52b-6a239e8bbd32',
        'idx': 'personality_customize',
        'title': 'New personality',
        'children': [
          {
            'id': '41947e7d-9590-4482-9693-2dda468a6010',
            'idx': 'ai_framework_eternal_ai',
            'title': 'Eternal AI',
            'children': [],
            'data': {
              'aiFrameworkId': 'ai_framework_eternal_ai',
            },
            'rect': {
              'position': {
                'x': 536.828125,
                'y': 375,
              },
            },
            'categoryIdx': 'ai_framework',
          },
          {
            'id': '4ddb7e2a-3eb1-4fbc-81b7-aff8d02b2c1f',
            'idx': 'blockchain_ape_chain',
            'title': 'Ape Chain',
            'children': [],
            'data': {
              'chainId': 'ape',
            },
            'rect': {
              'position': {
                'x': 393.828125,
                'y': 366,
              },
            },
            'categoryIdx': 'blockchain',
          },
          {
            'id': 'd9227871-6e69-4362-845c-14a570aa9450',
            'idx': 'decentralize_inference_intellect_1_10b',
            'title': 'INTELLECT 1 10B',
            'children': [],
            'data': {
              'decentralizeId': 'PrimeIntellect/INTELLECT-1-Instruct',
            },
            'rect': {
              'position': {
                'x': 363.828125,
                'y': 435,
              },
            },
            'categoryIdx': 'decentralized_inference',
          },
          {
            'id': '758ed68d-6e00-4e95-9b41-8fd1dd03144d',
            'idx': 'decentralize_inference_llama_3_1_405b',
            'title': 'Llama 3.1 405B',
            'children': [],
            'data': {
              'decentralizeId': 'neuralmagic/Meta-Llama-3.1-405B-Instruct-quantized.w4a16',
            },
            'rect': {
              'position': {
                'x': 470.828125,
                'y': 374,
              },
            },
            'categoryIdx': 'decentralized_inference',
          },
        ],
        'data': {},
        'rect': {
          'position': {
            'x': 360.828125,
            'y': 351,
          },
        },
        'categoryIdx': 'personality',
      },
    ],
    'viewport': {
      'x': 700,
      'y': -149,
      'zoom': 1,
    },
  } satisfies GraphData,
} satisfies StudioProps;

const meta: Meta<typeof Studio> = {
  title: 'Studio',
  component: Studio,
  args,
};

export const AgentStudio: Story = {
  render: function useTabs(args) {
    const ref = useRef<StudioRef>(null);
    const [cate, setCate] = useState(args.categories);

    return (
      <div style={{ width: 'calc(100vw - 3rem)', height: 'calc(100vh - 3rem)' }}>
        <Studio
          {...args}
          categories={cate}
          ref={ref}
          isDebugger
          onChange={(data) => {
            console.log('[Studio] onChange', data);
            // disable personality if have existed
          }}
        />
      </div>
    );
  },
};

export default meta;
