import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import { TradeAnalyticsOnDefiFormData } from "../types";
import FrequencyField from "./Fields/FrequencyField";
import TokenField from "./Fields/TokenField";
import CustomRendererBase from "../../../../components/CustomRendererBase";
import StudioFormWrapper from "../../../../components/form";

const CustomTradeAnalyticsOnDefiRenderer = ({
  id,
  formData,
  data,
  setFormFields,
}: StudioCategoryOptionRenderPayload<TradeAnalyticsOnDefiFormData>) => {
  const { frequency, token } = formData;

  return (
    <div>
      <CustomRendererBase tag="Ability" title="Trade Analytics">
        <StudioFormWrapper>
          <TokenField
            id={id}
            value={token as string}
            onChange={(v: string) => setFormFields({ token: v })}
          />
          <FrequencyField
            id={id}
            value={frequency as string}
            onChange={(v: string) => {
              setFormFields({ frequency: v });
            }}
          />
        </StudioFormWrapper>
      </CustomRendererBase>
    </div>
  );
};

export default CustomTradeAnalyticsOnDefiRenderer;
