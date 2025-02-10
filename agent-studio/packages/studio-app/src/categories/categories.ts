import { CATEGORY_OPTION_KEYS } from "../constants/category-option-keys";
import { AI_FRAME_WORK_VALIDATES } from "./ai-framework/validates";
import CustomBlockchainRendererOnBoard from "./blockchains/CustomRenderer";
import { BLOCKCHAIN_VALIDATES } from "./blockchains/validates";
import { DECENTRALIZE_INFERENCE_VALIDATES } from "./decentralize-inference/validates";
import CustomPostOnFarcasterRenderer from "./farcaster/postOnFarcaster/CustomRenderer";
import POST_ON_FARCASTER_VALIDATES from "./farcaster/postOnFarcaster/validates";
import CustomReplyOnFarcasterRenderer from "./farcaster/replyOnFarcaster/CustomRenderer";
import REPLY_ON_FARCASTER_VALIDATES from "./farcaster/replyOnFarcaster/validates";
import NEW_AGENT_VALIDATES from "./new-agent/validates";
import ImportFromNft from "./personalities/import-from-nft/ImportFromNft";
import CustomImportFromOrdinalsRenderer from "./personalities/import-from-ordinals/CustomRenderer";
import CustomImportFromTokenRenderer from "./personalities/import-from-token/CustomRenderer";
import ImportGenomic from "./personalities/import-genomics/CustomRenderer";
import CustomKnowledgeRenderer from "./personalities/knowledge/CustomRenderer";
import NewPersonality from "./personalities/new-personality/NewPersonality";
import { PERSONALITIES_VALIDATES } from "./personalities/validates";
import BaseTokenCustomizeOnBoard from "./tokens/CustomRenderer";
import { TOKEN_VALIDATES } from "./tokens/validates";
import CustomEngageOnXRenderer from "./x/engageOnX/CustomRenderer";
import ENGAGE_ON_X_VALIDATES from "./x/engageOnX/validates";
import CustomFollowOnXRenderer from "./x/followOnX/CustomRenderer";
import FOLLOW_ON_X_VALIDATES from "./x/followOnX/validates";
import CustomPostOnXRenderer from "./x/postOnX/CustomRenderer";
import POST_ON_X_VALIDATES from "./x/postOnX/validates";
import CustomReplyOnXRenderer from "./x/replyOnX/CustomRenderer";
import REPLY_ON_X_VALIDATES from "./x/replyOnX/validates";
import CustomTokenRendererOnBoard from "./tokens/CustomRenderer";
import CustomPostNewsOnXRenderer from "./x/postNewsOnX/CustomRenderer";
import POST_NEWS_ON_X_VALIDATES from "./x/postNewsOnX/validates";
import KNOWLEDGE_VALIDATES from "./personalities/knowledge/validates";
import ElizaFramework from "./ai-framework/eliza/ElizaFramework";
import { LegoComponentIcon } from "../components/icons/studio";
import CustomTradingOnXRenderer from "./x/tradingOnX/CustomRenderer";
import CustomPostFollowingOnXRenderer from "./x/postFollowingOnX/CustomRenderer";
import POST_FOLLOWING_ON_X_VALIDATES from "./x/postFollowingOnX/validates";
import CustomTradeAnalyticsOnDefiRenderer from "./defi/tradeAnalytics/CustomRenderer";
import TRADE_ANALYTICS_ON_DEFI_VALIDATES from "./defi/tradeAnalytics/validates";
import {
  StudioCategory,
  StudioCategoryOption,
  StudioCategoryType,
} from "@agent-studio/studio-dnd";
import {
  AGENT_CATEGORY_KEY,
  AI_FRAMEWORK_CATEGORY_KEY,
  BLOCKCHAIN_CATEGORY_KEY,
  DECENTRALIZED_INFERENCE_CATEGORY_KEY,
  MISSION_ON_DEFI_CATEGORY_KEY,
  MISSION_ON_FARCASTER_CATEGORY_KEY,
  MISSION_ON_X_CATEGORY_KEY,
  PERSONALITY_CATEGORY_KEY,
  TOKEN_CATEGORY_KEY,
} from "../constants/category-keys";
import {
  OPTION_ELIZA_ID,
  OPTION_ETERNAL_AI_ID,
  OPTION_MODEL_HERMES_3_70B_KEY,
  OPTION_MODEL_INTELLECT_1_INSTECT_KEY,
  OPTION_MODEL_LLAMA_3_1_405B_INSTECT_KEY,
  OPTION_MODEL_LLAMA_3_3_70B_INSTECT_KEY,
  OPTION_ZERE_PY_ID,
} from "../constants/option-values";
import {
  ELIZA_CONFIG_DEFAULT,
  MISSION_DEFI_DEFAULT_VALUES,
  MISSION_FARCASTER_DEFAULT_VALUES,
  MISSION_X_DEFAULT_VALUES,
} from "../constants/default-values";
import { ChainId, SUPPORT_NETWORKS } from "../constants/networks";
import { compareString } from "../utils/string";
import { categoryImageIcon } from "../components/icons/common";
import CustomRendererNoInput from "./shared/CustomRendererNoInput";
import CustomAIFrameworkRendererOnBoard from "./ai-framework/CustomRenderer";
import CustomDIRendererOnBoard from "./decentralize-inference/CustomRenderer";
import { tokens } from "../constants/tokens";

