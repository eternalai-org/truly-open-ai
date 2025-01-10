import path from "path";
import dotenv from "dotenv";

const __dirname = path.dirname(new URL(import.meta.url).pathname);
dotenv.config({ path: path.resolve(__dirname, "../../../.env") });

const getEnvironment = () => {
  const PRIVATE_KEY = process.env.OWNER_PRIVATE_KEY as string;
  const ETERNAL_AI_URL = process.env.ETERNAL_AI_URL as string;
  const TWITTER_CLIENT_ID = process.env.TWITTER_CLIENT_ID as string;

  const FARCASTER_CLIENT_ID = process.env.FARCASTER_CLIENT_ID as string;
  return {
    PRIVATE_KEY,
    ETERNAL_AI_URL,
    TWITTER: {
      CLIENT_ID: TWITTER_CLIENT_ID
    },
    FARCASTER: {
      CLIENT_ID: FARCASTER_CLIENT_ID
    }
  };
};

export {
  getEnvironment
};