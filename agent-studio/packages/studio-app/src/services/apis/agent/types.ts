import { Collection } from "../../../types/collection";

export interface ICollectionsResponse {
  collections: Collection[];
  next: string;
}
