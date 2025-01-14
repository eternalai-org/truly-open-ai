import { Collection, INFTInfo } from "../types/collection";
import { isEmptyString } from "./validator";

export const getAgentPromptByNFT = (
  selectedNFT: INFTInfo,
  selectedOption: Collection
) => {
  if (selectedNFT?.normalized_metadata?.attributes) {
    return selectedNFT?.normalized_metadata?.attributes?.reduce(
      (result, att) => {
        return result + ` ${att.trait_type}: ${att.value}.`;
      },
      `write a cohesive and comprehensive identity of an AI agent for the following NFT, including personality, lores and background, skills, quirks, what it loves and hates, especially cryptocurrency and blockchain technology, starting with "You are [40 char max agent name]!".
    Here is the NFT info: NFT Collection Name: ${selectedOption?.name}. NFT Collection description: ${selectedOption?.description}. Your NFT ID in the collection: ${selectedNFT?.token_id}. Your NFT unique attributes:
    `
    );
  }

  return `NFT Collection Name: ${selectedOption?.name}
            NFT Collection description: ${selectedOption?.description}
            Your NFT ID in the collection: ${selectedNFT?.token_id}
            `;
};

export const getAgentPromptByIdea = (idea?: string | null | unknown) => {
  if (isEmptyString(idea)) return null;

  return `write a cohesive and comprehensive identity of an AI agent for the following idea, including personality, lores and background, skills, quirks, what it loves and hates, especially cryptocurrency and blockchain technology, starting with "You are [40 char max agent name]!".

Here is the idea: ${idea}.`;
};

export const getAgentPromptByToken = (tokenInfo: any) => {
  return `write a cohesive and comprehensive identity of an AI agent for the following token, including personality , lores and background, skills, quirks, what it loves and hates, especially cryptocurrency and blockchain technology, starting with "You are [40 char max agent name]!".

Here is the token: Token name: ${tokenInfo?.baseToken?.name}, token description: ${tokenInfo?.description}
`;
};
