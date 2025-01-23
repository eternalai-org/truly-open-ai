import {
  StudioCategory,
  StudioCategoryOption,
  StudioCategoryOptionRenderPayload,
  StudioCategoryType,
} from "@agent-studio/studio-dnd";
import { LegoComponentIcon } from "../components/icons/studio";
import { CATEGORY_KEYS } from "../constants/category-keys";

import {
  OPTION_ETERNAL_AI_ID,
  OPTION_MODEL_HERMES_3_70B_KEY,
  OPTION_MODEL_INTELLECT_1_INSTECT_KEY,
  OPTION_MODEL_LLAMA_3_1_405B_INSTECT_KEY,
  OPTION_MODEL_LLAMA_3_3_70B_INSTECT_KEY,
} from "../constants/option-values";
import { CATEGORY_OPTION_KEYS } from "../constants/category-option-keys";
import { AgentChainId, AgentTokenChainId } from "@eternal-dagent/core";
import NEW_AGENT_VALIDATES from "./new-agent/validates";
import { PERSONALITIES_VALIDATES } from "./personalities/validates";
import { AI_FRAME_WORK_VALIDATES } from "./ai-framework/validates";
import { BLOCKCHAIN_VALIDATES } from "./blockchains/validates";
import { DECENTRALIZE_INFERENCE_VALIDATES } from "./decentralize-inference/validates";
import { TOKEN_VALIDATES } from "./tokens/validates";
import NewPersonality from "./personalities/renders/onflow/new-personality/NewPersonality";
import { JSX, ReactNode } from "react";
import X_VALIDATES from "./x/validates";
import CustomBaseXForm from "./x/renders/onflow/custom-base/CustomBaseXForm";
import CustomBaseFarcasterForm from "./farcaster/renders/onflow/custom-base/CustomBaseFarcasterForm";

type State = "create" | "update";
const getAgentCategory = (state: State) => {
  const validators =
    state === "create"
      ? NEW_AGENT_VALIDATES.create
      : NEW_AGENT_VALIDATES.update;
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

type CustomizeComponentType = <T>(
  data: StudioCategoryOptionRenderPayload<T>
) => ReactNode | JSX.Element;

const getPersonalitiesCategory = (state: State) => {
  const validators =
    state === "create"
      ? PERSONALITIES_VALIDATES.create
      : PERSONALITIES_VALIDATES.update;
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
        customizeRenderOnBoard: NewPersonality as CustomizeComponentType,
      },
    ],

    ...validators,
  } satisfies StudioCategory;
};

const getAiFrameworkCategory = (state: State) => {
  const validators =
    state === "create"
      ? AI_FRAME_WORK_VALIDATES.create
      : AI_FRAME_WORK_VALIDATES.update;
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
  const validators =
    state === "create"
      ? BLOCKCHAIN_VALIDATES.create
      : BLOCKCHAIN_VALIDATES.update;
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
            defaultValue: AgentChainId.Base,
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
            defaultValue: AgentChainId.Arbitrum,
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
            defaultValue: AgentChainId.BSC,
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
            defaultValue: AgentChainId.Bitcoin,
          },
        },
        // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard as any,
      },
      // {
      //   idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_bittensor,
      //   title: "Bittensor",
      //   icon: "/icons/blockchains/ic-tao.svg",
      //   data: {
      //     chainId: {
      //       type: "hidden",
      //       label: "Chain Id",
      //       placeholder: "Chain Id",
      //       defaultValue: "tao",
      //     },
      //   },
      //   // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard as any,
      // },
      // {
      //   idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_ape_chain,
      //   title: "Ape Chain",
      //   icon: "/icons/blockchains/ic-ape.png",
      //   data: {
      //     chainId: {
      //       type: "hidden",
      //       label: "Chain Id",
      //       placeholder: "Chain Id",
      //       defaultValue: "ape",
      //     },
      //   },
      //   // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard,
      // },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_solana,
        title: "Solana",
        icon: "/icons/blockchains/ic_solana.svg",
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: AgentChainId.Solana,
          },
        },
        // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard,
      },
      // {
      //   idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_polygon,
      //   title: "Polygon",
      //   icon: "/icons/blockchains/ic_polygon.svg",
      //   data: {
      //     chainId: {
      //       type: "hidden",
      //       label: "Chain Id",
      //       placeholder: "Chain Id",
      //       defaultValue: "polygon",
      //     },
      //   },
      //   // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard,
      // },
      // {
      //   idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_zksync_era,
      //   title: "ZKsync Era",
      //   icon: "/icons/blockchains/ic_zksync.svg",
      //   data: {
      //     chainId: {
      //       type: "hidden",
      //       label: "Chain Id",
      //       placeholder: "Chain Id",
      //       defaultValue: "zksync",
      //     },
      //   },
      //   // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard,
      // },
      // {
      //   idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_avalanche,
      //   title: "Avalanche C-Chain",
      //   icon: "/icons/blockchains/ic_avax.svg",
      //   data: {
      //     chainId: {
      //       type: "hidden",
      //       label: "Chain Id",
      //       placeholder: "Chain Id",
      //       defaultValue: "avax",
      //     },
      //   },
      //   // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard,
      // },
      // {
      //   idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_abstract_testnet,
      //   title: "Abstract Testnet",
      //   icon: "/icons/blockchains/ic-abs.png",
      //   data: {
      //     chainId: {
      //       type: "hidden",
      //       label: "Chain Id",
      //       placeholder: "Chain Id",
      //       defaultValue: "abs",
      //     },
      //   },
      //   // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard,
      // },
      // {
      //   idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_duck_chain,
      //   title: "DuckChain",
      //   icon: "/icons/blockchains/ic-duck.svg",
      //   data: {
      //     chainId: {
      //       type: "hidden",
      //       label: "Chain Id",
      //       placeholder: "Chain Id",
      //       defaultValue: "duck",
      //     },
      //   },
      //   // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard,
      // },
      {
        idx: CATEGORY_OPTION_KEYS.blockchains.blockchain_symbiosis,
        title: "Symbiosis",
        icon: "/icons/blockchains/ic_nbs.svg",
        data: {
          chainId: {
            type: "hidden",
            label: "Chain Id",
            placeholder: "Chain Id",
            defaultValue: AgentChainId.Symbiosis,
          },
        },
        // customizeRenderOnBoard: BaseBlockchainCustomizeOnBoard,
      },
    ],
    ...validators,
  } satisfies StudioCategory;
};

