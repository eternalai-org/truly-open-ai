import { agentAPIClient } from "../clients";
import {
  AgentDetail,
  CreateStudioAgentPayload,
  UpdateStudioAgentPayload,
} from "./types";

const StudioApi = {
  createAgent: async (
    payload: CreateStudioAgentPayload
  ): Promise<AgentDetail[]> => {
    try {
      const agent: any = await agentAPIClient.post(
        `/api/agent/create_agent_studio`,
        {
          ...payload,
        }
      );
      return agent;
    } catch (e) {
      throw e;
    }
  },
  updateAgent: async (
    id: string,
    payload: UpdateStudioAgentPayload
  ): Promise<any> => {
    try {
      const agent: any = await agentAPIClient.post(
        `/api/agent/update_agent_studio/${id}`,
        {
          ...payload,
        }
      );
      return agent;
    } catch (e) {
      throw e;
    }
  },
};

export default StudioApi;
