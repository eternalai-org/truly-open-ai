import {
  findDataByCategoryKey,
  findDataByOptionKey,
  GraphData,
  StudioDataNode,
} from "@agent-studio/studio-dnd";
import { useCallback, useMemo } from "react";
import { CATEGORY_OPTION_KEYS } from "../constants/category-option-keys";
import { removeToast, showValidateError } from "../utils/toast";
import {
  AI_FRAMEWORK_CATEGORY_KEY,
  BLOCKCHAIN_CATEGORY_KEY,
  DECENTRALIZED_INFERENCE_CATEGORY_KEY,
  MISSION_ON_DEFI_CATEGORY_KEY,
  MISSION_ON_FARCASTER_CATEGORY_KEY,
  MISSION_ON_X_CATEGORY_KEY,
  PERSONALITY_CATEGORY_KEY,
  TOKEN_CATEGORY_KEY,
} from "../constants/category-keys";
import { ImportGenomicFormData } from "../categories/personalities/import-genomics/types";
import { KnowledgeFormData } from "../categories/personalities/knowledge/types";
import { ElizaFrameworkFormData } from "../categories/ai-framework/eliza/types";
import { compareString, isJsonString } from "../utils/string";
import { NETWORK_SUPPORT_FOR_KNOWLEDGE_AGENT } from "../constants/validates";
import { SUPPORT_NETWORKS } from "../constants/networks";
import useCommonStore from "../stores/useCommonStore";
import { tokens as networkTokens } from "../constants/tokens";