type State = "create" | "update";

const getAgentCategory = (
  state: State,
  extendOptions: StudioCategoryOption[] = []
): StudioCategory => {
  const validators =
    state === "create"
      ? NEW_AGENT_VALIDATES.create
      : NEW_AGENT_VALIDATES.update;

  return {
    idx: AGENT_CATEGORY_KEY,
    title: "Agent",
    required: true,
    icon: LegoComponentIcon,
    multipleOption: false,
    disabled: state === "update",
    // hidden: state === 'update',
    order: 0,
    options: [
      {
        zIndex: 10,
        idx: CATEGORY_OPTION_KEYS.agent.agent_new,
        title: state === "create" ? "New Agent" : "Agent Name",
        data: {
          agentName: {
            type: "text",
            label: "Agent Name",
            placeholder: "Agent Name",
            defaultValue: "",
            disabled: state === "update",
          },
        },
        ...validators,
      } satisfies StudioCategoryOption,
      ...extendOptions,
    ],
  } satisfies StudioCategory;
};

const getPersonalitiesCategory = (
  state: State,
  extendOptions: StudioCategoryOption[] = []
): StudioCategory => {
  const validators =
    state === "create"
      ? PERSONALITIES_VALIDATES.create
      : PERSONALITIES_VALIDATES.update;

  return {
    idx: PERSONALITY_CATEGORY_KEY,
    title: "Personality",
    required: true,
    tooltip:
      "Create an agent for your NFT, Ordinals, token, —or start fresh with a new idea. This section defines your agent’s lore and backstory.",
    icon: LegoComponentIcon,
    color: "#12DAC2",
    order: 1,
    options: [
      {
        idx: CATEGORY_OPTION_KEYS.personalities.personality_customize,
        title: "New personality",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_custom.svg",
        customizeRenderOnBoard: NewPersonality as any,
      },
      {
        idx: CATEGORY_OPTION_KEYS.personalities.personality_nft,
        title: "Import from NFT",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_nft.svg",
        customizeRenderOnBoard: ImportFromNft as any,
        disabled: state === "update",
        // hidden: state === 'update',
      },
      {
        idx: CATEGORY_OPTION_KEYS.personalities.personality_ordinals,
        title: "Import from Ordinals",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_ordinals.svg",
        customizeRenderOnBoard: CustomImportFromOrdinalsRenderer as any,
        disabled: state === "update",
        // hidden: state === 'update',
      },
      {
        idx: CATEGORY_OPTION_KEYS.personalities.personality_token,
        title: "Import from Token",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_token.svg",
        customizeRenderOnBoard: CustomImportFromTokenRenderer as any,
        disabled: state === "update",
        // hidden: state === 'update',
      },
      {
        idx: CATEGORY_OPTION_KEYS.personalities.personality_knowledge,
        title: "Knowledge",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_knowledge.svg",
        customizeRenderOnBoard: CustomKnowledgeRenderer as any,
        disabled: state === "update",
        // hidden: state === 'update',
        ...KNOWLEDGE_VALIDATES.create,
      },
      {
        idx: CATEGORY_OPTION_KEYS.personalities.personality_genomics,
        title: "Genomic Labs",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_genomic.svg",
        customizeRenderOnBoard: ImportGenomic as any,
        disabled: state === "update",
        // hidden: state === 'update',
        // disabled: true,
      },
      ...extendOptions,
    ],

    ...validators,
  } satisfies StudioCategory;
};

