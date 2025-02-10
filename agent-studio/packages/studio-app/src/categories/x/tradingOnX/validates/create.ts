import {
  OnAddPayload,
  OnCreatePayload,
  OnLinkPayload,
} from "@agent-studio/studio-dnd";
import { xMissionLinkToAllowed } from "../../../shared/validators";

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

const onLinkValidate = (data: OnLinkPayload) => {
  return xMissionLinkToAllowed(data);
};

const CREATE_FLOW_TRADING_ON_X_VALIDATORS = {
  onAddValidate,
  onSnapValidate,
  onSplitValidate,
  onMergeValidate,
  onDropInValidate,
  onDropOutValidate,
  onLinkValidate,
};

export default CREATE_FLOW_TRADING_ON_X_VALIDATORS;
