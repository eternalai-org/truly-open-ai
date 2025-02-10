import {
  BaseDagent,
  IAgentCharacter,
  ICharacter,
} from "@eternalai-dagent/core";
import {
  AGENT_AI_URL,
  FARCASTER_CLIENT_ID,
  OWNER_PRIVATE_KEY,
  TWITTER_CLIENT_ID,
} from "../configs";

export const getBaseAgent = () => {
  const dagentCharacter: IAgentCharacter = {
    character: {} as ICharacter,
  };

  const baseAgent = new BaseDagent({
    dagentCharacter,
    environment: {
      PRIVATE_KEY: OWNER_PRIVATE_KEY,
      ETERNAL_AI_URL: AGENT_AI_URL,
      TWITTER: {
        CLIENT_ID: TWITTER_CLIENT_ID,
      },
      FARCASTER: {
        CLIENT_ID: FARCASTER_CLIENT_ID,
      },
    },
  });
  return baseAgent;
};
