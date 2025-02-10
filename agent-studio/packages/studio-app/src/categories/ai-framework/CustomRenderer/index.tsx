import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import CustomRendererNoInput from "../../shared/CustomRendererNoInput";

function CustomAIFrameworkRendererOnBoard(
  props: StudioCategoryOptionRenderPayload
) {
  return <CustomRendererNoInput {...props.option} />;
}

export default CustomAIFrameworkRendererOnBoard;
