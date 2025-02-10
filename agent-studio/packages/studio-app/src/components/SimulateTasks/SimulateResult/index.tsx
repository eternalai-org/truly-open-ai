import { Box, Flex, Grid, Image, Text } from "@chakra-ui/react";
import s from "./SimulateResult.module.scss";
import { useEffect, useRef, useState } from "react";
import useStudioAgentStore from "../../../stores/useStudioAgentStore";

type TSimulateResultItem = {
  avatar: string;
  name: string;
  content: string;
};

type Props = {
  upperData: TSimulateResultItem;
  lowerData: TSimulateResultItem;
  isLoading?: boolean;
};

const SimulateResult = ({ upperData, lowerData, isLoading = false }: Props) => {
  const { simulatePrompt } = useStudioAgentStore();
  const contentRef = useRef<HTMLPreElement>(null);

  const [contentHeight, setContentHeight] = useState(0);

  const renderLoadingText = () => {
    if (!simulatePrompt) return null;

    return (
      <Grid h="100%" placeItems={"center"}>
        <Flex alignItems={"center"} flexDir={"column"}>
          <Image
            src="/images/eai-loading-agent2.gif"
            w="120px"
            h="120px"
            filter={"invert(1)"}
          />
          <Text>Processing...</Text>
        </Flex>
      </Grid>
    );
  };

  useEffect(() => {
    if (contentRef.current) {
      setContentHeight(contentRef.current.scrollHeight);
    }
  }, [isLoading, simulatePrompt?.simulate_type]);

  if ((!upperData || !lowerData) && !isLoading) {
    return (
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
    );
  }
  return (
    <Box overflow={"hidden"} h={isLoading ? "500px" : "auto"}>
      {isLoading ? (
        renderLoadingText()
      ) : (
        <Box className={s.container}>
          <Box p="24px 22px">
            <div className={s.wrapper}>
              <Box className={s.left}>
                <div className={s.avatar}>
                  <Image
                    boxSize="44px"
                    rounded="50%"
                    bg="#000"
                    src={upperData?.avatar}
                    onError={(e) => {
                      e.currentTarget.src = "/images/eai-loading-agent2.gif";
                    }}
                  />
                </div>
                <div className={s.line}></div>
              </Box>
              <div className={s.right}>
                <div className={s.agent}>
                  <Text fontSize={"13px"} fontWeight={700}>
                    Instruction
                  </Text>
                </div>
                <div className={s.content_wrapper}>
                  <pre ref={contentRef} className={s.content}>
                    {upperData?.content}
                  </pre>
                </div>
              </div>
            </div>
            {/* {simulateType === OPTION_SOCIAL_TWITTER_REPLY_MENTION_KEY && ( */}
            <div className={s.wrapper}>
              <Box className={s.left}>
                <div className={s.avatar}>
                  <Image
                    boxSize="44px"
                    rounded="50%"
                    bg="#000"
                    src={lowerData?.avatar}
                    onError={(e) => {
                      e.currentTarget.src = "/images/eai-loading-agent2.gif";
                    }}
                  />
                </div>
              </Box>
              <div className={s.right}>
                <div className={s.agent}>
                  <Text fontSize={"15px"} fontWeight={700}>
                    {lowerData?.name}
                  </Text>
                </div>
                <div className={s.content_wrapper}>
                  <pre className={s.content}>{lowerData?.content}</pre>
                </div>
              </div>
            </div>
            {/* )} */}
          </Box>
        </Box>
      )}
    </Box>
  );
};

export default SimulateResult;
