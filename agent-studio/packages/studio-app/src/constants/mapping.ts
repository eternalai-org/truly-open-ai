import { CATEGORY_OPTION_KEYS } from "./category-option-keys";
import { EMissionToolSet } from "./toolset";

export type MissionXSupport =
  | EMissionToolSet.POST
  | EMissionToolSet.REPLY_MENTIONS
  | EMissionToolSet.REPLY_NON_MENTIONS
  | EMissionToolSet.FOLLOW;

export const MISSION_X_MAPPING: Record<string, MissionXSupport> = {
  [CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post]: EMissionToolSet.POST,
  [CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_reply]:
    EMissionToolSet.REPLY_MENTIONS,
  [CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_engage]:
    EMissionToolSet.REPLY_NON_MENTIONS,
  [CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_follow]: EMissionToolSet.FOLLOW,
};

export const MISSION_X_REVERSE_MAPPING: Record<MissionXSupport, string> = {
  [EMissionToolSet.POST]: CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post,
  [EMissionToolSet.REPLY_MENTIONS]:
    CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_reply,
  [EMissionToolSet.REPLY_NON_MENTIONS]:
    CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_engage,
  [EMissionToolSet.FOLLOW]: CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_follow,
};

export type MissionFarcasterSupport =
  | EMissionToolSet.POST_FARCASTER
  | EMissionToolSet.REPLY_MENTIONS_FARCASTER;

export const MISSION_FARCASTER_MAPPING: Record<
  string,
  MissionFarcasterSupport
> = {
  [CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_post]:
    EMissionToolSet.POST_FARCASTER,
  [CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_reply]:
    EMissionToolSet.REPLY_MENTIONS_FARCASTER,
};

export const MISSION_FARCASTER_REVERSE_MAPPING: Record<
  MissionFarcasterSupport,
  string
> = {
  [EMissionToolSet.POST_FARCASTER]:
    CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_post,
  [EMissionToolSet.REPLY_MENTIONS_FARCASTER]:
    CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_reply,
};
