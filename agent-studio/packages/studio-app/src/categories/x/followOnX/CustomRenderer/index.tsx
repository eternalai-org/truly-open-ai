import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import DetailsField from "./Fields/DetailsField";
import { FollowOnXFormData } from "../types";
import FrequencyField from "./Fields/FrequencyField";
import AiModelField from "./Fields/AiModelField";
import CustomRendererBase from "../../../../components/CustomRendererBase";
import StudioFormWrapper from "../../../../components/form";

const CustomFollowOnXRenderer = ({
  id,
  formData,
  setFormFields,
  data,
}: StudioCategoryOptionRenderPayload<FollowOnXFormData>) => {
  const { frequency, details, model } = formData;

  return (
    <CustomRendererBase tag="Ability" title="Follow">
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
      </StudioFormWrapper>
    </CustomRendererBase>
  );
};

export default CustomFollowOnXRenderer;
