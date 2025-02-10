import { useEffect } from "react";
import useCommonStore from "../../stores/useCommonStore";
import ChainAPI from "../../services/apis/chain";
import ModelAPI from "../../services/apis/model";
import useAgentServiceStore from "../../stores/useAgentServiceStore";

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

  useEffect(() => {
    useAgentServiceStore
      .getState()
      .baseDAgent.configAccessToken()
      .then((accessToken) => {
        useAgentServiceStore.setState({ accessToken });
      });

    useAgentServiceStore.setState({
      walletAddress: useAgentServiceStore
        .getState()
        .baseDAgent.getSignerAddress(),
    });
  }, []);

  return children;
}

export default GlobalDataProvider;
