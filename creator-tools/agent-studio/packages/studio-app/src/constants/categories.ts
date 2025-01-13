import {
  StudioCategory,
  StudioCategoryOption,
  StudioCategoryType,
} from "@agent-studio/studio-dnd";
import { LegoComponentIcon } from "../components/icons/studio";
import { CATEGORY_KEYS } from "../constants/category-keys";

import { ChainId } from "./network";
import {
  OPTION_ETERNAL_AI_ID,
  OPTION_MODEL_HERMES_3_70B_KEY,
  OPTION_MODEL_INTELLECT_1_INSTECT_KEY,
  OPTION_MODEL_LLAMA_3_1_405B_INSTECT_KEY,
  OPTION_MODEL_LLAMA_3_3_70B_INSTECT_KEY,
} from "./option-values";
import { CATEGORY_OPTION_KEYS } from "./category-option-keys";

type State = "create" | "update";
const getAgentCategory = (state: State) => {
  const validators = state === "create" ? {} : {};
  return {
    idx: CATEGORY_KEYS.agent,
    title: "Agent",
    required: true,
    icon: LegoComponentIcon,
    multipleOption: false,
    disabled: state === "update",
    options: [
      {
        idx: CATEGORY_OPTION_KEYS.agent.agent_new,
        title: "New Agent",
        tooltip: "",
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
    ],
  } satisfies StudioCategory;
};

const getPersonalitiesCategory = (state: State) => {
  const validators = state === "create" ? {} : {};
  return {
    idx: CATEGORY_KEYS.personalities,
    title: "Personality",
    required: true,
    tooltip:
      "Create an agent for your NFT, Ordinals, token, —or start fresh with a new idea. This section defines your agent’s lore and backstory.",
    icon: LegoComponentIcon,
    color: "#12DAC2",
    disabled: state === "update",
    options: [
      {
        idx: CATEGORY_OPTION_KEYS.personalities.personality_customize,
        title: "New personality",
        tooltip: "",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_custom.svg",
        // customizeRenderOnBoard: NewPersonality as any,
      },
      {
        idx: CATEGORY_OPTION_KEYS.personalities.personality_nft,
        title: "Import from NFT",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_nft.svg",
        // customizeRenderOnBoard: ImportFromNft as any,
      },
      {
        idx: CATEGORY_OPTION_KEYS.personalities.personality_ordinals,
        title: "Import from Ordinals",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_ordinals.svg",
        // customizeRenderOnBoard: CustomImportFromOrdinalsRenderer as any,
      },
      {
        idx: CATEGORY_OPTION_KEYS.personalities.personality_token,
        title: "Import from Token",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_personality_token.svg",
        // customizeRenderOnBoard: CustomImportFromTokenRenderer as any,
      },
    ],

    ...validators,
  } satisfies StudioCategory;
};

const getAiFrameworkCategory = (state: State) => {
  const validators = state === "create" ? {} : {};
  return {
    idx: CATEGORY_KEYS.aiFrameworks,
    title: "AI Framework",
    required: true,
    tooltip:
      "Pick the blockchain where your agent lives. Each option has unique deployment fees, performance, and ongoing costs. Choose what best fits your needs.",
    icon: LegoComponentIcon,
    color: "#12DAC2",
    disabled: state === "update",
    options: [
      {
        idx: CATEGORY_OPTION_KEYS.aiFrameworks.ai_framework_eternal_ai,
        title: "Eternal AI",
        tooltip: "",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio/ic_eternal_ai.svg",
        data: {
          aiFrameworkId: {
            type: "hidden",
            label: "AI Framework Id",
            placeholder: "AI Framework Id",
            defaultValue: OPTION_ETERNAL_AI_ID,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.aiFrameworks.ai_framework_eliza,
        title: "Eliza",
        disabled: true,
      },
      {
        idx: CATEGORY_OPTION_KEYS.aiFrameworks.ai_framework_zerepy,
        title: "ZerePy",
        disabled: true,
      },
    ],
    ...validators,
  } satisfies StudioCategory;
};

const getBlockchainCategory = (state: State) => {
  const validators = state === "create" ? {} : {};
  return {
    idx: CATEGORY_KEYS.blockchains,
    title: "Blockchains",
    required: true,
    icon: LegoComponentIcon,
    color: "#368cdc",
    tooltip: "",
    disabled: state === "update",
    options: [
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_base,
        title: "Base",
        tooltip: "",
        icon: "/icons/blockchains/ic_base.svg",
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: "base",
          },
        },
        // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard as any,
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_arbitrum,
        title: "Arbitrum",
        icon: "/icons/blockchains/ic_arbitrum.svg",
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: "arbitrum",
          },
        },
        // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard as any,
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_bnb,
        title: "BNB",
        icon: "/icons/blockchains/ic-bsc.png",
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: "bsc",
          },
        },
        // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard as any,
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_bitcoin,
        title: "Bitcoin",
        icon: "/icons/blockchains/ic_bitcoin.svg",
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: "bitcoin",
          },
        },
        // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard as any,
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_bittensor,
        title: "Bittensor",
        icon: "/icons/blockchains/ic-tao.svg",
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: "tao",
          },
        },
        // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard as any,
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_ape_chain,
        title: "Ape Chain",
        icon: "/icons/blockchains/ic-ape.png",
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: "ape",
          },
        },
        // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard,
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_solana,
        title: "Solana",
        icon: "/icons/blockchains/ic_solana.svg",
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: "solana",
          },
        },
        // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard,
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_polygon,
        title: "Polygon",
        icon: "/icons/blockchains/ic_polygon.svg",
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: "polygon",
          },
        },
        // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard,
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_zksync_era,
        title: "ZKsync Era",
        icon: "/icons/blockchains/ic_zksync.svg",
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: "zksync",
          },
        },
        // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard,
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_avalanche,
        title: "Avalanche C-Chain",
        icon: "/icons/blockchains/ic_avax.svg",
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: "avax",
          },
        },
        // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard,
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_abstract_testnet,
        title: "Abstract Testnet",
        icon: "/icons/blockchains/ic-abs.png",
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: "abs",
          },
        },
        // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard,
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_duck_chain,
        title: "DuckChain",
        icon: "/icons/blockchains/ic-duck.svg",
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: "duck",
          },
        },
        // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard,
      },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_symbiosis,
        title: "Symbiosis",
        icon: "/icons/blockchains/ic_nbs.svg",
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: "symbiosis",
          },
        },
        // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard,
      },
    ],
    ...validators,
  } satisfies StudioCategory;
};

