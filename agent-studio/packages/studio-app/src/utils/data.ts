import {
  AGENT_CATEGORY_KEY,
  AI_FRAMEWORK_CATEGORY_KEY,
  BLOCKCHAIN_CATEGORY_KEY,
  DECENTRALIZED_INFERENCE_CATEGORY_KEY,
  MISSION_ON_FARCASTER_CATEGORY_KEY,
  MISSION_ON_X_CATEGORY_KEY,
  PERSONALITY_CATEGORY_KEY,
  TOKEN_CATEGORY_KEY,
} from "../constants/category-keys";
import { v4 } from "uuid";
import getAgentModelCategories from "../categories/categories";
import { CATEGORY_OPTION_KEYS } from "../constants/category-option-keys";

import {
  MISSION_FARCASTER_REVERSE_MAPPING,
  MISSION_X_REVERSE_MAPPING,
  MissionFarcasterSupport,
  MissionXSupport,
} from "../constants/mapping";
import { EngageOnXFormData } from "../categories/x/engageOnX/types";

import cloneDeep from "lodash.clonedeep";
import { ChainId, SUPPORT_NETWORKS } from "../constants/networks";
import { PostNewsOnXFormData } from "../categories/x/postNewsOnX/types";

import { PostFollowingOnXFormData } from "../categories/x/postFollowingOnX/types";
import { getDynamicAiModelOptions } from "./category";
import { AgentDetail } from "../services/apis/studio/types";
import {
  createNodeData,
  findDataByCategoryKey,
  findDataByOptionKey,
  StudioDataNode,
} from "@agent-studio/studio-dnd";
import { tokens } from "../constants/tokens";
import { TokenSetupMode } from "@eternalai-dagent/core";
import { OPTION_ETERNAL_AI_ID } from "../constants/option-values";
import { compareString } from "./string";
import useCommonStore from "../stores/useCommonStore";
import { getHourFromSecond } from "./time";

const categories = getAgentModelCategories("create");

const findCategoryInCategories = (categoryKey: string) => {
  return categories.find((category) => category.idx === categoryKey);
};

const findCategoryOptionInCategories = (
  categoryKey: string,
  optionKey: string
) => {
  const category = findCategoryInCategories(categoryKey);
  if (category) {
    return category.options.find((option) => option.idx === optionKey);
  }
};

const findCategoryOptionInCategoriesByFieldValue = (
  categoryKey: string,
  dataField: string,
  dataValue: string | number
) => {
  const category = findCategoryInCategories(categoryKey);
  if (category) {
    return category.options.find(
      (option) =>
        `${option?.data?.[dataField].defaultValue}`.toLowerCase() ===
        `${dataValue}`.toLowerCase()
    );
  }
};

