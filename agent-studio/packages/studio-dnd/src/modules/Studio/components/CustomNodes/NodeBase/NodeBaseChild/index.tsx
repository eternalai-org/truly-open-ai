import { useMemo } from 'react';

import ProductAddon from '../../../DnD/ProductAddon';
import ProductPlaceholder from '../../../DnD/ProductPlaceholder';
import DraggingPlaceholder from '../../DraggingPlaceholder';
import DraggingFloating from '../DraggingFloating';
import LegoRender from '../LegoRender';

import useStudioCategoryStore from '@/modules/Studio/stores/useStudioCategoryStore';
import useStudioDndStore from '@/modules/Studio/stores/useStudioDndStore';
import { StudioCategoryOptionMapValue } from '@/modules/Studio/types/category';
import { DraggableData } from '@/modules/Studio/types/dnd';
import { StudioNode } from '@/modules/Studio/types/graph';

type Props = {
  data: StudioNode;
  items: StudioNode[];
  index: number;
  belongsTo: string;
  isHidden?: boolean;
  isLast?: boolean;
};

const NodeBaseChild = ({ data, items, index, belongsTo, isHidden, isLast = false }: Props) => {
  const categoryOptionMap = useStudioCategoryStore((state) => state.categoryOptionMap);
  const draggingOverId = useStudioDndStore((state) => state.draggingOverId);
  const draggingData = useStudioDndStore((state) => state.draggingData);

  const idx = data.data.metadata.idx;
  const option: StudioCategoryOptionMapValue | undefined = categoryOptionMap[idx];

  const productData: Omit<DraggableData, 'type'> = useMemo(
    () => ({
      optionKey: option.idx,
      belongsTo,
      childIndex: index,
      id: data.id,
    }),
    [index, option.idx, belongsTo, data.id],
  );

  const floatingItems = useMemo(() => items.slice(index), [items, index]);

  const isDraggingOver = useMemo(() => {
    return !!(draggingOverId && draggingOverId === data.id && draggingData?.id !== draggingOverId);
  }, [draggingOverId, data.id, draggingData?.id]);

  return (
    <div
      style={{
        position: 'relative',
        width: 'fit-content',
      }}
    >
      <ProductAddon
        id={data.id}
        data={productData}
        isHidden={isHidden}
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

      <ProductPlaceholder id={data.id} data={productData} hidden={isDraggingOver} />

      {isDraggingOver && <DraggingPlaceholder showDraggingFloating={isLast} />}
    </div>
  );
};

export default NodeBaseChild;
