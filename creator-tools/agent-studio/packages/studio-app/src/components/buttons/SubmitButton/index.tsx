import { Button, ButtonProps } from "@chakra-ui/react";
import { ButtonStyleMap } from "../styles";

const SubmitButton = (props: ButtonProps) => {
  return (
    <Button
      minW="92px"
      backgroundColor={"#4185EC"}
      color={"#fff"}
      {...(ButtonStyleMap.DEFAULT_STYLE as any)}
      {...(props as any)}
    >
      {props.title}
    </Button>
  );
};

export default SubmitButton;
