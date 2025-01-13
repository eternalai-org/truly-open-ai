import { useDroppable } from '@dnd-kit/core';
import React, { memo } from 'react';

import { DroppableData } from '@/modules/Studio/types/dnd';
import './StaticDroppable.scss';

type Props = React.HTMLAttributes<HTMLDivElement> & {
  id: string;
  data: DroppableData;
  disabled?: boolean; // means the droppable do not accept any droppable
};

const StaticDroppable = ({ id, data, disabled, children, ...props }: Props) => {
  const { setNodeRef } = useDroppable({
    id,
    data,
    disabled,
  });

  return (
    <div
      ref={setNodeRef}
      style={{
        width: '100%',
        height: '100%',
      }}
      {...props}
    >
      {children}
    </div>
  );
};

export default memo(StaticDroppable);
