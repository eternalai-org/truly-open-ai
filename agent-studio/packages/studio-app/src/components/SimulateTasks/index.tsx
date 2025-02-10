import { Box, Flex, Grid, Text } from "@chakra-ui/react";
import ChatSimulate from "./ChatSimulate";
import PostSimulate from "./PostSimulate";
import ReplySimulate from "./ReplySimulate";
import EngageSimulate from "./EngageSimulate";
import PostNewsSimulate from "./PostNewsSimulate";
import PostFollowingSimulate from "./PostFollowingSimulate";
import useStudioAgentStore from "../../stores/useStudioAgentStore";

type Props = {};

const SimulateTasks = (props: Props) => {
  const { simulatePrompt, setSimulatePrompt } = useStudioAgentStore();

  // const { id } = useParams();

  // const renderChat = () => {
  //   return (
  //     <Box w="100%" h="100%" position={'relative'} overflow={'hidden'}>
  //       <PreviewAgentContainer
  //         name={agentDetail?.agent_name || ''}
  //         personality={mergePersonality || ''}
  //         isPreviewable={isPreviewable}
  //         onUpdatePersonality={() => {
  //           //
  //         }}
  //       />
  //     </Box>
  //   );
  // };

  // const [show, setShow] = useState(false);

  if (!simulatePrompt) return null;

  return (
    <Box
      w="500px"
      // maxH={'100%'}
      height={"100%"}
      right="0"
      top="0"
      position={"absolute"}
      borderRadius={"0 12px 12px 0"}
      borderLeft={"1px solid #E5E7EB"}
      bgColor={"white"}
    >
      <Box
        p="8px 12px"
        borderBottom={"1px solid #E0E0E0"}
        background={"#FAFAFA"}
      >
        <Flex alignItems={"center"} justifyContent={"space-between"} mb="8px">
          <Text fontSize={"16px"} fontWeight={600}>
            Test your agent{" "}
            {!!simulatePrompt?.topics?.values &&
              `(Topic: ${simulatePrompt?.topics?.values})`}
          </Text>
          <Text
            onClick={() => {
              setSimulatePrompt(null);
            }}
            display={"flex"}
            alignItems={"center"}
            gap="8px"
            fontSize={"14px"}
            fontWeight={500}
            color="#555"
            cursor={"pointer"}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              viewBox="0 0 16 16"
              fill="none"
            >
              <g clip-path="url(#clip0_48247_1505)">
                <path
                  d="M7.99992 14.6667C11.6818 14.6667 14.6666 11.6819 14.6666 8.00004C14.6666 4.31814 11.6818 1.33337 7.99992 1.33337C4.31802 1.33337 1.33325 4.31814 1.33325 8.00004C1.33325 11.6819 4.31802 14.6667 7.99992 14.6667Z"
                  stroke="black"
                  stroke-width="1.33333"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                />
                <path
                  d="M10 6L6 10"
                  stroke="black"
                  stroke-width="1.33333"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                />
                <path
                  d="M6 6L10 10"
                  stroke="black"
                  stroke-width="1.33333"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                />
              </g>
              <defs>
                <clipPath id="clip0_48247_1505">
                  <rect width="16" height="16" fill="white" />
                </clipPath>
              </defs>
            </svg>
          </Text>
        </Flex>
      </Box>
      <Box p="0" h={"calc(100% - 49px)"} overflowY={"scroll"}>
        <Flex
          // gridTemplateRows={{
          //   base: '1fr auto',
          //   '2k': '1fr auto',
          // }}
          flexDirection={"column"}
          // h="100%"
          minH="100%"
          overflowY={"scroll"}
          className={"hide-scrollbar"}
        >
          <ChatSimulate />
          <PostSimulate />
          <ReplySimulate />
          <EngageSimulate />
          <PostNewsSimulate />
          <PostFollowingSimulate />
        </Flex>
      </Box>
    </Box>
  );
};

export default SimulateTasks;
