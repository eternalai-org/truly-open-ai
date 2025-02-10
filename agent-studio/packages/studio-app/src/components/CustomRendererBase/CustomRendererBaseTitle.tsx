import { Flex, Text } from "@chakra-ui/react";
import React from "react";
import { TextStyleMap } from "../form/styles";

type Props = {
  tag?: string;
  title?: string;
};

const CustomRendererBaseTitle = ({ tag, title }: Props) => {
  return (
    <Flex align={"center"} gap={"4px"}>
      {tag && (
        <Text {...(TextStyleMap.FORM_HEADING_TAG_STYLE as any)}>{tag}</Text>
      )}

      {title && (
        <Text {...(TextStyleMap.FORM_HEADING_STYLE as any)}>{title}</Text>
      )}
    </Flex>
  );
};

export default CustomRendererBaseTitle;
