import { useStoreApi } from '@xyflow/react';
import { useCallback } from 'react';
import { v4 } from 'uuid';

import useStudioDndStore from '../stores/useStudioDndStore';
import useStudioFlowStore from '../stores/useStudioFlowStore';
import useStudioFlowViewStore from '../stores/useStudioFlowViewStore';
import useStudioFormStore from '../stores/useStudioFormStore';
import { StudioCategory, StudioCategoryOption, StudioCategoryOptionMapValue } from '../types/category';
import { DraggableData } from '../types/dnd';
import { StudioNode } from '../types/graph';
import { noUndefinedElement } from '../utils/array';
import { cloneData, getFormDataFromCategoryOption } from '../utils/data';
import { createNewBaseEdge, generateSourceHandleId } from '../utils/edge';
import { createNewNodeBase } from '../utils/node';

import { isNil } from '@/utils/data';

const useDndAction = () => {
  const flowStore = useStoreApi();

  const getNewNodeInfo = useCallback((idx?: string, option?: StudioCategoryOption, existedId?: string) => {
    if (!idx || !option) return;

    const {
      transform: [transformX, transformY, zoomLevel],
    } = flowStore.getState();

    const { draggingPoint } = useStudioDndStore.getState();

    const mousePosition = useStudioFlowViewStore.getState().mousePosition;
    const transformedX = (mousePosition.x - transformX - (draggingPoint?.x || 0)) / zoomLevel;
    const transformedY = (mousePosition.y - transformY - (draggingPoint?.y || 0)) / zoomLevel;

    const id = existedId || v4();
    const position = {
      x: transformedX,
      y: transformedY,
    };

    const newNodeInfo = createNewNodeBase(id, position, {
      children: [],
      idx,
    });

    if (!existedId) {
      const defaultValues = getFormDataFromCategoryOption(option || {});
      useStudioFormStore.getState().addForm(newNodeInfo.id, {
        ...defaultValues,
      });
    }

    return newNodeInfo;
  }, []);

  const removePartOfPackage = useCallback((node?: StudioNode, index?: number) => {
    if (!node) return {};

    node.data.metadata.children = cloneData(node.data.metadata.children).filter((_, i) => i < (index || 0));

    return {
      sourceNode: node,
    };
  }, []);

  const addToPackage = useCallback((node?: StudioNode, products?: (StudioNode | undefined)[]) => {
    if (!node || !products) return {};

    node.data.metadata.children = [...node.data.metadata.children, ...noUndefinedElement(products)];

    return {
      targetNode: node,
    };
  }, []);

  const movePartOfPackage = useCallback((fromNode?: StudioNode, toNode?: StudioNode, fromData?: DraggableData) => {
    if (!fromNode || !toNode || !fromData) return {};

    const addons = cloneData(fromNode.data.metadata.children).filter((_, index) => index >= (fromData?.childIndex || 0));

    addToPackage(toNode, addons);
    removePartOfPackage(fromNode, fromData?.childIndex || 0);

    return {
      targetNode: toNode,
      sourceNode: fromNode,
    };
  }, []);

  const removeProduct = useCallback((nodeId?: string) => {
    if (!nodeId) return;

    useStudioFlowStore.getState().removeNode(nodeId);
  }, []);

  const addProduct = useCallback((rootNode?: StudioNode, fromData?: DraggableData, fromOption?: StudioCategoryOption) => {
    if (!fromData?.optionKey || !fromOption) return {};

    const newNode = getNewNodeInfo(fromData.optionKey, fromOption);
    if (!newNode) return {};

    if (rootNode) {
      const newEdge = createNewBaseEdge(rootNode.id, newNode.id, true);

      rootNode.data.sourceHandles.push(generateSourceHandleId(rootNode.id, newNode.id));

      useStudioFlowStore.getState().addLinkedNode(rootNode.id, newNode.id);
      useStudioFlowStore.getState().addEdge(newEdge);
    }

    useStudioFlowStore.getState().addNode(newNode);

    return { rootNode, targetNode: newNode };
  }, []);

  const splitPackage = useCallback(
    (rootNode?: StudioNode, fromNode?: StudioNode, fromData?: DraggableData, fromOption?: StudioCategoryOption) => {
      if (!fromNode || !fromData) return {};

      const childData = !isNil(fromData.childIndex) ? fromNode.data.metadata.children[fromData.childIndex as number] : null;

      const newNode = getNewNodeInfo(fromData.optionKey, fromOption, childData?.id);

      if (newNode) {
        newNode.data.metadata.children = cloneData(fromNode.data.metadata.children)
          .filter((_, index) => index > (fromData?.childIndex || 0))
          .map((child) => getNewNodeInfo(child.data.metadata.idx, fromOption, child.id))
          .filter((child) => !!child);

        if (rootNode) {
          const newEdge = createNewBaseEdge(rootNode.id, newNode.id, true);

          rootNode.data.sourceHandles.push(generateSourceHandleId(rootNode.id, newNode.id));

          useStudioFlowStore.getState().addEdge(newEdge);
          useStudioFlowStore.getState().addLinkedNode(rootNode.id, newNode.id);
        }

        useStudioFlowStore.getState().addNode(newNode);
      }

      fromNode.data.metadata.children = fromNode.data.metadata.children.filter((_, index) => index < (fromData?.childIndex || 0));

      return {
        rootNode,
        sourceNode: fromNode,
        targetNode: newNode,
      };
    },
    [],
  );

  const mergeProducts = useCallback((fromNode?: StudioNode, toNode?: StudioNode, fromData?: DraggableData) => {
    if (!fromNode || !toNode || !fromData) return {};

    const clonedFromNode = cloneData(fromNode);
    clonedFromNode.data.metadata.children = [];

    addToPackage(toNode, [clonedFromNode, ...fromNode.data.metadata.children]);
    removeProduct(fromNode.id);

    return {
      sourceNode: fromNode,
      targetNode: toNode,
    };
  }, []);

  const updateFieldValidate = useCallback(
    ({ fromCategory, fromOption }: { fromCategory: StudioCategory | undefined; fromOption: StudioCategoryOptionMapValue | undefined }) => {
      if (!fromCategory || !fromOption) return;
    },
    [],
  );

  return {
    removeProduct,
    addProduct,
    removePartOfPackage,
    addToPackage,
    movePartOfPackage,
    splitPackage,
    mergeProducts,
    getNewNodeInfo,
    updateFieldValidate,
  };
};

export default useDndAction;
