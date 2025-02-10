import s from "./styles.module.scss";
import { Button, Checkbox, Flex, Text } from "@chakra-ui/react";
import { PostNewsTopics } from "../../types";
import useStudioAgentStore from "../../../../../stores/useStudioAgentStore";
import {
  findById,
  findParentById,
  findPersonalityObj,
} from "../../../../../utils/process";
import StudioVerticalField from "../../../../../components/form/fields/StudioVerticalField";
import StudioInput from "../../../../../components/form/inputs/StudioInput";
import { StudioDataNode } from "@agent-studio/studio-dnd";

type Props = {
  id: string;
  value: PostNewsTopics;
  data: StudioDataNode[];
  onChange: (v: PostNewsTopics) => void;
  topic: string;
  setTopic: (v: string) => void;
};

const SearchTopicField = ({
  id,
  value,
  onChange,
  data,
  topic,
  setTopic,
}: Props) => {
  const { simulatePrompt, setSimulatePrompt } = useStudioAgentStore();

  const { values, bingSearch, twitterSearch } = value;

  const handleChangeTopics = (v: string) => {
    onChange({
      ...value,
      values: v,
    });
  };

  const handleUpdateBing = () => {
    onChange({
      ...value,
      bingSearch: !bingSearch,
    });
  };

  const handleUpdateTwitter = () => {
    onChange({
      ...value,
      twitterSearch: !twitterSearch,
    });
  };

  const getRandomValue = (arr: string[]): string => {
    const randomIndex = Math.floor(Math.random() * arr.length);
    return arr[randomIndex];
  };

  const getTopic = () => {
    const topics = values.split(",");
    const randomTopic = getRandomValue(topics);
    setTopic(randomTopic.trim());

    return randomTopic.trim();
  };

  const handleSearchNews = async (topic: string) => {
    const result = findById(data, id);
    const parentObj = findParentById(data, id);
    const personalityObj = findPersonalityObj(data, parentObj);

    // if (!personalityObj) {
    //   showError({
    //     message:
    //       'Personality not found. Please connect the ability to a personality',
    //   });
    //   return;
    // }

    setSimulatePrompt({
      id: [id],
      personality: personalityObj?.data?.personality,
      simulate_prompt: result?.data?.details || "",
      simulate_type: `${result?.idx}_topic`,
      topics: {
        values: topic,
        bingSearch: bingSearch,
        twitterSearch: twitterSearch,
      },
    });
  };

  return (
    <StudioVerticalField
      label="Topics "
      tooltip="List topics separated by commas. Example: crypto, AI, decentralized inference,..."
    >
      <StudioInput
        value={values}
        onChange={(e) => handleChangeTopics(e.target.value)}
        placeholder={`List topics separated by commas. Example: crypto, AI, decentralized inference,...`}
        width={"100% !important"}
      />
      <Flex
        gap="20px"
        alignItems={"center"}
        justifyContent={"flex-end"}
        mt="8px"
      >
        <Flex alignItems={"center"} gap="12px">
          <Text fontSize={"16px"}>Search tool:</Text>
          <Checkbox
            fontSize={"16px"}
            className={s.checkbox}
            defaultChecked={!!bingSearch}
            onChange={handleUpdateBing}
          >
            Bing
          </Checkbox>
          <Checkbox
            fontSize={"16px"}
            className={s.checkbox}
            defaultChecked={!!twitterSearch}
            onChange={handleUpdateTwitter}
          >
            X
          </Checkbox>
        </Flex>
        {/* {topic && <Text fontSize={'15px'}>Topic: {topic}</Text>} */}
        <Button
          h={"32px"}
          p={"8px 20px"}
          color="#FFF"
          fontSize={"16px"}
          fontWeight={400}
          borderRadius={"100px"}
          bg={"#000"}
          _hover={{
            background: "#000",
          }}
          disabled={!values || (!bingSearch && !twitterSearch)}
          onClick={() => {
            const _topic = getTopic();
            handleSearchNews(_topic);
          }}
          pointerEvents={!values ? "none" : "auto"}
          opacity={!values ? 0.5 : 1}
        >
          Search
        </Button>
      </Flex>
    </StudioVerticalField>
  );
};

export default SearchTopicField;
