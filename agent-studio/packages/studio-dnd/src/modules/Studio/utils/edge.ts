import { Edge, InternalNode, MarkerType, Node, Position, XYPosition } from '@xyflow/react';
import { v4 } from 'uuid';

import { EdgeType } from '@/enums/node-type';

/**
 * Generates a unique source handle ID
 * @param {string} source - Source node ID
 * @param {string} target - Target node ID
 * @returns {string} Generated source handle ID
 */
export const generateSourceHandleId = (source: string, target: string) => {
  return `${source}-s-${target}`;
};

/**
 * Generates a unique target handle ID
 * @param {string} source - Source node ID
 * @param {string} target - Target node ID
 * @returns {string} Generated target handle ID
 */
export const generateTargetHandleId = (source: string, target: string) => {
  return `${target}-t-${source}`;
};

/**
 * Creates a new base edge
 * @param {string} source - Source node ID
 * @param {string} target - Target node ID
 * @param {boolean} animated - Whether the edge is animated
 * @returns {Edge} New base edge
 */
export const createNewBaseEdge = (source: string, target: string, animated: boolean = false): Edge => {
  return {
    id: v4(),
    source,
    sourceHandle: 'c',
    target,
    targetHandle: 'a',
    type: EdgeType.EDGE_BASE,
    selectable: true,
    selected: false,
    focusable: false,
    label: '',
    animated,
    markerEnd: {
      type: MarkerType.Arrow,
      width: 20,
      height: 20,
    },
  } satisfies Edge;
};

export const getNodeIntersection = (intersectionNode: InternalNode<Node>, targetNode: InternalNode<Node>) => {
  // https://math.stackexchange.com/questions/1724792/an-algorithm-for-finding-the-intersection-point-between-a-center-of-vision-and-a
  const { width: intersectionNodeWidth, height: intersectionNodeHeight } = intersectionNode.measured;
  const intersectionNodePosition = intersectionNode.internals.positionAbsolute;
  const targetPosition = targetNode.internals.positionAbsolute;

  const w = (intersectionNodeWidth ?? 0) / 2;
  const h = (intersectionNodeHeight ?? 0) / 2;

  const x2 = intersectionNodePosition.x + w;
  const y2 = intersectionNodePosition.y + h;
  const x1 = targetPosition.x + (targetNode.measured?.width ?? 0) / 2;
  const y1 = targetPosition.y + (targetNode.measured?.height ?? 0) / 2;

  const xx1 = (x1 - x2) / (2 * w) - (y1 - y2) / (2 * h);
  const yy1 = (x1 - x2) / (2 * w) + (y1 - y2) / (2 * h);
  const a = 1 / (Math.abs(xx1) + Math.abs(yy1));
  const xx3 = a * xx1;
  const yy3 = a * yy1;
  const x = w * (xx3 + yy3) + x2;
  const y = h * (-xx3 + yy3) + y2;

  return { x, y };
};

export const getEdgePosition = (node: InternalNode<Node>, intersectionPoint: XYPosition) => {
  const n = { ...node.internals.positionAbsolute, ...node };
  const nx = Math.round(n.x);
  const ny = Math.round(n.y);
  const px = Math.round(intersectionPoint.x);
  const py = Math.round(intersectionPoint.y);

  if (px <= nx + 1) {
    return Position.Left;
  }
  if (px >= nx + (n.measured?.width ?? 0) - 1) {
    return Position.Right;
  }
  if (py <= ny + 1) {
    return Position.Top;
  }
  if (py >= n.y + (n.measured?.height ?? 0) - 1) {
    return Position.Bottom;
  }

  return Position.Top;
};

export const getEdgeParams = (source: InternalNode<Node> | undefined, target: InternalNode<Node> | undefined) => {
  if (!source || !target) {
    return {
      sx: 0,
      sy: 0,
      tx: 0,
      ty: 0,
      sourcePos: Position.Top,
      targetPos: Position.Top,
    };
  }

  const sourceIntersectionPoint = getNodeIntersection(source, target);
  const targetIntersectionPoint = getNodeIntersection(target, source);

  const sourcePos = getEdgePosition(source, sourceIntersectionPoint);
  const targetPos = getEdgePosition(target, targetIntersectionPoint);

  return {
    sx: sourceIntersectionPoint.x,
    sy: sourceIntersectionPoint.y,
    tx: targetIntersectionPoint.x,
    ty: targetIntersectionPoint.y,
    sourcePos,
    targetPos,
  };
};
