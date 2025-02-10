import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import CustomRendererNoInput from "../../shared/CustomRendererNoInput";

function CustomDIRendererOnBoard(props: StudioCategoryOptionRenderPayload) {
  return <CustomRendererNoInput {...props.option} />;
}

export default CustomDIRendererOnBoard;
