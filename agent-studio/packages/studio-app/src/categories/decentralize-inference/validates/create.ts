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
import { onlyOptionSameCategoryInTree } from "../../shared/validators";
import useCommonStore from "../../../stores/useCommonStore";
import { BLOCKCHAIN_CATEGORY_KEY } from "../../../constants/category-keys";
import { compareString } from "../../../utils/string";
import { tokens } from "../../../constants/tokens";
import { showValidateError } from "../../../utils/toast";

const validateNetworkForAiModel = (
  toNode: StudioNode,
  data: StudioDataNode[],
  aiModel: string
) => {
  if (aiModel) {
    try {
      // validate for ai models
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
          const chainId = network?.data?.chainId as string;
          if (chainId) {
            const chains = useCommonStore.getState().chains || [];
            const token = tokens.find((v) => compareString(v.id, chainId));
            if (token) {
              const selectChain = chains.find(
                (chain) => chain.chain_id === token?.chainId
              );

              if (selectChain) {
                const isSupported =
                  !!selectChain?.support_model_names?.[aiModel];
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

const onAddValidate = (data: OnAddPayload) => {
  if (!onlyOptionSameCategoryInTree(data)) {
    return false;
  }

  if (
    !validateNetworkForAiModel(
      data.toNode,
      data.data,
      data?.option?.data?.decentralizeId?.defaultValue as string
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
    !validateNetworkForAiModel(
      data.toNode,
      data.data,
      data?.option?.data?.decentralizeId?.defaultValue as string
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
    !validateNetworkForAiModel(
      data.toNode,
      data.data,
      data?.option?.data?.decentralizeId?.defaultValue as string
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

export const CREATE_DECENTRALIZE_INFERENCE = {
  onAddValidate,
  onSnapValidate,
  onSplitValidate,
  onMergeValidate,
  onDropInValidate,
  onDropOutValidate,
};
