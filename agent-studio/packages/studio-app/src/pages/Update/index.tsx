import { Flex } from "@chakra-ui/react";
import getAgentModelCategories from "../../categories";
import {
  GraphData,
  Studio,
  StudioDataNode,
  StudioRef,
} from "@agent-studio/studio-dnd";
import { useParams } from "react-router";
import { useEffect, useMemo, useRef, useState } from "react";
import useStudioAgentStore from "../../stores/useStudioAgentStore";
import { getAgentInstance } from "../../utils/agent";
import { IAgent } from "@eternalai-dagent/core";
import agentDatabase from "../../services/agent-database";
import { createGraphDataFromAgentDetail } from "../../utils/data";

const args = {
  dataSource: {},
  showConnectLine: true,
};

const updateAgentModelCategories = getAgentModelCategories("update");
function Update() {
  const { id } = useParams<{ id: string }>();
  const ref = useRef<StudioRef>(null);

  const [categories, setCategories] = useState(updateAgentModelCategories);

  const createGraphDataForNonLocal = (agentDetail: IAgent) => {
    return createGraphDataFromAgentDetail(agentDetail);
  };
  const reMapGraphData = (agentDetail: IAgent, data: StudioDataNode[]) => {
    return createGraphDataFromAgentDetail(agentDetail, data);
  };

  const fetchAgentLocalGraph = async (id: string, agentDetail: IAgent) => {
    // try {
    //   const data = await agentDatabase.getItem(id);
    //   if (data) {
    //     const parsedData = JSON.parse(data?.data || `[]`);
    //     // re-map remote data to local graph data
    //     const graphData = reMapGraphData(agentDetail, parsedData);
    //     useStudioAgentStore.getState().setData(graphData);
    //     if (ref.current) {
    //       ref.current.redraw(graphData);
    //     }
    //     // const parsedData = JSON.parse(data?.data || `[]`);
    //     // useStudioAgentStore.getState().setData(parsedData);
    //     // if (ref.current) {
    //     //   ref.current.redraw(parsedData);
    //     // }
    //   } else {
    //     const graphData = createGraphDataForNonLocal(agentDetail);
    //     useStudioAgentStore.getState().setData(graphData);
    //     if (ref.current) {
    //       ref.current.redraw(graphData);
    //     }
    //   }
    // } catch (e) {
    //   const graphData = createGraphDataForNonLocal(agentDetail);
    //   useStudioAgentStore.getState().setData(graphData);
    //   if (ref.current) {
    //     ref.current.redraw(graphData);
    //   }
    // }
  };

  const getAgentDetail = async () => {
    if (id) {
      const baseAgent = await getAgentInstance({} as any);
      const agentDetail: IAgent = await baseAgent.getAgent(id);
      if (agentDetail) {
        useStudioAgentStore.getState().setAgentDetail(agentDetail);
        fetchAgentLocalGraph(agentDetail.id, agentDetail);
      }
    }
  };

  useEffect(() => {
    getAgentDetail();
  }, [id]);

  useEffect(() => {
    useStudioAgentStore.getState().setIsDetail(true);
    setCategories(updateAgentModelCategories);

    return () => {
      // cleanup
      if (ref.current) {
        ref.current.cleanup();
      }
    };
  }, []);

  const dataGraph = useMemo(() => {
    return {
      data: [],
      viewport: { x: 0, y: 0, zoom: 1 },
    } satisfies GraphData;
  }, []);

  return (
    <Flex w="100%" h="100%" position={"relative"}>
      <Studio
        {...args}
        categories={categories}
        ref={ref}
        graphData={dataGraph}
        onChange={(data) => {
          // useStudioAgentStore.getState().setData(data);
        }}
      />
    </Flex>
  );
}

export default Update;
