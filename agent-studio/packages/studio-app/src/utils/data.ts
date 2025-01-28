import { IAgent } from "@eternal-dagent/core";
import getAgentModelCategories from "../categories";
import {
  createNodeData,
  findDataByCategoryKey,
  findDataByOptionKey,
  StudioDataNode,
} from "@agent-studio/studio-dnd";
import { CATEGORY_OPTION_KEYS } from "../constants/category-option-keys";
import { CATEGORY_KEYS } from "../constants/category-keys";
import { getHourFromSecond } from "./time";
import {
  MISSION_FARCASTER_REVERSE_MAPPING,
  MISSION_X_REVERSE_MAPPING,
  MissionFarcasterSupport,
  MissionXSupport,
} from "../constants/mapping";
import { v4 } from "uuid";
import cloneDeep from "lodash.clonedeep";
import { BaseFarcasterFormData } from "../categories/farcaster/renders/onflow/custom-base/types";
import { BaseXFormData } from "../categories/x/renders/onflow/custom-base/types";

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

export const createGraphDataFromAgentDetail = (
  agentDetail: IAgent,
  graph: StudioDataNode[] = []
): StudioDataNode[] => {
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
  let treeData: StudioDataNode;
  if (agentOptionData.length) {
    treeData = agentOptionData[0];
    agentOptionData[0].data = {
      ...agentOptionData[0].data,
      agentName,
    };
  } else {
    const defaultAgentOption = findCategoryOptionInCategories(
      CATEGORY_KEYS.agent,
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

  // @ts-ignore
  if (!treeData) {
    return newGraph;
  }

  // for personality
  const personality = agentDetail.system_content;

  // check existing personality
  const personalities = findDataByCategoryKey(
    CATEGORY_KEYS.personalities,
    graph
  );
  if (personalities.length) {
    //
  } else {
    // no existing personality
    let newOptionData;
    const defaultPersonalityOption = findCategoryOptionInCategories(
      CATEGORY_KEYS.personalities,
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

    if (newOptionData) {
      treeData.children.push(newOptionData);
    }
  }

  // for block chain
  const blockchainOptionKeys = Object.values(CATEGORY_OPTION_KEYS.blockchains);
  const aiBlockchainOptionData = treeData?.children.filter((item) =>
    blockchainOptionKeys.includes(item.idx)
  );

  if (aiBlockchainOptionData?.[0]) {
    aiBlockchainOptionData[0].data = {
      ...aiBlockchainOptionData[0].data,
      chainId: agentDetail.chain_id,
    };
  }

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
  }

  const decentralizeOptionKeys = Object.values(CATEGORY_OPTION_KEYS.tokens);
  const decentralizeOptionData = treeData?.children.filter((item) =>
    decentralizeOptionKeys.includes(item.idx)
  );

  if (decentralizeOptionData?.[0]) {
    decentralizeOptionData[0].data = {
      ...decentralizeOptionData[0].data,
      decentralizeId: agentDetail.agent_base_model,
    };
  }

  // mission
  const agentSnapshotMission = agentDetail.agent_info?.agent_snapshot_mission;

  // mission on x
  const missionOnXOptionData = findDataByCategoryKey(
    CATEGORY_KEYS.missionOnX,
    graph
  );

  const missionOnXOptionKeys = Object.values(CATEGORY_OPTION_KEYS.missionOnX);

  treeData.children = treeData.children.filter(
    (item) => !missionOnXOptionKeys.includes(item.idx)
  );

  if (agentSnapshotMission.length) {
    agentSnapshotMission.forEach((mission) => {
      const matchedMission = missionOnXOptionData.find(
        (item) => item?.data?.id === (mission as any).id
      );

      if (matchedMission) {
        matchedMission.data = {
          ...matchedMission.data,
          id: (mission as any).id as any,
          frequency: getHourFromSecond(mission.interval) as any,
          details: mission.user_prompt,
        } satisfies BaseXFormData;
        treeData.children.push(matchedMission);
      } else {
        if (MISSION_X_REVERSE_MAPPING[mission.tool_set as MissionXSupport]) {
          const missionOptionData = findCategoryOptionInCategories(
            CATEGORY_KEYS.missionOnX,
            MISSION_X_REVERSE_MAPPING[mission.tool_set as MissionXSupport]
          );
          if (missionOptionData) {
            const newOptionData = createNodeData(
              v4(),
              missionOptionData,
              [],
              {
                id: (mission as any).id as any,
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

  // for mission on farcaster
  const missionOnFarcasterOptionData = findDataByCategoryKey(
    CATEGORY_KEYS.missionOnFarcaster,
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
        (item) => item?.data?.id === (mission as any).id
      );

      if (matchedMission) {
        matchedMission.data = {
          ...matchedMission.data,
          id: (mission as any).id as any,
          frequency: getHourFromSecond(mission.interval) as any,
          details: mission.user_prompt,
        } satisfies BaseFarcasterFormData;
        treeData.children.push(matchedMission);
      } else {
        if (
          MISSION_FARCASTER_REVERSE_MAPPING[
            mission.tool_set as MissionFarcasterSupport
          ]
        ) {
          const missionOptionData = findCategoryOptionInCategories(
            CATEGORY_KEYS.missionOnFarcaster,
            MISSION_FARCASTER_REVERSE_MAPPING[
              mission.tool_set as MissionFarcasterSupport
            ]
          );

          if (missionOptionData) {
            const newOptionData = createNodeData(
              v4(),
              missionOptionData,
              [],
              {
                id: (mission as any).id as any,
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
};
