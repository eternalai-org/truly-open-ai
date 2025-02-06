import { useDraggable } from '@dnd-kit/core';
import { CSS } from '@dnd-kit/utilities';
import cx from 'clsx';
import { HTMLAttributes, useEffect, useMemo, useRef } from 'react';

import useStudioDndStore from '@/modules/Studio/stores/useStudioDndStore';
import { DraggableData } from '@/modules/Studio/types/dnd';
import { DomRect, TouchingPoint } from '@/modules/Studio/types/ui';
import { calculateTouchingPercentage } from '@/modules/Studio/utils/ui';

import './Draggable.scss';

type Props = HTMLAttributes<HTMLDivElement> & {
  id: string;
  data: DraggableData;
  disabled?: boolean;
  handleOnDrag: (data: DraggableData, touchingPoint: DomRect | null) => void;
  handleOnDrop: (data: DraggableData) => void;
  isHidden?: boolean;
};

const Draggable = ({
  id,
  data,
  disabled = false,
  children,
  className,
  handleOnDrag,
  handleOnDrop,
  isHidden,
  ...props
}: Props) => {
  const touchingPointRef = useRef<TouchingPoint>({ clientX: 0, clientY: 0 });

  const { attributes, listeners, setNodeRef, transform, isDragging, node } = useDraggable({
    id,
    disabled,
    data,
  });

  const styles = useMemo(() => {
    if (isDragging) {
      return {
        ...props.style,
        transform: CSS.Translate.toString(transform),
        opacity: isDragging || isHidden ? 0 : 1,
      };
    }

    return {
      ...props.style,
      opacity: isHidden ? 0 : 1,
    };
  }, [transform, isDragging, props.style, isHidden]);

  useEffect(() => {
    if (isDragging) {
      const touchingPointAt = node?.current
        ? calculateTouchingPercentage(node.current, touchingPointRef.current)
        : null;

      handleOnDrag(data, touchingPointAt);
    } else {
      handleOnDrop(data);
    }
  }, [isDragging, data]);

  useEffect(() => {
    if (isDragging) {
      useStudioDndStore.getState().setTransform(transform);
    }
  }, [transform, isDragging]);

  return (
    <div
      ref={setNodeRef}
      style={styles}
      {...props}
      {...listeners}
      {...attributes}
      className={cx('draggable', { 'draggable--disabled': disabled }, className)}
      onMouseMove={(e) => {
        touchingPointRef.current = {
          clientX: e.clientX,
          clientY: e.clientY,
        };
      }}
    >
      {children}
    </div>
  );
};

export default Draggable;
