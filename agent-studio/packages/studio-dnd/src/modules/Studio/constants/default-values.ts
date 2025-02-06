import { SECOND } from './time';
import { StudioCategoryType } from '../enums/category';
import { SidebarSide } from '../enums/side';
import { TabBehavior } from '../enums/tab';
import { BoardConfig } from '../types/config';

export const DEFAULT_DISABLED_CONNECTION = false;
export const DEFAULT_CATEGORY_TYPE = StudioCategoryType.INLINE;

export const DEFAULT_THROTTLE_NODES_DELAY = SECOND;
export const DEFAULT_THROTTLE_VIEW_DELAY = SECOND;
export const DEFAULT_THROTTLE_DATA_DELAY = SECOND / 2;

export const DEFAULT_TAB_BEHAVIOR = TabBehavior.SCROLL;
export const DEFAULT_SIDEBAR_SIDE = SidebarSide.LEFT;

export const DEFAULT_BOARD_CONFIG: BoardConfig = {
  minZoom: 0.5,
  maxZoom: 2,
  fitViewOptions: {
    padding: 1,
  },

  disabledDrag: false,
  disabledConnection: false,
  disabledZoom: false,
  disabledMiniMap: false,
  disabledControls: false,
  disabledBackground: false,
};
