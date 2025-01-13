import { Viewport, XYPosition } from '@xyflow/react';
import { create } from 'zustand';

const DEFAULT_VALUE = {
  mousePosition: { x: 0, y: 0 },
  view: { x: 0, y: 0, zoom: 1 },
};

type Store = {
  mousePosition: XYPosition;
  setMousePosition: (position: XYPosition) => void;

  view: Viewport;
  setView: (view: Viewport) => void;

  clear: () => void;
};

const useStudioFlowViewStore = create<Store>((set) => ({
  ...DEFAULT_VALUE,

  setMousePosition: (position) => set({ mousePosition: position }),
  setView: (view) => set({ view }),

  clear: () => set(DEFAULT_VALUE),
}));

export default useStudioFlowViewStore;