const getAiFrameworkCategory = (
  state: State,
  extendOptions: StudioCategoryOption[] = []
): StudioCategory => {
  const validators =
    state === "create"
      ? AI_FRAME_WORK_VALIDATES.create
      : AI_FRAME_WORK_VALIDATES.update;

  const category = {
    idx: AI_FRAMEWORK_CATEGORY_KEY,
    title: "AI Framework",
    required: true,
    tooltip:
      "Pick the blockchain where your agent lives. Each option has unique deployment fees, performance, and ongoing costs. Choose what best fits your needs.",
    icon: LegoComponentIcon,
    color: "#5B913B",
    disabled: state === "update",
    // hidden: state === 'update',
    order: state === "update" ? 1000 : 2,
    options: [
      {
        idx: CATEGORY_OPTION_KEYS.aiFrameworks.ai_framework_eternal_ai,
        title: "Eternal AI",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio/ic_eternal_ai.svg",
        customizeRenderOnBoard: CustomAIFrameworkRendererOnBoard as any,
        data: {
          aiFrameworkId: {
            type: "hidden",
            label: "AI Framework Id",
            placeholder: "AI Framework Id",
            defaultValue: OPTION_ETERNAL_AI_ID,
          },
        },
      } satisfies StudioCategoryOption,
      {
        idx: CATEGORY_OPTION_KEYS.aiFrameworks.ai_framework_eliza,
        title: "Eliza",
        customizeRenderOnBoard: ElizaFramework as any,
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio/ic_eliza_ai.png",
        data: {
          aiFrameworkId: {
            type: "hidden",
            label: "AI Eliaz Id",
            placeholder: "AI Eliaz Id",
            defaultValue: OPTION_ELIZA_ID,
          },
          config: {
            type: "hidden",
            label: "Config",
            placeholder: "Config",
            defaultValue: JSON.stringify(ELIZA_CONFIG_DEFAULT, null, 2),
          },
        },
      } satisfies StudioCategoryOption,
      {
        idx: CATEGORY_OPTION_KEYS.aiFrameworks.ai_framework_zerepy,
        title: "ZerePy",
        disabled: true,
        customizeRenderOnBoard: CustomAIFrameworkRendererOnBoard as any,
        data: {
          aiFrameworkId: {
            type: "hidden",
            label: "AI ZerePy Id",
            placeholder: "AI ZerePy Id",
            defaultValue: OPTION_ZERE_PY_ID,
          },
        },
      } satisfies StudioCategoryOption,
      ...extendOptions,
    ] as StudioCategoryOption[],
    ...validators,
  } satisfies StudioCategory;

  return category;
};

