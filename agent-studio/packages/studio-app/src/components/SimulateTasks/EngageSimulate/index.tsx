import React, { useEffect } from "react";
import useHandleSimulateTasks from "../hooks/useHandleSimulateTasks";
import SimulateResult from "../SimulateResult";
import { Box, Flex } from "@chakra-ui/react";
import { CATEGORY_OPTION_KEYS } from "../../../constants/category-option-keys";
import useStudioAgentStore from "../../../stores/useStudioAgentStore";

type Props = {};

const ENGAGE_OPTIONS = [
  `$GLS path is clear: VC round at $100M, launch at $4M, AI infra consistently hitting $30M+ this month infrastructure plays don't stay cheap long`,
  `$SUI ecosystem seeing major liquidity depth $LOFI now ranked #3 project behind only cetus and deepbook with 179m mcap`,
  `$EIGEN ecosystem showing highest dev growth in crypto for 2024, 2x faster than nearest competitor 40+ AVS projects building on top restaking sector dominance continues w/ $20B TVL`,
];

const EngageSimulate = (props: Props) => {
  const { simulatePrompt } = useStudioAgentStore();

  const { data, handleSubmitSimulate, isLoading } = useHandleSimulateTasks();

  // useEffect(() => {
  //   if (!simulatePrompt) return;
  //   if (
  //     simulatePrompt?.simulate_type ===
  //     CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_engage
  //   ) {
  //     const content = ``;

  //     handleSubmitSimulate(content || '');
  //   }
  // }, [JSON.stringify(simulatePrompt)]);

  const styleEngageOption = {
    p: "12px",
    borderRadius: "8px",
    border: "1px solid #898989",
    cursor: "pointer",
    background: "#fff",
    fontSize: "12px",
    width: "100%",
    flex: "1",
    _hover: {
      borderColor: "#5400FB",
    },
    whiteSpace: "nowrap",
    maxWidth: "475px",
    overflow: "hidden",
    textOverflow: "ellipsis",
  };

  if (
    simulatePrompt?.simulate_type !==
    CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_engage
  )
    return null;

  return (
    <>
      <Box flex="1">
        <SimulateResult
          upperData={data.upperData}
          lowerData={data.lowerData}
          isLoading={isLoading}
        />
      </Box>
      <Flex
        flexDir={"column"}
        gap={"8px"}
        alignItems={"center"}
        w="100%"
        bg="rgb(250, 250, 250)"
        py={"12px"}
      >
        {ENGAGE_OPTIONS.map((option, index) => (
          <Box
            {...styleEngageOption}
            key={index}
            onClick={() => {
              handleSubmitSimulate(option);
            }}
          >
            {option}
          </Box>
        ))}
      </Flex>
    </>
  );
};

export default EngageSimulate;
