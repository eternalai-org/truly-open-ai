import {
  OnAddPayload,
  OnCreatePayload,
  OnMergePayload,
  OnSnapPayload,
  findDataByCategoryKey,
  StudioDataNode,
  StudioNode,
  findDataById,
} from "@agent-studio/studio-dnd";
import { onlyOptionSameCategoryInTree } from "../../shared/validators";
import { SUPPORT_NETWORKS } from "../../../constants/networks";
import {
  DECENTRALIZED_INFERENCE_CATEGORY_KEY,
  PERSONALITY_CATEGORY_KEY,
} from "../../../constants/category-keys";
import { tokens } from "../../../constants/tokens";
import { compareString } from "../../../utils/string";
import useCommonStore from "../../../stores/useCommonStore";
import { showValidateError } from "../../../utils/toast";
import { CATEGORY_OPTION_KEYS } from "../../../constants/category-option-keys";
import { NETWORK_SUPPORT_FOR_KNOWLEDGE_AGENT } from "../../../constants/validates";

const validateAiModelOnNetwork = (
  toNode: StudioNode,
  data: StudioDataNode[],
  chainId: SUPPORT_NETWORKS
) => {
  if (chainId) {
    try {
      const targetData = findDataById(toNode.id, data);
      if (!targetData) {
        return true;
      }
      // validate for ai models
      const aiModels = findDataByCategoryKey(
        DECENTRALIZED_INFERENCE_CATEGORY_KEY,
        [targetData]
      );
      if (aiModels?.length) {
        const aiModel = aiModels[0];
        if (aiModel) {
          const currentAiModel = aiModel?.data?.decentralizeId as string;
          if (currentAiModel) {
            const chains = useCommonStore.getState().chains || [];
            const token = tokens.find((v) => compareString(v.id, chainId));
            if (token) {
              const selectChain = chains.find(
                (chain) => chain.chain_id === token?.chainId
              );

              if (selectChain) {
                const isSupported =
                  !!selectChain?.support_model_names?.[currentAiModel];
                if (!isSupported) {
                  showValidateError("Network doesn't support ai model");
                }
                return isSupported;
              }
            }
          }
        }
      }
    } catch (e) {
      //
    }
  }
  return true;
};

const validateKnowledgeOnNetwork = (
  toNode: StudioNode,
  data: StudioDataNode[],
  chainId: SUPPORT_NETWORKS
) => {
  try {
    const targetData = findDataById(toNode.id, data);
    if (!targetData) {
      return true;
    }
    const personalities = findDataByCategoryKey(PERSONALITY_CATEGORY_KEY, [
      targetData,
    ]);
    if (personalities?.length) {
      const personality = personalities[0];
      if (
        personality?.idx ===
        CATEGORY_OPTION_KEYS.personalities.personality_knowledge
      ) {
        const isSupported =
          NETWORK_SUPPORT_FOR_KNOWLEDGE_AGENT.includes(chainId);

        if (!isSupported) {
          showValidateError("Network doesn't support knowledge agent");
        }

        return isSupported;
      }
    }
  } catch (e) {
    //
  }

  return true;
};

const onAddValidate = (data: OnAddPayload) => {
  if (!onlyOptionSameCategoryInTree(data)) {
    return false;
  }

  if (
    !validateKnowledgeOnNetwork(
      data.toNode,
      data.data,
      data?.option?.data?.chainId?.defaultValue as SUPPORT_NETWORKS
    )
  ) {
    return false;
  }

  // validate for ai models
  if (
    !validateAiModelOnNetwork(
      data.toNode,
      data.data,
      data?.option?.data?.chainId?.defaultValue as SUPPORT_NETWORKS
    )
  ) {
    return false;
  }
  return true;
};

const onSnapValidate = (data: OnSnapPayload) => {
  if (!onlyOptionSameCategoryInTree(data)) {
    return false;
  }
  if (
    !validateKnowledgeOnNetwork(
      data.toNode,
      data.data,
      data?.option?.data?.chainId?.defaultValue as SUPPORT_NETWORKS
    )
  ) {
    return false;
  }
  if (
    !validateAiModelOnNetwork(
      data.toNode,
      data.data,
      data?.option?.data?.chainId?.defaultValue as SUPPORT_NETWORKS
    )
  ) {
    return false;
  }
  return true;
};

const onSplitValidate = () => {
  return true;
};

const onMergeValidate = (data: OnMergePayload) => {
  if (!onlyOptionSameCategoryInTree(data)) {
    return false;
  }
  if (
    !validateKnowledgeOnNetwork(
      data.toNode,
      data.data,
      data?.option?.data?.chainId?.defaultValue as SUPPORT_NETWORKS
    )
  ) {
    return false;
  }
  if (
    !validateAiModelOnNetwork(
      data.toNode,
      data.data,
      data?.option?.data?.chainId?.defaultValue as SUPPORT_NETWORKS
    )
  ) {
    return false;
  }
  return true;
};

const onDropInValidate = (data: OnCreatePayload) => {
  return true;
};

const onDropOutValidate = () => {
  return true;
};

export const CREATE_BLOCKCHAIN = {
  onAddValidate,
  onSnapValidate,
  onSplitValidate,
  onMergeValidate,
  onDropInValidate,
  onDropOutValidate,
};
