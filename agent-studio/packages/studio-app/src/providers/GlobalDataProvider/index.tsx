import { useEffect } from "react";
import useCommonStore from "../../stores/useCommonStore";
import ChainAPI from "../../services/apis/chain";
import ModelAPI from "../../services/apis/model";

function GlobalDataProvider({ children }: { children: any }) {
  useEffect(() => {
    ChainAPI.getChainList().then((chains) => {
      useCommonStore.setState({ chains });
    });
  }, []);

  useEffect(() => {
    ModelAPI.getModelDescription().then((res) => {
      useCommonStore.getState().setModelDescriptions(res as any);
    });
  }, []);

  return children;
}

export default GlobalDataProvider;
