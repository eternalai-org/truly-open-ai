/* eslint-disable @typescript-eslint/no-unused-vars */
import Dexie, { type EntityTable } from "dexie";
import { StudioDataNode } from "@agent-studio/studio-dnd";

type PersistedAgentItem = {
  id: string;
  data: string;
  createdAt?: string;
};

class AgentDatabase {
  private databaseName = "agent-database";
  private db;
  constructor() {
    try {
      this.db = new Dexie(this.databaseName) as Dexie & {
        agent: EntityTable<PersistedAgentItem, "id">;
      };

      // for version 1
      this.db.version(1).stores({
        agent: "id, data, createdAt",
      });
    } catch (e) {
      //
    }
  }

  public async getItem(id: string) {
    try {
      return await this.db?.agent.get(id);
    } catch (e) {
      //
    }

    return null;
  }

  public async getAllItems() {
    try {
      return await this.db?.agent.toArray();
    } catch (e) {
      //
    }

    return [];
  }

  private async addItem(newItem: PersistedAgentItem) {
    try {
      await this.db?.agent.add({
        ...newItem,
        createdAt: new Date().toISOString(),
      });

      return newItem;
    } catch (e) {
      //
    }

    return null;
  }

  private async updateItem(updatedItem: PersistedAgentItem) {
    try {
      await this.db?.agent.update(updatedItem.id, updatedItem);

      return updatedItem;
    } catch (e) {
      //
    }

    return null;
  }

  async upsertItem(item: PersistedAgentItem) {
    try {
      const persisted = await this.getItem(item.id);
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

  async deleteItem(id: string) {
    try {
      await this.db?.agent.delete(id);
    } catch (e) {
      //
    }
  }
}

const agentDatabase = new AgentDatabase();
export default agentDatabase;
