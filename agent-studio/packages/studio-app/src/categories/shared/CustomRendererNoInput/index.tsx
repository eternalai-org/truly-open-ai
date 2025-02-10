import { Flex, Text } from "@chakra-ui/react";
import { StudioCategoryOption } from "@agent-studio/studio-dnd";
import React from "react";
import TextRender, { convertToReactNode } from "../../../components/TextRender";
import IconRenderer from "../../../components/IconRenderer";
import { TextStyleMap } from "../../../components/form/styles";

type Props = StudioCategoryOption & {
  prefix?: React.ReactNode;
  postfix?: React.ReactNode;
};

const CustomRendererNoInput = ({ icon, title, prefix, postfix }: Props) => {
  const processedTitle = convertToReactNode(title);

  return (
    <Flex alignItems="center" gap={"8px"} h="28px" overflow="visible">
      <IconRenderer icon={icon} />
      <TextRender
        data={
          <Flex>
            {prefix ? <>{prefix} </> : ""}
            {processedTitle}
            {postfix ? <> {postfix}</> : ""}
          </Flex>
        }
        {...TextStyleMap.SIDEBAR_CATEGORY_OPTION_LABEL_STYLE}
      />
    </Flex>
  );
};

export default CustomRendererNoInput;
