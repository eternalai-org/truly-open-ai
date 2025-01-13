import { ReactNode } from 'react';
import { create } from 'zustand';

import { DraggableData } from '../types/dnd';
import { DomRect } from '../types/ui';

const DEFAULT_VALUE = {
  draggingElement: null,
  draggingData: null,
  draggingPoint: null,
};

type Store = {
  draggingElement: ReactNode | null;
  draggingData: DraggableData | null;
  draggingPoint: DomRect | null; // touching point of dragging element

  setDragging: (node?: ReactNode | null, data?: DraggableData | null, point?: DomRect | null) => void;

  clear: () => void;
};

const useStudioDndStore = create<Store>((set) => ({
  ...DEFAULT_VALUE,

  setDragging: (node, data, point) => {
    set({ draggingElement: node, draggingData: data, draggingPoint: point });
  },

  clear: () => {
    set(DEFAULT_VALUE);
  },
}));

export default useStudioDndStore;
