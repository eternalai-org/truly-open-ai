import StudioHorizontalField from "../../../../../components/form/fields/StudioHorizontalField";
import StudioInput from "../../../../../components/form/inputs/StudioInput";
import useStudioAgentStore from "../../../../../stores/useStudioAgentStore";

type Props = {
  id: string;
  value: string;
  onChange: (ct: string) => void;
};

const DescriptionView = ({ id, value, onChange }: Props) => {
  const { isDetail } = useStudioAgentStore();

  return (
    <StudioHorizontalField
      label="Description"
      tooltip="Explain what your knowledge agent is about, its main topics, and how it helps. Mention what it provides, like insights or guidance. Keep it simple and clear."
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

export default DescriptionView;
