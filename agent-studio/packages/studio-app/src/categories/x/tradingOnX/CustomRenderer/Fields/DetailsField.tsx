import StudioVerticalField from "../../../../../components/form/fields/StudioVerticalField";
import StudioTextArea from "../../../../../components/form/inputs/StudioTextArea";

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
        placeholder="Guide the agent on how to create the post by specifying the desired voice, tone, and style. Should it be professional, casual, witty, or inspiring? Provide this preference to shape the content effectively. 
If you have specific ideas, phrases, or facts you'd like included, include them in the input for a more tailored result. 
Mention a word count if you prefer the post to be concise or detailed. 
Test the feature by reviewing the generated post, and if it’s not exactly what you need, refine your input and try again until it aligns with your expectations."
      />
    </StudioVerticalField>
  );
};

export default DetailsField;
