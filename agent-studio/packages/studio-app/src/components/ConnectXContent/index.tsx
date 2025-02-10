import {
  Collapse,
  Flex,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Image,
  Input,
  Spinner,
  Text,
} from "@chakra-ui/react";
import cn from "classnames";
import cx from "classnames";
import React, {
  forwardRef,
  useEffect,
  useImperativeHandle,
  useMemo,
  useState,
} from "react";

import copy from "copy-to-clipboard";
import toast from "react-hot-toast";
import s from "./styles.module.scss";
import { showError } from "../../utils/toast";
import { AI_AGENT_NAME_UPPERCASE } from "../../constants/agent";
import AgentAPI from "../../services/apis/agent";
import useCommonStore from "../../stores/useCommonStore";
import { PERP_API_URL, TWITTER_CLIENT_ID } from "../../configs";
import useAgentServiceStore from "../../stores/useAgentServiceStore";

type IProps = {
  agentId: string;
  onSkip?: () => void;
  onNext?: () => void;
  isShowSkipBtn?: boolean;
  url?: string;
  openType?: "_blank" | "_self";
  showOnlyConnectBtn?: boolean;
  className?: string;
  ref?: any;
};

const ConnectXContent: React.FC<IProps> = forwardRef((props: IProps, ref) => {
  const {
    agentId,
    onSkip,
    onNext,
    isShowSkipBtn = true,
    url = `/`,
    openType = "_blank",
    showOnlyConnectBtn = false,
    className,
  } = props;

  const baseDAgent = useAgentServiceStore((state) => state.baseDAgent);
  const walletAddress = baseDAgent.getSignerAddress();

  const [isLoading, setIsLoading] = useState(false);
  const [twClientId, setTwClientId] = useState<string | null>(null);
  const [twClientSecret, setTwClientSecret] = useState<string | null>(null);
  const [useAdvanceMode, setUseAdvanceMode] = useState(false);
  const [twitterUser, setTwitterUser] = useState<any>(null);
  const [error, setError] = useState({
    twitter_client_id: "",
    twitter_client_secret: "",
  });

  const callbackUrl = `${PERP_API_URL}`;

  const getTwitterOauthUrl = useMemo(() => {
    const rootUrl = "https://twitter.com/i/oauth2/authorize";
    const URL = `${url}&address=${
      walletAddress || ""
    }&agent_id=${agentId}&client_id=${
      useAdvanceMode ? twClientId : TWITTER_CLIENT_ID
    }`;
    const options = {
      redirect_uri: `${PERP_API_URL}?callback=${URL}`,
      client_id: useAdvanceMode && twClientId ? twClientId : TWITTER_CLIENT_ID,
      state: "state",
      response_type: "code",
      code_challenge: "challenge",
      code_challenge_method: "plain",
      scope: [
        "offline.access",
        "tweet.read",
        "tweet.write",
        "users.read",
        "follows.write",
        "like.write",
        "like.read",
        "users.read",
      ].join(" "),
    };
    const qs = new URLSearchParams(options).toString();
    return `${rootUrl}?${qs}`;
  }, [walletAddress, agentId]);

  useImperativeHandle(
    ref,
    () => ({
      callSubmitFunction: handleSubmit,
    }),
    [getTwitterOauthUrl]
  );

  const getAgentDetails = async () => {
    try {
      const res = await AgentAPI.getAgentDetail(agentId);
      if (res && !!res?.agent_info?.twitter_info?.twitter_id) {
        setIsLoading(false);
        setTwitterUser({
          id: res?.agent_info?.twitter_info?.twitter_id,
          username:
            res?.agent_info?.twitter_info?.twitter_username ||
            res?.agent_info?.twitter_info?.twitter_name,
          avatar: res?.agent_info?.twitter_info?.twitter_avatar,
        });
        useCommonStore.getState().requestReload();
      }
    } catch (error) {
      console.error("getAgentDetails error", error);
    }
  };

  const handleSubmit = async () => {
    if (!useAdvanceMode) {
      window.open(getTwitterOauthUrl, openType);
      setIsLoading(true);
      return;
    }

    if (!twClientId) {
      setError((prev) => ({
        ...prev,
        twitter_client_id: "Client ID is required",
      }));
    }
    if (!twClientSecret) {
      setError((prev) => ({
        ...prev,
        twitter_client_secret: "Client Secret is required",
      }));
    }

    if (!twClientId || !twClientSecret) return;

    try {
      const res = await AgentAPI.getVerifyXAccount({
        agent_id: agentId,
        twitter_client_id: twClientId,
        twitter_client_secret: twClientSecret,
      });

      if (res === true) {
        window.open(getTwitterOauthUrl, openType);
        setIsLoading(true);
      } else {
        showError({
          message: "Your Agent is not ready. Please try again later",
        });
      }
    } catch (error) {
      toast.error("Connect to X failed");
      console.error("handleSubmit error", error);
    }
  };

  useEffect(() => {
    if (!isLoading) return;
    // call getAgentDetails every 5 seconds, max retry 60 times

    let retry = 0;

    const interval = setInterval(() => {
      if (retry >= 30) {
        clearInterval(interval);
        setIsLoading(false);
        showError({
          message: "Connect to X failed. Please try again later",
        });
        return;
      }
      getAgentDetails();
      retry++;
    }, 5000);

    return () => clearInterval(interval);
  }, [isLoading]);

  useEffect(() => {
    if (twClientId) setError((prev) => ({ ...prev, twitter_client_id: "" }));
    if (twClientSecret)
      setError((prev) => ({ ...prev, twitter_client_secret: "" }));
  }, [twClientId, twClientSecret]);

  const styleInput = {
    fontSize: "15px",
    lineHeight: "24px !important",
    fontFamily: "var(--font-inter) !important",
    fontWeight: "400",
    border: "1px solid #E5E7EB !important",
    background: "#F8F9FA !important",
    borderRadius: "8px",
  };

  const styleLabel = {
    fontSize: "13px",
    lineHeight: "22px !important",
    fontFamily: "var(--font-inter) !important",
    fontWeight: "500",
    color: "black !important",
  };

  if (showOnlyConnectBtn) {
    return (
      <Flex className={cx(s.container, className)} w={"fit-content"}>
        {!!twitterUser?.id ? (
          <></>
        ) : (
          <button
            className={cn(s.ConnectXBtn)}
            // type="submit"
            onClick={handleSubmit}
            disabled={isLoading}
          >
            {isLoading ? <Spinner /> : "Connect"}
          </button>
        )}
      </Flex>
    );
  }

  return (
    <Flex className={s.container} direction="column">
      {/* <Flex alignItems={'center'} justifyContent={'space-between'} mb="12px"> */}
      {/* <Flex gap="12px" alignItems="center"> */}
      {/* <Image src="/svg/ic-twitter.svg" /> */}
      {/* <Text className={s.title}>Connect to X</Text> */}
      {/* </Flex> */}
      {/* <UserBox /> */}
      {/* </Flex> */}
      <Image
        src="/svg/ic-twitter-big.svg"
        alt={"ic-twitter-big.svg"}
        w={"48px"}
        mb={"24px"}
      />
      <Text fontSize={"18px"} fontWeight={"500"} mb="16px" color="#000">
        Set your {AI_AGENT_NAME_UPPERCASE} free on X
      </Text>
      <Text className={s.desc} mb="20px">
        Link your X account to enable autonomous posting and engagement. Your{" "}
        {AI_AGENT_NAME_UPPERCASE} will grow its presence on Twitter by itself,
        posting freely and interacting with humans - completely autonomous and
        self-sustaining. You don’t have to do anything.
      </Text>
      <Flex
        alignItems={"center"}
        gap="12px"
        display="none"
        className={cn(s.mode_wrapper, {
          [s.loading]: isLoading,
        })}
        flexWrap="wrap"
      >
        <div
          className={cn(s.mode_item, { [s.active]: !useAdvanceMode })}
          onClick={() => setUseAdvanceMode(false)}
        >
          <p>
            <span>Standard</span>
          </p>
          <b>100 engagements/day</b>
        </div>
        <div
          className={cn(s.mode_item, { [s.active]: useAdvanceMode })}
          onClick={() => setUseAdvanceMode(true)}
        >
          <p>
            <span>Advance</span>
          </p>
          <b>Higher limit using X's API</b>
        </div>
      </Flex>
      <Collapse in={useAdvanceMode} animateOpacity>
        <Flex flexDir={"column"} gap="12px" mt="24px">
          <div className={s.card}>
            <Text mb="12px" color="black" fontWeight={600}>
              <Text as="span" color="#8142FF">
                Step 1.
              </Text>{" "}
              Enable Read, Write, and Direct Message permissions, then paste
              this link into the App Info section on your{" "}
              <a
                href="https://developer.x.com/en/portal/dashboard"
                target="_blank"
                rel="noopener noreferrer"
              >
                X Development Portal.
              </a>
            </Text>
            <Text fontWeight={500}>Callback URI / Redirect URL</Text>
            <div
              className={s.copy_link_wrapper}
              onClick={() => {
                copy(`${callbackUrl}`);
                toast.remove();
                toast.success("Copied");
              }}
            >
              <Text fontWeight={400}>{callbackUrl}</Text>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="20"
                height="20"
                viewBox="0 0 20 20"
                fill="none"
              >
                <path
                  d="M16.375 8.26083V14.5025C16.375 15.7508 15.7508 16.375 14.5025 16.375H8.26083C7.0125 16.375 6.38833 15.7508 6.38833 14.5025V8.26083C6.38833 7.0125 7.0125 6.38833 8.26083 6.38833H14.5025C15.7508 6.38833 16.375 7.0125 16.375 8.26083ZM12.63 3.535C12.63 3.23968 12.3903 3 12.095 3H5.4075C3.85457 3 3 3.85457 3 5.4075V12.095C3 12.3903 3.23968 12.63 3.535 12.63C3.83032 12.63 4.07 12.3903 4.07 12.095V5.4075C4.07 4.45734 4.45734 4.07 5.4075 4.07H12.095C12.3903 4.07 12.63 3.83032 12.63 3.535Z"
                  fill="white"
                />
              </svg>
            </div>
          </div>

          <div className={s.card}>
            <Text mb="12px" fontWeight={600}>
              <Text as="span" color="#8142FF">
                Step 2.
              </Text>{" "}
              Provide the OAuth 2.0 Client ID and Client Secret.{" "}
            </Text>
            <FormControl
              mb="16px"
              isRequired
              isInvalid={!!error.twitter_client_id}
            >
              <FormLabel htmlFor="twitter_client_id" {...styleLabel}>
                Client ID
              </FormLabel>
              <Input
                id="twitter_client_id"
                name="twitter_client_id"
                variant="filled"
                placeholder="Enter Client ID"
                autoComplete="nope"
                onBlur={(e) => setTwClientId(e.target.value)}
                {...styleInput}
              />
              <FormErrorMessage>{error.twitter_client_id}</FormErrorMessage>
            </FormControl>
            <FormControl isRequired isInvalid={!!error.twitter_client_secret}>
              <FormLabel htmlFor="twitter_client_secret" {...styleLabel}>
                Client Secret
              </FormLabel>
              <Input
                id="twitter_client_secret"
                name="twitter_client_secret"
                variant="filled"
                placeholder="Enter Client Secret"
                autoComplete="nope"
                onBlur={(e) => setTwClientSecret(e.target.value)}
                {...styleInput}
              />
              <FormErrorMessage>{error.twitter_client_secret}</FormErrorMessage>
            </FormControl>
          </div>
        </Flex>
      </Collapse>

      <Flex gap="10px" mt="32px">
        {isShowSkipBtn && (
          <button
            onClick={() => {
              typeof onNext === "function" && onNext();
              // router.push(RoutePathManager.YOUR_AGENT);
            }}
            className={s.btn}
          >
            Skip
          </button>
        )}

        {!!twitterUser?.id ? (
          <button
            onClick={() => {
              typeof onSkip === "function" && onSkip();
            }}
            className={cn(s.btn, s.connectBtn)}
          >
            Done
          </button>
        ) : (
          <button
            className={cn(s.btn, s.connectBtn)}
            // type="submit"
            onClick={handleSubmit}
            disabled={isLoading}
          >
            {isLoading ? <Spinner /> : "Connect"}
          </button>
        )}

        {/* <a
            className={cn(s.btn, s.connectBtn)}
            href={getTwitterOauthUrl}
            target="_blank"
            rel="noopener noreferrer"
          >
            Connect
          </a> */}
      </Flex>
    </Flex>
  );
});

export default ConnectXContent;
