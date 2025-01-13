import { memo } from 'react';

import Overlay from '@/modules/Studio/components/Overlay';
import useStudioDndStore from '@/modules/Studio/stores/useStudioDndStore';
import { StudioZone } from '@/modules/Studio/types/dnd';

const BoardOverlay = () => {
  const { draggingData, draggingElement } = useStudioDndStore();

  return (
    <Overlay active={draggingData?.type === StudioZone.ZONE_SOURCE && !!draggingElement}>
      <></>
    </Overlay>
  );
};

export default memo(BoardOverlay);
