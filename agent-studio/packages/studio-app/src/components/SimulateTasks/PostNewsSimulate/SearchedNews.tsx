import axios from "axios";
import React, { useEffect, useState } from "react";
import s from "./SearchedNews.module.scss";
import { UnorderedList } from "@chakra-ui/react";
import useStudioAgentStore from "../../../stores/useStudioAgentStore";
import { showError } from "../../../utils/toast";
import useAgentServiceStore from "../../../stores/useAgentServiceStore";
import { NEWS_QUERY_TOKEN } from "../../../configs";

type Props = {
  newsPrompt: string[] | null;
  setNewsPrompt: React.Dispatch<React.SetStateAction<string[] | null>>;
};

const SearchedNews = ({ newsPrompt, setNewsPrompt }: Props) => {
  const { simulatePrompt } = useStudioAgentStore();
  const [loading, setLoading] = useState(false);
  //   const [newsPrompt, setNewsPrompt] = useState<string[] | null>(null);

  const fetchNewsFromBing = async (_topic: string) => {
    if (!_topic) return;

    setLoading(true);

    const authToken = useAgentServiceStore.getState().accessToken;

    try {
      const res = await axios.get(
        `https://offchain-auto-agent-api.eternalai.org/api/bing-news/?query=${_topic}`,
        {
          headers: {
            Authorization: authToken,
            "X-Token": NEWS_QUERY_TOKEN,
          },
        }
      );

      return res.data.data;
    } catch (error: any) {
      showError({
        message: error.message,
      });
      setNewsPrompt(null);
    } finally {
      setLoading(false);
    }
  };

  const fetchNewsFromTwitter = async (_topic: string) => {
    if (!_topic) return;

    setLoading(true);
    const authToken = useAgentServiceStore.getState().accessToken;
    try {
      const res = await axios.get(
        `https://offchain-auto-agent-api.eternalai.org/api/twitter-news?query=${_topic}`,
        {
          headers: {
            Authorization: authToken,
            "X-Token": NEWS_QUERY_TOKEN,
          },
        }
      );

      return res.data.data;
    } catch (error: any) {
      showError({
        message: error.message,
      });
      setNewsPrompt(null);
    } finally {
      setLoading(false);
    }
  };

  const getNewsFromBing = async (_topic: string) => {
    setLoading(true);
    const res = await fetchNewsFromBing(_topic);

    if (res) {
      setNewsPrompt(res.splice(0, 3));
    }
  };

  const getNewsFromTwitter = async (_topic: string) => {
    setLoading(true);
    const res = await fetchNewsFromTwitter(_topic);
    if (res) {
      setNewsPrompt(res.splice(0, 3));
    }
  };

  const getAllNews = async (_topic: string) => {
    const [bingRes, twRes] = await Promise.all([
      fetchNewsFromBing(_topic),
      fetchNewsFromTwitter(_topic),
    ]);
    if (!bingRes && !twRes) return;

    if (bingRes && twRes) {
      setNewsPrompt([bingRes[0], bingRes[1], twRes[0]]);
      return;
    }
  };

  useEffect(() => {
    if (
      simulatePrompt?.topics?.bingSearch &&
      simulatePrompt?.topics?.twitterSearch
    ) {
      getAllNews(simulatePrompt.topics?.values);
    }

    if (
      simulatePrompt?.topics?.bingSearch &&
      !simulatePrompt?.topics?.twitterSearch
    ) {
      getNewsFromBing(simulatePrompt.topics?.values);
    }

    if (
      !simulatePrompt?.topics?.bingSearch &&
      simulatePrompt?.topics?.twitterSearch
    ) {
      getNewsFromTwitter(simulatePrompt.topics?.values);
    }
  }, [JSON.stringify(simulatePrompt)]);

  if (
    !simulatePrompt?.topics?.bingSearch &&
    !simulatePrompt?.topics?.twitterSearch
  )
    return null;

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

export default SearchedNews;
