import { AgentClient, BaseDagent } from "@eternalai-dagent/core";
import { create } from "zustand";
import { getBaseAgent } from "../utils/agent";

interface IStore {
  baseDAgent: BaseDagent;

  accessToken: string;
}

const useAgentServiceStore = create<IStore>((set) => ({
  baseDAgent: getBaseAgent(),

  accessToken: "",
}));

export default useAgentServiceStore;
