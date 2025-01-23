import { Flex } from "@chakra-ui/react";
import { StudioFieldBase } from "../../types";
import StudioFieldInputPlaceholder from "../StudioFieldInputPlaceholder";
import StudioFieldLabel from "../StudioFieldLabel";
import StudioFieldTooltip from "../StudioFieldTooltip";

type Props = StudioFieldBase;

const StudioVerticalField = ({
  label,
  tooltip,
  children,
  errorMessage,
}: Props) => {
  return (
    <Flex color={"white"} flexDir={"column"} gap={"4px"} w={"100%"}>
      <Flex flexDir={"row"} align={"center"} gap="8px" w={"100%"}>
        {label && <StudioFieldLabel>{label}</StudioFieldLabel>}

        {tooltip && <StudioFieldTooltip tooltip={tooltip} />}
      </Flex>

      <StudioFieldInputPlaceholder errorMessage={errorMessage}>
        {children}
      </StudioFieldInputPlaceholder>
    </Flex>
  );
};

export default StudioVerticalField;
