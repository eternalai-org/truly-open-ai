import {
  OnAddPayload,
  OnCreatePayload,
  OnMergePayload,
  OnSnapPayload,
} from "@agent-studio/studio-dnd";
import { onlyOptionSameCategoryInTree } from "../../../utils/validate";

const onAddValidate = (data: OnAddPayload) => {
  return false;
};
const onSnapValidate = (data: OnSnapPayload) => {
  return false;
};
const onSplitValidate = () => {
  return false;
};
const onMergeValidate = (data: OnMergePayload) => {
  return onlyOptionSameCategoryInTree(data);
};
const onDropInValidate = (data: OnCreatePayload) => {
  return false;
};
const onDropOutValidate = () => {
  return false;
};

export const UPDATE_PERSONALITY = {
  onAddValidate,
  onSnapValidate,
  onSplitValidate,
  onMergeValidate,
  onDropInValidate,
  onDropOutValidate,
};
