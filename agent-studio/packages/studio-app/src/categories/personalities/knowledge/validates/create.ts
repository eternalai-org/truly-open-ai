import {
  findDataByCategoryKey,
  findDataById,
  OnAddPayload,
  OnCreatePayload,
  OnMergePayload,
  OnSnapPayload,
  StudioDataNode,
  StudioNode,
} from "@agent-studio/studio-dnd";
import { onlyOptionSameCategoryInTree } from "../../../shared/validators";
import { CATEGORY_OPTION_KEYS } from "../../../../constants/category-option-keys";
import { SUPPORT_NETWORKS } from "../../../../constants/networks";
import { BLOCKCHAIN_CATEGORY_KEY } from "../../../../constants/category-keys";
import { NETWORK_SUPPORT_FOR_KNOWLEDGE_AGENT } from "../../../../constants/validates";
import { showValidateError } from "../../../../utils/toast";

const validateNetworkForKnowledge = (
  toNode: StudioNode,
  data: StudioDataNode[],
  personality: string
) => {
  try {
    if (
      personality !== CATEGORY_OPTION_KEYS.personalities.personality_knowledge
    ) {
      return true;
    }
    const targetData = findDataById(toNode.id, data);
    if (!targetData) {
      return true;
    }
    const networks = findDataByCategoryKey(BLOCKCHAIN_CATEGORY_KEY, [
      targetData,
    ]);
    if (networks?.length) {
      const network = networks[0];
      if (network) {
        const chainId = network?.data?.chainId as SUPPORT_NETWORKS;
        if (chainId) {
          const isSupported =
            NETWORK_SUPPORT_FOR_KNOWLEDGE_AGENT.includes(chainId);

          if (!isSupported) {
            showValidateError("Network doesn't support knowledge agent");
          }

          return isSupported;
        }
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

  if (!validateNetworkForKnowledge(data.toNode, data.data, data.option.idx)) {
    return false;
  }

  return true;
};
const onSnapValidate = (data: OnSnapPayload) => {
  if (!onlyOptionSameCategoryInTree(data)) {
    return false;
  }

  if (!validateNetworkForKnowledge(data.toNode, data.data, data.option.idx)) {
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

  if (!validateNetworkForKnowledge(data.toNode, data.data, data.option.idx)) {
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

export const CREATE_KNOWLEDGE_VALIDATORS = {
  onAddValidate,
  onSnapValidate,
  onSplitValidate,
  onMergeValidate,
  onDropInValidate,
  onDropOutValidate,
};
