import { Flex, Text } from "@chakra-ui/react";
import React, { memo, useRef, useState } from "react";
import cs from "classnames";
import s from "./styles.module.scss";
import { IChatMessage } from "./types";
import ScrollableFeed from "react-scrollable-feed";

type Props = {
  message: IChatMessage;
  isLast: boolean;
};

const ChatMessage = React.forwardRef<HTMLDivElement, Props>(
  ({ message, isLast }: Props, ref: React.ForwardedRef<HTMLDivElement>) => {
    const [animatedText, setAnimatedText] = React.useState("");
    const [animatedId, setAnimatedId] = useState<string[]>([]);
    const scrollableRef = useRef<ScrollableFeed | null>(null);
    const chatInputRef = useRef<any>(null);

    const setText = () => {
      const text = message?.msg;
      const isOld = animatedId.includes(message?.id);
      if (isLast && message?.is_reply && !isOld) {
        let index = 0;
        const animate = () => {
          if (index < text.length) {
            setAnimatedText(text.substring(0, index + 1));
            index++;
            requestAnimationFrame(animate);
            scrollableRef.current?.scrollToBottom();
          }
        };
        animate();
        chatInputRef?.current?.focus();
        setAnimatedId((values: string[]) => [...values, message?.id]);
      } else {
        setAnimatedText(text);
      }
    };
    React.useEffect(() => {
      setText();
    }, []);

    return (
      <div className={s.chat_message_wrapper} ref={ref}>
        {message?.is_reply && (
          <Flex
            direction="row"
            alignItems="center"
            gap="6px"
            width="100%"
            justifyContent={message?.is_reply ? "flex-start" : "flex-end"}
          >
            <Flex
              gap="8px"
              alignItems="center"
              width="100%"
              flexDirection={message?.is_reply ? "row" : "row-reverse"}
            >
              <Text fontSize="15px" fontWeight="600" width="fit-content">
                {message.name}
              </Text>
            </Flex>
          </Flex>
        )}
        <Flex
          className={cs(s.content, message?.is_reply ? s.reply : "")}
          alignSelf={message?.is_reply ? "flex-start" : "flex-end"}
          style={{
            display: "block",
            flexDirection: "column",
          }}
        >
          <pre
            className={cs(s.normalText, {
              [s.align_right]: !message?.is_reply,
            })}
          >
            {animatedText}
          </pre>
        </Flex>
      </div>
    );
  }
);

export default memo(ChatMessage);
