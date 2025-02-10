import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import DetailsField from "./Fields/DetailsField";
import { PostOnXFormData } from "../types";
import FrequencyField from "./Fields/FrequencyField";

import { Button, Flex } from "@chakra-ui/react";
import { useEffect, useMemo, useState } from "react";
import s from "../../styles.module.scss";
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

const CustomPostOnXRenderer = ({
  id,
  formData,
  data,
  setFormFields,
}: StudioCategoryOptionRenderPayload<PostOnXFormData>) => {
  const { frequency, details, model } = formData;

  const { setSimulatePrompt, simulatePrompt } = useStudioAgentStore();

  const handleSelectMissionPrompt = () => {
    const result = findById(data, id);
    const parentObj = findParentById(data, id);
    const personalityObj = findPersonalityObj(data, parentObj);

    if (!result.data.details) {
      showError({
        message: "Please fill the details field",
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

  const isActive = useMemo(() => {
    return simulatePrompt?.id.includes(id);
  }, [simulatePrompt]);

  return (
    <div data-lego-class={isActive && s.active}>
      <CustomRendererBase tag="Ability" title="Post (using CoT)">
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

export default CustomPostOnXRenderer;
