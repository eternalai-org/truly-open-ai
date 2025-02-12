import {
  AgentApps,
  FarcasterAgentSnapshotMission,
  IAgent,
  AgentSnapshotMissionVer2,
  TwitterAgentSnapshotMission
} from "../types";
import {
  ChatCompletionType,
  IDeployAgentTokenParams,
  IGetAgentsParams,
  IUpdateAgentParams,
  IGenerateText,
} from "../types";
import { IChainConnected } from "../types";
import { ICharacter } from "../types";
import { IAgentToken, IGetAgentTokensParams } from "../types/token";
import BaseAPI from "./base";
import { generationTextOpenAi } from "../utils/generation";
import dagentLogger from "../logger";

export interface IAgentClient {
  /** Create new agent */
  create: (params: ICharacter) => Promise<IAgent>;
  /** Updte agent */
  update: (params: IUpdateAgentParams) => Promise<IAgent>;
  /** Delete agent */
  delete: (agentId: string) => Promise<boolean>;
  /** Deploy token for agent */
  deployToken: (params: IDeployAgentTokenParams) => Promise<IAgent>;
  /** Check Rune ticker */
  checkRuneTicker: (ticker: string) => Promise<boolean>;
  /** Agent list  */
  list: (params: IGetAgentsParams) => Promise<IAgent[]>;
  /** Agent detail  */
  detail: (agentId: string) => Promise<IAgent>;
  /** Agent list  */
  listAgentTokens: (params: IGetAgentTokensParams) => Promise<IAgentToken[]>;
  /** Agent token detail  */
  getAgentById: (agentId: string) => Promise<IAgentToken>;
  /** Chat completions  */
  chatCompletions: (prevMessages: ChatCompletionType[]) => Promise<string>;
  /** Get chains  */
  getChainList: () => Promise<IChainConnected[]>;
  /** Setup missions */
  setupMissions: (params: { agentId: string, missions: Array<TwitterAgentSnapshotMission | FarcasterAgentSnapshotMission> }) => Promise<void>;

  agentApps: () => Promise<AgentApps[]>;
  appsInstalled: (params: { agentId: string }) => Promise<AgentApps[]>;
  getInstallCode: (params: { agentId: string, appId: string }) => Promise<string>;
}
export class AgentClient extends BaseAPI implements IAgentClient {
  create = async (params: ICharacter): Promise<IAgent> => {
    try {
      const agent: IAgent = await this.api.post(
        "/agent/create_agent_assistant",
          {
            ...params,
            chain_id: params?.chain_id ? Number(params?.chain_id) : undefined,
          }
      );
      return agent;
    } catch (e) {
      throw e;
    }
  };

  update = async (params: IUpdateAgentParams): Promise<IAgent> => {
    try {
      const agent: IAgent = await this.api.post(
        `/agent/update_agent_assistant`,
        {
          ...params,
          agent_id: params?.id,
          chain_id: Number(params?.chain_id),
        }
      );
      return agent;
    } catch (e) {
      throw e;
    }
  };

  deployToken = async (params: IDeployAgentTokenParams): Promise<IAgent> => {
    try {
      const agent: IAgent = await this.api.post(
        `/agent/update_agent_assistant`,
        { ...params, chain_id: Number(params?.chain_id) }
      );
      return agent;
    } catch (e) {
      throw e;
    }
  };

  list = async (params: IGetAgentsParams): Promise<IAgent[]> => {
    try {
      const res: IAgent[] = await this.api.get(`/agent/dojo/list`, {
        params,
      });
      return res as IAgent[];
    } catch (error) {
      throw error;
    }
  };

  detail = async (agentId: string) => {
    try {
      const res: IAgent = await this.api.get(`/agent/dojo/${agentId}`);
      return res;
    } catch (error) {
      throw error;
    }
  };

  delete = async (agentId: string) => {
    try {
      await this.api.delete(`/agent/${agentId}`);
      return true;
    } catch {
      return false;
    }
  };

  listAgentTokens = async (params: IGetAgentTokensParams) => {
    try {
      const response = (await this.api.get("/agent/dashboard", {
        params,
      })) as any;
      return response?.rows as IAgentToken[];
    } catch (error) {
      throw error;
    }
  };

  getAgentById = async (agentId: string): Promise<IAgentToken> => {
    try {
      const response = (await this.api.get(
        `/agent/${agentId}`
      )) as any;
      return response;
    } catch (error) {
      throw error;
    }
  };

  chatCompletions = async (prevMessages: ChatCompletionType[]) => {
    try {
      const res: string = await this.api.post("/agent/preview/v1", {
        messages: JSON.stringify(prevMessages),
      });
      return res;
    } catch (e) {
      throw e;
    }
  };

  getChainList = async () => {
    try {
      const chainList: IChainConnected[] = await this.api.get(
        "/chain-config/list"
      );
      return chainList;
    } catch (error) {
      throw error;
    }
  };

  checkRuneTicker = async (ticker: string) => {
    try {
      await this.api.get(`/agent/validate-ticker?ticker=${ticker}`);
      return true;
    } catch {
      return false;
    }
  };

  setupMissions = async (params: {
    agentId: string, missions: Array<any>
  }) => {
    dagentLogger.info(`Setting up missions for agent ${params.agentId}`, params.missions);
    await this.api.post(`/agent/mission/update/${params.agentId}`, params.missions);
  };

  setupMissionsVer2 = async (params: {
    agentId: string, missions: Array<AgentSnapshotMissionVer2>
  }) => {
    dagentLogger.info(`Setting up missions for agent ${params.agentId}`, params.missions);
    await this.api.post(`/agent/mission/update/${params.agentId}`, params.missions);
  }

  generateText = async (params: IGenerateText): Promise<string> => {
    let generatedText: string = "";
    switch (params.aiProvider) {
      case "openai": {
        generatedText = await generationTextOpenAi(params.options);
      }
    }
    return generatedText;
  }

  agentApps = async (): Promise<AgentApps[]> => {
    const apps = (await this.api.get(`/agent-store/list`)) as AgentApps[];
    return apps || [];
  }

  appsInstalled = async (params: { agentId: string }): Promise<AgentApps[]> => {
    const res = (await this.api.get(`/agent-store/install/list?agent_info_id=${params.agentId}`)) as AgentApps[];
    return res || [];
  }

  getInstallCode = async (params: { agentId: string, appId: string }): Promise<string> => {
    const code = (await this.api.get(
        `/agent-store/${params.appId}/install-code/${params.agentId}`,
    )) as string;
    return code;
  }
}
