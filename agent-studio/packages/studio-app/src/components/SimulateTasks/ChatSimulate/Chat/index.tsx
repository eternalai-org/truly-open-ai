import { Box, Flex, Image, Text } from "@chakra-ui/react";
import ChatList from "./ChatList";
import { TChatMessage } from "./types";
import useStudioAgentStore from "../../../../stores/useStudioAgentStore";

type Props = {
  messages: TChatMessage[];
  isLoading?: boolean;
};

const SimulateChat = ({ messages, isLoading = false }: Props) => {
  const { simulatePrompt } = useStudioAgentStore();

  if (!simulatePrompt || simulatePrompt?.simulate_type !== "chat") return null;

  return (
    <Flex
      direction={"column"}
      position={"relative"}
      py="20px"
      h="100%"
      overflow={"hidden"}
    >
      <Flex flexDirection={"column"} w={"100%"} h={"100%"} overflow={"hidden"}>
        <Box
          flex={1}
          flexDirection="column"
          gap="20px"
          position="relative"
          overflow={"hidden"}
          display={messages?.length === 0 ? "box" : "flex"}
          mt={messages?.length === 0 ? "3%" : "0"}
        >
          {/* <Back onClose={onClose} /> */}
          {messages?.length === 0 ? (
            <Flex
              direction={"column"}
              alignItems="center"
              justifyContent={"center"}

              // borderRadius={'100px'}
              // gap="4px"
            >
              <Image
                src={"/svg/ic-chat-feel.svg"}
                alt={"image feel"}
                h="100%"
                maxH="200px"
              />
              <Text
                color={"rgba(0, 0, 0, 0.6)"}
                fontSize="16px"
                fontWeight="500"
                // translateY={'-50px'}
                transform={"translateY(-50px)"}
                // marginTop={'-160px'}
              >
                Get a feel for your agent?
              </Text>
            </Flex>
          ) : (
            <Flex
              flexDirection="column"
              position="relative"
              pl="22px"
              flex={1}
              overflow={"hidden"}

              // maxHeight="50vh"
            >
              <ChatList messages={messages || []} isLoading={isLoading} />
            </Flex>
          )}
        </Box>
      </Flex>
    </Flex>
  );
};

export default SimulateChat;