const getDecentralizedInferenceCategory = (state: State) => {
  const validators = state === "create" ? {} : {};
  return {
    idx: CATEGORY_KEYS.decentralized,
    title: "Decentralize Inference",
    required: true,
    tooltip:
      "Create an agent for your NFT, Ordinals, token, —or start fresh with a new idea. This section defines your agent’s lore and backstory.",
    icon: LegoComponentIcon,
    color: "#15C888",
    disabled: state === "update",
    options: [
      {
        idx: CATEGORY_OPTION_KEYS.decentralized
          .decentralize_inference_hermes_3_70b,
        title: "Hermes 3 70B",
        tooltip: "",
        data: {
          decentralizeId: {
            type: "hidden",
            label: "Decentralize Id",
            placeholder: "Decentralize Id",
            defaultValue: OPTION_MODEL_HERMES_3_70B_KEY,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.decentralized
          .decentralize_inference_intellect_1_10b,
        title: "INTELLECT 1 10B",
        tooltip: "",
        data: {
          decentralizeId: {
            type: "hidden",
            label: "Decentralize Id",
            placeholder: "Decentralize Id",
            defaultValue: OPTION_MODEL_INTELLECT_1_INSTECT_KEY,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.decentralized
          .decentralize_inference_llama_3_1_405b,
        title: "Llama 3.1 405B",
        tooltip: "",
        data: {
          decentralizeId: {
            type: "hidden",
            label: "Decentralize Id",
            placeholder: "Decentralize Id",
            defaultValue: OPTION_MODEL_LLAMA_3_1_405B_INSTECT_KEY,
          },
        },
      },
      {
        idx: CATEGORY_OPTION_KEYS.decentralized
          .decentralize_inference_llama_3_3_470b,
        title: "Llama 3.3 470B",
        tooltip: "",
        disabled: true,
        data: {
          decentralizeId: {
            type: "hidden",
            label: "Decentralize Id",
            placeholder: "Decentralize Id",
            defaultValue: OPTION_MODEL_LLAMA_3_3_70B_INSTECT_KEY,
          },
        },
      },
    ],
    ...validators,
  } satisfies StudioCategory;
};

const getTokenCategory = (state: State) => {
  const validators = state === "create" ? {} : {};
  return {
    idx: CATEGORY_KEYS.tokens,
    title: "Tokens",
    required: false,
    tooltip:
      "Select the blockchain to issue your agent’s token. Consider factors like accessibility, liquidity, and trading volume.",
    icon: LegoComponentIcon,
    color: "#A041FF",
    disabled: state === "update",
    options: [
      {
        idx: CATEGORY_OPTION_KEYS.tokens.token_base,
        title: "Token on Base",
        tooltip: "",
        icon: "/icons/blockchains/ic_base.svg",
        data: {
          tokenId: {
            type: "hidden",
            label: "Token Name",
            placeholder: "Token Name",
            defaultValue: ChainId.Base,
          },
        },
        // customizeRenderOnBoard: BaseTokenCustomizeOnBoard as any,
      },
      {
        idx: CATEGORY_OPTION_KEYS.tokens.token_solana,
        title: "Token on Solana",
        tooltip: "",
        icon: "/icons/blockchains/ic_solana.svg",
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
        tooltip: "",
        icon: "/icons/blockchains/ic_arbitrum.svg",
        data: {
          tokenId: {
            type: "hidden",
            label: "Token Name",
            placeholder: "Token Name",
            defaultValue: ChainId.Arbitrum,
          },
        },
        // customizeRenderOnBoard: BaseTokenCustomizeOnBoard as any,
      },
      {
        idx: CATEGORY_OPTION_KEYS.tokens.token_bnb,
        title: "Token on BNB",
        tooltip: "",
        icon: "/icons/blockchains/ic-bsc.png",
        data: {
          tokenId: {
            type: "hidden",
            label: "Token Name",
            placeholder: "Token Name",
            defaultValue: ChainId.BSC,
          },
        },
        // customizeRenderOnBoard: BaseTokenCustomizeOnBoard as any,
      },
      {
        idx: CATEGORY_OPTION_KEYS.tokens.token_ape,
        title: "Token on ApeChain",
        tooltip: "",
        icon: "/icons/blockchains/ic-ape.png",
        data: {
          tokenId: {
            type: "hidden",
            label: "Token Name",
            placeholder: "Token Name",
            defaultValue: ChainId.Ape,
          },
        },
        // customizeRenderOnBoard: BaseTokenCustomizeOnBoard as any,
      },
      {
        idx: CATEGORY_OPTION_KEYS.tokens.token_avax,
        title: "Token on Avalance C-Chain",
        tooltip: "",
        icon: "/icons/blockchains/ic_avax.svg",
        data: {
          tokenId: {
            type: "hidden",
            label: "Token Name",
            placeholder: "Token Name",
            defaultValue: ChainId.Avax,
          },
        },
        // customizeRenderOnBoard: BaseTokenCustomizeOnBoard as any,
      },
    ],
    ...validators,
  } satisfies StudioCategory;
};

const getMissionOnXCategory = (state: State) => {
  return {
    idx: CATEGORY_KEYS.missionOnX,
    title: "X",
    required: false,
    tooltip:
      "Set your agent’s activities on X. Choose what it does (e.g., posts, replies) and how often it does them.",
    icon: LegoComponentIcon,
    color: "#6B39F9",
    options: [
      {
        idx: CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post,
        title: "Post on X",
        tooltip: "",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_post.svg",
        // customizeRenderOnBoard: CustomPostOnXRenderer as any,
        type: StudioCategoryType.LINK,
        // ...POST_ON_X_HANDLER.create,
      },
      {
        idx: CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_reply,
        title: "Reply on X",
        tooltip: "",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_reply.svg",
        // customizeRenderOnBoard: CustomReplyOnXRenderer as any,
        type: StudioCategoryType.LINK,
        // ...REPLY_ON_X_HANDLER.create,
      },
      {
        idx: CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_engage,
        title: "Engage on X",
        tooltip: "",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_engage.svg",
        // customizeRenderOnBoard: CustomEngageOnXRenderer as any,
        type: StudioCategoryType.LINK,
        // ...ENGAGE_ON_X_HANDLER.create,
      },
      {
        idx: CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_follow,
        title: "Follow on X",
        tooltip: "",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_follow.svg",
        // customizeRenderOnBoard: CustomFollowOnXRenderer as any,
        type: StudioCategoryType.LINK,
        // ...FOLLOW_ON_X_HANDLER.create,
      },
    ],
  } satisfies StudioCategory;
};

const getMissionOnFarcasterCategory = (state: State) => {
  return {
    idx: CATEGORY_KEYS.missionOnFarcaster,
    title: "Farcaster",
    required: false,
    tooltip:
      "Set your agent’s activities on X. Choose what it does (e.g., posts, replies) and how often it does them.",
    icon: LegoComponentIcon,
    color: "#6B39F9",
    options: [
      {
        idx: CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_post,
        title: "Post on Farcaster",
        tooltip: "",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_post.svg",
        // customizeRenderOnBoard: CustomPostOnFarcasterRenderer as any,
        type: StudioCategoryType.LINK,
        // ...POST_ON_FARCASTER_HANDLER.create,
      },
      {
        idx: CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_reply,
        title: "Reply on Farcaster",
        tooltip: "",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_reply.svg",
        // customizeRenderOnBoard: CustomReplyOnFarcasterRenderer as any,
        type: StudioCategoryType.LINK,
        // ...REPLY_ON_FARCASTER_HANDLER.create,
      },
    ],
  } satisfies StudioCategory;
};

export default function getAgentModelCategories(
  state: State
): StudioCategory[] {
  const AGENT_MODEL_CATEGORIES: StudioCategory[] = [
    getAgentCategory(state),
    getPersonalitiesCategory(state),
    getAiFrameworkCategory(state),
    getBlockchainCategory(state),
    getDecentralizedInferenceCategory(state),
    getTokenCategory(state),
    getMissionOnXCategory(state),
    getMissionOnFarcasterCategory(state),
  ];
  return AGENT_MODEL_CATEGORIES;
}
