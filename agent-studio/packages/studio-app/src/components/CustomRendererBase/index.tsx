import { Flex, Text } from "@chakra-ui/react";
import React, { PropsWithChildren } from "react";
import { TextStyleMap } from "../form/styles";

type Props = PropsWithChildren & {
  title?: string | React.ReactNode;
  actions?: React.ReactNode;
};

const CustomRendererBase = ({ title, children, actions }: Props) => {
  return (
    <Flex flexDir={"column"} gap={"4px"}>
      {title && (
        <Text {...(TextStyleMap.FORM_HEADING_STYLE as any)}>{title}</Text>
      )}

      {children}

      {actions && (
        <Flex align={"center"} justify={"flex-end"} gap={"10px"}>
          {actions}
        </Flex>
      )}
    </Flex>
  );
};

export default CustomRendererBase;
