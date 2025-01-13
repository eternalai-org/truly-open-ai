import { DataSource } from '../types/data-source';

export const DATA_SOURCE: Record<string, DataSource[]> = {
  'network-data-source': [
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
  'token-data-source': [
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
};
