import { Box, Flex } from "@chakra-ui/react";
import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import { KnowledgeFormData } from "../types";
import NameView from "./Fields/NameView";
import DescriptionView from "./Fields/DescriptionView";
import UploadFile from "./Fields/UploadFile";
import ListFiles from "./Fields/UploadFile/ListFiles";
import { useEffect, useState } from "react";
import useStudioAgentStore from "../../../../stores/useStudioAgentStore";
import { showValidateError } from "../../../../utils/toast";
import CustomRendererBase from "../../../../components/CustomRendererBase";
import StudioFormWrapper from "../../../../components/form";
import ResetButton from "../../../../components/buttons/ResetButton";
import SubmitButton from "../../../../components/buttons/SubmitButton";
import FileAPI from "../../../../services/apis/file";

const CustomKnowledgeRenderer = ({
  id,
  formData,
  setFormFields,
  resetFormData,
}: StudioCategoryOptionRenderPayload<KnowledgeFormData>) => {
  const [isLoading, setIsLoading] = useState(false);
  const { name, description, fileUpload, stepper } = formData;

  const { isDetail } = useStudioAgentStore();

  const [fileUploading, setFileUploading] = useState<any>(null);
  const [listFileUploadingTemp, setListFileUploadingTemp] = useState<any>([]);

  useEffect(() => {
    if (fileUpload && fileUpload?.length > 0) {
      setListFileUploadingTemp([...fileUpload]);
    }
  }, fileUpload);

  const isValidateForm = (): boolean => {
    if (!formData.name) {
      showValidateError("Knowledge name is required");
      return false;
    }

    if (!formData.description) {
      showValidateError("Knowledge description is required");
      return false;
    }

    if (listFileUploadingTemp?.length < 1) {
      showValidateError("Please upload at least one file");
      return false;
    }

    return true;
  };

  const asyncIterableFileUpload = async function* () {
    for await (const [index, file] of listFileUploadingTemp.entries()) {
      const uploadSize = 10_000_000;

      const url = await FileAPI.uploadFile(
        {
          file: file,
        },
        uploadSize,
        () => {},
        (process) => {
          setFileUploading({
            id: index,
            progress: process.progress,
          });
        }
      );

      if (!url) return;

      yield {
        name: file.name,
        size: file.size,
        url,
      };
    }
  };

  const handleUploadFilesToCloud = async () => {
    if (isLoading || isValidateForm() === false) {
      return;
    }
    try {
      setIsLoading(true);

      (Array as any)
        .fromAsync(asyncIterableFileUpload() as any)
        .then(async (urls: any) => {
          console.log(urls);
          setFormFields({
            stepper: 2,
            name: name as string,
            description: description as string,
            fileUpload: urls,
          });
        });
    } catch (error) {
      console.error("Error uploading files:", error);
    } finally {
      setIsLoading(false);
    }
  };

  const resetOnClickHandler = async () => {
    resetFormData();
    setListFileUploadingTemp([]);
  };

  return (
    <CustomRendererBase tag="Personality" title="Knowledge">
      <StudioFormWrapper>
        <NameView
          id={id}
          value={name as string}
          onChange={(value) => {
            setFormFields({ name: value });
          }}
        />

        <DescriptionView
          id={id}
          value={description as string}
          onChange={(value) => {
            setFormFields({ description: value });
          }}
        />
        {stepper !== 2 && (
          <>
            <UploadFile
              multiple
              labelText="Upload knowledge files here."
              maxSize={999999999}
              disabled={false}
              onChange={(file: FileList) => {
                const fileArr = Object.values(file);
                if (!listFileUploadingTemp) {
                  setListFileUploadingTemp(fileArr);
                  return;
                }

                setListFileUploadingTemp([
                  ...listFileUploadingTemp,
                  ...fileArr,
                ]);
              }}
              fileOrFiles={listFileUploadingTemp}
              acceptedFileType={[
                "asciidoc",
                "md",
                "docx",
                "html",
                "pdf",
                "pptx",
                "xlsx",
                "xml_pubmed",
                "xml_uspto",
              ]}
            />
            <Box h="8px" w="100%" />
          </>
        )}

        <ListFiles
          fileUpload={listFileUploadingTemp}
          fileUploading={fileUploading}
        />

        {!isDetail && (
          <Flex
            flexDir={"row"}
            align={"center"}
            justify={"flex-end"}
            gap={"10px"}
          >
            {stepper === 2 ? (
              <ResetButton onClick={resetOnClickHandler} title={"Retry"} />
            ) : (
              <SubmitButton
                disabled={isLoading}
                isLoading={isLoading}
                onClick={isLoading ? undefined : handleUploadFilesToCloud}
                title={"Next"}
              />
            )}
          </Flex>
        )}
      </StudioFormWrapper>
    </CustomRendererBase>
  );
};

export default CustomKnowledgeRenderer;
