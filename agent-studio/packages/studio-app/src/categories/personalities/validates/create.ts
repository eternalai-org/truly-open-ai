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

const onAddValidate = (data: OnAddPayload) => {
  if (!onlyOptionSameCategoryInTree(data)) {
    return false;
  }

  return true;
};
const onSnapValidate = (data: OnSnapPayload) => {
  if (!onlyOptionSameCategoryInTree(data)) {
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

  return true;
};
const onDropInValidate = (data: OnCreatePayload) => {
  return true;
};
const onDropOutValidate = () => {
  return true;
};

export const CREATE_PERSONALITY = {
  onAddValidate,
  onSnapValidate,
  onSplitValidate,
  onMergeValidate,
  onDropInValidate,
  onDropOutValidate,
};
