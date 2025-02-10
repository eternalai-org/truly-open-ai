import { ElizaFrameworkFormData } from "../types";
import { Box, Flex, Text } from "@chakra-ui/react";
import { useMemo } from "react";
import useStudioAgentStore from "../../../../stores/useStudioAgentStore";
import { ELIZA_CONFIG_DEFAULT } from "../../../../constants/default-values";
import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import { isJsonString } from "../../../../utils/string";
import CustomRendererBase from "../../../../components/CustomRendererBase";
import StudioVerticalField from "../../../../components/form/fields/StudioVerticalField";
import StudioTextArea from "../../../../components/form/inputs/StudioTextArea";

function ElizaFramework({
  formData,
  setFormFields,
  option,
}: StudioCategoryOptionRenderPayload<ElizaFrameworkFormData>) {
  const { config } = formData;
  const { isDetail } = useStudioAgentStore();

  const downloadJSON = () => {
    const jsonData = new Blob([JSON.stringify(ELIZA_CONFIG_DEFAULT)], {
      type: "application/json",
    });
    const jsonURL = URL.createObjectURL(jsonData);
    const link = document.createElement("a");
    link.href = jsonURL;
    link.download = "config.json";
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  };

  const errorMessage = useMemo(() => {
    if (config) {
      if (isJsonString(config)) {
        return undefined;
      }
      return "Config is not a valid JSON";
    }

    return "Config is required";
  }, [config]);

  return (
    <CustomRendererBase tag="AI Framework" title="Eliza">
      <StudioVerticalField>
        <StudioTextArea
          value={config}
          onChange={(e) => {
            setFormFields({ config: e.target.value });
          }}
          placeholder="Explain what your knowledge agent is about, its main topics, and how it helps. Mention what it provides, like insights or guidance. Keep it simple and clear."
          errorMessage={errorMessage}
          disabled={isDetail}
        />
      </StudioVerticalField>

      {!isDetail && (
        <Box>
          <Flex justifyContent={"flex-end"}>
            <Text
              fontSize={"13px"}
              textDecoration={"underline"}
              cursor={"pointer"}
              onClick={downloadJSON}
            >
              Download template
            </Text>
          </Flex>
        </Box>
      )}
    </CustomRendererBase>
  );
}

export default ElizaFramework;
