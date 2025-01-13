import { CATEGORY_OPTION_KEYS } from "./category-option-keys";

export enum EMissionToolSet {
  DEFAULT = "default",
  REPLY_MENTIONS = "reply_mentions",
  REPLY_NON_MENTIONS = "reply_non_mentions",
  FOLLOW = "follow",
  POST = "post",
  POST_SEARCH = "post_search",
  ENGAGE = "engage",
  CREATE_TOKEN = "create_token",
  TRADING = "trading",
  POST_FARCASTER = "post_farcaster",
  REPLY_MENTIONS_FARCASTER = "reply_mentions_farcaster",
  TRADING_FARCASTER = "trading_farcaster",
  POST_SEARCH_FARCASTER = "post_search_farcaster",
}

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