const getBlockchainCategory = (
  state: State,
  extendOptions: StudioCategoryOption[] = []
): StudioCategory => {
  const validators =
    state === "create"
      ? BLOCKCHAIN_VALIDATES.create
      : BLOCKCHAIN_VALIDATES.update;

  const getRequiredAmount = (chainId: string) => {
    const matchedToken = tokens.find((v) => compareString(v.id, chainId));
    if (matchedToken) {
      return `${matchedToken.required_amount} EAI/post`;
    }
    return "";
  };

  const category = {
    idx: BLOCKCHAIN_CATEGORY_KEY,
    title: "Blockchains",
    required: true,
    icon: LegoComponentIcon,
    color: "#368cdc",
    disabled: state === "update",
    // hidden: state === 'update',
    order: state === "update" ? 1000 : 3,
    options: [
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_base,
        title: "Base",
        icon: "/icons/blockchains/ic_base.svg",
        tooltip: getRequiredAmount(SUPPORT_NETWORKS.BASE),
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: SUPPORT_NETWORKS.BASE,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_arbitrum,
        title: "Arbitrum",
        icon: "/icons/blockchains/ic_arbitrum.svg",
        tooltip: getRequiredAmount(SUPPORT_NETWORKS.ARBITRUM),
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: SUPPORT_NETWORKS.ARBITRUM,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_bnb,
        title: "BNB",
        icon: "/icons/blockchains/ic-bsc.png",
        tooltip: getRequiredAmount(SUPPORT_NETWORKS.BSC),
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: SUPPORT_NETWORKS.BSC,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_eth,
        title: "Ethereum",
        icon: "https://eternalai.org/icons/blockchains/ic-ether.png",
        tooltip: getRequiredAmount(SUPPORT_NETWORKS.ETHEREUM),
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: SUPPORT_NETWORKS.ETHEREUM,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_bitcoin,
        title: "Bitcoin",
        icon: "/icons/blockchains/ic_bitcoin.svg",
        tooltip: getRequiredAmount(SUPPORT_NETWORKS.BITCOIN),
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: SUPPORT_NETWORKS.BITCOIN,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_bittensor,
        title: "Bittensor",
        icon: "/icons/blockchains/ic-tao.svg",
        tooltip: getRequiredAmount(SUPPORT_NETWORKS.TAO),
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: SUPPORT_NETWORKS.TAO,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_ape_chain,
        title: "Ape Chain",
        icon: "/icons/blockchains/ic-ape.png",
        tooltip: getRequiredAmount(SUPPORT_NETWORKS.APE),
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: SUPPORT_NETWORKS.APE,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_solana,
        title: "Solana",
        icon: "/icons/blockchains/ic_solana.svg",
        tooltip: getRequiredAmount(SUPPORT_NETWORKS.SOLANA),
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: SUPPORT_NETWORKS.SOLANA,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_polygon,
        title: "Polygon",
        icon: "/icons/blockchains/ic_polygon.svg",
        tooltip: getRequiredAmount(SUPPORT_NETWORKS.POLYGON),
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: SUPPORT_NETWORKS.POLYGON,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_zksync_era,
        title: "ZKsync Era",
        icon: "/icons/blockchains/ic_zksync.svg",
        tooltip: getRequiredAmount(SUPPORT_NETWORKS.ZKSYNC),
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: SUPPORT_NETWORKS.ZKSYNC,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_avalanche,
        title: "Avalanche C-Chain",
        icon: "/icons/blockchains/ic_avax.svg",
        tooltip: getRequiredAmount(SUPPORT_NETWORKS.AVAX),
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: SUPPORT_NETWORKS.AVAX,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_abstract_testnet,
        title: "Abstract Testnet",
        icon: "/icons/blockchains/ic-abs.png",
        tooltip: getRequiredAmount(SUPPORT_NETWORKS.ABS),
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: SUPPORT_NETWORKS.ABS,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_duck_chain,
        title: "DuckChain",
        icon: "/icons/blockchains/ic-duck.svg",
        tooltip: getRequiredAmount(SUPPORT_NETWORKS.DUCK),
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: SUPPORT_NETWORKS.DUCK,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_symbiosis,
        title: "Symbiosis",
        icon: "/icons/blockchains/ic_nbs.svg",
        tooltip: getRequiredAmount(SUPPORT_NETWORKS.SYMBIOSIS),
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: SUPPORT_NETWORKS.SYMBIOSIS,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_tron,
        title: "Tron",
        icon: "https://eternalai.org/icons/blockchains/ic-tron.svg",
        tooltip: getRequiredAmount(SUPPORT_NETWORKS.TRON),
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: SUPPORT_NETWORKS.TRON,
          },
        },
      },
      ...extendOptions,
    ],
    ...validators,
  } satisfies StudioCategory;

  category.options = category.options.map((option) => ({
    ...option,
    customizeRenderOnBoard: CustomBlockchainRendererOnBoard as any,
  }));

  return category;
};

