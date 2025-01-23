import StudioHorizontalField from "../../../../../components/form/fields/StudioHorizontalField";
import StudioInput from "../../../../../components/form/inputs/StudioInput";

type Props = {
  id: string;
  value: string;
  onChange: (v: string) => void;
};

const FrequencyField = ({ id, value, onChange }: Props) => {
  return (
    <StudioHorizontalField
      label="Frequency (hours)"
      tooltip="Frequency information..."
    >
      <StudioInput
        value={value}
        onChange={(e) => onChange(e.target.value)}
        placeholder="e.g 2"
      />
    </StudioHorizontalField>
  );
};

export default FrequencyField;
