import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import CustomRendererNoInput from "../../shared/CustomRendererNoInput";
import StudioFieldTooltip from "../../../components/form/fields/StudioFieldTooltip";

function CustomBlockchainRendererOnBoard(
  props: StudioCategoryOptionRenderPayload
) {
  return (
    <CustomRendererNoInput
      {...props.option}
      prefix={"Agent on"}
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

export default CustomBlockchainRendererOnBoard;
