import React, { memo, useMemo } from 'react';

import StaticDroppable from '../base/StaticDroppable';

import { StudioZone } from '@/modules/Studio/types/dnd';

type Props = Omit<React.HTMLAttributes<HTMLDivElement>, 'id'> & {
  disabled?: boolean;
};

const Distribution = ({ children, className, ...props }: Props) => {
  const data = useMemo(() => {
    return { type: StudioZone.ZONE_DISTRIBUTION };
  }, []);

  return (
    <StaticDroppable id={StudioZone.ZONE_DISTRIBUTION} data={data} className={className} {...props}>
      {children}
    </StaticDroppable>
  );
};

export default memo(Distribution);