const getDecentralizedInferenceCategory = (
  state: State,
  extendOptions: StudioCategoryOption[] = []
): StudioCategory => {
  const validators =
    state === "create"
      ? DECENTRALIZE_INFERENCE_VALIDATES.create
      : DECENTRALIZE_INFERENCE_VALIDATES.update;

  const category = {
    idx: DECENTRALIZED_INFERENCE_CATEGORY_KEY,
    title: "Decentralize Inference",
    required: true,
    tooltip:
      "Create an agent for your NFT, Ordinals, token, —or start fresh with a new idea. This section defines your agent’s lore and backstory.",
    icon: LegoComponentIcon,
    color: "#15C888",
    disabled: state === "update",
    // hidden: state === 'update',
    order: state === "update" ? 1000 : 4,
    options: [...extendOptions],
    ...validators,
  } satisfies StudioCategory;

  category.options = category.options.map((option) => ({
    ...option,
    customizeRenderOnBoard: CustomDIRendererOnBoard as any,
  }));

  return category;
};

const getTokenCategory = (
  state: State,
  extendOptions: StudioCategoryOption[] = []
): StudioCategory => {
  const validators =
    state === "create" ? TOKEN_VALIDATES.create : TOKEN_VALIDATES.update;

  const getTokenPrice = (chainId: string) => {
    const matchedToken = tokens.find((v) => compareString(v.id, chainId));
    if (matchedToken) {
      return `${matchedToken.deploy_token_amount} EAI`;
    }
    return "Free";
  };
  const category = {
    idx: TOKEN_CATEGORY_KEY,
    title: "Tokens",
    required: true,
    tooltip:
      "Select the blockchain to issue your agent’s token. Consider factors like accessibility, liquidity, and trading volume.",
    icon: LegoComponentIcon,
    color: "#A041FF",
    disabled: state === "update",
    // hidden: state === 'update',
    order: state === "update" ? 1000 : 5,
    options: [
      {
        idx: CATEGORY_OPTION_KEYS.tokens.token_base,
        title: "Token on Base",
        icon: "/icons/blockchains/ic_base.svg",
        tooltip: getTokenPrice(SUPPORT_NETWORKS.BASE),
        data: {
          tokenId: {
            type: "hidden",
            label: "Token Name",
            placeholder: "Token Name",
            defaultValue: ChainId.Base,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.tokens.token_solana,
        title: "Token on Solana",
        icon: "/icons/blockchains/ic_solana.svg",
        tooltip: getTokenPrice(SUPPORT_NETWORKS.SOLANA),
        data: {
          tokenId: {
            type: "hidden",
            label: "Token Name",
            placeholder: "Token Name",
            defaultValue: ChainId.Solana,
          },
        },
        // customizeRenderOnBoard: BaseTokenCustomizeOnBoard as any,
      },
      {
        idx: CATEGORY_OPTION_KEYS.tokens.token_arbitrum,
        title: "Token on Arbitrum",
        icon: "/icons/blockchains/ic_arbitrum.svg",
        tooltip: getTokenPrice(SUPPORT_NETWORKS.ARBITRUM),
        data: {
          tokenId: {
            type: "hidden",
            label: "Token Name",
            placeholder: "Token Name",
            defaultValue: ChainId.Arbitrum,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.tokens.token_bnb,
        title: "Token on BNB",
        icon: "/icons/blockchains/ic-bsc.png",
        tooltip: getTokenPrice(SUPPORT_NETWORKS.BSC),
        data: {
          tokenId: {
            type: "hidden",
            label: "Token Name",
            placeholder: "Token Name",
            defaultValue: ChainId.BSC,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.tokens.token_ape,
        title: "Token on ApeChain",
        icon: "/icons/blockchains/ic-ape.png",
        tooltip: getTokenPrice(SUPPORT_NETWORKS.APE),
        data: {
          tokenId: {
            type: "hidden",
            label: "Token Name",
            placeholder: "Token Name",
            defaultValue: ChainId.Ape,
          },
        },
        customizeRenderOnBoard: BaseTokenCustomizeOnBoard as any,
      },
      {
        idx: CATEGORY_OPTION_KEYS.tokens.token_avax,
        title: "Token on Avalance C-Chain",
        icon: "/icons/blockchains/ic_avax.svg",
        tooltip: getTokenPrice(SUPPORT_NETWORKS.AVAX),
        data: {
          tokenId: {
            type: "hidden",
            label: "Token Name",
            placeholder: "Token Name",
            defaultValue: ChainId.Avax,
          },
        },
      },

      ...extendOptions,
    ],
    ...validators,
  } satisfies StudioCategory;

  category.options = category.options.map((option) => ({
    ...option,
    customizeRenderOnBoard: CustomTokenRendererOnBoard as any,
  }));

  return category;
};

const getMissionOnXCategory = (
  state: State,
  extendOptions: StudioCategoryOption[] = []
): StudioCategory => {
  return {
    idx: MISSION_ON_X_CATEGORY_KEY,
    title: "X",
    required: false,
    tooltip:
      "Set your agent’s activities on X. Choose what it does (e.g., posts, replies) and how often it does them.",
    icon: categoryImageIcon("https://eternalai.org/icons/blockchains/ic_x.svg"),
    color: "#6B39F9",
    order: 6,
    options: [
      {
        idx: CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post,
        title: "Post (using CoT)",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_post.svg",
        customizeRenderOnBoard: CustomPostOnXRenderer as any,
        type: StudioCategoryType.LINK,
        tooltip: "Cost: 1 EAI per execution.",
        data: {
          toolset: {
            type: "hidden",
            label: "Tool set",
            placeholder: "Tool set",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post
              ].tool_set,
          },
          frequency: {
            type: "hidden",
            label: "Frequency",
            placeholder: "Frequency",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post
              ].time,
          },
          details: {
            type: "hidden",
            label: "Details",
            placeholder: "Details",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post
              ].description,
          },
        },
        ...POST_ON_X_VALIDATES.create,
      },
      {
        idx: CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_news,
        title: "Post (using Bing / X search) on X",
        tooltip: "",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_post.svg",
        customizeRenderOnBoard: CustomPostNewsOnXRenderer as any,
        type: StudioCategoryType.LINK,
        data: {
          toolset: {
            type: "hidden",
            label: "Tool set",
            placeholder: "Tool set",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_news
              ].tool_set,
          },
          frequency: {
            type: "hidden",
            label: "Frequency",
            placeholder: "Frequency",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_news
              ].time,
          },
          details: {
            type: "hidden",
            label: "Details",
            placeholder: "Details",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_news
              ].description,
          },
        },
        ...POST_NEWS_ON_X_VALIDATES.create,
      },
      {
        idx: CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_following,
        title: "Post (using the content from followings) on X",
        tooltip: "",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_post.svg",
        customizeRenderOnBoard: CustomPostFollowingOnXRenderer as any,
        type: StudioCategoryType.LINK,
        ...POST_FOLLOWING_ON_X_VALIDATES.create,
      },
      {
        idx: CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_reply,
        title: "Reply",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_reply.svg",
        customizeRenderOnBoard: CustomReplyOnXRenderer as any,
        type: StudioCategoryType.LINK,
        tooltip: "Cost: 1 EAI per execution.",
        data: {
          toolset: {
            type: "hidden",
            label: "Tool set",
            placeholder: "Tool set",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_reply
              ].tool_set,
          },
          frequency: {
            type: "hidden",
            label: "Frequency",
            placeholder: "Frequency",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_reply
              ].time,
          },
          details: {
            type: "hidden",
            label: "Details",
            placeholder: "Details",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_reply
              ].description,
          },
        },
        ...REPLY_ON_X_VALIDATES.create,
      },
      {
        idx: CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_engage,
        title: "Engage",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_engage.svg",
        customizeRenderOnBoard: CustomEngageOnXRenderer as any,
        type: StudioCategoryType.LINK,
        tooltip: "Cost: 1 EAI per execution.",
        data: {
          toolset: {
            type: "hidden",
            label: "Tool set",
            placeholder: "Tool set",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_engage
              ].tool_set,
          },
          frequency: {
            type: "hidden",
            label: "Frequency",
            placeholder: "Frequency",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_engage
              ].time,
          },
          details: {
            type: "hidden",
            label: "Details",
            placeholder: "Details",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_engage
              ].description,
          },
        },
        ...ENGAGE_ON_X_VALIDATES.create,
      },
      {
        idx: CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_follow,
        title: "Follow",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_follow.svg",
        customizeRenderOnBoard: CustomFollowOnXRenderer as any,
        type: StudioCategoryType.LINK,
        tooltip: "Cost: 1 EAI per execution.",
        data: {
          toolset: {
            type: "hidden",
            label: "Tool set",
            placeholder: "Tool set",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_follow
              ].tool_set,
          },
          frequency: {
            type: "hidden",
            label: "Frequency",
            placeholder: "Frequency",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_follow
              ].time,
          },
          details: {
            type: "hidden",
            label: "Details",
            placeholder: "Details",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_follow
              ].description,
          },
        },
        ...FOLLOW_ON_X_VALIDATES.create,
      },
      {
        idx: CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_trading,
        title: "Trading",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio/ic_abilities_news.svg",
        customizeRenderOnBoard: CustomTradingOnXRenderer as any,
        type: StudioCategoryType.LINK,
        tooltip: "Cost: 1 EAI per execution.",
        data: {
          toolset: {
            type: "hidden",
            label: "Tool set",
            placeholder: "Tool set",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_trading
              ].tool_set,
          },
          frequency: {
            type: "hidden",
            label: "Frequency",
            placeholder: "Frequency",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_trading
              ].time,
          },
          details: {
            type: "hidden",
            label: "Details",
            placeholder: "Details",
            defaultValue:
              MISSION_X_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_trading
              ].description,
          },
        },
        ...FOLLOW_ON_X_VALIDATES.create,
      },
      ...extendOptions,
    ],
  } satisfies StudioCategory;
};

