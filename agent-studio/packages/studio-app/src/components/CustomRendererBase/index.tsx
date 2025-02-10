import { Flex } from "@chakra-ui/react";
import { PropsWithChildren } from "react";
import CustomRendererBaseTitle from "./CustomRendererBaseTitle";

type Props = PropsWithChildren & {
  title?: string;
  tag?: string;
  actions?: React.ReactNode;
  isSelected?: boolean;
  onClick?: () => void;
};

const CustomRendererBase = ({
  title,
  tag,
  children,
  actions,
  isSelected = false,
  onClick,
}: Props) => {
  return (
    <Flex flexDir={"column"} gap={"4px"} paddingTop={".2em"} onClick={onClick}>
      <CustomRendererBaseTitle tag={tag} title={title} />

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
