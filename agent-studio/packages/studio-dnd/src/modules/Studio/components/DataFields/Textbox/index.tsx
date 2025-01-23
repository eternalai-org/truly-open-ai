import cs from 'clsx';
import './Textbox.scss';
import { useMemo } from 'react';

import NoDraggable from '../../DnD/base/NoDraggable';

import { useFormFunction } from '@/modules/Studio/hooks/useFormFunction';
import useStudioDataStore from '@/modules/Studio/stores/useStudioDataStore';
import useStudioFormStore from '@/modules/Studio/stores/useStudioFormStore';
import { DataSchema } from '@/modules/Studio/types/category';

type Props = Omit<React.ComponentPropsWithoutRef<'input'>, 'defaultValue'> & {
  formId: string;
  name: string;
  readonly?: boolean;
  schemaData?: DataSchema;
  fieldKey: string;
};

/**
 * Textbox component for form input
 * @param {Object} props - Component props
 * @param {string} props.formId - Unique identifier for the form
 * @param {string} props.name - Field name in the form
 * @param {boolean} props.readonly - Whether the field is read-only
 * @param {DataSchema} props.schemadata - Schema definition for the field
 * @param {string} props.fieldKey - Unique key for field validation
 * @returns {JSX.Element} Rendered textbox component
 */
function Textbox({ formId, placeholder, className, name, readonly, fieldKey, schemaData, ...rest }: Props) {
  const formFunctions = useFormFunction(fieldKey);

  const formMap = useStudioFormStore((state) => state.formMap);
  const setFormFields = useStudioFormStore((state) => state.setFormFields);

  const value = formMap[formId]?.[name] || '';
  const handleOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (!readonly) {
      setFormFields(formId, {
        [name]: e.target.value,
      });
    }
  };

  const isError = useMemo(() => {
    return !(
      formFunctions?.onFieldValidate?.(name, value, {
        formId,
        formData: formMap[formId],
        allFormData: formMap,
        data: useStudioDataStore.getState().data,
      }) ?? true
    );

    // No need to add onFieldValidate to dependencies
  }, [name, value]);

  const fieldData = schemaData?.[name];
  return (
    <NoDraggable>
      <input
        {...rest}
        onChange={handleOnChange}
        type="text"
        placeholder={placeholder}
        name={name}
        disabled={!!fieldData?.disabled}
        className={cs('studio-field-input', className, {
          'studio-field-input__error': isError,
        })}
        value={value as string}
      />
    </NoDraggable>
  );
}

export default Textbox;
