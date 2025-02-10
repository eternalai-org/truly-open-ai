import { ChatCompletionType } from "@eternalai-dagent/core";
import { ShortAgentToken } from "../../../stores/useCommonStore";
import { ISearchTwitterInfo } from "../../../types/agent";
import { INFTInfo } from "../../../types/collection";
import { agentAPIClient } from "../clients";
import { ICollectionsResponse } from "./types";
import axios from "axios";
import { AgentDetail } from "../studio/types";

let maxRetries = 3;

const AgentAPI = {
  getAgentTokens: async (): Promise<ShortAgentToken[] | undefined> => {
    try {
      const res: ShortAgentToken[] = await agentAPIClient.get(
        `/api/agent/mission/tokens`
      );
      return res;
    } catch {
      return undefined;
    }
  },

  getNFTCollections: async (params: any): Promise<ICollectionsResponse> => {
    try {
      const res: ICollectionsResponse = await agentAPIClient.get(
        `/api/nft/collection`,
        { params: { ...params, order_by: "seven_day_volume" } }
      );

      return res;
    } catch (e) {
      return { collections: [], next: "" };
    }
  },

  getNFTsByCollection: async ({
    contractAddress,
    nftId,
    cursor,
    inscription,
  }: {
    contractAddress: string;
    nftId?: string;
    cursor?: string;
    inscription?: boolean;
  }): Promise<INFTInfo[]> => {
    try {
      const res: INFTInfo[] = await agentAPIClient.get(
        `/api/nft/collection/${contractAddress}${nftId ? `/${nftId}` : ""}`,
        { params: { cursor: cursor, inscription } }
      );

      return res;
    } catch (e) {
      return [];
    }
  },

  getTwitterInfo: async (user_name: string): Promise<ISearchTwitterInfo[]> => {
    try {
      const res: ISearchTwitterInfo[] = await agentAPIClient.get(
        `/api/internal/twitter/user/search`,
        { params: { query: user_name } }
      );
      return res;
    } catch {
      return [];
    }
  },
  getAgentDetail: async (agentId: string) => {
    const res: AgentDetail = await agentAPIClient.get(
      `/api/agent/dojo/${agentId}`
    );
    return res;
  },
  getVerifyXAccount: async ({
    agent_id,
    twitter_client_id,
    twitter_client_secret,
  }: {
    agent_id: string;
    twitter_client_id: string;
    twitter_client_secret: string;
  }): Promise<boolean> => {
    try {
      const res: any = await agentAPIClient.post(
        `/api/agent/twitter-info/${agent_id}`,
        {
          twitter_client_id,
          twitter_client_secret,
        }
      );
      return res;
    } catch {
      return false;
    }
  },
  chatCompletions: async ({
    messages,
    agentId,
    kb_id,
    model_name,
  }: {
    messages: ChatCompletionType[];
    agentId?: number;
    kb_id?: string;
    model_name?: string;
  }): Promise<string> => {
    try {
      const res: string = await agentAPIClient.post(`/api/agent/preview/v1`, {
        messages: JSON.stringify(messages),
        agent_id: agentId,
        kb_id,
        model_name,
      });
      return res;
    } catch (e) {
      return "";
    }
  },
  generateFromPrompt: async (prompt: string, max_tokens: number) => {
    // retry 3 times

    console.log("generateFromPrompt", prompt);

    const payload = {
      model: "NousResearch/Hermes-3-Llama-3.1-70B-FP8",
      messages: [
        {
          role: "system",
          content: "You are a helpful assistant",
        },
        {
          role: "user",
          content: `${prompt}`,
        },
      ],
      max_tokens: max_tokens,
      temperature: 0.0,
      stream: false,
    };

    try {
      const res = await axios.post(
        "https://0yy8mvm59lqqqf-8000.proxy.runpod.net/v1/chat/completions",
        payload,
        {
          headers: {
            Authorization:
              "Bearer d50b6ba5169ea538a71fe7b0685b755823a3746934fa3cc4", // Another example header
          },
        }
      );

      return res.data;
    } catch (error: any) {
      maxRetries = maxRetries - 1;
      if (maxRetries > 0) {
        console.log("aaaa", this);
        // @ts-ignore
        this.generateFromPrompt(prompt);
      }

      return { error: error.message };
    }
  },
  getAgentTokenInfo: async (address: string): Promise<any> => {
    try {
      const res: any = await agentAPIClient.get(
        `/api/agent/token-info/${address}`,
        { params: {} }
      );
      return res;
    } catch {
      return undefined;
    }
  },
};

export default AgentAPI;
