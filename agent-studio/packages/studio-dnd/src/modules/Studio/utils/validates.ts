import { findAncestorNodeIdOfNodeId } from './node';
import useStudioCategoryStore from '../stores/useStudioCategoryStore';
import useStudioDataStore from '../stores/useStudioDataStore';
import useStudioFlowStore from '../stores/useStudioFlowStore';
import { StudioNode } from '../types';

const getOptionInLinkedNode = (nodeId: string, optionKey: string) => {
  try {
    const nodes = useStudioFlowStore.getState().nodes;
    const linkedNodes = useStudioFlowStore.getState().linkedNodes;
    const linkedNodesId = linkedNodes[nodeId] || [];

    if (linkedNodesId.length) {
      const linkedChildrenNodes = nodes.filter((node) => linkedNodesId.includes(node.id));
      const linkedChildrenOption = linkedChildrenNodes
        .map((linkedNode) => {
          // check children
          const directlyChildrenOptions = linkedNode.data.metadata.children.filter((node) => node.data.metadata.idx === optionKey);

          // check in linked node
          const linkedChildrenOptions = getOptionInLinkedNode(linkedNode.id, optionKey) as StudioNode[];

          return [...directlyChildrenOptions, ...linkedChildrenOptions];
        })
        .flat();

      return linkedChildrenOption;
    }

    return [] as StudioNode[];
  } catch (e) {
    return [];
  }
};

export const getOptionNodesExistInNode = (nodeId: string, optionKey: string, checkAncestor: boolean = false) => {
  try {
    const data = useStudioDataStore.getState().data;
    let targetNodeId = nodeId;
    if (checkAncestor) {
      const foundAncestor = findAncestorNodeIdOfNodeId(data, nodeId);
      if (foundAncestor) {
        targetNodeId = foundAncestor;
      }
    }

    const nodes = useStudioFlowStore.getState().nodes;

    const matchedNode = nodes.find((node) => node.id === targetNodeId);
    if (!matchedNode) {
      return [];
    }

    const foundItem = [];
    // check is self
    if (matchedNode.data.metadata.idx === optionKey) {
      foundItem.push(matchedNode);
    }

    // check children
    const childrenOption = matchedNode.data.metadata.children.filter((node) => node.data.metadata.idx === optionKey);
    foundItem.push(...childrenOption);

    // check in linked node
    const linkedChildrenOptions = getOptionInLinkedNode(matchedNode.id, optionKey) as StudioNode[];
    foundItem.push(...linkedChildrenOptions);

    return foundItem;
  } catch (e) {
    return [];
  }
};

export const getOptionNodesSameCategoryExistInNode = (nodeId: string, optionKey: string, checkAncestor: boolean = false) => {
  try {
    const data = useStudioDataStore.getState().data;
    let targetNodeId = nodeId;
    if (checkAncestor) {
      const foundAncestor = findAncestorNodeIdOfNodeId(data, nodeId);
      if (foundAncestor) {
        targetNodeId = foundAncestor;
      }
    }

    const option = useStudioCategoryStore.getState().categoryOptionMap[optionKey];
    if (!option || !option.parent) {
      return [];
    }

    const optionKeys = option.parent.options.map((option) => option.idx);

    const foundItem = optionKeys
      .map((option) => {
        return getOptionNodesExistInNode(targetNodeId, option, false);
      })
      .flat();

    return foundItem;
  } catch (e) {
    return [];
  }
};

export const getRelatedNodes = (nodeId: string) => {
  try {
    const linkedNodes = useStudioFlowStore.getState().linkedNodes;
    const relatedNodes = linkedNodes[nodeId] || [];

    const nodes = useStudioFlowStore.getState().nodes;

    return nodes.filter((node) => relatedNodes.includes(node.id));
  } catch (e) {
    return [];
  }
};
