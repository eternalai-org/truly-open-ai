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

  const link = useCallback((fromNode?: StudioNode, toNode?: StudioNode) => {
    if (!fromNode || !toNode) return;

    const newEdge = createNewBaseEdge(toNode.id, fromNode.id, true);

    toNode.data.sourceHandles.push(generateSourceHandleId(toNode.id, fromNode.id));

    useStudioFlowStore.getState().addEdge(newEdge);
    useStudioFlowStore.getState().addLinkedNode(toNode.id, fromNode.id);
  }, []);

  const unlink = useCallback((fromNode?: StudioNode, toNode?: StudioNode) => {
    if (!fromNode || !toNode) return;

    toNode.data.sourceHandles = toNode.data.sourceHandles.filter((handle) => handle !== generateSourceHandleId(toNode.id, fromNode.id));

    const edge = useStudioFlowStore.getState().edges.find((edge) => edge.source === fromNode.id && edge.target === toNode.id);

    if (edge) {
      useStudioFlowStore.getState().removeEdge(edge.id);
      useStudioFlowStore.getState().removeLinkedNode(fromNode.id, toNode.id);
    }
  }, []);

  const unlinkAll = useCallback((node?: StudioNode) => {
    if (!node) return;

    const edges = useStudioFlowStore.getState().edges.filter((edge) => edge.source === node.id || edge.target === node.id);
    edges.forEach((edge) => {
      useStudioFlowStore.getState().removeEdge(edge.id);
      useStudioFlowStore.getState().removeLinkedNode(edge.source, edge.target);
    });
  }, []);

  const removePartOfPackage = useCallback((node?: StudioNode, index?: number) => {
    if (!node) return {};

    node.data.metadata.children = cloneData(node.data.metadata.children).filter((_, i) => i < (index || 0));

    return {
      sourceNode: node,
    };
  }, []);

  const addToPackage = useCallback((node?: StudioNode, products?: (StudioNode | undefined)[], index = -1) => {
    if (!node || !products) return {};

    if (index === -1) {
      node.data.metadata.children = [...node.data.metadata.children, ...noUndefinedElement(products)];
    } else {
      const befores = node.data.metadata.children.slice(0, index);
      const afters = node.data.metadata.children.slice(index);
      node.data.metadata.children = [...befores, ...noUndefinedElement(products), ...afters];
    }

    return {
      targetNode: node,
    };
  }, []);

  const movePartOfPackage = useCallback((fromNode?: StudioNode, toNode?: StudioNode, fromData?: DraggableData, index = -1) => {
    if (!fromNode || !toNode || !fromData) return {};

    const addons = cloneData(fromNode.data.metadata.children).filter((_, index) => index >= (fromData?.childIndex || 0));

    addToPackage(toNode, addons, index);
    removePartOfPackage(fromNode, fromData?.childIndex || 0);

    return {
      targetNode: toNode,
      sourceNode: fromNode,
    };
  }, []);

  const sortPartOfPackage = useCallback((node?: StudioNode, moveTo = 0, fromTo = 0) => {
    if (!node || moveTo === -1 || fromTo === -1) return {};

    const children = cloneData(node.data.metadata.children);
    const befores = children.slice(0, moveTo);
    const sortItems = children.slice(fromTo);
    const afters = children.slice(moveTo, fromTo);

    node.data.metadata.children = [...befores, ...sortItems, ...afters];

    return {
      targetNode: node,
      sourceNode: node,
    };
  }, []);

  const removeProductAndAllBelong = useCallback((nodeId?: string) => {
    if (!nodeId) return;

    useStudioFlowStore.getState().removeNodeAndAllBelong(nodeId);
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
      link(newNode, rootNode);
    }

    useStudioFlowStore.getState().addNode({
      ...newNode,
      zIndex: fromOption?.zIndex || 0,
    });

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
          link(rootNode, newNode);
        }

        useStudioFlowStore.getState().addNode({
          ...newNode,
          zIndex: fromOption?.zIndex || 0,
        });
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

  const mergeProducts = useCallback((fromNode?: StudioNode, toNode?: StudioNode, fromData?: DraggableData, index = -1) => {
    if (!fromNode || !toNode || !fromData) return {};

    const clonedFromNode = cloneData(fromNode);
    clonedFromNode.data.metadata.children = [];

    const fromNodeLinkedNodes = useStudioFlowStore.getState().linkedNodes[fromNode.id];

    addToPackage(toNode, [clonedFromNode, ...fromNode.data.metadata.children], index);
    removeProduct(fromNode.id);

    fromNodeLinkedNodes?.forEach((linkedNodeId) => {
      link(
        useStudioFlowStore.getState().nodes.find((node) => node.id === linkedNodeId),
        toNode,
      );
    });

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
    removeProductAndAllBelong,
    removeProduct,
    addProduct,
    removePartOfPackage,
    addToPackage,
    movePartOfPackage,
    splitPackage,
    mergeProducts,
    getNewNodeInfo,
    updateFieldValidate,
    link,
    unlink,
    unlinkAll,
    sortPartOfPackage,
  };
};

export default useDndAction;
