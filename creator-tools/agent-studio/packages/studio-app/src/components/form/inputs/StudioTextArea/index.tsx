import { Textarea } from "@chakra-ui/react";
import cn from "classnames";
import { ChangeEventHandler, memo } from "react";
import { TextStyleMap } from "../../styles";

type Props = {
  className?: string;
  value: string;
  onChange: ChangeEventHandler<HTMLTextAreaElement>;
  placeholder?: string;
  disabled?: boolean;
};

const StudioTextArea = ({
  className,
  value,
  onChange,
  placeholder,
  disabled = false,
  ...rest
}: Props) => {
  return (
    <Textarea
      {...rest}
      className={cn("nowheel", className)}
      minH={"160px"}
      w={"100%"}
      minW={"460px"}
      p={"8px 12px"}
      borderRadius={"12px"}
      backgroundColor={"#fff"}
      border={"none"}
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
  );
};

export default memo(StudioTextArea);
