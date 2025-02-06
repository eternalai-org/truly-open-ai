import { Node, Viewport, XYPosition } from '@xyflow/react';
import { FunctionComponent } from 'react';

import { FormDataMap, Key } from './base';

export type StudioNodeMetadata = Record<string, unknown> & {
  children: StudioNode[];
  idx: Key;
};

export type StudioInternalDataNode = {
  sourceHandles: string[];
  targetHandles: string[];
  id: string;
  metadata: StudioNodeMetadata;
};
export type StudioNode = Node<StudioInternalDataNode>;

export type StudioDataNode = {
  id: string;
  idx: Key;
  categoryIdx?: Key;
  title: React.ReactNode | FunctionComponent;
  children: StudioDataNode[];
  data?: FormDataMap; // this field can be used to store additional data or form input data
  rect?: {
    position: XYPosition;
  };
};

export type GraphData = {
  data: StudioDataNode[];
  viewport: Viewport;
};
