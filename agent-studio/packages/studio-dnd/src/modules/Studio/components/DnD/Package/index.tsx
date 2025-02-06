import { useDroppable } from '@dnd-kit/core';
import React, { useMemo } from 'react';

import { DraggableData, StudioZone } from '@/modules/Studio/types/dnd';

import './Package.scss';

type Props = React.HTMLAttributes<HTMLDivElement> & {
  id: string;
  data: Omit<DraggableData, 'type'>;
  disabled?: boolean; // means the droppable do not accept any droppable
};

const Package = ({ id, data, disabled, children, ...props }: Props) => {
  const extendedData = useMemo(() => {
    return {
      ...data,
      type: StudioZone.ZONE_PACKAGE,
    } satisfies DraggableData;
  }, [data]);

  const { setNodeRef } = useDroppable({
    id: id + '-package',
    data: extendedData,
    disabled,
  });

  return (
    <div
      ref={setNodeRef}
      id={`node-base-${data.belongsTo}`}
      style={{
        padding: '16px',
      }}
      {...props}
    >
      {children}
    </div>
  );
};

export default Package;
