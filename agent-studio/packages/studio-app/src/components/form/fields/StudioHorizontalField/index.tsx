import { Box, Flex, Text } from "@chakra-ui/react";
import { StudioFieldBase } from "../../types";
import StudioFieldLabel from "../StudioFieldLabel";
import StudioFieldTooltip from "../StudioFieldTooltip";
import StudioFieldInputPlaceholder from "../StudioFieldInputPlaceholder";

type Props = StudioFieldBase;

const StudioHorizontalField = ({
  label,
  tooltip,
  children,
  action,
  errorMessage,
}: Props) => {
  return (
    <Flex color={"white"} gap={"8px"} p={"0"}>
      {label && <StudioFieldLabel>{label}</StudioFieldLabel>}

      {tooltip && <StudioFieldTooltip tooltip={tooltip} />}

      <StudioFieldInputPlaceholder errorMessage={errorMessage}>
        {children}
      </StudioFieldInputPlaceholder>

      <Box w={"28px"}>{action}</Box>
    </Flex>
  );
};

export default StudioHorizontalField;
