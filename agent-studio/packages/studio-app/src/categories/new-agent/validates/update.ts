import { OnAddPayload } from "@agent-studio/studio-dnd";

const onAddValidate = (data: OnAddPayload) => {
  // TODO: Block adding new personality to the root

  return false;
};

const onSnapValidate = () => {
  return false;
};

const onSplitValidate = () => {
  return false;
};

const onMergeValidate = () => {
  return false;
};

const onDropInValidate = () => {
  return false;
};

const onDropOutValidate = () => {
  return false;
};

const UPDATE_FLOW_NEW_AGENT_VALIDATORS = {
  onAddValidate,
  onSnapValidate,
  onSplitValidate,
  onMergeValidate,
  onDropInValidate,
  onDropOutValidate,
};

export default UPDATE_FLOW_NEW_AGENT_VALIDATORS;
