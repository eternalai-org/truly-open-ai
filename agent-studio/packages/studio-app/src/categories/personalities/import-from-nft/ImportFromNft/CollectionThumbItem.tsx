import { Flex, Text, useDisclosure, Image } from "@chakra-ui/react";
import CollectionModal from "./CollectionModal";
import { CollectionListView } from "./CollectionModal/CollectionListView";
import { CollectionNFTItem } from "./CollectionModal/CollectionNFTItem";
import { useNewStore } from "@agent-studio/studio-dnd";
import { ImportFromNftFormData, ImportFromNftState } from "../types";
import useStudioAgentStore from "../../../../stores/useStudioAgentStore";
import { Collection } from "../../../../types/collection";

type Props = {
  id: string;
  formData: ImportFromNftFormData;
  setFormFields: (fields: Partial<ImportFromNftFormData>) => void;
};

export const CollectionThumbItem = ({ id, formData, setFormFields }: Props) => {
  const { dataStore, addData } = useNewStore<ImportFromNftState>(id);
  const selectedOption = formData.selectedOption as Collection;
  const {
    isOpen: isOpenCollectionModal,
    onOpen: onOpenCollectionModal,
    onClose: onCloseCollectionModal,
  } = useDisclosure({ id: "NftCollectionModal" });
  const { isDetail } = useStudioAgentStore();

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
        h="30px"
        p={"5px"}
        align={"center"}
        justify={"flex-start"}
        minW={"250px"}
        bgColor={"#fff"}
        borderRadius={"16px"}
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
