/* eslint-disable @typescript-eslint/no-explicit-any */
import { XYPosition } from '@xyflow/react';

import { createNewBaseEdge, generateSourceHandleId } from './edge';
import { AREA_CLASS_NAMES } from '../constants/area-class-names';
import { StudioCategoryType } from '../enums/category';
import useStudioCategoryStore from '../stores/useStudioCategoryStore';
import useStudioFlowStore from '../stores/useStudioFlowStore';
import useStudioFormStore from '../stores/useStudioFormStore';
import { FormDataMap } from '../types';
import { StudioDataNode, StudioNode, StudioNodeMetadata } from '../types/graph';

import { NodeType } from '@/enums/node-type';

/**
 * Creates configuration for a new base node
 * @param {string} id - Unique node identifier
 * @param {XYPosition} position - Node position coordinates
 * @param {StudioNodeMetadata} metadata - Node metadata
 * @returns {StudioNode} New node configuration
 */
export const createNewNodeBase = (id: string, position: XYPosition, metadata: StudioNodeMetadata): StudioNode => {
  return {
    id,
    type: NodeType.NODE_BASE,
    position,
    data: {
      sourceHandles: [],
      targetHandles: [],
      id,
      metadata,
    },
    dragHandle: `.${AREA_CLASS_NAMES.DRAG_HANDLE}`,
  } satisfies StudioNode;
};

const createLinkedNode = (linkedChild: StudioDataNode, parentNode: StudioNode) => {
  const metadata = {
    children: [],
    idx: linkedChild.idx,
  } satisfies StudioNodeMetadata;

  const position = linkedChild.rect?.position || { x: 0, y: 0 };
  const linkedNode = createNewNodeBase(linkedChild.id, position, metadata);

  const newEdge = createNewBaseEdge(parentNode.id, linkedNode.id, true);
  parentNode.data.sourceHandles.push(generateSourceHandleId(parentNode.id, linkedNode.id));
  useStudioFlowStore.getState().addLinkedNode(parentNode.id, linkedNode.id);
  useStudioFlowStore.getState().addEdge(newEdge);

  if (linkedChild.children.length) {
    // eslint-disable-next-line @typescript-eslint/no-use-before-define
    const childrenOfLinkedNode = transformDataToNodes([linkedChild]);

    return childrenOfLinkedNode;
  }

  return [linkedNode];
};

export const transformDataToNodes = (data: StudioDataNode[]) => {
  const nodes: StudioNode[] = [];
  const categoryOptionMap = useStudioCategoryStore.getState().categoryOptionMap;

  data.forEach((item) => {
    if (item.idx) {
      const position = item.rect?.position || { x: 0, y: 0 };

      const childrenNode: StudioNode[] = [];
      if (item.children.length) {
        // for directly children
        const directlyChildren = item.children.filter((child) => (categoryOptionMap[child.idx].type as any) === StudioCategoryType.INLINE);

        childrenNode.push(...transformDataToNodes(directlyChildren));
      }

      const metadata = {
        children: childrenNode,
        idx: item.idx,
      } satisfies StudioNodeMetadata;

      const node = createNewNodeBase(item.id, position, metadata);
      nodes.push({
        ...node,
        zIndex: categoryOptionMap?.[item.idx]?.zIndex || 0,
      });

      if (item.children.length) {
        // for linked children
        const linkedChildren = item.children.filter((child) => categoryOptionMap[child.idx].type === StudioCategoryType.LINK);

        linkedChildren.forEach((linkedChild) => {
          const newLinkedNodes = createLinkedNode(linkedChild, node);
          nodes.push(...newLinkedNodes);
        });
      }
    }
  });

  return nodes;
};

export const findAncestorNodeIdOfNodeId = (graph: StudioDataNode[], nodeId: string) => {
  for (const node of graph) {
    if (node.id === nodeId) {
      return node.id;
    }
    if (node.children.length > 0) {
      const ancestorId = findAncestorNodeIdOfNodeId(node.children, nodeId);
      if (ancestorId) {
        return node.id;
      }
    }
  }

  return null;
};

export const updateNodeFormData = (nodeId: string, formData: FormDataMap) => {
  if (nodeId) {
    useStudioFormStore.getState().editForm(nodeId, formData);
  }
};
