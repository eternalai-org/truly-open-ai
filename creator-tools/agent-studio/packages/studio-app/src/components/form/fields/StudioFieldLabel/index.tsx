import { Text } from "@chakra-ui/react";
import { PropsWithChildren } from "react";
import { TextStyleMap } from "../../styles";

type Props = PropsWithChildren;

const StudioFieldLabel = ({ children }: Props) => {
  return (
    <Text {...(TextStyleMap.LABEL_STYLE as any)} color={"#ffffff"}>
      {children}
    </Text>
  );
};

export default StudioFieldLabel;
