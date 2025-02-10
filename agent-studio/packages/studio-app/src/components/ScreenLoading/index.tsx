import { Image } from "@chakra-ui/react";
import { Flex } from "@chakra-ui/react";
import useStudioAgentStore from "../../stores/useStudioAgentStore";
import cn from "classnames";
import s from "./styles.module.scss";

function ScreenLoading() {
  const { isLoading } = useStudioAgentStore();
  if (isLoading) {
    return (
      <Flex
        position={"absolute"}
        top={0}
        left={0}
        width={"100%"}
        height={"100%"}
        justifyContent={"center"}
        alignItems={"center"}
        zIndex={100}
        backdropFilter={"blur(1px)"}
        background={"rgba(0, 0, 0, 0.1)"}
      >
        <div className={cn(s.loading)}>
          <Image src={"/eai/eai-loading.gif"} alt={"loading"} w={"50px"} />
        </div>
      </Flex>
    );
  }
  return <></>;
}

export default ScreenLoading;
