import TwitterCloneTube from "./TwitterCloneTube";
import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import { ImportGenomicFormData } from "../types";
import CustomRendererBase from "../../../../components/CustomRendererBase";
import useStudioAgentStore from "../../../../stores/useStudioAgentStore";
import StudioFormWrapper from "../../../../components/form";

function ImportGenomic({
  id,
}: StudioCategoryOptionRenderPayload<ImportGenomicFormData>) {
  const { isDetail } = useStudioAgentStore();
  return (
    <CustomRendererBase
      tag="Personality"
      title="Genomic Labs"
      actions={!isDetail}
    >
      <StudioFormWrapper>
        <TwitterCloneTube formId={id} />
      </StudioFormWrapper>
    </CustomRendererBase>
  );
}

export default ImportGenomic;
