import { create } from 'zustand';

import { DEFAULT_BOARD_CONFIG, DEFAULT_SIDEBAR_SIDE, DEFAULT_TAB_BEHAVIOR } from '../constants/default-values';
import { StudioConfig } from '../types/config';

const DEFAULT_VALUE = {
  config: {
    sidebar: {
      side: DEFAULT_SIDEBAR_SIDE,
    },
    tab: {
      behavior: DEFAULT_TAB_BEHAVIOR,
    },
    board: DEFAULT_BOARD_CONFIG,
  } satisfies StudioConfig,
};

type Store = {
  config: StudioConfig;

  setConfig: (config?: StudioConfig) => void;

  clear: () => void;
};

const useStudioConfigStore = create<Store>((set) => ({
  ...DEFAULT_VALUE,

  setConfig: (config) => {
    const newConfig = { ...DEFAULT_VALUE.config, ...config };

    set({ config: newConfig });
  },

  clear: () => set(DEFAULT_VALUE),
}));

export default useStudioConfigStore;