export const createGraphDataFromAgentDetail = (
  agentDetail: AgentDetail,
  graph: StudioDataNode[] = []
): StudioDataNode[] => {
  try {
    console.log("___________________createGraphDataFromAgentDetail", {
      agentDetail,
      graph,
    });

    const POSITION_OFFSET = 550;
    const MAX_COLUMN = 2;
    let position = { x: 0, y: 0 };

    const updatePosition = () => {
      position.x += POSITION_OFFSET;

      if (position.x > POSITION_OFFSET * MAX_COLUMN) {
        position.x = POSITION_OFFSET;
        position.y += POSITION_OFFSET;
      }
    };

    const newGraph = JSON.parse(JSON.stringify(graph));
    // for agent
    const agentName = agentDetail.agent_name;
    const agentOptionData = findDataByOptionKey(
      CATEGORY_OPTION_KEYS.agent.agent_new,
      newGraph
    );
    let treeData: StudioDataNode = undefined as any;
    // agent name
    if (agentOptionData.length) {
      treeData = agentOptionData[0];
      agentOptionData[0].data = {
        ...agentOptionData[0].data,
        agentName,
      };
    } else {
      const defaultAgentOption = findCategoryOptionInCategories(
        AGENT_CATEGORY_KEY,
        CATEGORY_OPTION_KEYS.agent.agent_new
      );

      if (defaultAgentOption) {
        const newOptionData = createNodeData(
          v4(),
          defaultAgentOption,
          [],
          {
            agentName,
          },
          cloneDeep(position)
        );

        updatePosition();

        newGraph.push(newOptionData);
        treeData = newOptionData;
      }
    }

    if (!treeData) {
      return newGraph;
    }

    // for personality
    const personality = agentDetail.system_content;

    // check existing personality
    const personalities = findDataByCategoryKey(
      PERSONALITY_CATEGORY_KEY,
      graph
    );
    if (personalities.length) {
      //
    } else {
      // no existing personality
      let newOptionData;
      if (agentDetail.create_token_mode === TokenSetupMode.LINK_EXISTING) {
        // pump token
        const defaultPersonalityOption = findCategoryOptionInCategories(
          PERSONALITY_CATEGORY_KEY,
          CATEGORY_OPTION_KEYS.personalities.personality_token
        );
        if (defaultPersonalityOption) {
          newOptionData = createNodeData(
            v4(),
            defaultPersonalityOption,
            [],
            {
              personality: personality,
            },
            { x: 0, y: 0 }
          );
        }
      } else {
        if (agentDetail.nft_token_id) {
          if (agentDetail.nft_public_key) {
            // ordinals
            const defaultPersonalityOption = findCategoryOptionInCategories(
              PERSONALITY_CATEGORY_KEY,
              CATEGORY_OPTION_KEYS.personalities.personality_ordinals
            );
            if (defaultPersonalityOption) {
              newOptionData = createNodeData(
                v4(),
                defaultPersonalityOption,
                [],
                {
                  personality: personality,
                },
                { x: 0, y: 0 }
              );
            }
          } else {
            // nft
            const defaultPersonalityOption = findCategoryOptionInCategories(
              PERSONALITY_CATEGORY_KEY,
              CATEGORY_OPTION_KEYS.personalities.personality_nft
            );
            if (defaultPersonalityOption) {
              newOptionData = createNodeData(
                v4(),
                defaultPersonalityOption,
                [],
                {
                  personality: personality,
                },
                { x: 0, y: 0 }
              );
            }
          }
        } else {
          // custom
          const defaultPersonalityOption = findCategoryOptionInCategories(
            PERSONALITY_CATEGORY_KEY,
            CATEGORY_OPTION_KEYS.personalities.personality_customize
          );
          if (defaultPersonalityOption) {
            newOptionData = createNodeData(
              v4(),
              defaultPersonalityOption,
              [],
              {
                personality: personality,
                originalText: personality,
              },
              { x: 0, y: 0 }
            );
          }
        }
      }

      if (newOptionData) {
        treeData.children.push(newOptionData);
      }
    }

    // for ai framework
    const aiFrameworkOptionKeys = Object.values(
      CATEGORY_OPTION_KEYS.aiFrameworks
    );
    const aiFrameworkOptionData = treeData?.children.filter((item) =>
      aiFrameworkOptionKeys.includes(item.idx)
    );

    if (aiFrameworkOptionData?.[0]) {
      aiFrameworkOptionData[0].data = {
        ...aiFrameworkOptionData[0].data,
        aiFrameworkId: OPTION_ETERNAL_AI_ID,
      };
    } else {
      const defaultAiFrameworkOption = findCategoryOptionInCategories(
        AI_FRAMEWORK_CATEGORY_KEY,
        CATEGORY_OPTION_KEYS.aiFrameworks.ai_framework_eternal_ai
      );

      if (defaultAiFrameworkOption) {
        const newOptionData = createNodeData(
          v4(),
          defaultAiFrameworkOption,
          [],
          {
            aiFrameworkId: OPTION_ETERNAL_AI_ID,
          },
          cloneDeep(position)
        );

        updatePosition();

        treeData.children.push(newOptionData);
      }
    }

    // for block chain
    const blockchainOptionKeys = Object.values(
      CATEGORY_OPTION_KEYS.blockchains
    );
    const aiBlockchainOptionData = treeData?.children.filter((item) =>
      blockchainOptionKeys.includes(item.idx)
    );

    let networkId = "";
    if (aiBlockchainOptionData?.[0]) {
      networkId = agentDetail.chain_id;
      aiBlockchainOptionData[0].data = {
        ...aiBlockchainOptionData[0].data,
        chainId: agentDetail.chain_id,
      };
    } else {
      const networkName =
        agentDetail.agent_info.network_name || SUPPORT_NETWORKS.BASE;
      const defaultBlockchainOption =
        findCategoryOptionInCategoriesByFieldValue(
          BLOCKCHAIN_CATEGORY_KEY,
          "chainId",
          networkName
        );

      if (defaultBlockchainOption) {
        networkId = `${networkName}`.toLowerCase();
        const newOptionData = createNodeData(
          v4(),
          defaultBlockchainOption,
          [],
          {
            chainId: `${networkName}`.toLowerCase(),
          },
          cloneDeep(position)
        );

        updatePosition();

        treeData.children.push(newOptionData);
      }
    }

    const matchedToken = tokens.find((v) => compareString(v.id, networkId));

    const selectedChain = useCommonStore
      .getState()
      .chains?.find((v) => compareString(v.chain_id, matchedToken?.chainId));

    // for token
    const tokenOptionKeys = Object.values(CATEGORY_OPTION_KEYS.tokens);
    const tokenOptionData = treeData?.children.filter((item) =>
      tokenOptionKeys.includes(item.idx)
    );

    if (tokenOptionData?.[0]) {
      tokenOptionData[0].data = {
        ...tokenOptionData[0].data,
        tokenId: agentDetail.token_chain_id,
      };
    } else {
      const tokenChainId = agentDetail.token_chain_id || ChainId.Base;
      const defaultToken = findCategoryOptionInCategoriesByFieldValue(
        TOKEN_CATEGORY_KEY,
        "tokenId",
        tokenChainId
      );

      if (defaultToken) {
        const newOptionData = createNodeData(
          v4(),
          defaultToken,
          [],
          {
            tokenId: tokenChainId,
          },
          cloneDeep(position)
        );

        updatePosition();

        treeData.children.push(newOptionData);
      }
    }

    const decentralizeOptionKeys = Object.values(
      CATEGORY_OPTION_KEYS.decentralized
    );

    const decentralizeOptionData = treeData?.children.filter((item) =>
      decentralizeOptionKeys.includes(item.idx)
    );

    if (decentralizeOptionData?.[0]) {
      const aiModel = [
        agentDetail.agent_base_model,
        selectedChain?.support_model_names?.[
          agentDetail.agent_base_model || ""
        ] || "",
      ];
      decentralizeOptionData[0].data = {
        ...decentralizeOptionData[0].data,
        decentralizeId: aiModel[0],
        decentralizeKey: aiModel[1],
      };
    } else {
      const newOptions = getDynamicAiModelOptions(
        useCommonStore.getState().chains || [],
        {}
      );
      const aiModel = [
        agentDetail.agent_base_model,
        selectedChain?.support_model_names?.[
          agentDetail.agent_base_model || ""
        ] || "",
      ];

      const option = newOptions.find(
        (item) =>
          item.idx === `decentralize_inference_${agentDetail.agent_base_model}`
      );

      if (option) {
        const newOptionData = createNodeData(
          v4(),
          option,
          [],
          {
            decentralizeId: aiModel[0],
            decentralizeKey: aiModel[1],
          },
          cloneDeep(position)
        );

        updatePosition();

        treeData.children.push(newOptionData);
      }
    }

    // mission
    const agentSnapshotMission = agentDetail.agent_info?.agent_snapshot_mission;

    // mission on x
    const missionOnXOptionData = findDataByCategoryKey(
      MISSION_ON_X_CATEGORY_KEY,
      graph
    );

    const missionOnXOptionKeys = Object.values(CATEGORY_OPTION_KEYS.missionOnX);

    treeData.children = treeData.children.filter(
      (item) => !missionOnXOptionKeys.includes(item.idx)
    );

    if (agentSnapshotMission.length) {
      agentSnapshotMission.forEach((mission) => {
        const matchedMission = missionOnXOptionData.find(
          // @ts-ignore
          (item) => item?.data?.id === mission.id
        );

        if (matchedMission) {
          matchedMission.data = {
            ...matchedMission.data,
            // @ts-ignore
            id: mission.id as any,
            frequency: getHourFromSecond(mission.interval) as any,
            details: mission.user_prompt,
            toolset: "",
            model: "",
            modelName: "",
          } satisfies EngageOnXFormData;
          treeData.children.push(matchedMission);
        } else {
          // @ts-ignore
          if (MISSION_X_REVERSE_MAPPING[mission.tool_set as MissionXSupport]) {
            const missionOptionData = findCategoryOptionInCategories(
              MISSION_ON_X_CATEGORY_KEY,
              // @ts-ignore
              MISSION_X_REVERSE_MAPPING[mission.tool_set as MissionXSupport]
            );

            if (missionOptionData) {
              const firstModel = Object.entries(
                selectedChain?.support_model_names || {}
              )[0];
              // @ts-ignore
              const model =
                selectedChain?.support_model_names?.[
                  // @ts-ignore
                  mission.agent_base_model
                ] || firstModel[1];
              // @ts-ignore
              const modelName = mission.agent_base_model || firstModel[0];

              if (
                missionOptionData.idx ===
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_news
              ) {
                const { prompt } = JSON.parse(mission.user_prompt);

                const newOptionData = createNodeData(
                  v4(),
                  missionOptionData,
                  [],
                  {
                    // @ts-ignore
                    id: mission.id as any,
                    frequency: getHourFromSecond(mission.interval) as any,
                    details: prompt,
                    topics: {
                      // @ts-ignore
                      values: mission.topics || "",
                      // @ts-ignore
                      bingSearch: mission.is_bing_search || false,
                      // @ts-ignore
                      twitterSearch: mission.is_twitter_search || false,
                    },
                    model: model,
                    modelName: modelName,
                  } satisfies PostNewsOnXFormData,
                  cloneDeep(position)
                );

                updatePosition();

                treeData.children.push(newOptionData);
              } else if (
                missionOptionData.idx ===
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_following
              ) {
                const { prompt, cutoff_hour } = JSON.parse(mission.user_prompt);
                const newOptionData = createNodeData(
                  v4(),
                  missionOptionData,
                  [],
                  {
                    // @ts-ignore
                    id: mission.id as any,
                    // @ts-ignore
                    frequency: getHourFromSecond(mission.interval) as any,
                    details: prompt,
                    fetchPostsFrequency: cutoff_hour,
                    model: model,
                    // @ts-ignore
                    modelName: modelName,
                  } satisfies PostFollowingOnXFormData,
                  cloneDeep(position)
                );
                updatePosition();

                treeData.children.push(newOptionData);
              } else {
                const newOptionData = createNodeData(
                  v4(),
                  missionOptionData,
                  [],
                  {
                    // @ts-ignore
                    id: mission.id as any,
                    frequency: getHourFromSecond(mission.interval) as any,
                    details: mission.user_prompt,
                    model: model,
                    modelName: modelName,
                  },
                  cloneDeep(position)
                );

                updatePosition();

                treeData.children.push(newOptionData);
              }
            }
          }
        }
      });
    }

    // for mission on farcaster
    const missionOnFarcasterOptionData = findDataByCategoryKey(
      MISSION_ON_FARCASTER_CATEGORY_KEY,
      graph
    );

    const missionOnFarcasterOptionKeys = Object.values(
      CATEGORY_OPTION_KEYS.missionOnFarcaster
    );

    treeData.children = treeData.children.filter(
      (item) => !missionOnFarcasterOptionKeys.includes(item.idx)
    );

    if (agentSnapshotMission.length) {
      agentSnapshotMission.forEach((mission) => {
        const matchedMission = missionOnFarcasterOptionData.find(
          // @ts-ignore
          (item) => item?.data?.id === mission.id
        );

        if (matchedMission) {
          matchedMission.data = {
            ...matchedMission.data,
            // @ts-ignore
            id: mission.id as any,
            frequency: getHourFromSecond(mission.interval) as any,
            details: mission.user_prompt,
            toolset: "",
            model: "",
            modelName: "",
          } satisfies EngageOnXFormData;
          treeData.children.push(matchedMission);
        } else {
          if (
            MISSION_FARCASTER_REVERSE_MAPPING[
              // @ts-ignore
              mission.tool_set as MissionFarcasterSupport
            ]
          ) {
            const missionOptionData = findCategoryOptionInCategories(
              MISSION_ON_FARCASTER_CATEGORY_KEY,
              MISSION_FARCASTER_REVERSE_MAPPING[
                // @ts-ignore
                mission.tool_set as MissionFarcasterSupport
              ]
            );

            if (missionOptionData) {
              const newOptionData = createNodeData(
                v4(),
                missionOptionData,
                [],
                {
                  // @ts-ignore
                  id: mission.id as any,
                  frequency: getHourFromSecond(mission.interval) as any,
                  details: mission.user_prompt,
                },
                cloneDeep(position)
              );

              updatePosition();

              treeData.children.push(newOptionData);
            }
          }
        }
      });
    }

    return newGraph;
  } catch (e) {
    console.log("_____________createGraphDataFromAgentDetail", e);
    return [];
  }
};