const getMissionOnFarcasterCategory = (
  state: State,
  extendOptions: StudioCategoryOption[] = []
): StudioCategory => {
  return {
    idx: MISSION_ON_FARCASTER_CATEGORY_KEY,
    title: "Farcaster",
    required: false,
    tooltip:
      "Set your agent’s activities on X. Choose what it does (e.g., posts, replies) and how often it does them.",
    icon: categoryImageIcon(
      "https://eternalai.org/icons/blockchains/ic_farcaster.svg"
    ),
    color: "#344CB7",
    order: 7,
    options: [
      {
        idx: CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_post,
        title: "Post",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_post.svg",
        customizeRenderOnBoard: CustomPostOnFarcasterRenderer as any,
        type: StudioCategoryType.LINK,
        tooltip: "Cost: 1 EAI per execution.",
        data: {
          toolset: {
            type: "hidden",
            label: "Tool set",
            placeholder: "Tool set",
            defaultValue:
              MISSION_FARCASTER_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnFarcaster
                  .mission_on_farcaster_post
              ].tool_set,
          },
          frequency: {
            type: "hidden",
            label: "Frequency",
            placeholder: "Frequency",
            defaultValue:
              MISSION_FARCASTER_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnFarcaster
                  .mission_on_farcaster_post
              ].time,
          },
          details: {
            type: "hidden",
            label: "Details",
            placeholder: "Details",
            defaultValue:
              MISSION_FARCASTER_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnFarcaster
                  .mission_on_farcaster_post
              ].description,
          },
        },
        ...POST_ON_FARCASTER_VALIDATES.create,
      },
      {
        idx: CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_reply,
        title: "Reply",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_reply.svg",
        customizeRenderOnBoard: CustomReplyOnFarcasterRenderer as any,
        type: StudioCategoryType.LINK,
        tooltip: "Cost: 1 EAI per execution.",
        data: {
          toolset: {
            type: "hidden",
            label: "Tool set",
            placeholder: "Tool set",
            defaultValue:
              MISSION_FARCASTER_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnFarcaster
                  .mission_on_farcaster_reply
              ].tool_set,
          },
          frequency: {
            type: "hidden",
            label: "Frequency",
            placeholder: "Frequency",
            defaultValue:
              MISSION_FARCASTER_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnFarcaster
                  .mission_on_farcaster_reply
              ].time,
          },
          details: {
            type: "hidden",
            label: "Details",
            placeholder: "Details",
            defaultValue:
              MISSION_FARCASTER_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnFarcaster
                  .mission_on_farcaster_reply
              ].description,
          },
        },
        ...REPLY_ON_FARCASTER_VALIDATES.create,
      },
      ...extendOptions,
    ],
  } satisfies StudioCategory;
};

