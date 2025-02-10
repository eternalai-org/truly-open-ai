import StudioHorizontalField from "../../../../components/form/fields/StudioHorizontalField";
import { ImportFromNftFormData } from "../types";
import { CollectionThumbItem } from "./CollectionThumbItem";

type Props = {
  id: string;
  formData: ImportFromNftFormData;
  setFormFields: (fields: Partial<ImportFromNftFormData>) => void;
};

export const CollectionView = ({ id, formData, setFormFields }: Props) => {
  return (
    <StudioHorizontalField
      label="Collection"
      tooltip="Collection information..."
    >
      <CollectionThumbItem
        id={id}
        formData={formData as ImportFromNftFormData}
        setFormFields={setFormFields}
      />
    </StudioHorizontalField>
  );
};
