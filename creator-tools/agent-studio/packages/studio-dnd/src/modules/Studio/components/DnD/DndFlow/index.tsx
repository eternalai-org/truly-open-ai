import {
  closestCenter,
  DndContext,
  DragAbortEvent,
  DragCancelEvent,
  DragEndEvent,
  DragMoveEvent,
  DragOverEvent,
  DragPendingEvent,
  DragStartEvent,
  MouseSensor,
  rectIntersection,
  useSensor,
  useSensors,
} from '@dnd-kit/core';
import { applyNodeChanges, XYPosition } from '@xyflow/react';
import { PropsWithChildren, useCallback, useRef } from 'react';

import { StudioCategoryType } from '@/modules/Studio/enums/category';
import useDndAction from '@/modules/Studio/hooks/useDndAction';
import useDndInteraction from '@/modules/Studio/hooks/useDndInteraction';
import useStudioCategoryStore from '@/modules/Studio/stores/useStudioCategoryStore';
import useStudioDataStore from '@/modules/Studio/stores/useStudioDataStore';
import useStudioFlowStore from '@/modules/Studio/stores/useStudioFlowStore';
import useStudioFormStore from '@/modules/Studio/stores/useStudioFormStore';
import { StudioCategory, StudioCategoryMapValue, StudioCategoryOptionMapValue } from '@/modules/Studio/types/category';
import { DraggableData, StudioZone } from '@/modules/Studio/types/dnd';
import { StudioNode } from '@/modules/Studio/types/graph';

