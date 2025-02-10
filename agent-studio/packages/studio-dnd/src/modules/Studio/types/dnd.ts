export type DragAndDropData = {
  type: StudioZone;
};

export type DraggableData = DragAndDropData & {
  categoryKey?: string;
  optionKey?: string;
  belongsTo?: string;
  childIndex?: number;
  isRoot?: boolean;
  id?: string;
};

export type DroppableData = DragAndDropData & {};

export enum StudioZone {
  ZONE_FACTORY = 'factory',
  ZONE_SOURCE = 'source',
  ZONE_PRODUCT = 'product',
  ZONE_PRODUCT_ADDON = 'product_addon',
  ZONE_PACKAGE = 'package',
  ZONE_DISTRIBUTION = 'distribution',
}
