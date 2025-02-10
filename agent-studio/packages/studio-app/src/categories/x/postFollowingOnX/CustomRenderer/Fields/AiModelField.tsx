import { StudioDataNode } from "@agent-studio/studio-dnd";
import { useEffect, useMemo } from "react";
import useStudioAgentStore from "../../../../../stores/useStudioAgentStore";
import useAgentDataFieldChange from "../../../../../hooks/useAgentDataFieldChange";
import { tokens } from "../../../../../constants/tokens";
import { compareString } from "../../../../../utils/string";
import { RenameModels } from "../../../../../constants/models";
import StudioHorizontalField from "../../../../../components/form/fields/StudioHorizontalField";
import StudioDropdown from "../../../../../components/form/inputs/StudioDropdown";
import useCommonStore from "../../../../../stores/useCommonStore";

type Props = {
  id: string;
  value: string;
  onChange: (v: string) => void;
  data: StudioDataNode[];
};

const AiModelField = ({ id, value, onChange, data }: Props) => {
  const { isDetail } = useStudioAgentStore();
  const chains = useCommonStore((state) => state.chains);
  const { chainId } = useAgentDataFieldChange();

  const options = useMemo(() => {
    try {
      const matchedToken = tokens.find((v) => compareString(v.id, chainId));

      const selectedChain = chains?.find((v) =>
        compareString(v.chain_id, matchedToken?.chainId)
      );

      if (selectedChain?.support_model_names) {
        return Object.entries(selectedChain.support_model_names).map(
          (item) => ({
            label: RenameModels?.[item[0] as any] || item[0],
            value: item[1],
          })
        );
      }
    } catch (e) {
      ///
    }
    return [];
  }, [chainId]);

  useEffect(() => {
    if (!isDetail) {
      if (chainId) {
        if (options?.length && value) {
          const option =
            options.find((item) => item.value === value) || options[0];
          if (option) {
            onChange(option.value);
          }
        } else {
          const option = options[0];

          if (option) {
            onChange(option.value);
          }
        }
      } else {
        onChange("");
      }
    }
  }, [chainId, options?.length, value]);

  return (
    <StudioHorizontalField label="Model" tooltip="Model">
      <StudioDropdown
        width={"350px"}
        value={value}
        onChange={(v) => {
          onChange(v as string);
        }}
        placeholder="AI Model"
        options={options}
      />
    </StudioHorizontalField>
  );
};

export default AiModelField;
