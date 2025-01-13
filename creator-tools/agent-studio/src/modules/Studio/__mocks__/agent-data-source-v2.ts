import { DataSource } from '../types/data-source';

export const AGENT_DATA_SOURCE: Record<string, DataSource[]> = {
  'personality_data_source': [
    {
      value: 'custom',
      label: 'New Personality',
    },
    {
      value: 'nft',
      label: 'Import from NFT',
    },
    {
      value: 'ordinals',
      label: 'Import from Ordinals',
    },
    {
      value: 'token',
      label: 'Import from Token',
    },
  ],
  'network_data_source': [
    {
      value: 'base',
      label: 'Base',
    },
    {
      value: 'solana',
      label: 'Solana',
    },
    {
      value: 'bnb',
      label: 'BNB',
    },
  ],
  'token_data_source': [
    {
      value: 'base',
      label: 'Base',
    },
    {
      value: 'solana',
      label: 'Solana',
    },
    {
      value: 'bnb',
      label: 'BNB',
    },
  ],
  'ai_framework_data_source': [
    {
      value: 'eternal-ai',
      label: 'Eternal AI',
    },
    {
      value: 'eliza',
      label: 'Eliza',
      selectable: false,
    },
    {
      value: 'zerepy',
      label: 'ZerePy',
      selectable: false,
    },
  ],
};