const getDecentralizedInferenceCategory = (state: State) => {
  const validators =
    state === "create"
      ? DECENTRALIZE_INFERENCE_VALIDATES.create
      : DECENTRALIZE_INFERENCE_VALIDATES.update;
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
  const validators =
    state === "create" ? TOKEN_VALIDATES.create : TOKEN_VALIDATES.update;
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
            defaultValue: AgentTokenChainId.Base,
          },
        },
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
            defaultValue: AgentTokenChainId.Solana,
          },
        },
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
            defaultValue: AgentTokenChainId.Arbitrum,
          },
        },
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
            defaultValue: AgentTokenChainId.BSC,
          },
        },
      },
    ],
    ...validators,
  } satisfies StudioCategory;
};

const getMissionOnXCategory = (state: State) => {
  const validators =
    state === "create" ? X_VALIDATES.create : X_VALIDATES.update;
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
        customizeRenderOnBoard: CustomBaseXForm as CustomizeComponentType,
        type: StudioCategoryType.LINK,
      },
      {
        idx: CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_reply,
        title: "Reply on X",
        tooltip: "",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_reply.svg",
        customizeRenderOnBoard: CustomBaseXForm as CustomizeComponentType,
        type: StudioCategoryType.LINK,
      },
      {
        idx: CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_engage,
        title: "Engage on X",
        tooltip: "",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_engage.svg",
        customizeRenderOnBoard: CustomBaseXForm as CustomizeComponentType,
        type: StudioCategoryType.LINK,
      },
      {
        idx: CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_follow,
        title: "Follow on X",
        tooltip: "",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_follow.svg",
        customizeRenderOnBoard: CustomBaseXForm as CustomizeComponentType,
        type: StudioCategoryType.LINK,
      },
    ],
    ...validators,
  } satisfies StudioCategory;
};

const getMissionOnFarcasterCategory = (state: State) => {
  const validators =
    state === "create" ? X_VALIDATES.create : X_VALIDATES.update;
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
        customizeRenderOnBoard:
          CustomBaseFarcasterForm as CustomizeComponentType,
        type: StudioCategoryType.LINK,
      },
      {
        idx: CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_reply,
        title: "Reply on Farcaster",
        tooltip: "",
        icon: "https://storage.googleapis.com/eternal-ai/agent-studio-v2/ic_abilities_reply.svg",
        customizeRenderOnBoard:
          CustomBaseFarcasterForm as CustomizeComponentType,
        type: StudioCategoryType.LINK,
      },
    ],
    ...validators,
  } satisfies StudioCategory;
};

export default function getAgentModelCategories(
  state: State
): StudioCategory[] {
  const AGENT_MODEL_CATEGORIES: StudioCategory[] = [
    getAgentCategory(state),
    getPersonalitiesCategory(state),
    // getAiFrameworkCategory(state),
    getBlockchainCategory(state),
    // getDecentralizedInferenceCategory(state),
    getTokenCategory(state),
    getMissionOnXCategory(state),
    getMissionOnFarcasterCategory(state),
  ];
  return AGENT_MODEL_CATEGORIES;
}
