import {
  EFarcasterMissionToolSet,
  ETwitterMissionToolSet,
} from "@eternal-ai/core";
import { CATEGORY_OPTION_KEYS } from "./category-option-keys";

export type MissionXSupport =
  | ETwitterMissionToolSet.POST
  | ETwitterMissionToolSet.REPLY_MENTIONS
  | ETwitterMissionToolSet.REPLY_NON_MENTIONS
  | ETwitterMissionToolSet.FOLLOW;

export const MISSION_X_MAPPING: Record<string, MissionXSupport> = {
  [CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post]:
    ETwitterMissionToolSet.POST,
  [CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_reply]:
    ETwitterMissionToolSet.REPLY_MENTIONS,
  [CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_engage]:
    ETwitterMissionToolSet.REPLY_NON_MENTIONS,
  [CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_follow]:
    ETwitterMissionToolSet.FOLLOW,
};

export const MISSION_X_REVERSE_MAPPING: Record<MissionXSupport, string> = {
  [ETwitterMissionToolSet.POST]:
    CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post,
  [ETwitterMissionToolSet.REPLY_MENTIONS]:
    CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_reply,
  [ETwitterMissionToolSet.REPLY_NON_MENTIONS]:
    CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_engage,
  [ETwitterMissionToolSet.FOLLOW]:
    CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_follow,
};

export type MissionFarcasterSupport =
  | EFarcasterMissionToolSet.POST
  | EFarcasterMissionToolSet.REPLY_NON_MENTIONS;

export const MISSION_FARCASTER_MAPPING: Record<
  string,
  MissionFarcasterSupport
> = {
  [CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_post]:
    EFarcasterMissionToolSet.POST,
  [CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_reply]:
    EFarcasterMissionToolSet.REPLY_NON_MENTIONS,
};

export const MISSION_FARCASTER_REVERSE_MAPPING: Record<
  MissionFarcasterSupport,
  string
> = {
  [EFarcasterMissionToolSet.POST]:
    CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_post,
  [EFarcasterMissionToolSet.REPLY_NON_MENTIONS]:
    CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_reply,
};
