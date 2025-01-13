import { DragOverlay } from '@dnd-kit/core';
import { memo } from 'react';

import useStudioDndStore from '@/modules/Studio/stores/useStudioDndStore';

const DragMask = () => {
  const draggingElement = useStudioDndStore((state) => state.draggingElement);

  if (!draggingElement) return null;

  return <DragOverlay zIndex={10000}>{draggingElement}</DragOverlay>;
};

export default memo(DragMask);
