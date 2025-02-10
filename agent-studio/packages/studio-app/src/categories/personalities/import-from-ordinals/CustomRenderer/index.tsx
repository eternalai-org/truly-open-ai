import {
  StudioCategoryOptionRenderPayload,
  useNewStore,
} from "@agent-studio/studio-dnd";
import { useEffect, useMemo, useState } from "react";
import { updateRelatedAgentData } from "../../helpers";
import { ImportFromOrdinalsFormData, ImportFromOrdinalsState } from "../types";
import CollectionView from "./Fields/CollectionView";
import PersonalityView from "./Fields/PersonalityView";
import TokenIdView from "./Fields/TokenIdView";
import useStudioAgentStore from "../../../../stores/useStudioAgentStore";
import { useDebounce } from "../../../../hooks/useDebounce";
import { EIdeaOption } from "../../../../types/agent";
import ResetButton from "../../../../components/buttons/ResetButton";
import SubmitButton from "../../../../components/buttons/SubmitButton";
import CustomRendererBase from "../../../../components/CustomRendererBase";
import StudioFormWrapper from "../../../../components/form";
import { Collection, INFTInfo } from "../../../../types/collection";
import useGeneratePersonality from "../../../../hooks/useGeneratePersonality";
import AgentAPI from "../../../../services/apis/agent";

const CustomImportFromOrdinalsRenderer = ({
  id,
  formData,
  setFormFields,
  resetFormData,
  data,
}: StudioCategoryOptionRenderPayload<ImportFromOrdinalsFormData>) => {
  const { dataStore, addData } = useNewStore<ImportFromOrdinalsState>(id);

  const [isLoading, setIsLoading] = useState(false);
  const [isGettingNFTInfo, setIsGettingNFTInfo] = useState(false);
  const { isDetail } = useStudioAgentStore();
  const { stepper } = formData;
  const selectedNFT = formData.selectedNFT as INFTInfo;
  const selectedOption = formData.selectedOption as Collection;

  const debouncedId = useDebounce((formData.tokenId || "") as string, 200);

  const { generatePersonalityByNFT } = useGeneratePersonality({
    onStartGenerating: () => setIsLoading(true),
    onFinishGenerating: () => setIsLoading(false),
  });

  const getCollections = async (type: EIdeaOption) => {
    try {
      const rs = await AgentAPI.getNFTCollections({
        inscription: type === EIdeaOption.ordinals,
        limit: type === EIdeaOption.ordinals ? 20 : 100,
      });

      addData({
        collections: rs?.collections || [],
      });
    } catch (e) {
      //
    }
  };

  useEffect(() => {
    getCollections(EIdeaOption.ordinals);
  }, []);

  const getNftInfo = async () => {
    try {
      setFormFields({
        selectedNFT: undefined,
      });
      setIsGettingNFTInfo(true);

      const nft = await AgentAPI.getNFTsByCollection({
        contractAddress: selectedOption?.contracts[0]?.address as string,
        nftId: debouncedId,
        inscription: true,
      });

      setFormFields({
        selectedNFT: nft[0],
      });

      if (nft) {
        if (nft.length > 0) {
          setFormFields({
            selectedNFT: nft[0],
            tokenIdErrorMessage: undefined,
          });
        } else {
          setFormFields({
            tokenIdErrorMessage: "Can not get NFT info",
          });
        }
      }
    } catch (error) {
      console.error("Error getting NFT info:", error);
    } finally {
      setIsGettingNFTInfo(false);
    }
  };

  useEffect(() => {
    if (debouncedId && selectedOption) {
      getNftInfo();
    }
  }, [debouncedId, selectedOption]);

  const handleGenerateIdea = async () => {
    if (isLoading || !selectedNFT) {
      return;
    }

    const { agent, personality } = await generatePersonalityByNFT(
      selectedNFT,
      selectedOption
    );

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
  };

  const Actions = useMemo(
    () =>
      stepper === 2 ? (
        <ResetButton onClick={resetOnClickHandler} title={"Retry"} />
      ) : (
        <SubmitButton
          isLoading={!!isLoading || !!isGettingNFTInfo}
          onClick={isLoading ? undefined : handleGenerateIdea}
          title={"Next"}
        />
      ),
    [
      stepper,
      isLoading,
      isGettingNFTInfo,
      handleGenerateIdea,
      resetOnClickHandler,
    ]
  );

  return (
    <CustomRendererBase
      tag="Personality"
      title="Import from Ordinals"
      actions={!isDetail && Actions}
    >
      <StudioFormWrapper>
        <CollectionView
          id={id}
          formData={formData as ImportFromOrdinalsFormData}
          setFormFields={setFormFields}
        />

        <TokenIdView
          id={id}
          formData={formData as ImportFromOrdinalsFormData}
          tokenIdStr={(formData.tokenId || "") as string}
          setTokenIdStr={(tokenId) => {
            setFormFields({ tokenId });
          }}
        />

        {stepper === 2 && (
          <PersonalityView
            id={id}
            personality={formData?.personality as string}
            onChange={(personality) => {
              setFormFields({ personality });
            }}
          />
        )}
      </StudioFormWrapper>
    </CustomRendererBase>
  );
};

export default CustomImportFromOrdinalsRenderer;
