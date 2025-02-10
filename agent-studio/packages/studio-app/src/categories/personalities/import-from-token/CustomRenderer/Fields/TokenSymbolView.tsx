import StudioHorizontalField from "../../../../../components/form/fields/StudioHorizontalField";
import StudioInput from "../../../../../components/form/inputs/StudioInput";
import useStudioAgentStore from "../../../../../stores/useStudioAgentStore";

type Props = {
  id: string;
  value: string;
  onChange: (ct: string) => void;
};

const TokenSymbolView = ({ id, value, onChange }: Props) => {
  const { isDetail } = useStudioAgentStore();

  return (
    <StudioHorizontalField
      label="Token Symbol"
      tooltip="Token Symbol information..."
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

export default TokenSymbolView;
