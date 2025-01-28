import {
  OnAddPayload,
  OnCreatePayload,
  OnMergePayload,
  OnSnapPayload,
} from "@agent-studio/studio-dnd";

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
  return false;
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
