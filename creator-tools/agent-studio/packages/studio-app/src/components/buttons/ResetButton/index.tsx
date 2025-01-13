import { Flex, FlexProps, Text } from "@chakra-ui/react";
import { ButtonStyleMap } from "../styles";

interface IProps extends FlexProps {
  title?: string;
}

const ResetButton = (props: IProps) => {
  return (
    <Flex
      backgroundColor={"#F2F2F2"}
      color={"#555555"}
      onClick={props.onClick}
      {...(ButtonStyleMap.DEFAULT_STYLE as any)}
      {...(props as any)}
    >
      <Text>{props.title || "Retry"}</Text>

      <svg
        width="16"
        height="16"
        viewBox="0 0 16 16"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M6 10L2 6L6 2"
          stroke="#555555"
          stroke-width="1.2"
          stroke-linecap="round"
          stroke-linejoin="round"
        />
        <path
          d="M5.99984 12.6667H10.6665C11.5506 12.6667 12.3984 12.3155 13.0235 11.6904C13.6486 11.0652 13.9998 10.2174 13.9998 9.33333V9.33333C13.9998 8.44928 13.6486 7.60143 13.0235 6.97631C12.3984 6.35119 11.5506 6 10.6665 6H2.6665"
          stroke="#555555"
          stroke-width="1.2"
          stroke-linecap="round"
          stroke-linejoin="round"
        />
      </svg>
    </Flex>
  );
};

export default ResetButton;
