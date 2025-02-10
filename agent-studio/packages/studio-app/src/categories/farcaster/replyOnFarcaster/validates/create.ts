import {
  OnAddPayload,
  OnCreatePayload,
  OnLinkPayload,
} from "@agent-studio/studio-dnd";
import { farcasterMissionLinkToAllowed } from "../../../shared/validators";

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
  return farcasterMissionLinkToAllowed(data);
};

const CREATE_FLOW_FARCASTER_ON_X_VALIDATORS = {
  onAddValidate,
  onSnapValidate,
  onSplitValidate,
  onMergeValidate,
  onDropInValidate,
  onDropOutValidate,
  onLinkValidate,
};

export default CREATE_FLOW_FARCASTER_ON_X_VALIDATORS;
