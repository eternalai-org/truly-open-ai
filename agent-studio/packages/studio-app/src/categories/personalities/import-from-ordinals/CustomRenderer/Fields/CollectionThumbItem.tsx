import { Flex, Image, Text, useDisclosure } from "@chakra-ui/react";
import CollectionModal from "./CollectionModal";
import { CollectionListView } from "./CollectionModal/CollectionListView";
import { CollectionNFTItem } from "./CollectionModal/CollectionNFTItem";
import {
  ImportFromOrdinalsFormData,
  ImportFromOrdinalsState,
} from "../../types";
import { useNewStore } from "@agent-studio/studio-dnd";
import useStudioAgentStore from "../../../../../stores/useStudioAgentStore";
import { Collection } from "../../../../../types/collection";

type Props = {
  id: string;
  formData: ImportFromOrdinalsFormData;
  setFormFields: (fields: Partial<ImportFromOrdinalsFormData>) => void;
};

const CollectionThumbItem = ({ id, formData, setFormFields }: Props) => {
  const { dataStore, addData } = useNewStore<ImportFromOrdinalsState>(id);
  const selectedOption = formData.selectedOption as Collection;
  const { isDetail } = useStudioAgentStore();

  const {
    isOpen: isOpenCollectionModal,
    onOpen: onOpenCollectionModal,
    onClose: onCloseCollectionModal,
  } = useDisclosure({ id: "NftCollectionModal" });

  const renderItemContent = () => {
    return (
      <Flex flexDir={"row"} align={"center"} gap="8px" p="5px">
        <Image
          src={selectedOption?.image_url}
          w={"24px"}
          h={"24px"}
          borderRadius={"100%"}
          overflow={"hidden"}
        />
        <Text
          fontSize={"14px"}
          lineHeight={"calc(20 / 14)"}
          color={"#000"}
          fontWeight={"500"}
          fontFamily={"var(--font-SFProDisplay)"}
        >
          {selectedOption?.name}
        </Text>
      </Flex>
    );
  };

  return (
    <>
      <Flex
        w="100%"
        h="28px"
        align={"center"}
        justify={"flex-start"}
        minW={"246px"}
        padding={"0 12px"}
        bgColor={"#fff"}
        borderRadius={"999px"}
        _hover={{
          cursor: isDetail ? "not-allowed" : "pointer",
          opacity: isDetail ? 0.7 : 1,
        }}
        onClick={isDetail ? undefined : onOpenCollectionModal}
        opacity={isDetail ? 0.7 : 1}
      >
        {!selectedOption ? (
          <Text
            fontSize={"18px"}
            lineHeight={"calc(28 / 18)"}
            color={"#777777"}
            fontWeight={"500"}
            fontFamily={"var(--font-SFProDisplay)"}
            textAlign={"center"}
          >
            Select Collection
          </Text>
        ) : (
          renderItemContent()
        )}
      </Flex>

      {!isDetail && isOpenCollectionModal && (
        <CollectionModal
          isOpen={isOpenCollectionModal}
          onClose={onCloseCollectionModal}
          title="Collection list"
          description="Choose one of those collections"
          dataListView={
            <CollectionListView
              dataList={dataStore?.collections || []}
              renderItem={(item, index) => {
                return (
                  <CollectionNFTItem
                    item={item}
                    index={index}
                    itemOnClick={(item) => {
                      setFormFields({
                        selectedOption: item,
                      });

                      onCloseCollectionModal();
                    }}
                  />
                );
              }}
            />
          }
        />
      )}
    </>
  );
};

export default CollectionThumbItem;
