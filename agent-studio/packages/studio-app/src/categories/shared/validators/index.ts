import {
  getOptionNodesExistInNode,
  getOptionNodesSameCategoryExistInNode,
  OnLinkPayload,
  StudioCategoryOption,
  StudioNode,
} from "@agent-studio/studio-dnd";
import { showValidateError } from "../../../utils/toast";
import { CATEGORY_OPTION_KEYS } from "../../../constants/category-option-keys";

export const onlyOptionSameCategoryInTree = (data: {
  toNode: StudioNode;
  option: StudioCategoryOption;
}) => {
  const { toNode, option } = data;
  if (!toNode) {
    showValidateError("Cannot add new personality to the root");
    return false;
  }

  const optionNodesExistInToNode = getOptionNodesExistInNode(
    toNode.id,
    option.idx
  );
  if (optionNodesExistInToNode.length) {
    showValidateError("This option already exist");
    return false;
  }

  const optionNodesSameCategoryInToNode = getOptionNodesSameCategoryExistInNode(
    toNode.id,
    option.idx
  );

  if (optionNodesSameCategoryInToNode.length) {
    showValidateError("This option same category already exist");
    return false;
  }
  return true;
};

const ALL_MISSIONS = [
  ...Object.values(CATEGORY_OPTION_KEYS.missionOnX),
  ...Object.values(CATEGORY_OPTION_KEYS.missionOnFarcaster),
  ...Object.values(CATEGORY_OPTION_KEYS.missionOnDefi),
];

export const xMissionLinkToAllowed = ({ toNode }: OnLinkPayload) => {
  if (ALL_MISSIONS.includes(toNode.data.metadata.idx)) {
    showValidateError("Cannot link to mission");
    return false;
  }
  return true;
};

export const farcasterMissionLinkToAllowed = ({ toNode }: OnLinkPayload) => {
  if (ALL_MISSIONS.includes(toNode.data.metadata.idx)) {
    showValidateError("Cannot link to mission");
    return false;
  }
  return true;
};

export const defiMissionLinkToAllowed = ({ toNode }: OnLinkPayload) => {
  if (ALL_MISSIONS.includes(toNode.data.metadata.idx)) {
    showValidateError("Cannot link to mission");
    return false;
  }
  return true;
};
