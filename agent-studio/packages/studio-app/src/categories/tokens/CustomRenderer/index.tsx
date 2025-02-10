import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import CustomRendererNoInput from "../../shared/CustomRendererNoInput";
import StudioFieldTooltip from "../../../components/form/fields/StudioFieldTooltip";

function CustomTokenRendererOnBoard(props: StudioCategoryOptionRenderPayload) {
  // return <CustomRendererNoInput {...props.option} />;
  return (
    <CustomRendererNoInput
      {...props.option}
      postfix={
        props.option.tooltip ? (
          <StudioFieldTooltip tooltip={props.option.tooltip} />
        ) : (
          <></>
        )
      }
    />
  );
}

export default CustomTokenRendererOnBoard;
