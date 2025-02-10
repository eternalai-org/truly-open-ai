import React, { useEffect } from "react";
import useHandleSimulateTasks from "../hooks/useHandleSimulateTasks";
import SimulateResult from "../SimulateResult";
import useStudioAgentStore from "../../../stores/useStudioAgentStore";
import { CATEGORY_OPTION_KEYS } from "../../../constants/category-option-keys";

type Props = {};

const ReplySimulate = (props: Props) => {
  const { simulatePrompt } = useStudioAgentStore();

  const { data, handleSubmitSimulate, isLoading } = useHandleSimulateTasks();

  useEffect(() => {
    if (!simulatePrompt) return;
    if (
      simulatePrompt?.simulate_type ===
        CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_reply ||
      simulatePrompt?.simulate_type ===
        CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_reply
    ) {
      const content = `We've been working on CryptoAgents!

The first-ever PFP collection for AI agents.

Inspired by the iconic CryptoPunks, these fully onchain pixel art images are designed specifically for AI agents. Living among us, they need PFPs, too!

Like punks, it's free to claim. Stay tuned!`;

      handleSubmitSimulate(content || "");
    }
  }, [JSON.stringify(simulatePrompt)]);

  if (
    simulatePrompt?.simulate_type !==
      CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_reply &&
    simulatePrompt?.simulate_type !==
      CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_reply
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

export default ReplySimulate;
