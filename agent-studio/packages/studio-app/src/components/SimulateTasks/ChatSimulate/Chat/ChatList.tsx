import React, { RefObject, useEffect, useMemo, useRef, useState } from "react";
import ScrollableFeed from "react-scrollable-feed";

import cn from "classnames";
import { TChatMessage } from "./types";
import { Box, Flex, Spinner } from "@chakra-ui/react";
import ChatMessage from "./ChatMessage";
import LoadingRow from "../../../LoadingRow";
import s from "./styles.module.scss";

type Props = {
  messages: any;
  isLoading: boolean;
};

const ChatList = ({ messages, isLoading }: Props) => {
  const topMostRef = useRef<HTMLDivElement | null>(null);
  const containerRef = useRef<HTMLDivElement | null>(null);
  const scrollableRef = useRef<ScrollableFeed | null>(null);

  const [isFetching, setIsFetching] = useState(false);

  const prevCountRef = useRef<number>(messages.length);

  const hasMore = useMemo(() => {
    return false;
  }, [messages.length]);

  useEffect(() => {
    return () => {
      // setMessages([]);
      // setisLoading(true);
    };
  }, []);

  useEffect(() => {
    if (prevCountRef.current < messages.length) {
      prevCountRef.current = messages.length;
    }
  }, [messages]);

  useEffect(() => {
    setTimeout(() => {
      scrollableRef.current?.scrollToBottom();
      // containerRef?.current?.scrollToBottom();
      // scrollableRef.current.scrollIntoView({ behavior: "smooth" });
    }, 2000);
  }, []);

  const isElementVisibleInContainer = (
    ele: RefObject<HTMLElement>,
    container: RefObject<HTMLElement>
  ) => {
    const rect = ele.current?.getBoundingClientRect();
    const containerRect = container.current?.getBoundingClientRect();

    if (!rect || !containerRect) {
      return false;
    }

    return (
      rect.top >= containerRect.top &&
      rect.left >= containerRect.left &&
      rect.bottom <= containerRect.bottom &&
      rect.right <= containerRect.right
    );
  };

  const renderBody = () => {
    // if (isLoading) {
    //   return (
    //     <div className={cn(s.MessageList, s.MessageList__Loading)}>
    //       <SkeletonMessage amount={15} />
    //     </div>
    //   );
    // }

    return (
      <ScrollableFeed
        className={s.scroll}
        ref={scrollableRef}
        onScroll={handleIsAtTop}
        // className="hahahamna"
        forceScroll={true}
      >
        {messages.map((message: TChatMessage, index: number) => {
          const isLast = index === messages.length - 1;
          if (index === 0) {
            return (
              <Box key={message.id} ref={topMostRef}>
                <ChatMessage message={message} isLast={isLast} />
              </Box>
            );
          }

          return (
            <ChatMessage key={message.id} message={message} isLast={isLast} />
          );
        })}
        {isLoading && (
          <Flex paddingLeft="12px">
            <LoadingRow />
          </Flex>
        )}
      </ScrollableFeed>
    );
  };

  const handleIsAtTop = () => {
    if (!hasMore) return;

    if (
      topMostRef &&
      isElementVisibleInContainer(topMostRef as any, containerRef as any)
    ) {
      setIsFetching(true);
    }

    setTimeout(() => {
      if (
        topMostRef &&
        isElementVisibleInContainer(topMostRef as any, containerRef as any)
      ) {
        topMostRef.current?.scrollIntoView(
          // add offset
          {
            block: "start",
            inline: "start",
          }
        );
        setIsFetching(false);
      }
    }, 1000);
  };

  return (
    <div
      className={cn(s.wrapper, {
        [s.empty_list as any]: messages.length === 0 && !isLoading,
      })}
      ref={containerRef}
    >
      {isFetching && (
        <Spinner
          size="md"
          position={"absolute"}
          left={"50%"}
          transform={"translateX(-50%)"}
        />
      )}

      {renderBody()}
    </div>
  );
};

export default ChatList;
