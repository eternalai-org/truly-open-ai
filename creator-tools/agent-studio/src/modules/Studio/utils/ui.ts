import { DomRect, TouchingPoint } from '../types/ui';

export const adjustColorShade = (hexColor: string, percent: number): string => {
  const hex = hexColor.replace('#', '');

  const r = parseInt(hex.substring(0, 2), 16);
  const g = parseInt(hex.substring(2, 4), 16);
  const b = parseInt(hex.substring(4, 6), 16);

  const amount = Math.round(2.55 * percent);

  const getNewColorValue = (color: number): number => {
    const newValue = color + amount;

    return Math.min(255, Math.max(0, newValue));
  };

  const newR = getNewColorValue(r);
  const newG = getNewColorValue(g);
  const newB = getNewColorValue(b);

  const toHex = (n: number): string => {
    const hex = n.toString(16);

    return hex.length === 1 ? '0' + hex : hex;
  };

  return `#${toHex(newR)}${toHex(newG)}${toHex(newB)}`;
};

export const calculateTouchingPercentage = (element: HTMLElement, point: TouchingPoint): DomRect => {
  const rect = element.getBoundingClientRect();
  const relativeX = point.clientX - rect.left;
  const relativeY = point.clientY - rect.top;

  return {
    x: relativeX,
    y: relativeY,
  };
};
