import { BaseDagent } from "@eternalai-dagent/core";
import { create } from "zustand";
import { getBaseAgent } from "../utils/agent";

interface IStore {
  baseDAgent: BaseDagent;

  accessToken: string;

  walletAddress: string;
}

const useAgentServiceStore = create<IStore>((set) => ({
  baseDAgent: getBaseAgent(),

  accessToken: "",

  walletAddress: "",
}));

export default useAgentServiceStore;
