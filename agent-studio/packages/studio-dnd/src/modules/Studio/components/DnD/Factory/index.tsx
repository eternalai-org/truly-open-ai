import React, { memo, useMemo } from 'react';

import StaticDroppable from '../base/StaticDroppable';

import { StudioZone } from '@/modules/Studio/types/dnd';

type Props = Omit<React.HTMLAttributes<HTMLDivElement>, 'id'> & {
  disabled?: boolean; // means the droppable do not accept any droppable
};

const Factory = ({ children, ...props }: Props) => {
  const data = useMemo(() => {
    return { type: StudioZone.ZONE_FACTORY };
  }, []);

  return (
    <StaticDroppable id={StudioZone.ZONE_FACTORY} data={data} {...props}>
      {children}
    </StaticDroppable>
  );
};

export default memo(Factory);
