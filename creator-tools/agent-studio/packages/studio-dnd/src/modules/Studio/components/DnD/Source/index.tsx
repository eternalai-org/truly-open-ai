import { HTMLAttributes, memo, useCallback, useMemo } from 'react';

import Draggable from '../base/Draggable';

import useStudioDndStore from '@/modules/Studio/stores/useStudioDndStore';
import { DraggableData, StudioZone } from '@/modules/Studio/types/dnd';
import { DomRect } from '@/modules/Studio/types/ui';
import useStudioConfigStore from '@/modules/Studio/stores/useStudioConfigStore';

type Props = HTMLAttributes<HTMLDivElement> & {
  id: string;
  data: Omit<DraggableData, 'type'>;
  disabled?: boolean;
};

const Source = ({ id, data, disabled = false, children, ...props }: Props) => {
  const disabledDrag = useStudioConfigStore((state) => state.config.board.disabledDrag);

  const extendedData = useMemo(() => {
    return {
      ...data,
      type: StudioZone.ZONE_SOURCE,
    } satisfies DraggableData;
  }, [data, id]);

  const handleOnDrag = useCallback(
    (_data: DraggableData, touchingPoint: DomRect | null) => {
      useStudioDndStore.getState().setDragging(children, extendedData, touchingPoint);
    },
    [extendedData],
  );

  const handleOnDrop = useCallback((_data: DraggableData) => {
    useStudioDndStore.getState().setDragging();
  }, []);

  return (
    <Draggable
      {...props}
      id={id}
      data={extendedData}
      disabled={disabled || disabledDrag}
      handleOnDrag={handleOnDrag}
      handleOnDrop={handleOnDrop}
    >
      {children}
    </Draggable>
  );
};

export default memo(Source);
