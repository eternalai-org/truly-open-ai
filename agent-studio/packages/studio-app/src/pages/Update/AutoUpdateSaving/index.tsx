import { useEffect, useRef } from "react";
import useStudioAgentStore from "../../../stores/useStudioAgentStore";
import { getAllItemData, GraphData } from "@agent-studio/studio-dnd";
import { useDataValidates } from "../../../hooks/useDataValidates";
import StudioApi from "../../../services/apis/studio";
import { compareString } from "../../../utils/string";
import { Box, Text } from "@chakra-ui/react";
import useAgentServiceStore from "../../../stores/useAgentServiceStore";

function AutoUpdateSaving() {
  const walletAddress = useAgentServiceStore((state) => state.walletAddress);

  const { graphData, agentDetail, isLoading, setIsLoading } =
    useStudioAgentStore();

  const refGraphData = useRef<GraphData>(null);
  const refGraphDataPrev = useRef<GraphData>(null);

  const { runUpdateValidate } = useDataValidates();

  const refLocking = useRef(false);
  const retryCallUpdatePayload = useRef<{
    id: string;
    updateData: string;
  } | null>(null);
  const callApiUpdate = async (id: string, updateData: string) => {
    if (refLocking.current) {
      retryCallUpdatePayload.current = { id, updateData };
      return;
    }
    if (isLoading) {
      return;
    }

    try {
      setIsLoading(true);
      refLocking.current = true;

      await StudioApi.updateAgent(id, {
        graph_data: updateData,
      });

      setIsLoading(false);
    } catch (error) {
      throw error;
    } finally {
      setIsLoading(false);
      setTimeout(() => {
        refLocking.current = false;
        if (retryCallUpdatePayload.current) {
          callApiUpdate(
            retryCallUpdatePayload.current.id,
            retryCallUpdatePayload.current.updateData
          );
          retryCallUpdatePayload.current = null;
        }
      }, 5000);
    }
  };

  const updateAgentOnClick = async () => {
    if (!agentDetail) {
      return;
    }

    if (
      !walletAddress ||
      !agentDetail?.creator ||
      !compareString(walletAddress, agentDetail?.creator)
    ) {
      return;
    }

    if (!refGraphDataPrev.current || !refGraphData.current) {
      return;
    }

    const dataUpdateStr = JSON.stringify(refGraphData.current);

    if (dataUpdateStr === agentDetail.graph_data) {
      return;
    }

    if (JSON.stringify(refGraphDataPrev.current) === dataUpdateStr) {
      return;
    }

    try {
      const newAllData = getAllItemData(refGraphData.current.data);
      const prevAllData = getAllItemData(refGraphDataPrev.current.data);
      const isShowToast =
        JSON.stringify(newAllData) !== JSON.stringify(prevAllData);

      if (!runUpdateValidate(refGraphData.current, isShowToast)) {
        return;
      }

      callApiUpdate(agentDetail.id, dataUpdateStr);
    } catch (error) {
      //
    } finally {
      //
    }
  };

  useEffect(() => {
    if (graphData && graphData.data.length) {
      refGraphData.current = graphData;
      updateAgentOnClick();

      return () => {
        if (graphData && graphData.data.length) {
          refGraphDataPrev.current = graphData;
        }
      };
    }
  }, [graphData]);

  if (isLoading) {
    return (
      <Box
        pos={"absolute"}
        zIndex={10}
        bottom="10"
        left="50%"
        transform={"translate(-50%)"}
        opacity={0.5}
      >
        <Text>Saving...</Text>
      </Box>
    );
  }
  return <></>;
}

export default AutoUpdateSaving;
