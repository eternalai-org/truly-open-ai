import useStudioAgentStore from "../stores/useStudioAgentStore";

const useAgentDataFieldChange = () => {
  const { agentTreeData, agentData } = useStudioAgentStore();

  return {
    agentTreeData,
    agentName: agentData.agentName,
    chainId: agentData.chainId,
  };
};

export default useAgentDataFieldChange;
