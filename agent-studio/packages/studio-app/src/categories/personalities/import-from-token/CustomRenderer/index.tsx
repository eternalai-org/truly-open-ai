import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import { useMemo, useState } from "react";
import { updateRelatedAgentData } from "../../helpers";
import { ImportFromTokenFormData } from "../types";
import ContractAddressView from "./Fields/ContractAddressView";
import PersonalityView from "./Fields/PersonalityView";
import TokenNameView from "./Fields/TokenNameView";
import TokenSymbolView from "./Fields/TokenSymbolView";
import useStudioAgentStore from "../../../../stores/useStudioAgentStore";
import useGeneratePersonality from "../../../../hooks/useGeneratePersonality";
import ResetButton from "../../../../components/buttons/ResetButton";
import SubmitButton from "../../../../components/buttons/SubmitButton";
import CustomRendererBase from "../../../../components/CustomRendererBase";
import StudioFormWrapper from "../../../../components/form";

const CustomImportFromTokenRenderer = ({
  id,
  formData,
  setFormFields,
  resetFormData,
  data,
}: StudioCategoryOptionRenderPayload<ImportFromTokenFormData>) => {
  const [isLoading, setIsLoading] = useState(false);
  const { contractAddressStr, tokenName, tokenSymbol, personality, stepper } =
    formData;

  const { isDetail } = useStudioAgentStore();

  const { generatePersonalityByToken } = useGeneratePersonality({
    onStartGenerating: () => setIsLoading(true),
    onFinishGenerating: () => setIsLoading(false),
  });

  const handleGenerateIdea = async () => {
    if (isLoading || !formData.contractAddressStr) {
      return;
    }

    const { agent, personality, tokenInfo } = await generatePersonalityByToken(
      contractAddressStr as string
    );

    setFormFields({
      contractAddressStr: contractAddressStr,
      tokenName: tokenInfo?.baseToken?.name || "",
      tokenSymbol: tokenInfo?.baseToken?.symbol || "",
      tokenImage: tokenInfo?.baseToken?.imageUrl || "",
      tokenNetwork: tokenInfo?.chainId || "",
    });

    if (personality) {
      setFormFields({
        personality,
        stepper: 2,
      });
    }

    if (agent) {
      updateRelatedAgentData(data, id, agent.agent_name);
    }
  };

  const resetOnClickHandler = async () => {
    resetFormData();
    // setAgentDetail(undefined);
    // setIsShowAgentChatInteraction(false);
    // reset();
  };

  const Actions = useMemo(() => {
    return stepper === 2 ? (
      <ResetButton onClick={resetOnClickHandler} title={"Retry"} />
    ) : (
      <SubmitButton
        isLoading={!!isLoading}
        onClick={isLoading ? undefined : handleGenerateIdea}
        title={"Next"}
      />
    );
  }, [stepper, isLoading, handleGenerateIdea, resetOnClickHandler]);

  return (
    <CustomRendererBase
      tag="Personality"
      title="Import from Token"
      actions={!isDetail && Actions}
    >
      <StudioFormWrapper>
        <ContractAddressView
          id={id}
          value={contractAddressStr as string}
          onChange={(ct) => {
            setFormFields({ contractAddressStr: ct });
          }}
        />

        {stepper === 2 && (
          <>
            <TokenNameView
              id={id}
              value={tokenName as string}
              onChange={(ct) => {
                setFormFields({ tokenName: ct });
              }}
            />

            <TokenSymbolView
              id={id}
              value={tokenSymbol as string}
              onChange={(ct) => {
                setFormFields({ tokenSymbol: ct });
              }}
            />

            <PersonalityView
              id={id}
              value={personality as string}
              onChange={(ct) => {
                setFormFields({ personality: ct });
              }}
            />
          </>
        )}
      </StudioFormWrapper>
    </CustomRendererBase>
  );
};

export default CustomImportFromTokenRenderer;
