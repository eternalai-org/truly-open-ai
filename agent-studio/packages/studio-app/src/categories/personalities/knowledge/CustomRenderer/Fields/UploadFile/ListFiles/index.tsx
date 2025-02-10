import { Box, Flex, Text } from "@chakra-ui/react";

import s from "./styles.module.scss";
import NumberCounter from "../../../../../../../components/NumberCounter";
import { prettyPrintBytes } from "../../../../../../../utils/file";

export interface IProps {
  fileUpload?: File[];
  fileUploading: {
    id: number;
    progress: number;
  };
}

const ListFiles: React.FC<IProps> = ({ fileUpload, fileUploading }: IProps) => {
  const renderStatus = (index: number) => {
    if (index < fileUploading?.id) {
      return <Text color="#4F43E2">Uploaded</Text>;
    }

    if (fileUploading?.id === index) {
      if (fileUploading?.progress === 1) {
        return <Text color="#4F43E2">Uploaded</Text>;
      }

      return (
        <Text color="#00AA6C" display={"flex"}>
          Uploading (
          <NumberCounter
            counter={Math.floor(fileUploading?.progress) * 100}
            delay={0}
          />
          %)
        </Text>
      );
    }

    return <Text color="#898989">Pending</Text>;
  };

  if (!fileUpload || fileUpload?.length < 1) return null;

  return (
    <Flex className={s.listFiles}>
      {fileUpload?.map((file: File, index: number) => {
        return (
          <Flex className={s.item} key={index}>
            <Flex
              alignItems={"center"}
              justifyContent={"space-between"}
              gap="8px"
            >
              <Text className={s.name}>{file?.name}</Text>
              <Text className={s.size}>({prettyPrintBytes(file?.size)})</Text>
            </Flex>
            {!!fileUploading && (
              <Box fontSize={"15px"} fontWeight={500}>
                {renderStatus(index)}
              </Box>
            )}
          </Flex>
        );
      })}
    </Flex>
  );
};

export default ListFiles;
