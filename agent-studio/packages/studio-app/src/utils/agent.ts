import { BaseDagent, IAgentCharacter, InitAgent } from "@eternalai-dagent/core";
import {
  ETERNAL_AI_URL,
  FARCASTER_CLIENT_ID,
  OWNER_PRIVATE_KEY,
  TWITTER_CLIENT_ID,
} from "../configs";

export const getAgentInstance = async ({
  dagentCharacter,
}: {
  dagentCharacter: IAgentCharacter;
}) => {
  const baseAgent = new BaseDagent({
    dagentCharacter,
    environment: {
      PRIVATE_KEY: OWNER_PRIVATE_KEY,
      ETERNAL_AI_URL: ETERNAL_AI_URL,
      TWITTER: {
        CLIENT_ID: TWITTER_CLIENT_ID,
      },
      FARCASTER: {
        CLIENT_ID: FARCASTER_CLIENT_ID,
      },
    },
  });
  await baseAgent.configAccessToken();
  return baseAgent;
};
