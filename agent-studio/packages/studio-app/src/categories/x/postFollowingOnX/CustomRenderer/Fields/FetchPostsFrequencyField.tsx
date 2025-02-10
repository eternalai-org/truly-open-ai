import { Box, Button, Flex, Spinner } from "@chakra-ui/react";
import axios from "axios";
import { StudioDataNode } from "@agent-studio/studio-dnd";
import React, { useRef, useState } from "react";
import useStudioAgentStore from "../../../../../stores/useStudioAgentStore";
import { showError } from "../../../../../utils/toast";
import {
  findById,
  findParentById,
  findPersonalityObj,
} from "../../../../../utils/process";
import StudioVerticalField from "../../../../../components/form/fields/StudioVerticalField";
import StudioInput from "../../../../../components/form/inputs/StudioInput";
import ConnectXContent from "../../../../../components/ConnectXContent";
import useAgentServiceStore from "../../../../../stores/useAgentServiceStore";
import { NEWS_QUERY_TOKEN } from "../../../../../configs";

type Props = {
  id: string;
  value: string;
  onChange: (v: string) => void;
  data: StudioDataNode[];
};

const FetchPostsFrequencyField = ({ id, value, onChange, data }: Props) => {
  const { agentInfo, agentDetail, simulatePrompt, setSimulatePrompt } =
    useStudioAgentStore();

  const [loading, setLoading] = useState(false);

  const connectXRef = useRef(null);

  const isLinkedX = React.useMemo(() => {
    return (agentInfo?.twitter_info_id &&
      agentInfo?.twitter_info?.re_link === false) as boolean;
  }, [agentInfo?.twitter_info_id, agentInfo?.twitter_info?.re_link]);

  const handleConnectXSubmit = async () => {
    if (connectXRef.current) {
      // @ts-ignore
      return connectXRef.current.callSubmitFunction();
    }
  };

  const fetchTweetFromFollowingTwitter = async (
    twitter_username: string,
    hours: number
  ) => {
    setLoading(true);
    const authToken = useAgentServiceStore.getState().accessToken;
    try {
      const res = await axios.get(
        `https://offchain-auto-agent-api.eternalai.org/api/post-v3-sample-content/?twitter_username=${twitter_username}&&cutoff_hour=${hours}`,
        {
          headers: {
            Authorization: authToken,
            "X-Token": NEWS_QUERY_TOKEN,
          },
        }
      );

      if (res.data.data.length === 0) {
        showError({
          message: "No tweets to show. Follow more accounts on X.",
        });
        return;
      }

      //   setFetchFollowingPosts(res.data.data);

      const result = findById(data, id);
      const parentObj = findParentById(data, id);
      const personalityObj = findPersonalityObj(data, parentObj);

      if (!personalityObj) {
        showError({
          message:
            "Personality not found. Please connect the ability to a personality",
        });
        return;
      }

      setSimulatePrompt({
        id: [id],
        personality: personalityObj?.data?.personality,
        simulate_prompt: result?.data?.details || "",
        simulate_type: `${result?.idx}_following`,
        fetchTwPosts: res.data.data,
      });

      //   return res.data.data;
    } catch (error: any) {
      showError({
        message: error.message,
      });
    } finally {
      setLoading(false);

      //   setLoadedFetchTweet(true);
    }
  };

  return (
    <>
      <StudioVerticalField
        label="Fetch a tweet from the most engaged posts by users you follow in last"
        tooltip="Frequency information..."
      >
        <Flex alignItems={"center"} gap="8px">
          <StudioInput
            value={value}
            onChange={(e) => onChange(e.target.value)}
            placeholder="e.g 2"
          />
          <Button
            borderRadius="100px"
            bg="black"
            color="white"
            isDisabled={loading}
            w="50px"
            onClick={() => {
              if (loading) return;

              if (!agentDetail) {
                showError({
                  message: "Need to create an agent first",
                });
                return;
              }

              if (isLinkedX && agentDetail) {
                // @ts-ignore

                fetchTweetFromFollowingTwitter(
                  agentDetail?.agent_info?.twitter_info?.twitter_username || "",
                  Number(value)
                );
              } else {
                handleConnectXSubmit();
              }
            }}
            _hover={{
              background: "black",
              opacity: 0.8,
            }}
          >
            {loading ? <Spinner color="white" size="sm"></Spinner> : "Fetch"}
          </Button>
        </Flex>
      </StudioVerticalField>
      <Box position={"absolute"} right="0" top="0" opacity={0}>
        <ConnectXContent
          agentId={`${agentInfo?.id}` || ""}
          isShowSkipBtn={false}
          showOnlyConnectBtn={true}
          ref={connectXRef}
          // className={s.connectX}
        />
      </Box>
    </>
  );
};

export default FetchPostsFrequencyField;
