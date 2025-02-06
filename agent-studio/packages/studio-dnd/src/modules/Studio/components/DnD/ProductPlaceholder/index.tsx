import { useDroppable } from '@dnd-kit/core';
import cx from 'clsx';
import React, { useMemo } from 'react';

import useStudioDataStore from '@/modules/Studio/stores/useStudioDataStore';
import useStudioDndStore from '@/modules/Studio/stores/useStudioDndStore';
import { DraggableData, StudioZone } from '@/modules/Studio/types/dnd';
import './Package.scss';

type Props = React.HTMLAttributes<HTMLDivElement> & {
  id: string;
  data: Omit<DraggableData, 'type'>;
  disabled?: boolean; // means the droppable do not accept any droppable
};

const ProductPlaceholderComponent = ({ id, data, disabled, className, children, ...props }: Props) => {
  const draggingElement = useStudioDndStore((state) => state.draggingElement);

  const extendedData = useMemo(() => {
    return {
      ...data,
      type: StudioZone.ZONE_PACKAGE,
    } satisfies DraggableData;
  }, [data]);

  const { setNodeRef, isOver, active } = useDroppable({
    id: id + '-placeholder',
    data: extendedData,
    disabled,
  });

  const isSelf = useMemo(() => {
    if (isOver) {
      return active?.id === id;
    }

    return isOver;
  }, [active, isOver, id]);

  const isParent = useMemo(() => {
    const matchedData = useStudioDataStore.getState().data.find((item) => item.id === id);
    if (matchedData && matchedData.children.length > 0) {
      return !!matchedData.children.find((item) => item.id === active?.id);
    }

    return false;
  }, [id, active?.id]);

  const isShowDropMask = useMemo(() => {
    return !isParent && !isSelf && isOver;
  }, [isSelf, isOver, isParent]);

  const extendedStyle = useMemo(() => {
    if (isShowDropMask) {
      return {
        zIndex: 2,
      };
    }

    return {};
  }, [isShowDropMask]);

  return (
    <div
      className={cx('droppable-node', className)}
      ref={setNodeRef}
      style={{
        width: '100%',
        minHeight: '100%',
        ...extendedStyle,
      }}
      {...props}
    >
      <div className="droppable-node__container">{children}</div>

      {isShowDropMask && <div className="droppable-node__mask">{draggingElement}</div>}
    </div>
  );
};

export default function ProductPlaceholder(props: Props) {
  const draggingData = useStudioDndStore((state) => state.draggingData);

  if (draggingData?.belongsTo === props.id) {
    return <></>;
  }

  return <ProductPlaceholderComponent {...props} />;
}
