import StudioVerticalField from "../../../../../components/form/fields/StudioVerticalField";
import StudioTextArea from "../../../../../components/form/inputs/StudioTextArea";
import { DEFAULT_ABILITY_PLACEHOLDER } from "../../../../../constants/default-values";

type Props = {
  id: string;
  value: string;
  onChange: (v: string) => void;
};

const DetailsField = ({ id, value, onChange }: Props) => {
  return (
    <StudioVerticalField
      label="Detailed instructions"
      tooltip="Detailed instructions information..."
    >
      <StudioTextArea
        value={value}
        onChange={(e) => onChange(e.target.value)}
        placeholder={DEFAULT_ABILITY_PLACEHOLDER}
      />
    </StudioVerticalField>
  );
};

export default DetailsField;
