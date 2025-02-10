import axios from "axios";
import React, { useEffect, useState } from "react";
import s from "./SearchedFollowing.module.scss";
import { UnorderedList } from "@chakra-ui/react";
import useStudioAgentStore from "../../../stores/useStudioAgentStore";

type Props = {
  newsPrompt: string[] | null;
  setNewsPrompt: React.Dispatch<React.SetStateAction<string[] | null>>;
};

const SearchedFollowing = ({ newsPrompt, setNewsPrompt }: Props) => {
  const { simulatePrompt } = useStudioAgentStore();
  const [loading, setLoading] = useState(false);
  //   const [newsPrompt, setNewsPrompt] = useState<string[] | null>(null);

  // const fetchTweetFromFollowingTwitter = async (
  //   twitter_username: string,
  //   hours: number,
  // ) => {
  //   setLoading(true);
  //   const authToken =
  //     LocalStorage.getItem(STORAGE_KEYS.NAKA_WALLET_API_ACCESS_TOKEN) ||
  //     LocalStorage.getItem(STORAGE_KEYS.NAKA_WALLET_AUTHEN);
  //   try {
  //     const res = await axios.get(
  //       `https://offchain-auto-agent-api.eternalai.org/api/post-v3-sample-content/?twitter_username=${twitter_username}&&cutoff_hour=${hours}`,
  //       {
  //         headers: {
  //           Authorization: authToken,
  //           'X-Token': NEWS_QUERY_TOKEN,
  //         },
  //       },
  //     );

  //     if (res.data.data.length === 0) {
  //       showError({
  //         message: 'No tweets to show. Follow more accounts on X.',
  //       });
  //     }

  //     setNewsPrompt(res.data.data);

  //     //   return res.data.data;
  //   } catch (error: any) {
  //     showError({
  //       message: error.message,
  //     });
  //     setNewsPrompt(null);
  //   } finally {
  //     setLoading(false);

  //     //   setLoadedFetchTweet(true);
  //   }
  // };

  useEffect(() => {
    if (!!simulatePrompt?.fetchTwPosts) {
      setNewsPrompt(simulatePrompt.fetchTwPosts);
    }
  }, [JSON.stringify(simulatePrompt)]);

  return (
    <div className={s.wrapper}>
      <UnorderedList display={"flex"} flexDirection={"column"} gap="12px">
        {loading
          ? "Loading..."
          : newsPrompt?.map((news, index) => <li key={index}>{news}</li>)}
      </UnorderedList>
    </div>
  );
};

export default SearchedFollowing;
