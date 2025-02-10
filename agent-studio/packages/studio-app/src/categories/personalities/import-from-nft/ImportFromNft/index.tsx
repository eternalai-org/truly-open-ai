import {
  StudioCategoryOptionRenderPayload,
  useNewStore,
} from "@agent-studio/studio-dnd";
import { useEffect, useMemo, useState } from "react";
import { updateRelatedAgentData } from "../../helpers";
import { ImportFromNftFormData, ImportFromNftState } from "../types";
import { CollectionView } from "./CollectionView";
import { PersonalityView } from "./PersonalityView";
import { TokenIdView } from "./TokenIdView";
import useStudioAgentStore from "../../../../stores/useStudioAgentStore";
import { useDebounce } from "../../../../hooks/useDebounce";
import useGeneratePersonality from "../../../../hooks/useGeneratePersonality";
import { EIdeaOption } from "../../../../types/agent";
import ResetButton from "../../../../components/buttons/ResetButton";
import SubmitButton from "../../../../components/buttons/SubmitButton";
import CustomRendererBase from "../../../../components/CustomRendererBase";
import StudioFormWrapper from "../../../../components/form";
import { Collection, INFTInfo } from "../../../../types/collection";
import AgentAPI from "../../../../services/apis/agent";

export default function ImportFromNft({
  id,
  formData,
  setFormFields,
  resetFormData,
  data,
}: StudioCategoryOptionRenderPayload<ImportFromNftFormData>) {
  const { addData } = useNewStore<ImportFromNftState>(id);
  const [isLoading, setIsLoading] = useState(false);
  const [isGettingNFTInfo, setIsGettingNFTInfo] = useState(false);
  const { stepper } = formData;
  const selectedNFT = formData.selectedNFT as INFTInfo;
  const selectedOption = formData.selectedOption as Collection;

  const { isDetail } = useStudioAgentStore();
  const debounceId = useDebounce((formData.tokenId || "") as string, 200);

  const { generatePersonalityByNFT } = useGeneratePersonality({
    onStartGenerating: () => setIsLoading(true),
    onFinishGenerating: () => setIsLoading(false),
  });

  const getCollections = async (type: EIdeaOption) => {
    try {
      const res = await AgentAPI.getNFTCollections({
        inscription: type === EIdeaOption.ordinals,
        limit: type === EIdeaOption.ordinals ? 20 : 100,
      });

      addData({
        collections: res?.collections || [],
      });
    } catch (e) {
      //
    }
  };

  useEffect(() => {
    getCollections(EIdeaOption.nft);
  }, []);

  const getNftInfo = async () => {
    try {
      setFormFields({
        selectedNFT: undefined,
      });
      setIsGettingNFTInfo(true);
      const nft = await AgentAPI.getNFTsByCollection({
        contractAddress: selectedOption?.contracts[0]?.address as string,
        nftId: debounceId,
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
    if (debounceId && selectedOption) {
      getNftInfo();
    }
  }, [debounceId, selectedOption]);

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

  const Actions = useMemo(() => {
    return stepper === 2 ? (
      <ResetButton onClick={resetOnClickHandler} title={"Retry"} />
    ) : (
      <SubmitButton
        isLoading={!!isLoading || !!isGettingNFTInfo}
        onClick={isLoading ? undefined : handleGenerateIdea}
        title={"Next"}
      />
    );
  }, [
    stepper,
    isLoading,
    isGettingNFTInfo,
    handleGenerateIdea,
    resetOnClickHandler,
  ]);

  return (
    <CustomRendererBase
      tag="Personality"
      title="Import from NFT"
      actions={!isDetail && Actions}
    >
      <StudioFormWrapper>
        <CollectionView
          id={id}
          formData={formData as ImportFromNftFormData}
          setFormFields={setFormFields}
        />

        <TokenIdView
          id={id}
          formData={formData as ImportFromNftFormData}
          tokenIdStr={(formData.tokenId || "") as string}
          setTokenIdStr={(tokenId) => {
            setFormFields({ tokenId });
            if (tokenId) {
              setFormFields({
                selectedNFT: undefined,
              });
            }
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
}
