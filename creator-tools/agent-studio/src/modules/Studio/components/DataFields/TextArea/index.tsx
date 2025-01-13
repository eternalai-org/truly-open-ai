import cs from 'clsx';
import './TextArea.scss';
import { useMemo } from 'react';

import NoDraggable from '../../DnD/base/NoDraggable';

import { useFormFunction } from '@/modules/Studio/hooks/useFormFunction';
import useStudioDataStore from '@/modules/Studio/stores/useStudioDataStore';
import useStudioFormStore from '@/modules/Studio/stores/useStudioFormStore';
import { DataSchema } from '@/modules/Studio/types/category';

type Props = Omit<React.ComponentPropsWithoutRef<'textarea'>, 'defaultValue'> & {
  formId: string;
  name: string;
  readonly?: boolean;
  schemaData?: DataSchema;
  fieldKey: string;
};

function TextArea({ formId, placeholder, className, name, readonly, fieldKey, schemaData, ...rest }: Props) {
  const formFunctions = useFormFunction(fieldKey);

  const formMap = useStudioFormStore((state) => state.formMap);
  const setFormFields = useStudioFormStore((state) => state.setFormFields);

  const value = formMap[formId]?.[name] || '';
  const handleOnChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
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
      <textarea
        {...rest}
        onChange={handleOnChange}
        placeholder={placeholder}
        disabled={fieldData?.disabled}
        name={name}
        className={cs('studio-field-text-area', className, {
          'studio-field-text-area__error': isError,
        })}
        value={value as string}
        rows={4}
      />
    </NoDraggable>
  );
}

export default TextArea;
