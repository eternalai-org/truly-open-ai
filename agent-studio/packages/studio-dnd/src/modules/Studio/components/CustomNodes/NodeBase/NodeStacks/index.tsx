import { useMemo } from 'react';

import Product from '../../../DnD/Product';
import DraggingPlaceholder from '../../DraggingPlaceholder';
import DraggingFloating from '../DraggingFloating';
import LegoRender from '../LegoRender';
import NodeBaseChild from '../NodeBaseChild';
import NodeBaseConnection from '../NodeBaseConnection';
import NodeBaseReadOnly from '../NodeBaseReadOnly';
import NodeBaseWrapper from '../NodeBaseWrapper';
import { NodeBaseProps } from '../types';

import useNodeSelected from '@/modules/Studio/hooks/useNodeSelected';
import useStudioCategoryStore from '@/modules/Studio/stores/useStudioCategoryStore';
import useStudioDndStore from '@/modules/Studio/stores/useStudioDndStore';
import { StudioCategoryOption } from '@/modules/Studio/types';
import { DraggableData } from '@/modules/Studio/types/dnd';
import './NodeStacks.scss';

type Props = NodeBaseProps;

const NodeStacks = ({ data, ...rest }: Props) => {
  const { isSelected } = useNodeSelected({ id: data.id });

  const draggingData = useStudioDndStore((state) => state.draggingData);
  const draggingOverId = useStudioDndStore((state) => state.draggingOverId);

  const categoryMap = useStudioCategoryStore((state) => state.categoryMap);
  const categoryOptionMap = useStudioCategoryStore((state) => state.categoryOptionMap);
  const children = data?.metadata?.children;

  const idx = data.metadata.idx;
  const option: StudioCategoryOption | undefined = categoryOptionMap[idx];
  const category = Object.values(categoryMap).find((category) => category.options.some((option) => option.idx === idx));
  const schemaData = option?.data;

  const productData: Omit<DraggableData, 'type'> = useMemo(
    () => ({
      optionKey: option.idx,
      id: data.id,
      belongsTo: data.id,
      categoryKey: category?.idx,
    }),
    [data.id, option.idx, category?.idx],
  );

  const renderChildren = useMemo(() => {
    if (!draggingData || draggingData.childIndex === undefined || draggingData.belongsTo !== data.id) return children;

    return children.slice(0, draggingData.childIndex + 1);
  }, [draggingData, data.id, children]);

  const highlightColor = useMemo(
    () => option?.highlightColor || category?.highlightColor || option?.color,
    [option, category],
  );

  const isDraggingOver = useMemo(() => {
    return !!(draggingOverId && draggingOverId === data.id && draggingData?.id !== draggingOverId);
  }, [draggingOverId, data.id, draggingData?.id]);

  return (
    <NodeBaseWrapper data={data} id={data.id} option={option}>
      <div className="node-base">
        <div
          style={{
            position: 'relative',
            width: 'fit-content',
          }}
        >
          <Product
            id={data.id}
            data={productData}
            draggingFloating={
              <div>
                <NodeBaseReadOnly {...rest} data={data} />
                {renderChildren.map((item) => (
                  <DraggingFloating key={`dragging-floating-${data.id}-${item.id}`} data={item} />
                ))}
              </div>
            }
          >
            <LegoRender
              background={isSelected ? highlightColor : option.color}
              icon={option.icon}
              title={option.title}
              id={data.id}
              schemaData={schemaData}
              idx={option.idx}
              render={option.customizeRenderOnBoard}
            />
          </Product>
          {/* <ProductPlaceholder id={data.id} data={productData} hidden={isDraggingOver} /> */}
        </div>

        {isDraggingOver && <DraggingPlaceholder />}

        {renderChildren?.map((item, index) => (
          <NodeBaseChild
            index={index}
            key={`node-base-child-${data.id}-${item.id}`}
            data={item}
            items={children}
            belongsTo={data.id}
            isHidden={draggingData?.belongsTo === data.id && draggingData?.optionKey === option?.idx}
            isLast={index === renderChildren.length - 1}
          />
        ))}

        <NodeBaseConnection />
      </div>
    </NodeBaseWrapper>
  );
};

export default NodeStacks;
