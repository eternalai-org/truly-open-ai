import { Flex, Textarea, Text } from "@chakra-ui/react";
import cn from "classnames";
import { ChangeEventHandler, memo } from "react";
import { TextStyleMap } from "../../styles";

type Props = {
  className?: string;
  value: string;
  onChange: ChangeEventHandler<HTMLTextAreaElement>;
  placeholder?: string;
  disabled?: boolean;
  errorMessage?: string;
};

const StudioTextArea = ({
  className,
  value,
  onChange,
  placeholder,
  disabled = false,
  errorMessage,
  ...rest
}: Props) => {
  return (
    <Flex direction={"column"}>
      <Textarea
        {...rest}
        className={cn("nowheel", className)}
        minH={"160px"}
        w={"100%"}
        minW={"460px"}
        p={"8px 12px"}
        borderRadius={"12px"}
        backgroundColor={"#fff"}
        border={errorMessage ? "1px solid #F56565" : "1px solid transparent"}
        _placeholder={TextStyleMap.TEXTAREA_PLACEHOLDER_STYLE}
        {...(TextStyleMap.TEXTAREA_INPUT_STYLE as any)}
        onMouseDown={(event) => {
          event.stopPropagation();
        }}
        value={value}
        onChange={onChange}
        placeholder={placeholder}
        disabled={disabled}
      />
      {errorMessage && (
        <Text color={"red.400"} fontSize={"12px"}>
          {errorMessage}
        </Text>
      )}
    </Flex>
  );
};

export default memo(StudioTextArea);
