import { Flex, Tooltip } from "@chakra-ui/react";
import React from "react";
import { InfoIcon } from "../../../icons/common";
import { TextStyleMap } from "../../styles";

type Props = {
  tooltip: string;
};

const StudioFieldTooltip = ({ tooltip }: Props) => {
  return (
    <Tooltip
      label={tooltip}
      borderRadius={"8px"}
      p={"6px 12px"}
      {...(TextStyleMap.TOOLTIP_STYLE as any)}
    >
      <Flex w={"28px"} h={"28px"} align={"center"} justify={"center"}>
        <InfoIcon />
      </Flex>
    </Tooltip>
  );
};

export default StudioFieldTooltip;
