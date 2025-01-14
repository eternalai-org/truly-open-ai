import { OnAddPayload } from "@agent-studio/studio-dnd";

const onAddValidate = (data: OnAddPayload) => {
  return true;
};

const onSnapValidate = () => {
  return true;
};

const onSplitValidate = () => {
  return true;
};

const onMergeValidate = () => {
  return true;
};

const onDropInValidate = () => {
  return true;
};

const onDropOutValidate = () => {
  return true;
};

const UPDATE_FARCASTER_VALIDATORS = {
  onAddValidate,
  onSnapValidate,
  onSplitValidate,
  onMergeValidate,
  onDropInValidate,
  onDropOutValidate,
};

export default UPDATE_FARCASTER_VALIDATORS;
