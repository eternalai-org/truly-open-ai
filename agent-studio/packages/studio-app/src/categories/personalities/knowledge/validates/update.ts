import {
  OnAddPayload,
  OnLinkPayload,
  OnSnapPayload,
  OnMergePayload,
  OnCreatePayload,
} from "@agent-studio/studio-dnd";
import { onlyOptionSameCategoryInTree } from "../../../shared/validators";

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
  return true;
};

const onDropOutValidate = () => {
  return true;
};

const onLinkValidate = (data: OnLinkPayload) => {
  return true;
};

const UPDATE_KNOWLEDGE_VALIDATORS = {
  onAddValidate,
  onSnapValidate,
  onSplitValidate,
  onMergeValidate,
  onDropInValidate,
  onDropOutValidate,
  onLinkValidate,
};

export default UPDATE_KNOWLEDGE_VALIDATORS;
