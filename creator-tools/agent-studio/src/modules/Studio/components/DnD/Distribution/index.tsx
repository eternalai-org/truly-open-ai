import React, { memo } from 'react';

import StaticDroppable from '../base/StaticDroppable';

import { StudioZone } from '@/modules/Studio/types/dnd';

type Props = Omit<React.HTMLAttributes<HTMLDivElement>, 'id'> & {
  disabled?: boolean;
};

const Distribution = ({ children, className, ...props }: Props) => {
  return (
    <StaticDroppable
      id={StudioZone.ZONE_DISTRIBUTION}
      data={{ type: StudioZone.ZONE_DISTRIBUTION }}
      className={className}
      {...props}
    >
      {children}
    </StaticDroppable>
  );
};

export default memo(Distribution);
