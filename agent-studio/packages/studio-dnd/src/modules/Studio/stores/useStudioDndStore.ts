import { Transform } from '@dnd-kit/utilities';
import { ReactNode } from 'react';
import { create } from 'zustand';

import { DraggableData } from '../types/dnd';
import { DomRect } from '../types/ui';

const DEFAULT_VALUE = {
  draggingElement: null,
  draggingData: null,
  draggingPoint: null,
  transform: null,
};

type Store = {
  draggingElement: ReactNode | null;
  draggingData: DraggableData | null;
  draggingPoint: DomRect | null; // touching point of dragging element
  transform: Transform | null;

  setDragging: (node?: ReactNode | null, data?: DraggableData | null, point?: DomRect | null) => void;
  setTransform: (transform: Transform | null) => void;

  clear: () => void;
};

const useStudioDndStore = create<Store>((set, get) => ({
  ...DEFAULT_VALUE,

  setDragging: (node, data, point) => {
    set({
      draggingElement: node,
      draggingData: data,
      draggingPoint: point,
      transform: node ? get().transform : null,
    });
  },

  clear: () => {
    set(DEFAULT_VALUE);
  },

  setTransform: (t) => {
    set({ transform: t });
  },
}));

export default useStudioDndStore;
