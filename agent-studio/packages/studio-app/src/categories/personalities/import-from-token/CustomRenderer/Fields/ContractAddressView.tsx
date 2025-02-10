import StudioHorizontalField from "../../../../../components/form/fields/StudioHorizontalField";
import StudioInput from "../../../../../components/form/inputs/StudioInput";
import useStudioAgentStore from "../../../../../stores/useStudioAgentStore";

type Props = {
  id: string;
  value: string;
  onChange: (ct: string) => void;
};

const ContractAddressView = ({ id, value, onChange }: Props) => {
  const { isDetail } = useStudioAgentStore();
  return (
    <StudioHorizontalField
      label="Contract Address"
      tooltip="Contract Address information..."
    >
      <StudioInput
        value={value}
        placeholder="e.g YTKN"
        onChange={(e) => onChange(e.target.value)}
        disabled={isDetail}
      />
    </StudioHorizontalField>
  );
};

export default ContractAddressView;
