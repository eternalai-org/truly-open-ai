import { EdgeTypes, NodeTypes } from '@xyflow/react';

import EdgeBase from '../components/CustomEdges/EdgeBase';
import NodeBase from '../components/CustomNodes/NodeBase';

import { EdgeType, NodeType } from '@/enums/node-type';

export const CATEGORY_KEY_MAPPER = {};

export const DEFAULT_NODE_TYPES: NodeTypes = {
  [NodeType.NODE_BASE]: NodeBase,
};

export const DEFAULT_EDGE_TYPES: EdgeTypes = {
  [EdgeType.EDGE_BASE]: EdgeBase,
};
