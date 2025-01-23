import {
  getOptionNodesExistInNode,
  getOptionNodesSameCategoryExistInNode,
  OnAddPayload,
  StudioCategoryOption,
  StudioNode,
} from "@agent-studio/studio-dnd";

export const onlyOptionSameCategoryInTree = (data: {
  toNode: StudioNode;
  option: StudioCategoryOption;
}) => {
  const { toNode, option } = data;
  if (!toNode) {
    return false;
  }

  const optionNodesExistInToNode = getOptionNodesExistInNode(
    toNode.id,
    option.idx
  );
  if (optionNodesExistInToNode.length) {
    return false;
  }

  const optionNodesSameCategoryInToNode = getOptionNodesSameCategoryExistInNode(
    toNode.id,
    option.idx
  );

  if (optionNodesSameCategoryInToNode.length) {
    return false;
  }
  return true;
};
