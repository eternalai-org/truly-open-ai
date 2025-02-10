import StudioHorizontalField from "../../../../../components/form/fields/StudioHorizontalField";
import { ImportFromOrdinalsFormData } from "../../types";
import CollectionThumbItem from "./CollectionThumbItem";

type Props = {
  id: string;
  formData: ImportFromOrdinalsFormData;
  setFormFields: (fields: Partial<ImportFromOrdinalsFormData>) => void;
};

const CollectionView = ({ id, formData, setFormFields }: Props) => {
  return (
    <StudioHorizontalField
      label="Collection"
      tooltip="Collection information..."
    >
      <CollectionThumbItem
        id={id}
        formData={formData as ImportFromOrdinalsFormData}
        setFormFields={setFormFields}
      />
    </StudioHorizontalField>
  );
};

export default CollectionView;
