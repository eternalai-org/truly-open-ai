import {
  Box,
  Button,
  Modal,
  ModalContent,
  ModalOverlay,
} from "@chakra-ui/react";
import { useState } from "react";
import SelectTwitterModal from "../SelectTwitterModal";
import { useStudio } from "@agent-studio/studio-dnd";
import { ImportGenomicFormData } from "../../types";
import TwitterClonedView from "../TwitterClonedView";
import useStudioAgentStore from "../../../../../stores/useStudioAgentStore";
import { ISearchTwitterInfo } from "../../../../../types/agent";

type Props = {
  formId: string;
};

const TwitterCloneTube = ({ formId }: Props) => {
  const { getFormDataById, setFormFields } = useStudio();
  const { twitterInfos } = getFormDataById<ImportGenomicFormData>(formId);
  const { isDetail } = useStudioAgentStore();
  const [showTwModal, setShowTwModal] = useState(false);

  const handleDeleteTwitter = (twInfo: ISearchTwitterInfo) => {
    if (isDetail) return;
    if (!twInfo) return;

    if (twitterInfos) {
      const foundTw = twitterInfos.find((item) => item?.id === twInfo?.id);

      const filteredTw = twitterInfos.filter(
        (item) => item?.id !== foundTw?.id
      );
      setFormFields<ImportGenomicFormData>(formId, {
        twitterInfos: [...filteredTw],
      });
    }
  };

  return (
    <div>
      {twitterInfos?.map((twInfo) => (
        <TwitterClonedView
          key={twInfo.id}
          info={twInfo}
          handleDeleteTwitter={handleDeleteTwitter}
        />
      ))}

      {!isDetail && (
        <>
          <Box marginTop={"8px"}>
            <Button
              width={"100%"}
              onClick={() => {
                setShowTwModal(true);
              }}
            >
              Pick the DNA
            </Button>
          </Box>
          {/* Modal */}
          <Modal
            isOpen={!!showTwModal}
            onClose={() => setShowTwModal(false)}
            size={"xl"}
            isCentered
          >
            <ModalOverlay />
            <ModalContent minW="600px">
              <SelectTwitterModal
                formId={formId}
                closeModal={() => setShowTwModal(false)}
              />
            </ModalContent>
          </Modal>
        </>
      )}
    </div>
  );
};

export default TwitterCloneTube;