const getMissionOnDefiCategory = (
  state: State,
  extendOptions: StudioCategoryOption[] = []
): StudioCategory => {
  return {
    idx: MISSION_ON_DEFI_CATEGORY_KEY,
    title: "Defi",
    required: false,
    tooltip:
      "Set your agent’s activities on Defi. Choose what it does (e.g., posts, replies) and how often it does them.",
    icon: categoryImageIcon(
      "https://eternalai.org/icons/blockchains/ic-defi.png"
    ),
    color: "#37AFE1",
    order: 8,
    options: [
      {
        idx: CATEGORY_OPTION_KEYS.missionOnDefi.mission_on_defi_trade_analytics,
        title: "Trade Analytics",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio/ic_abilities_news.svg",
        customizeRenderOnBoard: CustomTradeAnalyticsOnDefiRenderer as any,
        type: StudioCategoryType.LINK,
        tooltip: "Cost: 1 EAI per execution.",
        data: {
          toolset: {
            type: "hidden",
            label: "Tool set",
            placeholder: "Tool set",
            defaultValue:
              MISSION_DEFI_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnDefi
                  .mission_on_defi_trade_analytics
              ].tool_set,
          },
          frequency: {
            type: "hidden",
            label: "Frequency",
            placeholder: "Frequency",
            defaultValue:
              MISSION_FARCASTER_DEFAULT_VALUES[
                CATEGORY_OPTION_KEYS.missionOnFarcaster
                  .mission_on_farcaster_reply
              ].time,
          },
        },
        ...TRADE_ANALYTICS_ON_DEFI_VALIDATES.create,
      },
      ...extendOptions,
    ],
  } satisfies StudioCategory;
};

