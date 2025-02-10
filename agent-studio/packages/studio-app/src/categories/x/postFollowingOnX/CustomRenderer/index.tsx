import { Button, Flex, Spinner } from "@chakra-ui/react";
import { useMemo, useState } from "react";
import s from "../../styles.module.scss";
import { PostFollowingOnXFormData } from "../types";
import AiModelField from "./Fields/AiModelField";
import FrequencyField from "./Fields/FrequencyField";
import FetchPostsFrequencyField from "./Fields/FetchPostsFrequencyField";

import axios from "axios";
import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import useStudioAgentStore from "../../../../stores/useStudioAgentStore";
import { showError } from "../../../../utils/toast";
import {
  findById,
  findParentById,
  findPersonalityObj,
} from "../../../../utils/process";
import CustomRendererBase from "../../../../components/CustomRendererBase";
import StudioFormWrapper from "../../../../components/form";
import DetailsField from "./Fields/DetailsField";
import { NEWS_QUERY_TOKEN } from "../../../../configs";
import useAgentServiceStore from "../../../../stores/useAgentServiceStore";

const CustomPostFollowingOnXRenderer = ({
  id,
  formData,
  data,
  setFormFields,
}: StudioCategoryOptionRenderPayload<PostFollowingOnXFormData>) => {
  const { frequency, details, fetchPostsFrequency, model } = formData;

  const [loading, setLoading] = useState(false);

  const { setSimulatePrompt, simulatePrompt, agentDetail } =
    useStudioAgentStore();

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

      return res.data.data;
    } catch (error: any) {
      showError({
        message: error.message,
      });
    } finally {
      setLoading(false);

      //   setLoadedFetchTweet(true);
    }
  };

  const handleSelectMissionPrompt = async () => {
    const result = findById(data, id);
    const parentObj = findParentById(data, id);
    const personalityObj = findPersonalityObj(data, parentObj);

    if (!result?.data?.details) {
      showError({
        message: "Please fill the details field",
      });
      return;
    }

    if (!personalityObj) {
      showError({
        message:
          "Personality not found. Please connect the prompt to a personality",
      });
      return;
    }

    const fetchPosts = await fetchTweetFromFollowingTwitter(
      agentDetail?.agent_info?.twitter_info?.twitter_username || "",
      Number(fetchPostsFrequency)
    );

    if (!!fetchPosts && fetchPosts.length > 0) {
      setSimulatePrompt({
        id: [personalityObj.id, result.id],
        personality: personalityObj?.data?.personality,
        simulate_prompt: result.data.details,
        simulate_type: result.idx,
        fetchTwPosts: fetchPosts,
      });
    }
  };

  const isActive = useMemo(() => {
    return simulatePrompt?.id.includes(id);
  }, [simulatePrompt]);

  return (
    <div data-lego-class={isActive && s.active}>
      <CustomRendererBase
        tag="Ability"
        title="Generate posts from engaged tweet of your followings"
      >
        <StudioFormWrapper>
          <AiModelField
            id={id}
            value={model as string}
            onChange={(v: string) => {
              setFormFields({ model: v });
            }}
            data={data}
          />

          <FrequencyField
            id={id}
            value={frequency as string}
            onChange={(v: string) => {
              setFormFields({ frequency: v });
            }}
          />
          <FetchPostsFrequencyField
            id={id}
            value={fetchPostsFrequency as string}
            onChange={(v: string) => {
              setFormFields({ fetchPostsFrequency: v });
            }}
            data={data}
          />

          <DetailsField
            id={id}
            value={details as string}
            onChange={(v: string) => {
              setFormFields({ details: v });
            }}
          />
          <Flex justifyContent={"flex-end"}>
            <Button
              borderRadius={"100px"}
              width={"fit-content"}
              mb="8px"
              onClick={handleSelectMissionPrompt}
              isDisabled={loading}
              w="55px"
              _hover={{
                background: "black",
                opacity: 0.8,
              }}
            >
              {loading ? <Spinner color="white" size="sm" /> : "Simulate"}
            </Button>
          </Flex>
        </StudioFormWrapper>
      </CustomRendererBase>
    </div>
  );
};

export default CustomPostFollowingOnXRenderer;
