import { Flex, Box, Text } from "@chakra-ui/react";
import React, { PropsWithChildren } from "react";
import { TextStyleMap } from "../../styles";

type Props = PropsWithChildren & {
  errorMessage?: string;
};

const StudioFieldInputPlaceholder = ({ children, errorMessage }: Props) => {
  return (
    <Flex flexDir={"column"} gap={"4px"}>
      {children}

      <Box height={"12px"}>
        {errorMessage && (
          <Text {...(TextStyleMap.ERROR_STYLE as any)} color={"red"}>
            {errorMessage}
          </Text>
        )}
      </Box>
    </Flex>
  );
};

export default StudioFieldInputPlaceholder;
