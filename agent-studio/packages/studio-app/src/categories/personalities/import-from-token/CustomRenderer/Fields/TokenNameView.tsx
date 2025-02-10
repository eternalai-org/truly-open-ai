import StudioHorizontalField from "../../../../../components/form/fields/StudioHorizontalField";
import StudioInput from "../../../../../components/form/inputs/StudioInput";
import useStudioAgentStore from "../../../../../stores/useStudioAgentStore";

type Props = {
  id: string;
  value: string;
  onChange: (ct: string) => void;
};

const TokenNameView = ({ id, value, onChange }: Props) => {
  const { isDetail } = useStudioAgentStore();

  return (
    <StudioHorizontalField
      label="Token Name"
      tooltip="Token Name information..."
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

export default TokenNameView;
