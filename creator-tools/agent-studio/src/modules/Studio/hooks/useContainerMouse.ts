import { XYPosition } from '@xyflow/react';
import React from 'react';

type Props = {
  ref: React.RefObject<unknown | null>;
  handleOnTick: (contentRect: DOMRect, mousePosition: XYPosition, previousMousePosition: XYPosition) => void;
};

const useContainerMouse = ({ ref, handleOnTick }: Props) => {
  const frameRef = React.useRef<number | null>(null);
  const mousePositionRef = React.useRef({ x: 0, y: 0 });
  const previousMousePositionRef = React.useRef({ x: 0, y: 0 });

  const handleMouseMove = (event: MouseEvent) => {
    const rect = (ref.current as HTMLDivElement)?.getBoundingClientRect();
    if (!rect) return;

    const x = event.clientX - rect.left;
    const y = event.clientY - rect.top;

    mousePositionRef.current = { x, y };
  };

  const tick = () => {
    // const deltaMouseX = mousePositionRef.current.x - previousMousePositionRef.current.x;
    // const deltaMouseY = mousePositionRef.current.y - previousMousePositionRef.current.y;

    const contentRect = document.documentElement.getBoundingClientRect();
    handleOnTick(contentRect, mousePositionRef.current, previousMousePositionRef.current);

    previousMousePositionRef.current = {
      x: mousePositionRef.current.x,
      y: mousePositionRef.current.y,
    };

    frameRef.current = window.requestAnimationFrame(tick);
  };

  const addListeners = () => {
    window.removeEventListener('mousemove', handleMouseMove);
    if (frameRef.current) {
      window.cancelAnimationFrame(frameRef.current);
    }

    window.addEventListener('mousemove', handleMouseMove);
    tick();
  };

  const removeListeners = () => {
    window.removeEventListener('mousemove', handleMouseMove);
    if (frameRef.current) {
      window.cancelAnimationFrame(frameRef.current);
    }

    mousePositionRef.current = { x: 0, y: 0 };
    previousMousePositionRef.current = { x: 0, y: 0 };
  };

  return { addListeners, removeListeners };
};

export default useContainerMouse;
