import { Text, Flex, Spinner } from "@chakra-ui/react";
import useStudioAgentStore from "../../../stores/useStudioAgentStore";
import LaunchIcon from "../../../components/icons/LaunchIcon";
import { useDataValidates } from "../../../hooks/useDataValidates";
import { showError, showValidateError } from "../../../utils/toast";
import agentDatabase from "../../../services/agent-database";
import { NEW_AGENT_ID } from "../../../constants/agent";
import StudioApi from "../../../services/apis/studio";
import useAgentServiceStore from "../../../stores/useAgentServiceStore";
import { useNavigate } from "react-router";

function CreateButton() {
  const {
    graphData,
    isDisabled,
    isLoading,
    setIsLoading,
    setAgentDetail,
    setIsReviewAgentModalOpen,
  } = useStudioAgentStore();

  let navigate = useNavigate();
  const { runCreateValidate } = useDataValidates();

  const handleCreateAgent = async () => {
    const isValid = runCreateValidate(graphData, true);
    if (!isValid) {
      return;
    }

    if (isDisabled) {
      return;
    }

    if (isLoading) return;

    try {
      setIsLoading(true);

      // TODO: check authen
      if (!useAgentServiceStore.getState().accessToken) {
        // login required
        showValidateError("Login is required");
        return;
      }

      const agents = await StudioApi.createAgent({
        graph_data: JSON.stringify(graphData),
      });
      agentDatabase.deleteItem(NEW_AGENT_ID);
      if (agents.length && agents[0]?.graph_data) {
        agentDatabase.deleteItem(NEW_AGENT_ID);
        useStudioAgentStore
          .getState()
          .setGraphData(JSON.parse(agents[0]?.graph_data));
        setAgentDetail(agents[0]);

        agentDatabase.deleteItem(NEW_AGENT_ID);
        setIsReviewAgentModalOpen(true);
        navigate(`/${agents[0]?.id}`);
        setIsLoading(false);
      }
    } catch (error) {
      showError({
        message: "Something went wrong while creating agent!",
      });
      throw error;
    } finally {
      setIsLoading(false);
    }
  };

  const createAgentOnClick = async () => {
    await handleCreateAgent();
  };

  return (
    <>
      <Flex
        flexDir={"row"}
        align={"center"}
        justify={"center"}
        gap={"24px"}
        bgColor={"#000"}
        p={"7px 24px"}
        borderRadius={"100px"}
        onClick={createAgentOnClick}
        _hover={{
          cursor: "pointer",
          opacity: 0.7,
        }}
        style={{
          opacity: isLoading ? 0.5 : 1,
        }}
      >
        <Text fontSize={"20px"} fontWeight={500} color={"#fff"}>
          Create Agent
        </Text>

        {isLoading ? <Spinner color="#fff" size="md" /> : <LaunchIcon />}
      </Flex>
    </>
  );
}

export default CreateButton;
