import React, { memo } from 'react';

import StaticDroppable from '../base/StaticDroppable';

import { StudioZone } from '@/modules/Studio/types/dnd';

type Props = Omit<React.HTMLAttributes<HTMLDivElement>, 'id'> & {
  disabled?: boolean; // means the droppable do not accept any droppable
};

const Factory = ({ children, ...props }: Props) => {
  return (
    <StaticDroppable id={StudioZone.ZONE_FACTORY} data={{ type: StudioZone.ZONE_FACTORY }} {...props}>
      {children}
    </StaticDroppable>
  );
};

export default memo(Factory);
