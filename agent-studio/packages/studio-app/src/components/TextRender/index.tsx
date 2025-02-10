import React, { FunctionComponent, ReactNode } from "react";
import { Text, TextProps } from "@chakra-ui/react";

type Props = {
  data?: ReactNode | FunctionComponent;
  style?: TextProps;
};

// Convert FunctionComponent | ReactNode to ReactNode
export const convertToReactNode = (
  data: ReactNode | FunctionComponent
): ReactNode => {
  if (typeof data === "function") {
    return data({}) as any;
  }
  return data;
};

function TextRender({ data, ...props }: Props) {
  if (!data) {
    return null;
  }

  const content = convertToReactNode(data);

  return (
    <Text width={"max-content"} {...(props as any)}>
      {content}
    </Text>
  );
}

export default TextRender;
