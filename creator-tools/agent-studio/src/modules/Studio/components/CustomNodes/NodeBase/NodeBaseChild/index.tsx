import { useMemo } from 'react';

import ProductAddon from '../../../DnD/ProductAddon';
import DraggingFloating from '../DraggingFloating';
import LegoRender from '../LegoRender';

import useStudioCategoryStore from '@/modules/Studio/stores/useStudioCategoryStore';
import { StudioCategoryOptionMapValue } from '@/modules/Studio/types/category';
import { DraggableData } from '@/modules/Studio/types/dnd';
import { StudioNode } from '@/modules/Studio/types/graph';

type Props = {
  data: StudioNode;
  items: StudioNode[];
  index: number;
  belongsTo: string;
};

const NodeBaseChild = ({ data, items, index, belongsTo }: Props) => {
  const categoryOptionMap = useStudioCategoryStore((state) => state.categoryOptionMap);

  const idx = data.data.metadata.idx;
  const option: StudioCategoryOptionMapValue | undefined = categoryOptionMap[idx];

  const productData: Omit<DraggableData, 'type'> = useMemo(
    () => ({ optionKey: option.idx, belongsTo, childIndex: index }),
    [belongsTo, index, option.idx],
  );

  const floatingItems = useMemo(() => items.slice(index), [items, index]);

  return (
    <ProductAddon
      id={data.id}
      data={productData}
      draggingFloating={
        <div>
          {floatingItems.map((item) => (
            <DraggingFloating key={`dragging-floating-${data.id}-${item.id}`} data={item} />
          ))}
        </div>
      }
    >
      <LegoRender
        background={option.color}
        icon={option.icon}
        title={option.title}
        id={data.id}
        schemaData={option.data}
        idx={option.idx}
        render={option.customizeRenderOnBoard}
      />
    </ProductAddon>
  );
};

export default NodeBaseChild;
