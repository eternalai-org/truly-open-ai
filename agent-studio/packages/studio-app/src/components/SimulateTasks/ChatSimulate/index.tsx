import { Box, Button, Textarea } from "@chakra-ui/react";
import React, { useEffect, useState } from "react";
import SimulateChat from "./Chat";
import { TChatMessage } from "./Chat/types";
import { v4 } from "uuid";
// import { compareString } from "@/utils/string";
// import AgentAPI from "@/services/api/agent";
import useStudioAgentStore from "../../../stores/useStudioAgentStore";
import SvgInset from "../../SvgInset";
import { compareString } from "../../../utils/string";
import AgentAPI from "../../../services/apis/agent";

const ChatSimulate = () => {
  const { setSimulatePrompt, simulatePrompt } = useStudioAgentStore();

  const [isFocus, setIsFocus] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [input, setInput] = useState("");
  const [messages, setMessages] = useState<TChatMessage[]>([]);

  const handleSendMessage = async () => {
    if (!input || !simulatePrompt || isLoading) return;

    const newMessage: TChatMessage = {
      id: v4(),
      msg: input,
      is_reply: false,
      name: "You",
    };
    setIsLoading(true);
    setMessages((messages) => [...messages, newMessage]);
    const historyMessages = [
      {
        role: "system",
        content: simulatePrompt?.personality || "",
      },
      ...messages.map((item) => ({
        role: compareString(item.name, "You") ? "user" : "assistant",
        content: item.msg,
      })),
      { role: "user", content: newMessage.msg },
    ];
    try {
      setInput("");

      const responseMessage = await AgentAPI.chatCompletions({
        messages: historyMessages,
      });

      setMessages((messages) => [
        ...messages,
        {
          id: v4(),
          msg: responseMessage || "Something went wrong!",
          is_reply: true,
          name: simulatePrompt?.agent_name || "Agent",
        },
      ]);
    } catch (error) {
      console.log("error", error);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    setMessages([]);
  }, [simulatePrompt?.id]);

  if (simulatePrompt?.simulate_type !== "chat") return null;

  return (
    <>
      <Box flex="1">
        <SimulateChat messages={messages} isLoading={isLoading} />
      </Box>
      <Box
        position={"relative"}
        w="100%"
        h="100%"
        minH="150px"
        border={"1px solid"}
        borderRadius={"0 0 8px 0"}
        overflow={"hidden"}
        p="12px"
        bg="#fafafa"
        borderColor={"#E0E0E0"}
      >
        <Textarea
          // h="calc(100% - 46px)"
          className="nowheel hide-scrollbar"
          placeholder={"Say anything and see how your agent chats back."}
          resize={"none"}
          disabled={!simulatePrompt?.personality}
          border={"1px solid"}
          borderColor={isFocus ? "#5400FB" : "transparent"}
          h="120px"
          _focus={{
            outline: "none !important",
          }}
          _focusVisible={{
            boxShadow: "none",
          }}
          transition={"all 0.3s ease"}
          p="12px"
          pb="0px"
          fontSize={"1em"}
          onBlur={(e) => {
            setIsFocus(false);
            setInput(e.target.value);
          }}
          onFocus={() => {
            setIsFocus(true);
            //   if (isChangePersonalityAgent) {
            //     onUpdatePersonality?.(true);
            //     setTimeout(() => {
            //       handleUpdatePersonalityAgent();
            //       onUpdatePersonality?.(false);
            //     }, 600);
            //   }
          }}
          onMouseDown={(event) => {
            event.stopPropagation();
          }}
          onChange={(e) => {
            setInput(e.target.value);
          }}
          value={input}
          onKeyDown={(e) => {
            //if chat, allow enter to send message
            if (e.key === "Enter" && simulatePrompt?.simulate_type === "chat") {
              e.preventDefault();
              handleSendMessage();
              // onSendMessage(prompt, 'chat');
            }
          }}
          bg="#fff"
        />
        <Button
          onClick={handleSendMessage}
          disabled={false}
          position={"absolute"}
          zIndex={2}
          right="20px"
          bottom="20px"
          w="fit-content"
          color="#fff"
          fontSize={"14px"}
          fontWeight={500}
          p="0 16px"
          borderRadius={"20px"}
          bg="#000"
          display={"flex"}
          gap="8px"
          h="32px"
          _hover={{
            bg: "#999",
          }}
        >
          <SvgInset svgUrl="/ic-send.svg" size={16} />
          Submit
        </Button>
      </Box>
    </>
  );
};

export default ChatSimulate;
