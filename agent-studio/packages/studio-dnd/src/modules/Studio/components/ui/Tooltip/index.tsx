import React, { useEffect, useRef, useState } from 'react';
import { createPortal } from 'react-dom';

import './Tooltip.scss';
import { TooltipPlacement } from '@/modules/Studio/enums/side';

type Props = {
  children?: React.ReactNode;
  label?: string;
  placement?: TooltipPlacement;
  delay?: number;
  backgroundColor?: string;
  color?: string;
};

const MAX_TOOLTIP_WIDTH = 200;
const TOOLTIP_OFFSET = 10;
const VIEWPORT_PADDING = 10;

const Tooltip = ({
  children,
  label,
  placement = TooltipPlacement.BOTTOM,
  delay = 200,
  backgroundColor = '#333',
  color = 'white',
}: Props) => {
  const [isVisible, setIsVisible] = useState(false);
  const [position, setPosition] = useState({ x: 0, y: 0 });

  const triggerRef = useRef<HTMLDivElement>(null);
  const tooltipRef = useRef<HTMLDivElement>(null);

  const timeoutId = useRef<NodeJS.Timeout | null>(null);

  const calculatePosition = () => {
    if (!triggerRef.current || !tooltipRef.current) return;

    const triggerRect = triggerRef.current.getBoundingClientRect();
    const tooltipRect = tooltipRef.current.getBoundingClientRect();
    const viewportWidth = window.innerWidth;
    const viewportHeight = window.innerHeight;

    const getInitialPosition = (place: typeof placement) => {
      const positions = {
        top: {
          x: triggerRect.left + (triggerRect.width - tooltipRect.width) / 2,
          y: triggerRect.top - tooltipRect.height - TOOLTIP_OFFSET,
        },
        bottom: {
          x: triggerRect.left + (triggerRect.width - tooltipRect.width) / 2,
          y: triggerRect.bottom + TOOLTIP_OFFSET,
        },
        left: {
          x: triggerRect.left - tooltipRect.width - TOOLTIP_OFFSET,
          y: triggerRect.top + (triggerRect.height - tooltipRect.height) / 2,
        },
        right: {
          x: triggerRect.right + TOOLTIP_OFFSET,
          y: triggerRect.top + (triggerRect.height - tooltipRect.height) / 2,
        },
      };

      return positions[place];
    };

    const adjustPosition = (pos: { x: number; y: number }, place: typeof placement) => {
      let newPos = { ...pos };
      let newPlacement = place;

      if (pos.x < VIEWPORT_PADDING) {
        if (place === TooltipPlacement.LEFT) {
          newPos = getInitialPosition(TooltipPlacement.RIGHT);
          newPlacement = TooltipPlacement.RIGHT;
        } else {
          newPos.x = VIEWPORT_PADDING;
        }
      } else if (pos.x + tooltipRect.width > viewportWidth - VIEWPORT_PADDING) {
        if (place === TooltipPlacement.RIGHT) {
          newPos = getInitialPosition(TooltipPlacement.LEFT);
          newPlacement = TooltipPlacement.LEFT;
        } else {
          newPos.x = viewportWidth - tooltipRect.width - VIEWPORT_PADDING;
        }
      }

      if (pos.y < VIEWPORT_PADDING) {
        if (place === TooltipPlacement.TOP) {
          newPos = getInitialPosition(TooltipPlacement.BOTTOM);
          newPlacement = TooltipPlacement.BOTTOM;
        } else {
          newPos.y = VIEWPORT_PADDING;
        }
      } else if (pos.y + tooltipRect.height > viewportHeight - VIEWPORT_PADDING) {
        if (place === TooltipPlacement.BOTTOM) {
          newPos = getInitialPosition(TooltipPlacement.TOP);
          newPlacement = TooltipPlacement.TOP;
        } else {
          newPos.y = viewportHeight - tooltipRect.height - VIEWPORT_PADDING;
        }
      }

      return { position: newPos, placement: newPlacement };
    };

    const initialPosition = getInitialPosition(placement);
    const { position: adjustedPosition } = adjustPosition(initialPosition, placement);

    setPosition(adjustedPosition);
  };

  const handleMouseEnter = () => {
    timeoutId.current = setTimeout(() => {
      setIsVisible(true);

      requestAnimationFrame(calculatePosition);
    }, delay);
  };

  const handleMouseLeave = () => {
    if (timeoutId.current) clearTimeout(timeoutId.current);
    setIsVisible(false);
  };

  useEffect(() => {
    return () => {
      if (timeoutId.current) clearTimeout(timeoutId.current);
    };
  }, []);

  return (
    <>
      <div
        ref={triggerRef}
        onMouseEnter={handleMouseEnter}
        onMouseLeave={handleMouseLeave}
        className="studio-tooltip-trigger"
      >
        {children}
      </div>

      {createPortal(
        label && isVisible && (
          <div
            ref={tooltipRef}
            className={`studio-tooltip-content tooltip-${placement}`}
            style={{
              backgroundColor,
              color,
              left: position.x,
              top: position.y,
              maxWidth: MAX_TOOLTIP_WIDTH,
            }}
          >
            {label}
            {/* <div className={`tooltip-arrow tooltip-arrow-${placement}`} /> */}
          </div>
        ),
        document.body,
      )}
    </>
  );
};

export default Tooltip;
