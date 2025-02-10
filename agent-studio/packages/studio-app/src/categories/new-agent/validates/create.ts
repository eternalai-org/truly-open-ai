import { OnAddPayload } from "@agent-studio/studio-dnd";
import { showValidateError } from "../../../utils/toast";

const onAddValidate = (data: OnAddPayload) => {
  // TODO: Block adding new personality to the root

  return false;
};

const onSnapValidate = () => {
  showValidateError("The root item cannot be snapped");
  return false;
};

const onSplitValidate = () => {
  return false;
};

const onMergeValidate = () => {
  showValidateError("The root item cannot be snapped");
  return false;
};

const onDropInValidate = () => {
  return true;
};

const onDropOutValidate = () => {
  return true;
};

const CREATE_FLOW_NEW_AGENT_VALIDATORS = {
  onAddValidate,
  onSnapValidate,
  onSplitValidate,
  onMergeValidate,
  onDropInValidate,
  onDropOutValidate,
};

export default CREATE_FLOW_NEW_AGENT_VALIDATORS;
