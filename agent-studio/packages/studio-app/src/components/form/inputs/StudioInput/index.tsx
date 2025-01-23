import { Input } from "@chakra-ui/react";
import cn from "classnames";
import { ChangeEventHandler, memo } from "react";
import { TextStyleMap } from "../../styles";

type Props = {
  className?: string;
  value: string;
  onChange: ChangeEventHandler<HTMLInputElement>;
  placeholder?: string;
  disabled?: boolean;
};

const StudioInput = ({
  className,
  value,
  onChange,
  placeholder,
  disabled = false,
  ...rest
}: Props) => {
  return (
    <Input
      {...rest}
      className={cn("nowheel", className)}
      height={"28px"}
      width={"246px"}
      borderRadius={"999px"}
      border={"none"}
      bgColor={"#fff"}
      {...(TextStyleMap.INPUT_STYLE as any)}
      _placeholder={TextStyleMap.PLACEHOLDER_STYLE}
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

export default memo(StudioInput);
