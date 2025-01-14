import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import { NewPersonalityFormData } from "../types";
import useStudioAgentStore from "../../../../../../stores/useStudioAgentStore";

import CustomRendererBase from "../../../../../../components/CustomRendererBase";
import StudioVerticalField from "../../../../../../components/form/fields/StudioVerticalField";
import StudioTextArea from "../../../../../../components/form/inputs/StudioTextArea";

export default function NewPersonality({
  id,
  formData,
  setFormFields,
  data,
}: StudioCategoryOptionRenderPayload<NewPersonalityFormData>) {
  const { personality } = formData;

  const { isDetail } = useStudioAgentStore();

  return (
    <CustomRendererBase title="Brainstorm">
      <StudioVerticalField
        label="Personality"
        tooltip="Personality information..."
      >
        <StudioTextArea
          disabled={isDetail}
          value={personality}
          onChange={(e) => setFormFields({ personality: e.target.value })}
          placeholder="Tell us your agent idea, and our AI assistant will bring it to life."
        />
      </StudioVerticalField>
    </CustomRendererBase>
  );
}
