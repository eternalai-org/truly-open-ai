import {
  findAncestorNodeIdOfNodeId,
  findDataById,
  findDataByOptionKey,
  StudioDataNode,
  updateNodeFormData,
} from "@agent-studio/studio-dnd";
import { CATEGORY_OPTION_KEYS } from "../../constants/category-option-keys";

const findRelatedAgentData = (graph: StudioDataNode[], fromNodeId: string) => {
  const ancestorId = findAncestorNodeIdOfNodeId(graph, fromNodeId);
  if (ancestorId) {
    const ancestorData = findDataById(ancestorId, graph);
    // find agent node
    if (ancestorData) {
      return findDataByOptionKey(CATEGORY_OPTION_KEYS.agent.agent_new, [
        ancestorData,
      ]);
    }
  }
  return;
};

export const updateRelatedAgentData = (
  graph: StudioDataNode[],
  fromNodeId: string,
  agentName: string
) => {
  const agentData = findRelatedAgentData(graph, fromNodeId);
  if (agentData) {
    // update agent data
    agentData.forEach((agent) => {
      updateNodeFormData(agent.id, {
        agentName: agentName,
      });
    });
  }
};
