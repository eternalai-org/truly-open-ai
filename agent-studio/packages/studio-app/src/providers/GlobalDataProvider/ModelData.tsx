import { GraphData, Studio, StudioCategory } from "@agent-studio/studio-dnd";
import { CATEGORY_OPTION_KEYS } from "../../constants/category-option-keys";
import getAgentModelCategories from "../../categories/categories";
import { DECENTRALIZED_INFERENCE_CATEGORY_KEY } from "../../constants/category-keys";
import {
  getAllNewDecentralizedInferenceKeys,
  getDynamicAiModelOptions,
} from "../../utils/category";
import { useEffect, useMemo } from "react";
import useCommonStore from "../../stores/useCommonStore";

const args = {
  dataSource: {},
  showConnectLine: true,
};

function ModelData({
  children,
  state,
}: {
  children: (categories: StudioCategory[]) => React.ReactNode;
  state: "create" | "update";
}) {
  const { modelDescriptions, chains } = useCommonStore();

  useEffect(() => {
    const newDecentralizedKeys = getAllNewDecentralizedInferenceKeys(chains);

    CATEGORY_OPTION_KEYS.decentralized = {
      ...CATEGORY_OPTION_KEYS.decentralized,
      ...newDecentralizedKeys,
    };
  }, [chains]);

  const renderCategories = useMemo(() => {
    return getAgentModelCategories(state, {
      [DECENTRALIZED_INFERENCE_CATEGORY_KEY]: getDynamicAiModelOptions(
        chains || [],
        modelDescriptions || {}
      ),
    });
  }, [chains, modelDescriptions, state]);

  const dataGraph = useMemo(() => {
    return {
      data: [],
      viewport: { x: 0, y: 0, zoom: 1 },
    } satisfies GraphData;
  }, []);

  if (chains?.length && Object.keys(modelDescriptions).length) {
    return children(renderCategories) as any;
  }

  return (
    <Studio
      {...args}
      categories={[]}
      graphData={dataGraph}
      sidebarWidth={"430px"}
    />
  );
}

export default ModelData;
