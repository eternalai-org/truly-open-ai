import { Flex, Text, Image, Box } from "@chakra-ui/react";
import SvgInset from "../../../../../components/SvgInset";
import useStudioAgentStore from "../../../../../stores/useStudioAgentStore";
import { ISearchTwitterInfo } from "../../../../../types/agent";
import { getUrlAvatarTwitter } from "../../../../../utils/twitter";

type Props = {
  info: ISearchTwitterInfo;
  handleDeleteTwitter: (info: ISearchTwitterInfo) => void;
};

const CloseIcon = (props: React.ComponentPropsWithRef<"div">) => {
  return (
    <Box
      {...props}
      position={"absolute"}
      right={0}
      top={0}
      zIndex={1}
      cursor={"pointer"}
    >
      <SvgInset size={30} svgUrl={`/icons/ic_close_message.svg`} />
    </Box>
  );
};

function TwitterClonedView({ info, handleDeleteTwitter }: Props) {
  const { isDetail } = useStudioAgentStore();
  return (
    <Flex
      align={"center"}
      gap={"8px"}
      mt="8px"
      fontSize={"16px"}
      fontWeight={500}
      position={"relative"}
    >
      {!isDetail && <CloseIcon onClick={() => handleDeleteTwitter(info)} />}

      <Image
        w="42px"
        h={"42px"}
        src={getUrlAvatarTwitter(info?.profile_image_url, "medium")}
        objectFit={"cover"}
        borderRadius={"50%"}
        overflow={"hidden"}
      />
      <Flex direction={"column"}>
        <Text
          whiteSpace={"nowrap"}
          overflow={"hidden"}
          textOverflow={"ellipsis"}
          color="#000"
        >
          {info?.name}
        </Text>
        <Text
          color="#5B5B5B"
          fontSize="14px"
          fontStyle="normal"
          fontWeight="400"
          lineHeight="120%"
        >
          @{info?.username}
        </Text>
      </Flex>
    </Flex>
  );
}

export default TwitterClonedView;
