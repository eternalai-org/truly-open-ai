import { Box, Flex } from "@chakra-ui/react";
import {
  GraphData,
  Studio,
  StudioCategory,
  StudioRef,
} from "@agent-studio/studio-dnd";
import { useNavigate, useParams } from "react-router";
import { useEffect, useMemo, useRef, useState } from "react";
import useStudioAgentStore from "../../stores/useStudioAgentStore";
import AgentAPI from "../../services/apis/agent";
import { createGraphDataFromAgentDetail } from "../../utils/data";
import ModelData from "../../providers/GlobalDataProvider/ModelData";
import SimulateTasks from "../../components/SimulateTasks";
import AutoUpdateSaving from "./AutoUpdateSaving";
import useAgentServiceStore from "../../stores/useAgentServiceStore";
import { compareString } from "../../utils/string";

const args = {
  dataSource: {},
  showConnectLine: true,
};

function Update() {
  const walletAddress = useAgentServiceStore((state) => state.walletAddress);
  const params = useParams();
  const ref = useRef<StudioRef>(null);
  const { agentDetail } = useStudioAgentStore();

  const agentId = params?.id as string;
  const [isMounted, setIsMounted] = useState(false);

  const navigate = useNavigate();

  const getAgentInfo = async (id: string) => {
    let res: any = undefined;

    if (isNaN(Number(id))) {
      res = await AgentAPI.getAgent(id as string);
    } else {
      res = await AgentAPI.getAgentByImagine(id as string);
    }

    return res;
  };

  useEffect(() => {
    useStudioAgentStore.getState().setIsDetail(true);

    return () => {
      // cleanup
      if (ref.current) {
        ref.current.cleanup();
      }
    };
  }, []);

  useEffect(() => {
    const fetchRemoteAgent = async () => {
      try {
        const res = await AgentAPI.getAgentDetail(agentId as string);
        const agentInfoRes = await getAgentInfo(agentId as string);

        useStudioAgentStore.getState().setAgentDetail(res);
        useStudioAgentStore.getState().setAgentInfo(agentInfoRes);

        if (res.graph_data) {
          const dataGraph = JSON.parse(res.graph_data);
          useStudioAgentStore.getState().setGraphData(dataGraph);

          ref.current?.redraw(dataGraph);
        } else {
          const dataGraph = {
            data: createGraphDataFromAgentDetail(res),
            viewport: { x: 0, y: 0, zoom: 1 },
          };
          useStudioAgentStore.getState().setGraphData(dataGraph);
          ref.current?.redraw(dataGraph);
        }
      } catch (e) {
        navigate("/");
      } finally {
        setIsMounted(true);
      }
    };

    setIsMounted(false);
    fetchRemoteAgent();
  }, [agentId]);

  const dataGraph = useMemo(() => {
    return {
      data: [],
      viewport: { x: 0, y: 0, zoom: 1 },
    } satisfies GraphData;
  }, []);

  const renderStudio = (categories: StudioCategory[]) => {
    return (
      <Studio
        {...args}
        categories={categories}
        ref={ref}
        graphData={dataGraph}
        sidebarWidth={"430px"}
        onChange={(graph: GraphData) => {
          if (isMounted) {
            useStudioAgentStore.getState().setGraphData(graph);
          }
        }}
      />
    );
  };

  return (
    <ModelData state="update">
      {(categories) => (
        <>
          {/* <ReviewAgentModal /> */}
          <Flex w="100%" h="100%" position={"relative"}>
            {/* <TopWorkArea /> */}
            <AutoUpdateSaving />
            {walletAddress &&
              agentDetail?.creator &&
              compareString(walletAddress, agentDetail?.creator) && (
                <Flex
                  // pos={'absolute'}
                  // zIndex={10}
                  // top="10"
                  // right="10"
                  flexDir={"row"}
                  align={"center"}
                  gap={"24px"}
                  justifyContent={"flex-end"}
                >
                  {/* <TopUpButton /> */}
                </Flex>
              )}

            <Box w={"100%"} h={"100%"}>
              {renderStudio(categories)}
              <SimulateTasks />
            </Box>
          </Flex>
        </>
      )}
    </ModelData>
  );
}

export default Update;
