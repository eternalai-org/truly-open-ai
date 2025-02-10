import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import { useEffect, useMemo, useState } from "react";
import { updateRelatedAgentData } from "../../helpers";
import { NewPersonalityFormData } from "../types";
import s from "./styles.module.scss";
import { Button, Flex } from "@chakra-ui/react";
import useStudioAgentStore from "../../../../stores/useStudioAgentStore";
import useGeneratePersonality from "../../../../hooks/useGeneratePersonality";
import { findParentById } from "../../../../utils/process";
import SubmitButton from "../../../../components/buttons/SubmitButton";
import CustomRendererBase from "../../../../components/CustomRendererBase";
import StudioVerticalField from "../../../../components/form/fields/StudioVerticalField";
import StudioTextArea from "../../../../components/form/inputs/StudioTextArea";

export default function NewPersonality({
  id,
  formData,
  setFormFields,
  data,
}: StudioCategoryOptionRenderPayload<NewPersonalityFormData>) {
  const [isLoading, setIsLoading] = useState(false);
  const { stepper, originalText, personality: personalityValue } = formData;

  const { isDetail, setSimulatePrompt, simulatePrompt } = useStudioAgentStore();
  const { generatePersonalityByIdea } = useGeneratePersonality({
    onStartGenerating: () => setIsLoading(true),
    onFinishGenerating: () => setIsLoading(false),
  });

  const isActive = useMemo(() => {
    return simulatePrompt?.id.includes(id);
  }, [simulatePrompt]);

  useEffect(() => {
    if (personalityValue) {
      setFormFields({
        stepper: 2,
      });
    } else {
      setFormFields({
        stepper: 1,
      });
    }
  }, []);

  const handleGenerateIdea = async () => {
    if (isLoading || !personalityValue) {
      return;
    }

    let text;

    console.log("@JACKIE", stepper);
    if (stepper === 1) {
      text = personalityValue;
      console.log("@JACKIE", personalityValue);
      setFormFields({
        originalText: personalityValue as string,
      });
    } else {
      console.log("@JACKIE2", originalText);
      text = originalText;
    }

    console.log("[GenerateIdea] text =>  ", text);

    const { agent, personality } = await generatePersonalityByIdea(text);

    if (personality) {
      setFormFields({
        personality,
        stepper: 2,
      });
    }

    if (agent) {
      updateRelatedAgentData(data, id, agent.agent_name);
    }
  };

  const handleSelectSystemPrompt = () => {
    const dataObj = findParentById(data, id);

    const prompt = {
      id: [id],
      personality: personalityValue,
      simulate_type: "chat",
      agent_name: dataObj?.data?.agentName || "Agent",
    };

    setSimulatePrompt(prompt);
  };
  const Actions = useMemo(
    () => (
      <Flex alignItems={"center"} gap="4px">
        {!!personalityValue && (
          <Button
            h="24px"
            borderRadius="100px"
            onClick={handleSelectSystemPrompt}
            isDisabled={isLoading}
          >
            Chat
          </Button>
        )}

        <SubmitButton
          isLoading={isLoading}
          onClick={isLoading ? undefined : handleGenerateIdea}
          title={stepper === 2 ? "Re-generate" : "Next"}
        />
      </Flex>
    ),
    [stepper, isLoading, handleGenerateIdea]
  );

  return (
    <div data-lego-class={isActive && s.active}>
      <CustomRendererBase
        tag="Personality"
        title="Brainstorm"
        actions={Actions}
      >
        <StudioVerticalField>
          <StudioTextArea
            // disabled={isDetail}
            value={personalityValue}
            onChange={(e) => {
              setFormFields({ personality: e.target.value });
            }}
            placeholder="Tell us your agent idea, and our AI assistant will bring it to life."
          />
        </StudioVerticalField>
      </CustomRendererBase>
    </div>
  );
}
