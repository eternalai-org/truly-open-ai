import { Button, ButtonProps } from "@chakra-ui/react";

import { ButtonStyleMap } from "../styles";

type Props = ButtonProps & import("react").RefAttributes<HTMLButtonElement>;

const SubmitButton = (props: Props) => {
  return (
    <Button
      minW="92px"
      backgroundColor={"#4185EC"}
      color={"#fff"}
      isLoading={props.isLoading}
      {...(ButtonStyleMap.DEFAULT_STYLE as any)}
      {...(props as any)}
    >
      {props.title}
    </Button>
  );
};

export default SubmitButton;
