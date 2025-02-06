import { Viewport } from '@xyflow/react';
import { v4 } from 'uuid';
import { create } from 'zustand';

import { DEFAULT_DISABLED_CONNECTION } from '../constants/default-values';
import { StudioDataNode } from '../types/graph';

type Store = {
  entry: StudioDataNode | null;
  setEntry: (entry: StudioDataNode | null) => void;

  data: StudioDataNode[];
  setData: (data: StudioDataNode[]) => void;

  viewport: Viewport;
  setViewport: (viewport: Viewport) => void;

  disabledConnection?: boolean;
  setDisabledConnection: (disabledConnection: boolean) => void;

  clear: () => void;
};

const useStudioDataStore = create<Store>((set) => ({
  entry: null,
  setEntry: (entry) => {
    set({ entry });
  },

  data: [],
  setData: (data) => {
    const processingData = (data || []).map((item) => {
      if (item.id) {
        return item;
      }

      return {
        ...item,
        id: v4(),
      };
    });
    set({ data: processingData });
  },

  viewport: {
    x: 0,
    y: 0,
    zoom: 1,
  },
  setViewport: (viewport) => set({ viewport }),

  disabledConnection: DEFAULT_DISABLED_CONNECTION,
  setDisabledConnection: (disabledConnection) => {
    set({ disabledConnection });
  },

  clear: () => {
    set({ data: [], entry: null });
  },
}));

export default useStudioDataStore;
