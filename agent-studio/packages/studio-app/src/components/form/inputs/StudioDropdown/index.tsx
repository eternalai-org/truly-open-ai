import React, { useState, useRef } from "react";
import { Box, Flex, Text, IconButton } from "@chakra-ui/react";
import { ChevronDownIcon } from "@chakra-ui/icons";
import { TextStyleMap } from "../../styles";

type DropdownOption<T = any> = {
  value: string | number;
  label: string;
  icon?: React.ReactNode;
  extraData?: T;
};

type Props<T = any> = {
  options: DropdownOption<T>[];
  value?: string | number;
  onChange?: (value: string | number) => void;
  placeholder?: string;
  width?: string | number;
  beforeItemRenderer?: () => React.ReactNode;
  itemRenderer?: (option: DropdownOption) => React.ReactNode;
  inputRenderer?: (option?: DropdownOption) => React.ReactNode;
};

const StudioDropdown = <T,>({
  options,
  value,
  onChange,
  placeholder = "Select option",
  width = "246px",
  itemRenderer,
  inputRenderer,
  beforeItemRenderer,
}: Props<T>) => {
  const [isOpen, setIsOpen] = useState(false);
  const dropdownRef = useRef<HTMLDivElement>(null);

  const selectedOption = options.find((opt) => opt.value === value);

  const handleClickOutside = (e: MouseEvent) => {
    if (
      dropdownRef.current &&
      !dropdownRef.current.contains(e.target as Node)
    ) {
      setIsOpen(false);
    }
  };

  React.useEffect(() => {
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  return (
    <Box
      position="relative"
      width={width}
      ref={dropdownRef}
      onMouseDown={(event) => {
        event.stopPropagation();
      }}
    >
      <Flex
        onClick={() => setIsOpen(!isOpen)}
        align="center"
        justify="space-between"
        borderRadius="9999px"
        bg="white"
        cursor="pointer"
        p="0 12px"
        _hover={{ borderColor: "gray.300" }}
        overflow={"hidden"}
      >
        {inputRenderer ? (
          <Flex align="center" gap="8px" width={"calc(100% - 14px)"}>
            {inputRenderer(selectedOption)}
          </Flex>
        ) : (
          <Flex align="center" gap="8px" width={"calc(100% - 14px)"}>
            {selectedOption?.icon}

            <Text
              {...(TextStyleMap.INPUT_STYLE as any)}
              color={selectedOption ? "black" : "#777777"}
              whiteSpace={"nowrap"}
              overflow={"hidden"}
            >
              {selectedOption?.label || placeholder}
            </Text>
          </Flex>
        )}

        <IconButton
          aria-label="toggle dropdown"
          icon={<ChevronDownIcon />}
          variant="ghost"
          size="md"
          transform={isOpen ? "rotate(180deg)" : undefined}
          transition="transform 0.2s"
        />
      </Flex>

      {isOpen && (
        <Box
          position="absolute"
          top="100%"
          left="0"
          right="0"
          mt="2px"
          maxH="200px"
          overflowY="auto"
          bg="white"
          borderRadius="md"
          boxShadow="lg"
          zIndex={1000}
          onMouseDown={(event) => {
            event.stopPropagation();
          }}
        >
          {beforeItemRenderer && beforeItemRenderer()}

          {options.map((option) => {
            if (itemRenderer) {
              return (
                <Flex
                  key={option.value}
                  p="10px"
                  align="center"
                  gap="8px"
                  bg={option.value === value ? "gray.50" : "white"}
                  cursor="pointer"
                  _hover={{ bg: "gray.50" }}
                  onClick={() => {
                    onChange?.(option.value);
                    setIsOpen(false);
                  }}
                  whiteSpace={"nowrap"}
                  overflow={"hidden"}
                >
                  {itemRenderer(option)}
                </Flex>
              );
            }
            return (
              <Flex
                key={option.value}
                p="10px"
                align="center"
                gap="8px"
                bg={option.value === value ? "gray.50" : "white"}
                cursor="pointer"
                _hover={{ bg: "gray.50" }}
                onClick={() => {
                  onChange?.(option.value);
                  setIsOpen(false);
                }}
                whiteSpace={"nowrap"}
                overflow={"hidden"}
              >
                {option.icon}
                <Text
                  fontSize={"12px"}
                  lineHeight={"calc(20 / 12)"}
                  fontWeight={"500"}
                  color={"black"}
                  fontFamily={"var(--font-SFProDisplay)"}
                >
                  {option.label}
                </Text>
              </Flex>
            );
          })}
        </Box>
      )}
    </Box>
  );
};

export default StudioDropdown;
