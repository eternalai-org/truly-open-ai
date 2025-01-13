import LegoRender from '../LegoRender';
import { NodeBaseProps } from '../types';

import useStudioCategoryStore from '@/modules/Studio/stores/useStudioCategoryStore';
import { StudioCategoryOptionMapValue } from '@/modules/Studio/types/category';

type Props = NodeBaseProps;

const NodeBaseReadOnly = ({ data }: Props) => {
  const categoryOptionMap = useStudioCategoryStore((state) => state.categoryOptionMap);

  const idx = data.metadata.idx;
  const option: StudioCategoryOptionMapValue | undefined = categoryOptionMap[idx];
  const schemaData = option?.data;

  return (
    <LegoRender
      background={option?.color}
      icon={option?.icon}
      title={option?.title}
      id={data.id}
      schemaData={schemaData}
      idx={option?.idx}
      readonly
      render={option?.customizeRenderOnBoard}
    />
  );
};

export default NodeBaseReadOnly;
