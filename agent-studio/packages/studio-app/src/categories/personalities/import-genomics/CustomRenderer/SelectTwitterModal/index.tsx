import {
  Box,
  Button,
  Flex,
  FormControl,
  Image,
  Input,
  Text,
} from "@chakra-ui/react";
import React, { useState } from "react";
import s from "./SelectTwitterModal.module.scss";
import { useStudio } from "@agent-studio/studio-dnd";
import { ImportGenomicFormData } from "../../types";
import { useDebounce } from "../../../../../hooks/useDebounce";
import { getUrlAvatarTwitter } from "../../../../../utils/twitter";
import AgentAPI from "../../../../../services/apis/agent";
import { useQuery } from "@tanstack/react-query";

type Props = {
  formId: string;
  closeModal: () => void;
};

const SelectTwitterModal = ({ formId, closeModal }: Props) => {
  const { getFormDataById, setFormFields } = useStudio();
  const { twitterInfos } = getFormDataById<ImportGenomicFormData>(formId);

  const [values, setValues] = useState("");
  const [errorMsg, setErrorMsg] = useState("");

  const debounceTwitterUsername = useDebounce(values, 1000);

  const { data: twitterData, error } = useQuery({
    queryKey: ["getTwitterInfo", debounceTwitterUsername],
    queryFn: () => {
      return AgentAPI.getTwitterInfo(
        debounceTwitterUsername.replace(/^@+/, "") as string
      );
    },
  });

  const styleInput = {
    height: "48px ",
    padding: "6px 16px",
    fontSize: "15px !important",
    lineHeight: "24px !important",
    fontFamily: "var(--font-inter) !important",
    fontWeight: "400 !important",
    border: "1px solid rgba(255, 255, 255, 0.07)",
    background: "transparent",
    borderRadius: "8px",
    marginBottom: "20px",
  };

  const handleAddTwitter = () => {
    if (!twitterData || !twitterData[0]) return;
    if (twitterInfos && twitterInfos.length > 0) {
      const foundTw = twitterInfos.find(
        (item) => item?.id === twitterData[0].id
      );
      if (foundTw) {
        setErrorMsg("This account is already added");
        return;
      }
    }

    if (!twitterInfos) {
      setFormFields<ImportGenomicFormData>(formId, {
        twitterInfos: [twitterData[0]],
      });
    } else {
      setFormFields<ImportGenomicFormData>(formId, {
        twitterInfos: [...twitterInfos, twitterData[0]],
      });
    }

    closeModal();
  };

  return (
    <Box p="28px">
      <Flex direction={"column"} gap={"12px"} mb={"20px"}>
        <Text fontSize={"20px"} fontWeight={"500"}>
          Pick the DNA of degens, thinkers, and Twitter’s most fascinating
          personalities to craft your clone!
        </Text>
      </Flex>
      <Flex flexDir={"column"} gap="16px" position="relative">
        <FormControl flex={1}>
          <Input
            fontSize={"13px"}
            id="twitter_username"
            name="twitter_username"
            variant="filled"
            placeholder="X username"
            autoComplete="nope"
            autoFocus
            onChange={(e) => {
              setValues(e.target.value);
              setErrorMsg("");
            }}
            style={{
              ...styleInput,
              border: `1px solid ${!!error ? "#ff4747" : "#e5e7eb"}`,
            }}
          />
          {errorMsg && (
            <Text fontSize={"14px"} color={"#ff4747"} mb="4px">
              {errorMsg}
            </Text>
          )}
        </FormControl>
      </Flex>
      {twitterData && twitterData[0] && (
        <Flex
          direction={"column"}
          // background={'#FFF'}
          className={s.listSearch}
        >
          <Flex
            className={s.nftInfo}
            gap={"12px"}
            justifyContent={"space-between"}
            cursor={"pointer"}
            onClick={() => handleAddTwitter()}
          >
            <Flex gap={"12px"} alignItems={"center"}>
              <Image
                src={getUrlAvatarTwitter(
                  twitterData[0].profile_image_url,
                  "medium"
                )}
                w={"62px"}
                h={"62px"}
                borderRadius={"4px"}
              />
              <Flex direction={"column"} gap={"5px"}>
                <Text fontSize={"16px"} fontWeight={500}>
                  {twitterData[0].name}
                </Text>
                <Text fontSize={"14px"} fontWeight={400} color={"#6B7280"}>
                  @{twitterData[0].username}
                </Text>
              </Flex>
            </Flex>
          </Flex>
        </Flex>
      )}
      <Text fontSize={"14px"} fontWeight={400} mt={"16px"}>
        <Text as={"span"} fontWeight={600}>
          Note:
        </Text>{" "}
        Each DNA sequence costs 300 EAI, and the cloning might take a few hours.
      </Text>
      <Button
        mt="32px"
        w={"100%"}
        h="48px"
        bg="#5400FB"
        color={"#fff"}
        borderRadius="12px"
        boxShadow={" 0px 2px 12px 0px rgba(114, 43, 255, 0.25)"}
        _hover={{
          bg: "#5400FB",
          opacity: 0.8,
        }}
        fontSize={"15px"}
        fontWeight={500}
        onClick={() => {
          handleAddTwitter();
        }}
      >
        Add
      </Button>
    </Box>
  );
};

export default SelectTwitterModal;