export const useDataValidates = () => {
  const runValidateAgent = useCallback(
    (agentData: StudioDataNode, showToast = false) => {
      showToast && removeToast();

      if (!agentData) {
        showToast && showValidateError("Agent not found");
        return false;
      }

      if (!(agentData?.data?.agentName as string)?.trim()) {
        showToast && showValidateError("Agent name is required");
        return false;
      }

      // Personality
      const personalities = findDataByCategoryKey(
        PERSONALITY_CATEGORY_KEY,
        [agentData],
        agentData.id
      );

      if (!personalities?.length) {
        showToast && showValidateError("Personality is required");
        return false;
      }

      const personality = personalities[0];

      if (
        personality.idx ===
        CATEGORY_OPTION_KEYS.personalities.personality_customize
      ) {
        if (!personality?.data?.personality) {
          showToast && showValidateError("Personality name is required");
          return false;
        }
      }
      if (
        personality.idx === CATEGORY_OPTION_KEYS.personalities.personality_nft
      ) {
        if (!personality?.data?.personality) {
          showToast && showValidateError("Personality name is required");
          return false;
        }
      }
      if (
        personality.idx ===
        CATEGORY_OPTION_KEYS.personalities.personality_ordinals
      ) {
        if (!personality?.data?.personality) {
          showToast && showValidateError("Personality name is required");
          return false;
        }
      }

      if (
        personality.idx === CATEGORY_OPTION_KEYS.personalities.personality_token
      ) {
        if (!personality?.data?.personality) {
          showToast && showValidateError("Personality name is required");
          return false;
        }
      }

      if (
        personality.idx ===
        CATEGORY_OPTION_KEYS.personalities.personality_genomics
      ) {
        if (
          !(personality?.data as ImportGenomicFormData)?.twitterInfos?.length
        ) {
          showToast && showValidateError("DNA is required");
          return false;
        }
      }

      if (
        personality.idx ===
        CATEGORY_OPTION_KEYS.personalities.personality_knowledge
      ) {
        if (!(personality?.data as KnowledgeFormData)?.name) {
          showToast && showValidateError("Knowledge name is required");
          return false;
        }

        if (!(personality?.data as KnowledgeFormData)?.description) {
          showToast && showValidateError("Knowledge description is required");
          return false;
        }

        if (!(personality?.data as KnowledgeFormData)?.fileUpload) {
          showToast && showValidateError("Upload knowledge files are required");
          return false;
        }
        const fileUpload = (personality?.data as KnowledgeFormData)?.fileUpload;
        const allFileUpdated = (fileUpload || []).every(
          (file: { url: string }) => {
            return !!file?.url;
          }
        );
        if (!allFileUpdated) {
          showToast && showValidateError("Waiting for upload knowledge files");
          return false;
        }
      }

      // ai framework
      const aiFramework = findDataByCategoryKey(
        AI_FRAMEWORK_CATEGORY_KEY,
        [agentData],
        agentData.id
      );
      if (!aiFramework?.length) {
        showToast && showValidateError("AI Framework is required");
        return false;
      }

      if (
        aiFramework[0].idx ===
        CATEGORY_OPTION_KEYS.aiFrameworks.ai_framework_eliza
      ) {
        if (!(aiFramework[0]?.data as ElizaFrameworkFormData)?.config) {
          showToast && showValidateError("Config is required");
          return false;
        }
        if (
          (aiFramework[0]?.data as ElizaFrameworkFormData)?.config &&
          !isJsonString(
            (aiFramework[0]?.data as ElizaFrameworkFormData)?.config
          )
        ) {
          showToast && showValidateError("Config is not a valid JSON");
          return false;
        }
      }

      // blockchains
      const blockchains = findDataByCategoryKey(
        BLOCKCHAIN_CATEGORY_KEY,
        [agentData],
        agentData.id
      );
      if (!blockchains?.length) {
        showToast && showValidateError("Blockchain is required");
        return false;
      }

      // decentralize inference
      const decentralizeInferences = findDataByCategoryKey(
        DECENTRALIZED_INFERENCE_CATEGORY_KEY,
        [agentData],
        agentData.id
      );
      if (!decentralizeInferences?.length) {
        showToast && showValidateError("Decentralize inference is required");
        return false;
      }

      // tokens
      const tokens = findDataByCategoryKey(
        TOKEN_CATEGORY_KEY,
        [agentData],
        agentData.id
      );
      if (!tokens?.length) {
        showToast && showValidateError("Token is required");
        return false;
      }

      // validate knowledge personality
      if (
        personality.idx ===
        CATEGORY_OPTION_KEYS.personalities.personality_knowledge
      ) {
        const network = blockchains[0];
        if (network) {
          if (
            !NETWORK_SUPPORT_FOR_KNOWLEDGE_AGENT.includes(
              network?.data?.chainId as SUPPORT_NETWORKS
            )
          ) {
            showToast &&
              showValidateError("Network doesn't support knowledge agent");
            return;
          }
        }
      }

      // validate ai model
      if (decentralizeInferences[0]) {
        const network = blockchains[0];
        if (network) {
          const aiModel = decentralizeInferences[0];
          const chainId = network?.data?.chainId as string;
          const currentAiModel = aiModel?.data?.decentralizeId as string;

          if (currentAiModel) {
            const chains = useCommonStore.getState().chains || [];
            const token = networkTokens.find((v) =>
              compareString(v.id, chainId)
            );
            if (token) {
              const selectChain = chains.find(
                (chain) => chain.chain_id === token?.chainId
              );

              if (selectChain) {
                const isSupported =
                  !!selectChain?.support_model_names?.[currentAiModel];
                if (!isSupported) {
                  showToast &&
                    showValidateError("Network doesn't support ai model");
                  return false;
                }
              }
            }
          }
        }
      }

      // check x mission if any
      const xMissions = findDataByCategoryKey(
        MISSION_ON_X_CATEGORY_KEY,
        [agentData],
        agentData.id
      );

      if (xMissions.length) {
        const everyXMissionValid = xMissions.every((xMission) => {
          if (
            !xMission?.data?.frequency ||
            !xMission?.data?.details ||
            !xMission?.data?.model
          ) {
            return false;
          }
          return true;
        });

        if (!everyXMissionValid) {
          // showToast && showValidateError('Mission X is invalid');
          if (showToast) {
            // show error for model
            const missingModel = xMissions.find(
              (xMission) => !xMission?.data?.model
            );
            if (missingModel) {
              showValidateError("X mission model is required");
              return;
            }

            const missingDetail = xMissions.find(
              (xMission) => !xMission?.data?.details
            );
            if (missingDetail) {
              showValidateError("X mission instruction is required");
              return;
            }

            const missingFrequency = xMissions.find(
              (xMission) => !xMission?.data?.frequency
            );
            if (missingFrequency) {
              showValidateError("X mission frequency is required");
              return;
            }
          }
          return;
        }
      }

      // check farcaster mission if any
      const farcasterMissions = findDataByCategoryKey(
        MISSION_ON_FARCASTER_CATEGORY_KEY,
        [agentData],
        agentData.id
      );
      if (farcasterMissions.length) {
        const everyFarcasterMissionValid = farcasterMissions.every(
          (farcasterMission) => {
            if (
              !farcasterMission?.data?.frequency ||
              !farcasterMission?.data?.details ||
              !farcasterMission?.data?.model
            ) {
              return false;
            }
            return true;
          }
        );

        if (!everyFarcasterMissionValid) {
          // showToast && showValidateError('Mission Farcaster is invalid');
          if (showToast) {
            // show error for model
            const missingModel = farcasterMissions.find(
              (farcasterMission) => !farcasterMission?.data?.model
            );
            if (missingModel) {
              showValidateError("Farcaster mission model is required");
              return;
            }

            const missingDetail = farcasterMissions.find(
              (farcasterMission) => !farcasterMission?.data?.details
            );
            if (missingDetail) {
              showValidateError("Farcaster mission instruction is required");
              return;
            }

            const missingFrequency = farcasterMissions.find(
              (farcasterMission) => !farcasterMission?.data?.frequency
            );
            if (missingFrequency) {
              showValidateError("Farcaster mission frequency is required");
              return;
            }
          }
          return;
        }
      }

      // check defi mission if any
      const defiMissions = findDataByCategoryKey(
        MISSION_ON_DEFI_CATEGORY_KEY,
        [agentData],
        agentData.id
      );
      if (defiMissions.length) {
        const everyDefiMissionValid = defiMissions.every((defiMission) => {
          if (!defiMission?.data?.frequency || !defiMission?.data?.token) {
            return false;
          }
          return true;
        });

        if (!everyDefiMissionValid) {
          // showToast && showValidateError('Mission Defi is invalid');
          if (showToast) {
            // show error for model
            const missingToken = defiMissions.find(
              (defiMission) => !defiMission?.data?.token
            );
            if (missingToken) {
              showValidateError("Defi mission token is required");
              return;
            }

            const missingFrequency = defiMissions.find(
              (defiMission) => !defiMission?.data?.frequency
            );
            if (missingFrequency) {
              showValidateError("Defi mission frequency is required");
              return;
            }
          }
          return;
        }
      }

      return true;
    },
    []
  );

  const runDataValidate = useCallback(
    (graphData: GraphData, showToast = false) => {
      try {
        if (graphData.data.length) {
          // find tree has agent_new

          // Agent
          const treeWithNewAgent = graphData.data.find(
            (item) =>
              findDataByOptionKey(
                CATEGORY_OPTION_KEYS.agent.agent_new,
                graphData.data,
                item.id
              )?.length
          );

          if (!treeWithNewAgent) {
            showToast && showValidateError("Agent not found");
            return false;
          }

          return runValidateAgent(treeWithNewAgent, showToast);
        } else {
          showToast && showValidateError("Agent not found");
          return false;
        }
      } catch (error) {
        showToast && showValidateError("Validate error");
        return false;
      }
      return false;
    },
    []
  );

  const runCreateValidate = useCallback(
    (graphData: GraphData, showToast = false) => {
      return runDataValidate(graphData, showToast);
    },
    []
  );

  const runUpdateValidate = useCallback(
    (graphData: GraphData, showToast = false) => {
      return runDataValidate(graphData, showToast);
    },
    []
  );

  const memorizedValue = useMemo(() => {
    return {
      runCreateValidate,
      runUpdateValidate,
    };
  }, [runCreateValidate, runUpdateValidate]);

  return memorizedValue;
};
