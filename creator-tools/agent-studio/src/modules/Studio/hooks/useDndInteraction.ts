import { applyNodeChanges } from '@xyflow/react';
import { useCallback } from 'react';

import useStudioFlowStore from '../stores/useStudioFlowStore';
import { StudioNode } from '../types/graph';
import { noUndefinedElement } from '../utils/array';

const useDndInteraction = () => {
  const updateNodes = useCallback((nodes: (StudioNode | undefined)[]) => {
    const shouldUpdatedNodes = noUndefinedElement(nodes);
    if (shouldUpdatedNodes.length) {
      const updatedNodes = applyNodeChanges(
        shouldUpdatedNodes.map((node) => ({ id: node.id, type: 'position', position: node.position, positionAbsolute: node.position })),
        shouldUpdatedNodes,
      );

      useStudioFlowStore.getState().updateNodes(updatedNodes);
    }
  }, []);

  return { updateNodes };
};

export default useDndInteraction;
