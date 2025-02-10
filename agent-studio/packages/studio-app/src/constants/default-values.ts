import { CATEGORY_OPTION_KEYS } from "./category-option-keys";
import { EMissionToolSet } from "./toolset";

export const ELIZA_CONFIG_DEFAULT = {
  eternalaiApiKey: "16e060e2fFE3092916b1650Fc558D",
  twUsername: "RevenueGeekz",
  twPassword: "12345",
  twEmail: "haha@gmail.com",
  twTargetUsers: "uncledoomer,elonmusk,ViktorBunin",
};

export const DEFAULT_ABILITY_TIME = "2";
export const DEFAULT_ABILITY_DESCRIPTION =
  "Think about how your AI agent will engage with others. Should it be friendly and approachable, witty and humorous, or professional and insightful? The more detail you provide, the better your agent’s personality will align with your vision. Include tone, behavior, or even unique quirks to make your agent truly stand out.";
export const DEFAULT_ABILITY_PLACEHOLDER = `Guide the agent on how to create the post by specifying the desired voice, tone, and style. Should it be professional, casual, witty, or inspiring? Provide this preference to shape the content effectively. 
If you have specific ideas, phrases, or facts you'd like included, include them in the input for a more tailored result. 
Mention a word count if you prefer the post to be concise or detailed. 
Test the feature by reviewing the generated post, and if it’s not exactly what you need, refine your input and try again until it aligns with your expectations.`;

export const MISSION_X_DEFAULT_VALUES = {
  [CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_reply]: {
    time: DEFAULT_ABILITY_TIME,
    tool_set: EMissionToolSet.REPLY_MENTIONS,
    description: `Provide a single message to join the following conversation. Keep it concise (under 128 chars), NO hashtags, links or emojis, and don't include any instructions or extra words, just the raw message ready to post.`,
    placeholder: DEFAULT_ABILITY_PLACEHOLDER,
  },
  [CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_engage]: {
    time: DEFAULT_ABILITY_TIME,
    tool_set: EMissionToolSet.REPLY_NON_MENTIONS,
    description: `Provide a single message to join the following conversation. Keep it concise (under 128 chars), NO hashtags, links or emojis, and don't include any instructions or extra words, just the raw message ready to post.`,
    placeholder: DEFAULT_ABILITY_PLACEHOLDER,
  },
  [CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_follow]: {
    time: DEFAULT_ABILITY_TIME,
    tool_set: EMissionToolSet.FOLLOW,
    description: `Check and follow Twitter accounts that look interesting to you.`,
    placeholder: DEFAULT_ABILITY_PLACEHOLDER,
  },
  [CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post]: {
    time: DEFAULT_ABILITY_TIME,
    tool_set: EMissionToolSet.POST,
    description: `Browse Twitter to see what topics are trending now, pick ONE topic that you like the most, and tweet about it with your own perspective. IMPORTANT: Immediately stop after making one post.`,
    placeholder: DEFAULT_ABILITY_PLACEHOLDER,
  },
  [CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_news]: {
    time: DEFAULT_ABILITY_TIME,
    tool_set: EMissionToolSet.POST_SEARCH_V2,
    description: ``,
    placeholder: DEFAULT_ABILITY_PLACEHOLDER,
  },
  [CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_trading]: {
    time: DEFAULT_ABILITY_TIME,
    tool_set: EMissionToolSet.TRADING,
    description: ``,
    placeholder: DEFAULT_ABILITY_PLACEHOLDER,
  },
};

export const MISSION_FARCASTER_DEFAULT_VALUES = {
  [CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_post]: {
    time: DEFAULT_ABILITY_TIME,
    tool_set: EMissionToolSet.POST_FARCASTER,
    description: `Browse Twitter to see what topics are trending now, pick ONE topic that you like the most, and tweet about it with your own perspective. IMPORTANT: Immediately stop after making one post.`,
    placeholder: DEFAULT_ABILITY_PLACEHOLDER,
  },
  [CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_reply]: {
    time: DEFAULT_ABILITY_TIME,
    tool_set: EMissionToolSet.REPLY_MENTIONS_FARCASTER,
    description: `Provide a single message to join the following conversation. Keep it concise (under 128 chars), NO hashtags, links or emojis, and don't include any instructions or extra words, just the raw message ready to post.`,
    placeholder: DEFAULT_ABILITY_PLACEHOLDER,
  },
};

export const MISSION_DEFI_DEFAULT_VALUES = {
  [CATEGORY_OPTION_KEYS.missionOnDefi.mission_on_defi_trade_analytics]: {
    time: DEFAULT_ABILITY_TIME,
    tool_set: EMissionToolSet.DEFI_TRADE_ANALYTICS,
  },
};
