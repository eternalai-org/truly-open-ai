import { useEffect, useMemo, useState } from "react";
import {
  Flex,
  Image,
  Text,
  Input,
  InputGroup,
  InputLeftElement,
} from "@chakra-ui/react";
import StudioHorizontalField from "../../../../../components/form/fields/StudioHorizontalField";
import useCommonStore, {
  ShortAgentToken,
} from "../../../../../stores/useCommonStore";
import StudioDropdown from "../../../../../components/form/inputs/StudioDropdown";
import AgentAPI from "../../../../../services/apis/agent";

type Props = {
  id: string;
  value: string;
  onChange: (v: string) => void;
};

const TokenField = ({ id, value, onChange }: Props) => {
  const { agentTokens, setAgentTokens } = useCommonStore();

  const [searchValue, setSearchValue] = useState("");

  useEffect(() => {
    if (!agentTokens?.length) {
      fetchData();
    }
  }, [agentTokens?.length]);

  const fetchData = async () => {
    try {
      const rs: any = await AgentAPI.getAgentTokens();
      setAgentTokens(rs);
    } catch (error) {}
  };

  const options = useMemo(() => {
    if (searchValue) {
      return agentTokens
        .map((v) => ({
          label: v.name,
          value: v.symbol,
          extraData: v,
        }))
        .filter(
          (v) =>
            v.value.toUpperCase().includes(searchValue.toUpperCase()) ||
            v.label.toUpperCase().includes(searchValue.toUpperCase())
        );
    }
    return agentTokens.map((v) => ({
      label: v.name,
      value: v.symbol,
      extraData: v,
    }));
  }, [agentTokens, searchValue]);

  return (
    <StudioHorizontalField label="Token">
      <StudioDropdown<ShortAgentToken>
        value={value}
        onChange={(v) => {
          onChange(v as string);
        }}
        placeholder="Token"
        options={options}
        beforeItemRenderer={() => (
          <InputGroup mb={"6px"} alignItems={"center"}>
            <InputLeftElement height={"100%"} pointerEvents="none">
              <Image src="/icons/pump/ic-search.svg" />
            </InputLeftElement>
            <Input
              value={searchValue}
              style={{
                paddingLeft: "30px",
                height: "32px",
              }}
              type="text"
              color="black"
              placeholder="Search tokens"
              onChange={(e) => {
                setSearchValue(e.target.value);
              }}
            />
          </InputGroup>
        )}
        inputRenderer={(option) => (
          <>
            {option ? (
              <Flex color={"black"} alignItems={"center"} gap={"6px"}>
                <Image
                  src={option.extraData?.image_url}
                  width={"20px"}
                  height={"20px"}
                />
                <Text fontSize={"14px"} lineHeight={"140%"} fontWeight={"500"}>
                  {option.extraData?.name}
                </Text>
                <Text fontSize={"12px"} opacity={"0.7"} lineHeight={"140%"}>
                  {option.extraData?.symbol}
                </Text>
              </Flex>
            ) : (
              <Text color={"black"} fontSize={"14px"} lineHeight={"140%"}>
                Select Token
              </Text>
            )}
          </>
        )}
        itemRenderer={(option) => (
          <Flex
            cursor={"pointer"}
            key={option.value}
            alignItems={"center"}
            gap={"6px"}
            color={"black"}
          >
            <Image
              src={option.extraData?.image_url}
              width={"20px"}
              height={"20px"}
            />
            <Flex alignItems={"center"} gap={"2px"}>
              <Text fontSize={"14px"} lineHeight={"140%"} fontWeight={"500"}>
                {option.extraData?.name}
              </Text>
              <Text fontSize={"12px"} opacity={"0.7"} lineHeight={"140%"}>
                {option.extraData?.symbol}
              </Text>
            </Flex>
          </Flex>
        )}
      />
    </StudioHorizontalField>
  );
};

export default TokenField;
