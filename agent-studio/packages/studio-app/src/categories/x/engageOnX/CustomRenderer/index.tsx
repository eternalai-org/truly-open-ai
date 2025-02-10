import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import DetailsField from "./Fields/DetailsField";
import FrequencyField from "./Fields/FrequencyField";
import { EngageOnXFormData } from "../types";
import s from "../../styles.module.scss";
import { useEffect, useMemo } from "react";
import { Button, Flex } from "@chakra-ui/react";
import AiModelField from "./Fields/AiModelField";
import useStudioAgentStore from "../../../../stores/useStudioAgentStore";
import {
  findById,
  findParentById,
  findPersonalityObj,
} from "../../../../utils/process";
import { showError } from "../../../../utils/toast";
import CustomRendererBase from "../../../../components/CustomRendererBase";
import StudioFormWrapper from "../../../../components/form";

const CustomEngageOnXRenderer = ({
  id,
  formData,
  data,
  setFormFields,
}: StudioCategoryOptionRenderPayload<EngageOnXFormData>) => {
  const { frequency, details, model } = formData;
  const { setSimulatePrompt, simulatePrompt } = useStudioAgentStore();

  const isActive = useMemo(() => {
    return simulatePrompt?.id.includes(id);
  }, [simulatePrompt]);

  const handleSelectMissionPrompt = () => {
    const result = findById(data, id);
    const parentObj = findParentById(data, id);
    const personalityObj = findPersonalityObj(data, parentObj);

    if (!result) {
      showError({
        message: "Unable to find the selected prompt",
      });
      return;
    }

    if (!personalityObj) {
      showError({
        message:
          "Personality not found. Please connect the prompt to a personality",
      });
      return;
    }

    setSimulatePrompt({
      id: [personalityObj.id, result.id],
      personality: personalityObj?.data?.personality,
      simulate_prompt: result.data.details,
      simulate_type: result.idx,
    });
  };

  return (
    <div data-lego-class={isActive && s.active}>
      <CustomRendererBase
        tag="Ability"
        title="Engage (use CoT to choose a tweet to reply)"
      >
        <StudioFormWrapper>
          <AiModelField
            id={id}
            value={model as string}
            onChange={(id: string, name: string) => {
              setFormFields({ model: id, modelName: name });
            }}
            data={data}
          />
          <FrequencyField
            id={id}
            value={frequency as string}
            onChange={(v: string) => {
              setFormFields({ frequency: v });
            }}
          />

          <DetailsField
            id={id}
            value={details as string}
            onChange={(v: string) => {
              setFormFields({ details: v });
            }}
          />
          <Flex justifyContent={"flex-end"}>
            <Button
              borderRadius={"100px"}
              width={"fit-content"}
              mb="8px"
              onClick={handleSelectMissionPrompt}
            >
              Simulate
            </Button>
          </Flex>
        </StudioFormWrapper>
      </CustomRendererBase>
    </div>
  );
};

export default CustomEngageOnXRenderer;
