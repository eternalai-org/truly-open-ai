import { SidebarSide } from '../enums/side';
import { TabBehavior } from '../enums/tab';

export type BoardConfig = {
  minZoom?: number;
  maxZoom?: number;
  fitViewOptions?: {
    padding?: number;
  };

  disabledDrag?: boolean;
  disabledConnection?: boolean;
  disabledZoom?: boolean;
  disabledMiniMap?: boolean;
  disabledControls?: boolean;
  disabledBackground?: boolean;
};

export type SidebarConfig = {
  side: SidebarSide;
  width?: string | number;
};

export type TabConfig = {
  behavior: TabBehavior;
};

export type StudioConfig = {
  sidebar: SidebarConfig;
  tab: TabConfig;
  board: BoardConfig;
};
