import {
  findAncestorNodeIdOfNodeId,
  findDataByCategoryKey,
  findDataById,
  StudioDataNode,
} from "@agent-studio/studio-dnd";
import { useEffect, useState } from "react";
import useAgentDataFieldChange from "./useAgentDataFieldChange";
import { BLOCKCHAIN_CATEGORY_KEY } from "../constants/category-keys";
import useStudioAgentStore from "../stores/useStudioAgentStore";
import { useThrottleValue } from "./useThrottleValue";

export const useDetectChainIdForMissionModel = (nodeId: string) => {
  const { graphData } = useStudioAgentStore();
  const { chainId } = useAgentDataFieldChange();

  const [ancestorChainId, setAncestorChainId] = useState<string>();
  const [ancestorData, setAncestorData] = useState<StudioDataNode>();

  const data = useThrottleValue(graphData.data, 500);

  useEffect(() => {
    const id = findAncestorNodeIdOfNodeId(data, nodeId);
    if (id) {
      const foundData = findDataById(id, data);
      setAncestorData(foundData ?? undefined);
    } else {
      setAncestorData(undefined);
    }
  }, [data, nodeId]);

  useEffect(() => {
    try {
      if (ancestorData) {
        const blockchains = findDataByCategoryKey(
          BLOCKCHAIN_CATEGORY_KEY,
          [ancestorData],
          ancestorData.id
        );
        const network = blockchains[0];
        const chainId = network?.data?.chainId;
        setAncestorChainId(chainId as string);
      } else {
        setAncestorChainId(undefined);
      }
    } catch (e) {
      setAncestorChainId(undefined);
    }
  }, [ancestorData]);

  return ancestorChainId || chainId;
};