function DndFlow({ children }: PropsWithChildren) {
  const sensors = useSensors(useSensor(MouseSensor, { activationConstraint: { distance: 5 } }));

  const movingNodeRef = useRef<StudioNode>(null);

  const {
    addProduct,
    movePartOfPackage,
    removeProduct,
    removePartOfPackage,
    addToPackage,
    splitPackage,
    mergeProducts,
    getNewNodeInfo,
    updateFieldValidate,
  } = useDndAction();
  const { updateNodes } = useDndInteraction();

  const handleDragStart = useCallback((_event: DragStartEvent) => {
    movingNodeRef.current = null;
  }, []);

  const handleDragEnd = useCallback((event: DragEndEvent) => {
    console.log('[DndContainer] handleDragEnd', {
      categoryMap: useStudioCategoryStore.getState().categoryMap,
      categoryOptionMap: useStudioCategoryStore.getState().categoryOptionMap,
    });

    const { active, over } = event;

    const data = useStudioDataStore.getState().data;

    const rootCategory = useStudioCategoryStore.getState().rootCategory;
    const rootOptionKey = rootCategory?.options?.map((item) => item?.idx);

    const rootData = data.find((item) => item.idx === rootCategory?.idx || rootOptionKey?.includes(item.idx));
    const rootNode = useStudioFlowStore.getState().nodes.find((node) => node.id === rootData?.id);

    const fromData = active?.data?.current as DraggableData;
    const from = fromData?.type;
    const fromCategoryKey = fromData?.categoryKey;
    const fromOptionKey = fromData?.optionKey;
    const fromCategory: StudioCategory | undefined =
      useStudioCategoryStore.getState().categoryMap[fromCategoryKey || ''];
    const fromOption: StudioCategoryOptionMapValue | undefined =
      useStudioCategoryStore.getState().categoryOptionMap[fromOptionKey || ''];
    const fromNode = useStudioFlowStore.getState().nodes.find((node) => node.id === fromData?.belongsTo);

    const toData = over?.data?.current as DraggableData;
    const to = toData?.type;
    const toCategoryKey = toData?.categoryKey;
    const toOptionKey = toData?.optionKey;
    const toCategory: StudioCategoryMapValue | undefined =
      useStudioCategoryStore.getState().categoryMap[toCategoryKey || ''];
    const toOption: StudioCategoryOptionMapValue | undefined =
      useStudioCategoryStore.getState().categoryOptionMap[toOptionKey || ''];
    const toNode = useStudioFlowStore.getState().nodes.find((node) => node.id === toData?.belongsTo);

    const isTheSameNode = fromNode?.id === toNode?.id;

    const allFormData = useStudioFormStore.getState().formMap;
    const currentFormData = allFormData[fromData.belongsTo || ''];

    const parentOption = fromOption?.parent;

    console.log('[DndContainer] handleDragEnd from to', {
      from,
      to,
    });

    console.log('[DndContainer] handleDragEnd option', {
      fromOption,
      toOption,
    });

    console.log('[DndContainer] handleDragEnd data', {
      fromData,
      toData,
    });

    console.log('[DndContainer] handleDragEnd category', {
      fromCategory,
      toCategory,
    });

    console.log('[DndContainer] handleDragEnd root', {
      rootCategory,
      rootNode,
      rootData,
      parentOption,
    });

    console.log('[DndContainer] handleDragEnd node', {
      fromNode,
      toNode,
    });

    if (to === StudioZone.ZONE_DISTRIBUTION) {
      // Create
      if (from === StudioZone.ZONE_SOURCE && fromOption?.type !== StudioCategoryType.LINK) {
        const isValid =
          fromOption?.onDropInValidate?.({
            option: fromOption,
            parentOption,
            formData: currentFormData,
            allFormData,
            data,
          }) ?? true;

        if (!isValid) return;

        addProduct(rootNode, fromData, fromOption);
        updateNodes([rootNode]);

        return;
      }

      // Split
      if (
        from === StudioZone.ZONE_PRODUCT_ADDON &&
        !isTheSameNode &&
        fromNode &&
        fromOption?.type !== StudioCategoryType.LINK
      ) {
        if (!fromData.belongsTo) return;

        const isValid =
          fromOption.onSplitValidate?.({
            id: fromData.belongsTo,
            option: fromOption,
            parentOption,
            formData: currentFormData,
            allFormData,
            fromNode,
            data,
          }) ?? true;

        if (!isValid) return;

        splitPackage(rootNode, fromNode, fromData, fromOption);
        updateNodes([rootNode, fromNode]);

        return;
      }
    }

    if (to === StudioZone.ZONE_PACKAGE && toNode) {
      // Link
      if (from === StudioZone.ZONE_SOURCE && fromOption?.type === StudioCategoryType.LINK) {
        console.log('[DndContainer] handleDragEnd link', {
          fromOption,
          toOption,
          toNode,
        });

        if (!toOption || !toNode) return;

        const isValid =
          fromOption.onLinkValidate?.({
            option: fromOption,
            parentOption,
            formData: currentFormData,
            allFormData,
            data,
            toNode,
            toOption,
            toCategory,
          }) ?? true;

        if (!isValid) return;

        // toNode become a root for now
        addProduct(toNode, fromData, fromOption);
        updateNodes([toNode]);

        return;
      }

      // Add
      if (
        from === StudioZone.ZONE_SOURCE &&
        fromOption?.type !== StudioCategoryType.LINK &&
        toOption?.type !== StudioCategoryType.LINK
      ) {
        const isValid =
          fromOption?.onAddValidate?.({
            option: fromOption,
            parentOption,
            formData: currentFormData,
            allFormData,
            data,
            toNode,
            toOption,
            toCategory,
          }) ?? true;

        if (!isValid) return;

        const newNode = getNewNodeInfo(fromData.optionKey, fromOption);

        addToPackage(toNode, [newNode]);
        updateNodes([toNode]);

        return;
      }

      // Merge
      if (
        from === StudioZone.ZONE_PRODUCT &&
        !isTheSameNode &&
        toOption?.type !== StudioCategoryType.LINK &&
        fromOption?.type !== StudioCategoryType.LINK
      ) {
        if (!fromData.belongsTo || !fromNode || !toNode) return;

        const isValid =
          fromOption.onMergeValidate?.({
            id: fromData.belongsTo,
            option: fromOption,
            parentOption,
            formData: currentFormData,
            allFormData,
            data,
            fromNode,
            toNode,
            toOption,
            toCategory,
          }) ?? true;

        if (!isValid) return;

        mergeProducts(fromNode, toNode, fromData);
        updateNodes([fromNode, toNode]);

        return;
      }

      // Move
      if (from === StudioZone.ZONE_PRODUCT_ADDON && !isTheSameNode && fromNode) {
        if (!fromData.belongsTo || !fromNode || !toNode) return;

        const isValid =
          fromOption.onSnapValidate?.({
            id: fromData.belongsTo,
            option: fromOption,
            parentOption,
            toOption,
            formData: currentFormData,
            allFormData,
            fromNode,
            toNode,
            data,
            toCategory,
          }) ?? true;

        if (!isValid) return;

        movePartOfPackage(fromNode, toNode, fromData);
        updateNodes([fromNode, toNode]);

        return;
      }
    }

    if (to === StudioZone.ZONE_FACTORY && !fromCategory?.isRoot) {
      // Remove the whole node
      if (from === StudioZone.ZONE_PRODUCT) {
        if (!fromData.belongsTo || !fromNode) return;

        const isValid =
          fromOption.onDropOutValidate?.({
            id: fromData.belongsTo,
            option: fromOption,
            parentOption,
            formData: currentFormData,
            allFormData,
            fromNode,
            data,
          }) ?? true;

        if (!isValid) return;

        removeProduct(fromData?.belongsTo);

        return;
      }

      // Remove the node's children
      if (from === StudioZone.ZONE_PRODUCT_ADDON && !isTheSameNode && fromNode) {
        if (!fromData.belongsTo) return;

        const isValid =
          fromOption.onDropOutValidate?.({
            id: fromData.belongsTo,
            option: fromOption,
            parentOption,
            formData: currentFormData,
            allFormData,
            data,
            fromNode,
          }) ?? true;

        if (!isValid) return;

        removePartOfPackage(fromNode, fromData?.childIndex || 0);
        updateNodes([fromNode]);

        return;
      }
    }

    useStudioFlowStore.getState().reloadFlow();

    movingNodeRef.current = null;

    updateFieldValidate({
      fromCategory,
      fromOption,
    });
  }, []);

  const handleDragMove = useCallback((event: DragMoveEvent) => {
    const { active, delta } = event;

    if (active) {
      const id = active.id as string;
      let movingNode = movingNodeRef.current;

      if (!movingNode) {
        const nodes = useStudioFlowStore.getState().nodes;
        const movingNodeIndex = nodes.findIndex((node) => node.id === id);

        movingNode = movingNodeRef.current || nodes[movingNodeIndex];

        movingNodeRef.current = movingNode;
      }

      if (movingNode) {
        movingNodeRef.current = movingNode;

        const newPosition: XYPosition = {
          x: movingNode.position.x + delta.x,
          y: movingNode.position.y + delta.y,
        };

        const updatedNode = applyNodeChanges(
          [
            {
              id,
              type: 'position',
              position: newPosition,
              positionAbsolute: newPosition,
              dragging: true,
            },
          ],
          [movingNode],
        );

        useStudioFlowStore.getState().updateNode(updatedNode[0]);
      }
    }
  }, []);

  const handleDragOver = useCallback((_event: DragOverEvent) => {
    console.log('[DndContainer] handleDragOver', _event.over?.id);
  }, []);

  const handleDragCancel = useCallback((_event: DragCancelEvent) => {}, []);

  const handleDragAbort = useCallback((_event: DragAbortEvent) => {}, []);

  const handleDragPending = useCallback((_event: DragPendingEvent) => {}, []);

  return (
    <DndContext
      sensors={sensors}
      collisionDetection={rectIntersection}
      onDragStart={handleDragStart}
      onDragMove={handleDragMove}
      onDragOver={handleDragOver}
      onDragEnd={handleDragEnd}
      onDragCancel={handleDragCancel}
      onDragAbort={handleDragAbort}
      onDragPending={handleDragPending}
    >
      {children}
    </DndContext>
  );
}

export default DndFlow;
