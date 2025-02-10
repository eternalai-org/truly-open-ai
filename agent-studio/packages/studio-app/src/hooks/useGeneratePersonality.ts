import {
  getAgentPromptByIdea,
  getAgentPromptByNFT,
  getAgentPromptByToken,
} from "../utils/prompt";
import { isEmptyString } from "../utils/validator";
import { getAgentNameFromPersonality } from "../utils/process";
import { BEGIN_OF_GREETING } from "../utils/process";
import { Collection, INFTInfo } from "../types/collection";
import AgentAPI from "../services/apis/agent";
import { AgentDetail } from "../services/apis/studio/types";

const DEFAULT_RETURN_VALUE = {};

type Props = {
  onStartGenerating?: () => void;
  onFinishGenerating?: () => void;
};

type GeneratePersonalityReturnBase = {
  agent?: AgentDetail;
  personality?: string;
};

type GeneratePersonalityFromTokenReturn = GeneratePersonalityReturnBase & {
  tokenInfo?: any;
};

const useGeneratePersonality = ({
  onStartGenerating,
  onFinishGenerating,
}: Props) => {
  const generatePersonalityByNFT = async (
    selectedNFT: INFTInfo,
    selectedOption: Collection
  ): Promise<GeneratePersonalityReturnBase> => {
    if (!selectedNFT) return DEFAULT_RETURN_VALUE;

    try {
      onStartGenerating?.();

      const prompt = getAgentPromptByNFT(selectedNFT, selectedOption);
      const rs = await AgentAPI.generateFromPrompt(prompt as string, 512);

      if (!rs || !rs?.choices) {
        return DEFAULT_RETURN_VALUE;
      }

      const result = rs.choices[0].message.content;
      const personality = result;

      if (personality && personality?.startsWith(BEGIN_OF_GREETING)) {
        const name = getAgentNameFromPersonality(personality);
        const agent: AgentDetail = {
          agent_name: name,
          system_content: personality,
        } as AgentDetail;

        return { agent, personality };
      }

      return DEFAULT_RETURN_VALUE;
    } catch (error) {
      return DEFAULT_RETURN_VALUE;
    } finally {
      onFinishGenerating?.();
    }
  };

  const generatePersonalityByIdea = async (
    idea?: string | null | unknown
  ): Promise<GeneratePersonalityReturnBase> => {
    if (isEmptyString(idea)) {
      return DEFAULT_RETURN_VALUE;
    }

    try {
      onStartGenerating?.();

      const prompt = getAgentPromptByIdea(idea);
      const rs = await AgentAPI.generateFromPrompt(prompt as string, 512);

      if (!rs || !rs?.choices) {
        return DEFAULT_RETURN_VALUE;
      }

      const result = rs.choices[0].message.content;
      const personality = result;

      if (personality && personality?.startsWith(BEGIN_OF_GREETING)) {
        const name = getAgentNameFromPersonality(personality);
        const agent: AgentDetail = {
          agent_name: name,
          system_content: personality,
        } as AgentDetail;

        return { agent, personality };
      }

      return DEFAULT_RETURN_VALUE;
    } catch (error) {
      return DEFAULT_RETURN_VALUE;
    } finally {
      onFinishGenerating?.();
    }
  };

  const generatePersonalityByToken = async (
    contractAddressStr: string
  ): Promise<GeneratePersonalityFromTokenReturn> => {
    try {
      onStartGenerating?.();

      const tokenInfo = await AgentAPI.getAgentTokenInfo(contractAddressStr);

      const prompt = getAgentPromptByToken(tokenInfo);

      const rs = await AgentAPI.generateFromPrompt(prompt as string, 512);

      if (!rs || !rs?.choices) {
        return DEFAULT_RETURN_VALUE;
      }

      const result = rs.choices[0].message.content;
      const personality = result;

      if (personality && personality?.startsWith(BEGIN_OF_GREETING)) {
        const name = getAgentNameFromPersonality(personality);
        const agent: AgentDetail = {
          agent_name: name,
          system_content: personality,
        } as AgentDetail;

        return { agent, personality, tokenInfo };
      }

      return DEFAULT_RETURN_VALUE;
    } catch (error) {
      return DEFAULT_RETURN_VALUE;
    } finally {
      onFinishGenerating?.();
    }
  };

  return {
    generatePersonalityByNFT,
    generatePersonalityByIdea,
    generatePersonalityByToken,
  };
};

export default useGeneratePersonality;
