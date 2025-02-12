import { IDeployAgentTokenParams } from "./api";
import { AgentChainId } from "./chain";
import {FarcasterAgentSnapshotMission, AgentSnapshotMissionVer2, TwitterAgentSnapshotMission} from "./agent";

export interface ModelConfiguration {
  maxSteps?: number;
  temperature?: number;
  maxTokens?: number;
  frequencyPenalty?: number;
  presencePenalty: number;
  maxInputTokens?: number;
  maxResponseLength?: number;
}

export interface ICharacter {
  /** Chain id */
  chain_id: AgentChainId;

  /** Character name */
  agent_name: string;

  /** system prompt */
  system_content: string;

  /** Optional Character biography */
  bio?: string[];

  /** Optional  Character background lore */
  lore?: string[];

  /** Optional Knowledge base */
  knowledge?: string[];

  /** Optional Example messages */
  messageExamples?: IMessageExample[][];

  /** Optional Post messages */
  postExamples?: string[];

  /** Known topics */
  topics?: string[];

  /** Writing style guides */
  style?: {
    all?: string[];
    chat?: string[];
    post?: string[];
  };

  /** Character traits */
  adjectives?: string[];

  /** Hermes 3 70B, INTELLECT-1 10B, Llama 3.1 405B  */
  agent_base_model?: string;
}

export type AiProvider = "openai";

export interface IAgentCharacter {
  character: ICharacter;
  deployToken?: IDeployAgentTokenParams;
  twitterMissions?: TwitterAgentSnapshotMission[];
  farcasterMissions?: FarcasterAgentSnapshotMission[];
  agentMissions?: AgentSnapshotMissionVer2[]
  settings?: {
    aiProvider: AiProvider;
    modelConfig?: ModelConfiguration;
  }
}

export interface GenerateTextOptions {
  apiKey: string;
  baseURL: string;
  model: string;
  systemContent: string;
  prompt: string;
  modelConfig?: ModelConfiguration;
}

export interface IMessageExample {
  user: string;
  content: {
    text: string;
  };
}

export interface IGenerateText {
  aiProvider: AiProvider;
  options: GenerateTextOptions;
}