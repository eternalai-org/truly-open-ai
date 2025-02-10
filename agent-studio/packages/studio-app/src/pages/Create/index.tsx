import {
  GraphData,
  Studio,
  StudioCategory,
  StudioRef,
} from "@agent-studio/studio-dnd";
import { Box, Button, Flex } from "@chakra-ui/react";
import CreateButton from "./CreateButton";
import { useEffect, useMemo, useRef, useState } from "react";
import ModelData from "../../providers/GlobalDataProvider/ModelData";
import useStudioAgentStore from "../../stores/useStudioAgentStore";
import agentDatabase from "../../services/agent-database";
import { NEW_AGENT_ID } from "../../constants/agent";
import ScreenLoading from "../../components/ScreenLoading";
import SimulateTasks from "../../components/SimulateTasks";
import useCommonStore from "../../stores/useCommonStore";

const args = {
  categories: [],
  dataSource: {},
  showConnectLine: true,
};

function Create() {
  const ref = useRef<StudioRef>(null);
  const [isMounted, setIsMounted] = useState(false);

  const { isLoading } = useStudioAgentStore();

  const dataGraph = useMemo(() => {
    return {
      data: [],
      viewport: { x: 0, y: 0, zoom: 1 },
    } satisfies GraphData;
  }, []);

  useEffect(() => {
    useStudioAgentStore.getState().setIsDetail(false);
    const fetchData = async (id: string) => {
      try {
        const data = await agentDatabase.getItem(id);

        const parsedData = JSON.parse(data?.data || JSON.stringify(dataGraph));
        useStudioAgentStore.getState().setGraphData(parsedData);
        if (ref.current) {
          ref.current.redraw(parsedData satisfies GraphData);
        }
      } catch (e) {
        useStudioAgentStore.getState().setGraphData(dataGraph);
        if (ref.current) {
          ref.current.redraw(dataGraph);
        }
      } finally {
        setIsMounted(true);
      }
    };

    fetchData(NEW_AGENT_ID);

    return () => {
      // cleanup
      if (ref.current) {
        ref.current.cleanup();
      }
    };
  }, [ref.current]);

  const renderStudio = (categories: StudioCategory[]) => {
    return (
      <Studio
        {...args}
        categories={categories}
        ref={ref}
        graphData={dataGraph}
        sidebarWidth={"430px"}
        onChange={(graph) => {
          if (isMounted) {
            useStudioAgentStore.getState().setGraphData(graph);
            agentDatabase.upsertItem({
              id: NEW_AGENT_ID,
              data: JSON.stringify(graph),
            });
          }
        }}
      />
    );
  };

  return (
    <ModelData state="create">
      {(categories) => (
        <Flex w="100%" h="100%" position={"relative"}>
          <Box position={"absolute"} zIndex={1} top={"24px"} right={"24px"}>
            <Flex
              flexDir={"row"}
              align={"center"}
              gap={"24px"}
              justifyContent={"flex-end"}
            >
              <Button
                disabled={isLoading}
                onClick={() => {
                  if (isLoading) {
                    return;
                  }
                  useStudioAgentStore.getState().setGraphData(dataGraph);
                  if (ref.current) {
                    ref.current.cleanup();
                  }

                  agentDatabase.upsertItem({
                    id: NEW_AGENT_ID,
                    data: JSON.stringify(dataGraph),
                  });
                }}
              >
                Clear
              </Button>
              <CreateButton />
            </Flex>
          </Box>

          <Box w={"100%"} h={"100%"}>
            <ScreenLoading />
            {renderStudio(categories)}
            <SimulateTasks />
          </Box>
        </Flex>
      )}
    </ModelData>
  );
}

export default Create;
