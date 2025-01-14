import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import CustomRendererBase from "../../../../../components/CustomRendererBase";
import StudioFormWrapper from "../../../../../components/form";
import FrequencyField from "./FrequencyField";
import DetailsField from "./DetailsField";
import { BaseXFormData } from "./types";
import TextRender from "../../../../../components/TextRender";

const CustomBaseXForm = ({
  id,
  formData,
  option,
  setFormFields,
}: StudioCategoryOptionRenderPayload<BaseXFormData>) => {
  const { frequency, details } = formData;

  return (
    <CustomRendererBase
      title={
        <span style={{ color: "#fff" }}>
          <TextRender data={option.title} />
        </span>
      }
    >
      <StudioFormWrapper>
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

export default CustomBaseXForm;
