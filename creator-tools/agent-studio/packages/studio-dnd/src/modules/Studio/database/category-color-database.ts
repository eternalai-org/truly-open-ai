import Dexie, { type EntityTable } from 'dexie';

type PersistedStudioCategoryColorItem = {
  idx: string;
  color: string;
  createdAt?: string;
};

class CategoryColorDatabase {
  private databaseName = 'studio-category-color';
  private db;
  constructor() {
    try {
      this.db = new Dexie(this.databaseName) as Dexie & {
        category: EntityTable<PersistedStudioCategoryColorItem, 'idx'>;
      };

      // for version 1
      this.db.version(1).stores({
        category: 'idx, color, createdAt',
      });
    } catch (e) {
      //
    }
  }

  private async getItem(idx: string) {
    try {
      return await this.db?.category.get(idx);
    } catch (e) {
      //
    }

    return null;
  }

  public async getAllItemsToMap() {
    try {
      const allItems = await this.db?.category.toArray();
      if (allItems) {
        return allItems.reduce((acc: Record<string, string>, item) => {
          acc[item.idx] = item.color;

          return acc;
        }, {});
      }
    } catch (e) {
      //
    }

    return {};
  }

  private async addItem(newItem: PersistedStudioCategoryColorItem) {
    try {
      await this.db?.category.add({
        ...newItem,
        createdAt: new Date().toISOString(),
      });

      return newItem;
    } catch (e) {
      //
    }

    return null;
  }

  private async updateItem(updatedItem: PersistedStudioCategoryColorItem) {
    try {
      await this.db?.category.update(updatedItem.idx, updatedItem);

      return updatedItem;
    } catch (e) {
      //
    }

    return null;
  }

  async upsertItem(item: PersistedStudioCategoryColorItem) {
    try {
      const persisted = await this.getItem(item.idx);
      if (persisted) {
        return this.updateItem(item);
      } else {
        return this.addItem(item);
      }
    } catch (e) {
      //
    }

    return null;
  }
}

const categoryColorDatabase = new CategoryColorDatabase();
export default categoryColorDatabase;
