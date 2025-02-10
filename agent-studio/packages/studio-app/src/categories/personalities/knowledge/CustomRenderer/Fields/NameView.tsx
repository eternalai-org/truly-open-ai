import StudioHorizontalField from "../../../../../components/form/fields/StudioHorizontalField";
import StudioInput from "../../../../../components/form/inputs/StudioInput";
import useStudioAgentStore from "../../../../../stores/useStudioAgentStore";

type Props = {
  id: string;
  value: string;
  onChange: (ct: string) => void;
};

const NameView = ({ id, value, onChange }: Props) => {
  const { isDetail } = useStudioAgentStore();
  return (
    <StudioHorizontalField
      label="Name"
      tooltip="Enter the name of your knowledge agent."
    >
      <StudioInput
        value={value}
        placeholder="Enter"
        onChange={(e) => onChange(e.target.value)}
        disabled={isDetail}
      />
    </StudioHorizontalField>
  );
};

export default NameView;
