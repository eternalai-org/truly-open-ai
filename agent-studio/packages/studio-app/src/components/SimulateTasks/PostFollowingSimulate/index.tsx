import { useEffect, useState } from "react";
import useHandleSimulateTasks from "../hooks/useHandleSimulateTasks";
import SimulateResult from "../SimulateResult";
import SearchedNews from "./SearchedFollowing";
import SearchedFollowing from "./SearchedFollowing";
import useStudioAgentStore from "../../../stores/useStudioAgentStore";
import { CATEGORY_OPTION_KEYS } from "../../../constants/category-option-keys";

const PostFollowingSimulate = () => {
  const { simulatePrompt } = useStudioAgentStore();

  const [newsPrompt, setNewsPrompt] = useState<string[] | null>(null);

  const { data, handleSubmitSimulate, isLoading } = useHandleSimulateTasks();

  useEffect(() => {
    if (!simulatePrompt || !newsPrompt) return;
    if (
      simulatePrompt?.simulate_type ===
      CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_following
    ) {
      handleSubmitSimulate(simulatePrompt?.simulate_prompt || "", newsPrompt);
    }
  }, [JSON.stringify(simulatePrompt), newsPrompt]);

  return (
    <>
      {[
        `${CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_following}_following`,
        CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_following,
      ].includes(simulatePrompt?.simulate_type || "") && (
        <SearchedFollowing
          setNewsPrompt={setNewsPrompt}
          newsPrompt={newsPrompt}
        />
      )}

      {simulatePrompt?.simulate_type ===
        CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_following && (
        <SimulateResult
          upperData={data.upperData}
          lowerData={data.lowerData}
          isLoading={isLoading}
        />
      )}
    </>
  );
};

export default PostFollowingSimulate;
