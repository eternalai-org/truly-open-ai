import { DragOverlay } from '@dnd-kit/core';
import { useReactFlow } from '@xyflow/react';
import { memo } from 'react';

import useStudioDndStore from '@/modules/Studio/stores/useStudioDndStore';

const DragMask = () => {
  const { getZoom } = useReactFlow();

  const draggingElement = useStudioDndStore((state) => state.draggingElement);

  if (!draggingElement) return null;

  const zoom = getZoom();

  return (
    <DragOverlay zIndex={10000} style={{ opacity: 0.7 }}>
      <div
        style={{
          transform: `scale(${zoom})`,
          transformOrigin: 'top left',
        }}
      >
        {draggingElement}
      </div>
    </DragOverlay>
  );
};

export default memo(DragMask);
