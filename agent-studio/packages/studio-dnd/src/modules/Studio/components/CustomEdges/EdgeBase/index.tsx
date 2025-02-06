import { BaseEdge, EdgeLabelRenderer, EdgeProps, getSmoothStepPath, useInternalNode } from '@xyflow/react';
import React, { memo, useMemo } from 'react';

import useDndAction from '@/modules/Studio/hooks/useDndAction';
import useEdgeSelected from '@/modules/Studio/hooks/useEdgeSelected';
import useStudioCategoryStore from '@/modules/Studio/stores/useStudioCategoryStore';
import useStudioDataStore from '@/modules/Studio/stores/useStudioDataStore';
import useStudioFlowStore from '@/modules/Studio/stores/useStudioFlowStore';
import useStudioFormStore from '@/modules/Studio/stores/useStudioFormStore';
import type { StudioCategoryOptionMapValue } from '@/modules/Studio/types';
import { getEdgeParams } from '@/modules/Studio/utils/edge';
import './EdgeBase.scss';

const EdgeBase = ({ id, source, target, markerEnd, style }: EdgeProps) => {
  const sourceNode = useInternalNode(source);
  const targetNode = useInternalNode(target);
  const { unlink } = useDndAction();
  const { sx, sy, tx, ty, sourcePos, targetPos } = getEdgeParams(sourceNode, targetNode);

  const fromNode = useMemo(() => useStudioFlowStore.getState().nodes.find((node) => node.id === source), [source]);
  const toNode = useMemo(() => useStudioFlowStore.getState().nodes.find((node) => node.id === target), [target]);
  const categoryMap = useStudioCategoryStore((state) => state.categoryMap);
  const categoryOptionMap = useStudioCategoryStore((state) => state.categoryOptionMap);

  const fromIdx = fromNode?.data.metadata.idx || '';
  const fromOption: StudioCategoryOptionMapValue | undefined = categoryOptionMap[fromIdx];
  const fromCategory = Object.values(categoryMap).find((category) =>
    category.options.some((option) => option.idx === fromIdx),
  );

  const toIdx = toNode?.data.metadata.idx || '';
  const toOption: StudioCategoryOptionMapValue | undefined = categoryOptionMap[toIdx];
  const toCategory = Object.values(categoryMap).find((category) =>
    category.options.some((option) => option.idx === toIdx),
  );

  const { isSelected } = useEdgeSelected({ id });

  const [edgePath, labelX, labelY] = getSmoothStepPath({
    sourceX: sx,
    sourceY: sy,
    targetX: tx,
    targetY: ty,
    sourcePosition: sourcePos,
    targetPosition: targetPos,
  });

  const handleOnClick = () => {
    const data = useStudioDataStore.getState().data;

    const allFormData = useStudioFormStore.getState().formMap;
    const currentFormData = allFormData[fromNode?.id || ''];

    const parentOption = fromOption?.parent;

    if (!fromNode || !toNode || !fromCategory || !toCategory) return;

    const isValid =
      fromOption?.onUnlinkValidate?.({
        id: fromNode?.id,
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

    unlink(fromNode, toNode);
  };

  return (
    <React.Fragment>
      <BaseEdge id={id} path={edgePath} markerEnd={markerEnd} className="edge" style={{ ...style }} />

      <EdgeLabelRenderer>
        <div
          className="edge__actions"
          style={{
            transform: `translate(-50%, -50%) translate(${labelX}px,${labelY}px)`,
            opacity: isSelected ? 1 : 0,
            pointerEvents: 'auto',
          }}
        >
          {fromOption?.customizeRenderLabel ? (
            fromOption.customizeRenderLabel({ fromOption, toOption, fromCategory, toCategory })
          ) : (
            <button className="edge__actions__button" onClick={handleOnClick}>
              Unlink
            </button>
          )}
        </div>
      </EdgeLabelRenderer>
    </React.Fragment>
  );
};

export default memo(EdgeBase);
