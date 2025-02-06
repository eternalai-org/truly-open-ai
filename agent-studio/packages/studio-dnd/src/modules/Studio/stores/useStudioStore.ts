import { create } from 'zustand';

const DEFAULT_VALUE = {};

type Store = {
  clear: () => void;
};

const useStudioStore = create<Store>((set) => ({
  ...DEFAULT_VALUE,

  clear: () => set(DEFAULT_VALUE),
}));

export default useStudioStore;
