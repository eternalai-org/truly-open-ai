import {
  OnAddPayload,
  OnCreatePayload,
  OnMergePayload,
  OnSnapPayload,
} from "@agent-studio/studio-dnd";
import { onlyOptionSameCategoryInTree } from "../../shared/validators";

const onAddValidate = (data: OnAddPayload) => {
  return onlyOptionSameCategoryInTree(data);
};
const onSnapValidate = (data: OnSnapPayload) => {
  return onlyOptionSameCategoryInTree(data);
};
const onSplitValidate = () => {
  return true;
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

export const UPDATE_TOKEN = {
  onAddValidate,
  onSnapValidate,
  onSplitValidate,
  onMergeValidate,
  onDropInValidate,
  onDropOutValidate,
};
