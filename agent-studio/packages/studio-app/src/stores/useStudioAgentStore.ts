import {
  findDataByCategoryKey,
  findDataByOptionKey,
  GraphData,
  StudioDataNode,
} from "@agent-studio/studio-dnd";
import { AgentInfo } from "@eternalai-dagent/core";
import { create } from "zustand";
import { PostNewsTopics } from "../categories/x/postNewsOnX/types";
import { CATEGORY_OPTION_KEYS } from "../constants/category-option-keys";
import { BLOCKCHAIN_CATEGORY_KEY } from "../constants/category-keys";
import { AgentDetail } from "../services/apis/studio/types";

export interface AgentSimulate {
  id: string[];
  agent_name?: string;
  personality: string;
  simulate_prompt?: string;
  simulate_type: string;
  fetchTwPosts?: string[];
  topics?: PostNewsTopics;
}

type AgentData = {
  agentName?: string;
  chainId?: string;
};

interface IStore {
  isDetail: boolean;
  setIsDetail: (isDetail: boolean) => void;

  graphData: GraphData;
  setGraphData: (data: GraphData) => void;

  isLoading: boolean;
  setIsLoading: (isLoading: boolean) => void;

  isDisabled: boolean;
  setIsDisabled: (isDisabled: boolean) => void;

  agentInfo?: AgentInfo;
  setAgentInfo: (info: AgentInfo | undefined) => void;

  agentDetail?: AgentDetail;
  setAgentDetail: (agent: AgentDetail | undefined) => void;

  simulatePrompt: AgentSimulate | null;
  setSimulatePrompt: (prompt: AgentSimulate | null) => void;

  isReviewAgentModalOpen: boolean;
  setIsReviewAgentModalOpen: (value: boolean) => void;

  agentTreeData: StudioDataNode | undefined;
  agentData: AgentData;
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

    let treeWithNewAgent;
    try {
      treeWithNewAgent = data.data.find(
        (item) =>
          findDataByOptionKey(
            CATEGORY_OPTION_KEYS.agent.agent_new,
            data.data,
            item.id
          )?.length
      );
    } catch (e) {
      //
    }

    if (treeWithNewAgent) {
      const agentData: AgentData = {};
      set({
        agentTreeData: treeWithNewAgent,
      });

      agentData.agentName = treeWithNewAgent.data?.agentName as string;

      try {
        if (treeWithNewAgent) {
          const blockchains = findDataByCategoryKey(
            BLOCKCHAIN_CATEGORY_KEY,
            [treeWithNewAgent],
            treeWithNewAgent.id
          );
          const network = blockchains[0];
          const chainId = network?.data?.chainId;

          if (chainId) {
            agentData.chainId = chainId as string;
          }
        }
      } catch (e) {
        ///
      }

      set({ agentData });
    } else {
      set({ agentTreeData: undefined, agentData: {} });
    }
  },

  isLoading: false,
  setIsLoading: (isLoading) => set({ isLoading }),

  isDisabled: false,
  setIsDisabled: (isDisabled) => set({ isDisabled }),

  agentInfo: undefined,
  setAgentInfo: (info) => set({ agentInfo: info }),

  agentDetail: undefined,
  setAgentDetail: (agent) => set({ agentDetail: agent }),

  simulatePrompt: null,

  setSimulatePrompt: (prompt: AgentSimulate | null) =>
    set({ simulatePrompt: prompt }),

  isReviewAgentModalOpen: false,
  setIsReviewAgentModalOpen: (value) => set({ isReviewAgentModalOpen: value }),

  agentTreeData: undefined,
  agentData: {},
}));

export default useStudioAgentStore;
