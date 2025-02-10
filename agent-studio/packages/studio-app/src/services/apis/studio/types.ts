import { IAgent } from "@eternalai-dagent/core";

export type CreateStudioAgentPayload = {
  graph_data: string;
};

export type UpdateStudioAgentPayload = CreateStudioAgentPayload;

export type AgentDetail = IAgent & {
  graph_data?: string;
};
