import { Studio, StudioRef } from "@agent-studio/studio-dnd";
import getAgentModelCategories from "../../categories";
import { Flex } from "@chakra-ui/react";
import CreateButton from "./CreateButton";
import useStudioAgentStore from "../../stores/useStudioAgentStore";
import { useEffect, useRef } from "react";
import agentDatabase from "../../services/agent-database";
import { NEW_AGENT_ID } from "../../constants/agent";

const args = {
  categories: getAgentModelCategories("create"),
  dataSource: {},
  showConnectLine: true,
};

function Create() {
  const ref = useRef<StudioRef>(null);

  useEffect(() => {
    useStudioAgentStore.getState().setIsDetail(false);

    const fetchData = async (id: string) => {
      try {
        const data = await agentDatabase.getItem(id);

        const parsedData = JSON.parse(data?.data || `[]`);
        useStudioAgentStore.getState().setData(parsedData);
        if (ref.current) {
          ref.current.redraw(parsedData);
        }
      } catch (e) {
        useStudioAgentStore.getState().setData([]);
        if (ref.current) {
          ref.current.redraw([]);
        }
      }
    };

    fetchData(NEW_AGENT_ID);

    return () => {
      // cleanup
      if (ref.current) {
        ref.current.cleanup();
      }
    };
  }, []);

  return (
    <Flex w="100%" h="100%" position={"relative"}>
      <CreateButton />
      <Studio
        {...args}
        ref={ref}
        data={[]}
        onChange={(data) => {
          if (
            JSON.stringify(data) !==
            JSON.stringify(useStudioAgentStore.getState().data)
          ) {
            useStudioAgentStore.getState().setData(data);
            agentDatabase.upsertItem({
              id: NEW_AGENT_ID,
              data: JSON.stringify(data),
            });
          }
        }}
      />
    </Flex>
  );
}

export default Create;