const processOption = (
  category: StudioCategory,
  option: StudioCategoryOption
): StudioCategoryOption => {
  return {
    ...option,
    // customizeRenderOnBoard: onlyLabelCategory.includes(category.idx) ? CustomRendererNoInput as any : CustomRendererOnBoard as any,
    customizeRenderOnSideBar: CustomRendererNoInput as any,
  };
};

export default function getAgentModelCategories(
  state: State,
  extendOptions: Record<string, StudioCategoryOption[]> = {}
): StudioCategory[] {
  const AGENT_MODEL_CATEGORIES: StudioCategory[] = [
    getAgentCategory(state, extendOptions[AGENT_CATEGORY_KEY]),
    getPersonalitiesCategory(state, extendOptions[PERSONALITY_CATEGORY_KEY]),
    getAiFrameworkCategory(state, extendOptions[AI_FRAMEWORK_CATEGORY_KEY]),
    getBlockchainCategory(state, extendOptions[BLOCKCHAIN_CATEGORY_KEY]),
    getDecentralizedInferenceCategory(
      state,
      extendOptions[DECENTRALIZED_INFERENCE_CATEGORY_KEY]
    ),
    getTokenCategory(state, extendOptions[TOKEN_CATEGORY_KEY]),
    getMissionOnXCategory(state, extendOptions[MISSION_ON_X_CATEGORY_KEY]),
    getMissionOnFarcasterCategory(
      state,
      extendOptions[MISSION_ON_FARCASTER_CATEGORY_KEY]
    ),
    getMissionOnDefiCategory(
      state,
      extendOptions[MISSION_ON_DEFI_CATEGORY_KEY]
    ),
  ];

  return AGENT_MODEL_CATEGORIES.map((category) => ({
    ...category,
    options: category.options.map((option) => processOption(category, option)),
  }));
}
