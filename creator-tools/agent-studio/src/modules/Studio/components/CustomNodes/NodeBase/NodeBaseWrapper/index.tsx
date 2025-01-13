import cs from 'clsx';
import React, { useMemo } from 'react';

import './NodeBaseWrapper.scss';
import TextRender from '../../../Render/TextRender';

import useStudioFlowStore from '@/modules/Studio/stores/useStudioFlowStore';
import { StudioInternalDataNode } from '@/modules/Studio/types';
import { StudioCategoryOption } from '@/modules/Studio/types/category';

type Props = {
  id: string;
  data: StudioInternalDataNode;
  children: React.ReactNode;
  option: StudioCategoryOption;
};

function NodeBaseWrapper({ id, data, option, children }: Props) {
  // check if have linked children
  const linkedNodes = useStudioFlowStore((state) => state.linkedNodes);
  const isHaveLinkedChildren = useMemo(() => {
    return !!data?.metadata?.children?.length && linkedNodes[id]?.length > 0;
  }, [linkedNodes, id, data?.metadata?.children?.length]);

  if (option?.boxWrapper) {
    if (option.boxWrapper.render) {
      return option.boxWrapper.render(children, option);
    }

    if (option.boxWrapper.title) {
      return (
        <div className="node-base-wrapper">
          <div className="node-base-wrapper__title">
            <TextRender data={option.boxWrapper.title} />
          </div>
          <div className="node-base-wrapper__content">{children}</div>
        </div>
      );
    }
  }

  if (isHaveLinkedChildren) {
    return (
      <div className={cs('node-base-wrapper', { 'node-base-wrapper--linked': isHaveLinkedChildren })}>
        <div className="node-base-wrapper__content">{children}</div>
      </div>
    );
  }

  return <>{children}</>;
}

export default NodeBaseWrapper;
