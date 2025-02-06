import { GraphData } from "@agent-studio/studio-dnd";
import { IAgent } from "@eternalai-dagent/core";
import { create } from "zustand";

interface IStore {
  isDetail: boolean;
  setIsDetail: (isDetail: boolean) => void;

  graphData: GraphData;
  setGraphData: (data: GraphData) => void;

  isLoading: boolean;
  setIsLoading: (isLoading: boolean) => void;

  isDisabled: boolean;
  setIsDisabled: (isDisabled: boolean) => void;

  agentDetail?: IAgent;
  setAgentDetail: (agent: IAgent | undefined) => void;
}

const useStudioAgentStore = create<IStore>((set) => ({
  isDetail: false,
  setIsDetail: (isDetail) => set({ isDetail }),

  graphData: {
    data: [],
    viewport: { x: 0, y: 0, zoom: 1 },
  },
  setGraphData: (data) => {
    set({ graphData: data });
  },

  isLoading: false,
  setIsLoading: (isLoading) => set({ isLoading }),

  isDisabled: true,
  setIsDisabled: (isDisabled) => set({ isDisabled }),

  agentDetail: undefined,
  setAgentDetail: (agent) => set({ agentDetail: agent }),
}));

export default useStudioAgentStore;
