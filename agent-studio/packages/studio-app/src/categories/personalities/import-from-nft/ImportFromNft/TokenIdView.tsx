import StudioHorizontalField from "../../../../components/form/fields/StudioHorizontalField";
import StudioInput from "../../../../components/form/inputs/StudioInput";
import useStudioAgentStore from "../../../../stores/useStudioAgentStore";
import ReviewNft from "../../ReviewNft";
import { ImportFromNftFormData } from "../types";

type Props = {
  setTokenIdStr: (v: string) => void;
  tokenIdStr: string;
  id: string;
  formData: ImportFromNftFormData;
};

export const TokenIdView = ({
  id,
  setTokenIdStr,
  tokenIdStr,
  formData,
}: Props) => {
  const { selectedNFT, tokenIdErrorMessage } = formData;
  const { isDetail } = useStudioAgentStore();

  return (
    <StudioHorizontalField
      label="Token ID"
      tooltip="Token ID information..."
      errorMessage={tokenIdErrorMessage}
      action={<ReviewNft selectedNFT={selectedNFT} />}
    >
      <StudioInput
        value={tokenIdStr}
        placeholder="e.g 12345"
        onChange={(e) => setTokenIdStr(e.target.value)}
        disabled={isDetail}
      />
    </StudioHorizontalField>
  );
};
