import { StudioCategoryOption } from "@agent-studio/studio-dnd";
import { IChainConnected } from "@eternalai-dagent/core";
import { RenameModels } from "../constants/models";

export const getDynamicAiModelOptions = (
  chains: IChainConnected[] | undefined = [],
  descriptions: Record<string, string> = {}
): StudioCategoryOption[] => {
  try {
    const allSupportedAiModels =
      chains?.reduce(
        (acc: Record<string, string>, chain: IChainConnected) => ({
          ...acc,
          ...chain?.support_model_names,
        }),
        {}
      ) || {};
    const options = Object.entries(allSupportedAiModels).map(([id, key]) => {
      return {
        idx: `decentralize_inference_${id}`,
        title: RenameModels?.[id as any] || id,
        tooltip: descriptions?.[id as any] || "",
        data: {
          decentralizeId: {
            type: "hidden",
            label: "Decentralize Id",
            placeholder: "Decentralize Id",
            defaultValue: id,
          },
          decentralizeKey: {
            type: "hidden",
            label: "Decentralize key",
            placeholder: "Decentralize key",
            defaultValue: key,
          },
        },
      };
    });
    return options as StudioCategoryOption[];
  } catch (e) {
    return [];
  }
};

export const getAllNewDecentralizedInferenceKeys = (
  chains: IChainConnected[] | undefined = []
): Record<string, string> => {
  const newOptionKeys: Record<string, string> = {};
  try {
    const allSupportedAiModels =
      chains?.reduce(
        (acc: Record<string, string>, chain: IChainConnected) => ({
          ...acc,
          ...chain?.support_model_names,
        }),
        {}
      ) || {};
    Object.entries(allSupportedAiModels).map(([id, key]) => {
      newOptionKeys[`decentralize_inference_${id}`] =
        `decentralize_inference_${id}`;
    });

    return newOptionKeys;
  } catch (e) {
    return newOptionKeys;
  }
};
