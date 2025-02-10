import { useEffect, useState } from "react";
import useHandleSimulateTasks from "../hooks/useHandleSimulateTasks";
import SimulateResult from "../SimulateResult";
import SearchedNews from "./SearchedNews";
import useStudioAgentStore from "../../../stores/useStudioAgentStore";
import { CATEGORY_OPTION_KEYS } from "../../../constants/category-option-keys";

const PostNewsSimulate = () => {
  const { simulatePrompt } = useStudioAgentStore();

  const [newsPrompt, setNewsPrompt] = useState<string[] | null>(null);

  const { data, handleSubmitSimulate, isLoading } = useHandleSimulateTasks();

  useEffect(() => {
    if (!simulatePrompt || !newsPrompt) return;
    if (
      simulatePrompt?.simulate_type ===
      CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_news
    ) {
      handleSubmitSimulate(simulatePrompt?.simulate_prompt || "", newsPrompt);
    }
  }, [JSON.stringify(simulatePrompt), newsPrompt]);

  return (
    <>
      {[
        `${CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_news}_topic`,
        CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_news,
      ].includes(simulatePrompt?.simulate_type || "") && (
        <SearchedNews setNewsPrompt={setNewsPrompt} newsPrompt={newsPrompt} />
      )}

      {simulatePrompt?.simulate_type ===
        CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_news && (
        <SimulateResult
          upperData={data.upperData}
          lowerData={data.lowerData}
          isLoading={isLoading}
        />
      )}

      {/* <Button
        onClick={handleSubmitSimulate}
        disabled={false}
        position={'absolute'}
        zIndex={2}
        right="12px"
        bottom="12px"
        w="fit-content"
        color="#fff"
        fontSize={'14px'}
        fontWeight={500}
        p="0 16px"
        borderRadius={'20px'}
        bg="#000"
        display={'flex'}
        gap="8px"
        h="32px"
        _hover={{
          bg: '#999',
        }}
      >
        Simulate
      </Button> */}
    </>
  );
};

export default PostNewsSimulate;
