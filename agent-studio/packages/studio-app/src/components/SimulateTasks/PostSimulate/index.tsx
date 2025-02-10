import { useEffect } from "react";
import useHandleSimulateTasks from "../hooks/useHandleSimulateTasks";
import SimulateResult from "../SimulateResult";
import useStudioAgentStore from "../../../stores/useStudioAgentStore";
import { CATEGORY_OPTION_KEYS } from "../../../constants/category-option-keys";

const PostSimulate = () => {
  const { simulatePrompt } = useStudioAgentStore();

  const { data, handleSubmitSimulate, isLoading } = useHandleSimulateTasks();

  useEffect(() => {
    if (!simulatePrompt) return;
    if (
      simulatePrompt?.simulate_type ===
        CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post ||
      simulatePrompt?.simulate_type ===
        CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_post
    ) {
      handleSubmitSimulate(simulatePrompt?.simulate_prompt || "");
    }
  }, [JSON.stringify(simulatePrompt)]);

  if (
    simulatePrompt?.simulate_type !==
      CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post &&
    simulatePrompt?.simulate_type !==
      CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_post
  )
    return null;

  return (
    <div>
      <SimulateResult
        upperData={data.upperData}
        lowerData={data.lowerData}
        isLoading={isLoading}
      />
    </div>
  );
};

export default PostSimulate;
