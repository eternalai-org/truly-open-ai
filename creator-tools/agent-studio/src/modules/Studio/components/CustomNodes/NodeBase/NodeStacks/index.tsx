import { useMemo } from 'react';

import Package from '../../../DnD/Package';
import Product from '../../../DnD/Product';
import DraggingFloating from '../DraggingFloating';
import LegoRender from '../LegoRender';
import NodeBaseChild from '../NodeBaseChild';
import NodeBaseConnection from '../NodeBaseConnection';
import NodeBaseReadOnly from '../NodeBaseReadOnly';
import NodeBaseWrapper from '../NodeBaseWrapper';
import { NodeBaseProps } from '../types';

import useStudioCategoryStore from '@/modules/Studio/stores/useStudioCategoryStore';
import useStudioDndStore from '@/modules/Studio/stores/useStudioDndStore';
import { StudioCategoryOption } from '@/modules/Studio/types';
import { DraggableData } from '@/modules/Studio/types/dnd';
import './NodeStacks.scss';

type Props = NodeBaseProps;

const NodeStacks = ({ data, ...rest }: Props) => {
  const draggingData = useStudioDndStore((state) => state.draggingData);
  const categoryMap = useStudioCategoryStore((state) => state.categoryMap);
  const categoryOptionMap = useStudioCategoryStore((state) => state.categoryOptionMap);
  const children = data?.metadata?.children;

  const idx = data.metadata.idx;
  const option: StudioCategoryOption | undefined = categoryOptionMap[idx];
  const category = Object.values(categoryMap).find((category) => category.options.some((option) => option.idx === idx));
  const schemaData = option?.data;

  const productData: Omit<DraggableData, 'type'> = useMemo(
    () => ({ optionKey: option.idx, belongsTo: data.id, categoryKey: category?.idx }),
    [data.id, option.idx, category?.idx],
  );

  const renderChildren = useMemo(() => {
    if (!draggingData || draggingData.childIndex === undefined || draggingData.belongsTo !== data.id) return children;

    return children.slice(0, draggingData.childIndex + 1);
  }, [draggingData, data.id, children]);

  return (
    <NodeBaseWrapper data={data} id={data.id} option={option}>
      <div className="node-base">
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
            background={option.color}
            icon={option.icon}
            title={option.title}
            id={data.id}
            schemaData={schemaData}
            idx={option.idx}
            render={option.customizeRenderOnBoard}
          />
        </Product>

        {renderChildren?.map((item, index) => (
          <NodeBaseChild
            index={index}
            key={`node-base-child-${data.id}-${item.id}`}
            data={item}
            items={children}
            belongsTo={data.id}
          />
        ))}

        <Package id={data.id} data={productData} />
        <NodeBaseConnection />
      </div>
    </NodeBaseWrapper>
  );
};

export default NodeStacks;
