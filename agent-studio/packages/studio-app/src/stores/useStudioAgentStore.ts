import { StudioDataNode } from "@agent-studio/studio-dnd";
import { IAgent } from "@eternal-ai/core";
import { create } from "zustand";

interface IStore {
  isDetail: boolean;
  setIsDetail: (isDetail: boolean) => void;

  data: StudioDataNode[];
  setData: (data: StudioDataNode[]) => void;

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

  data: [],
  setData: (data) => set({ data }),

  isLoading: false,
  setIsLoading: (isLoading) => set({ isLoading }),

  isDisabled: true,
  setIsDisabled: (isDisabled) => set({ isDisabled }),

  agentDetail: undefined,
  setAgentDetail: (agent) => set({ agentDetail: agent }),
}));

export default useStudioAgentStore;
