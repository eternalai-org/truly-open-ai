import { Flex } from "@chakra-ui/react";
import React, { PropsWithChildren } from "react";

type Props = PropsWithChildren;

const StudioFormWrapper = ({ children }: Props) => {
  return (
    <Flex flexDir={"column"} gap={"4px"}>
      {children}
    </Flex>
  );
};

export default StudioFormWrapper;
