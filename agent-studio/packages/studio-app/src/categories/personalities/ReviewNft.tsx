import { Flex, Image, Text, Tooltip } from "@chakra-ui/react";
import React from "react";
import { INFTInfo } from "../../types/collection";
import { getImageIPFSCreateAgent } from "../../utils/common";

type Props = {
  selectedNFT?: INFTInfo;
};

const ReviewNft = ({ selectedNFT }: Props) => {
  if (!selectedNFT?.normalized_metadata?.image) return null;

  return (
    <Tooltip
      label={
        <Flex flexDir={"row"} align={"center"} gap="10px">
          <Image
            w={"32px"}
            h={"32px"}
            borderRadius={"50%"}
            src={getImageIPFSCreateAgent(
              selectedNFT?.normalized_metadata?.image
            )}
          />
          <Text fontSize={"14px"} fontWeight={500} color={"#fff"}>
            {`${selectedNFT?.name} #${selectedNFT?.token_id}`}
          </Text>
        </Flex>
      }
      fontSize={"14px"}
      borderRadius={"8px"}
      p={"6px 12px"}
    >
      <Text
        fontSize={"14px"}
        color="blue"
        textDecoration={"underline"}
        height={"28px"}
        display={"flex"}
        alignItems={"center"}
      >
        View
      </Text>
    </Tooltip>
  );
};

export default ReviewNft;
