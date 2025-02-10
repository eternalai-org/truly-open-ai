import { IChainConnected } from "@eternalai-dagent/core";
import { create } from "zustand";

export type ShortAgentToken = {
  image_url: string;
  name: string;
  symbol: string;
};

interface IStore {
  chains: IChainConnected[];
  setChains: (chains: IChainConnected[]) => void;

  modelDescriptions: Record<string, string>;
  setModelDescriptions: (descriptions: Record<string, string>) => void;

  agentTokens: ShortAgentToken[];
  setAgentTokens: (tokens: ShortAgentToken[]) => void;

  needReload: number;
  requestReload: () => void;
}

const useCommonStore = create<IStore>((set) => ({
  chains: [],
  setChains: (chains) => set({ chains }),

  modelDescriptions: {},
  setModelDescriptions: (descriptions) =>
    set({ modelDescriptions: descriptions }),

  agentTokens: [],
  setAgentTokens: (tokens) => set({ agentTokens: tokens || [] }),

  needReload: 0,
  requestReload: () => set((state) => ({ needReload: state.needReload + 1 })),
}));

export default useCommonStore;
