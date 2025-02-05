import { Box, Button } from "@chakra-ui/react";
import useStudioAgentStore from "../../../stores/useStudioAgentStore";
import { useEffect } from "react";
import {
  findDataByCategoryKey,
  findDataByOptionKey,
  StudioDataNode,
} from "@agent-studio/studio-dnd";
import { CATEGORY_OPTION_KEYS } from "../../../constants/category-option-keys";
import { CATEGORY_KEYS } from "../../../constants/category-keys";
import {
  AgentChainId,
  AgentTokenChainId,
  EFarcasterMissionToolSet,
  ETwitterMissionToolSet,
  FarcasterAgentSnapshotMission,
  IAgentCharacter,
  MissionTypeEnum,
  TokenSetupMode,
  TwitterAgentSnapshotMission,
  IAgent,
} from "@eternal-ai/core";
import { BaseXFormData } from "../../../categories/x/renders/onflow/custom-base/types";
import { getSecondsFromHour } from "../../../utils/time";
import {
  MISSION_FARCASTER_MAPPING,
  MISSION_X_MAPPING,
} from "../../../constants/mapping";
import { BaseFarcasterFormData } from "../../../categories/farcaster/renders/onflow/custom-base/types";
import { useNavigate } from "react-router";
import {
  ETERNAL_AI_URL,
  FARCASTER_CLIENT_ID,
  OWNER_PRIVATE_KEY,
  TWITTER_CLIENT_ID,
} from "../../../configs";
import toast from "react-hot-toast";
import { getAgentInstance } from "../../../utils/agent";
import agentDatabase from "../../../services/agent-database";
import { NEW_AGENT_ID } from "../../../constants/agent";

function CreateButton() {
  let navigate = useNavigate();
  const {
    data,
    isDisabled,
    setIsDisabled,
    isLoading,
    setIsLoading,
    setAgentDetail,
  } = useStudioAgentStore();

  useEffect(() => {
    const runDateValidate = () => {
      try {
        let updatedDisabled = true;
        if (data.length) {
          // find tree has agent_new
          const treeWithNewAgent = data.find((item) =>
            findDataByOptionKey(
              CATEGORY_OPTION_KEYS.agent.agent_new,
              data,
              item.id
            )
          );

          if (
            treeWithNewAgent &&
            (treeWithNewAgent?.data?.agentName as string)?.trim()
          ) {
            // check agent has any personality
            const personalities = findDataByCategoryKey(
              CATEGORY_KEYS.personalities,
              [treeWithNewAgent],
              treeWithNewAgent.id
            );

            if (personalities.length) {
              // create personality
              const firstPersonality = personalities[0];
              if ((firstPersonality?.data?.personality as string)?.trim()) {
                updatedDisabled = false;
              }
            }
          }
        }
        setIsDisabled(updatedDisabled);
      } catch (error) {
        setIsDisabled(true);
      }
    };

    runDateValidate();
  }, [data]);

  const getMissionOnX = (nodeData: StudioDataNode) => {
    const missions = findDataByCategoryKey(
      CATEGORY_KEYS.missionOnX,
      [nodeData],
      nodeData.id
    );
    return missions.map(
      (item) =>
        ({
          user_prompt: (item.data as BaseXFormData).details,
          interval: getSecondsFromHour(
            Number((item.data as BaseXFormData).frequency)
          ),
          tool_set: MISSION_X_MAPPING[item.idx] as ETwitterMissionToolSet,
          agent_type: MissionTypeEnum.CHAT,
        }) as TwitterAgentSnapshotMission
    );
  };

  const getMissionOnFarcaster = (nodeData: StudioDataNode) => {
    const missions = findDataByCategoryKey(
      CATEGORY_KEYS.missionOnFarcaster,
      [nodeData],
      nodeData.id
    );

    return missions.map(
      (item) =>
        ({
          user_prompt: (item.data as BaseFarcasterFormData).details,
          interval: getSecondsFromHour(
            Number((item.data as BaseFarcasterFormData).frequency)
          ),
          tool_set: MISSION_FARCASTER_MAPPING[
            item.idx
          ] as EFarcasterMissionToolSet,
          agent_type: MissionTypeEnum.CHAT,
        }) as FarcasterAgentSnapshotMission
    );
  };

  const handleOnClick = async () => {
    if (isDisabled) {
      return;
    }

    if (isLoading) return;
    try {
      setIsLoading(true);
      const treeWithNewAgent = data.find((item) =>
        findDataByOptionKey(CATEGORY_OPTION_KEYS.agent.agent_new, data, item.id)
      );
      if (!treeWithNewAgent?.data?.agentName) {
        return;
      }
      const agentName = treeWithNewAgent.data.agentName as string;
      const personalities = findDataByCategoryKey(
        CATEGORY_KEYS.personalities,
        [treeWithNewAgent],
        treeWithNewAgent.id
      );

      if (!personalities?.[0].data?.personality) {
        return;
      }

      let systemPrompt = (personalities?.[0].data?.personality as string) || "";

      const tokens = findDataByCategoryKey(
        CATEGORY_KEYS.tokens,
        [treeWithNewAgent],
        treeWithNewAgent.id
      );

      const tokenId = tokens?.[0]?.data?.tokenId as AgentTokenChainId;

      const blockchains = findDataByCategoryKey(
        CATEGORY_KEYS.blockchains,
        [treeWithNewAgent],
        treeWithNewAgent.id
      );

      const chainId = blockchains?.[0]?.data?.chainId as AgentChainId;

      const dagentCharacter: IAgentCharacter = {
        character: {
          chain_id: chainId,
          agent_name: agentName,
          system_content: systemPrompt,
          bio: [],
          lore: [],
          knowledge: [],
          postExamples: [],
          topics: [],
        },
        deployToken: {
          agent_id: "",
          ticker: "",
          create_token_mode: TokenSetupMode.CREATE_TOKEN,
          chain_id: chainId,
          token_chain_id: tokenId,
        },
        twitterMissions: getMissionOnX(treeWithNewAgent),
        farcasterMissions: getMissionOnFarcaster(treeWithNewAgent),
      };

      const baseAgent = await getAgentInstance({ dagentCharacter });
      const res: IAgent = (await baseAgent.createAgent()) as IAgent;

      if (res) {
        agentDatabase.deleteItem(NEW_AGENT_ID);
        agentDatabase.upsertItem({
          id: res.id,
          data: JSON.stringify([treeWithNewAgent]),
        });
        agentDatabase.deleteItem(NEW_AGENT_ID);
        useStudioAgentStore.getState().setAgentDetail(res);
        // await baseAgent.deployToken(res.id);
        navigate(`/${res?.id}`);
      }
      toast.success("Agent created successfully");
    } catch (error) {
      toast.error("Failed to create agent");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Box position={"absolute"} top={4} right={4} zIndex={1}>
      <Button
        colorScheme="blue"
        disabled={isDisabled}
        onClick={handleOnClick}
        isLoading={isLoading}
      >
        Create
      </Button>
    </Box>
  );
}

export default CreateButton;
