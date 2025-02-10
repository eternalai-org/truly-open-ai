import StudioVerticalField from "../../../../../components/form/fields/StudioVerticalField";
import StudioTextArea from "../../../../../components/form/inputs/StudioTextArea";

type Props = {
  id: string;
  value: string;
  onChange: (ct: string) => void;
};

const PersonalityView = ({ id, value, onChange }: Props) => {
  // const { isDetail } = useStudioAgentStore();

  return (
    <StudioVerticalField
      label="Personality"
      tooltip="Personality information..."
    >
      <StudioTextArea
        // disabled={isDetail}
        value={value}
        onChange={(e) => onChange(e.target.value)}
        placeholder="Tell us your agent idea, and our AI assistant will bring it to life."
      />
    </StudioVerticalField>
  );
};

export default PersonalityView;
