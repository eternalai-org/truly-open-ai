import cs from 'clsx';
import './Select.scss';
import { useMemo } from 'react';

import NoDraggable from '../../DnD/base/NoDraggable';

import { useFormFunction } from '@/modules/Studio/hooks/useFormFunction';
import useStudioDataSourceStore from '@/modules/Studio/stores/useStudioDataSourceStore';
import useStudioDataStore from '@/modules/Studio/stores/useStudioDataStore';
import useStudioFormStore from '@/modules/Studio/stores/useStudioFormStore';
import { DataSchema } from '@/modules/Studio/types/category';

type Props = Omit<React.ComponentPropsWithoutRef<'select'>, 'defaultValue'> & {
  formId: string;
  name: string;
  placeholder?: string;
  dataSourceKey?: string;
  readonly?: boolean;
  schemaData?: DataSchema;
  fieldKey: string;
};

function Select({
  formId,
  className,
  name,
  placeholder = 'Select',
  dataSourceKey,
  readonly,
  fieldKey,
  schemaData,
  ...rest
}: Props) {
  // const data = useStudioDataStore((state) => state.data);
  const formFunctions = useFormFunction(fieldKey);
  const formMap = useStudioFormStore((state) => state.formMap);
  const setFormFields = useStudioFormStore((state) => state.setFormFields);

  const { dataSource } = useStudioDataSourceStore();

  const options = useMemo(() => {
    if (dataSourceKey) {
      return dataSource[dataSourceKey] || [];
    }

    return [];
  }, [dataSource, dataSourceKey]);

  const value = formMap[formId]?.[name] || '';
  const handleOnChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    if (!readonly) {
      const selectable = options.find((op) => op.value === e.target.value)?.selectable;

      if (selectable) {
        setFormFields(formId, {
          [name]: e.target.value,
        });
      }
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

    // do not update dependencies
  }, [name, value]);

  const fieldData = schemaData?.[name];

  return (
    <NoDraggable>
      <select
        {...rest}
        onChange={handleOnChange}
        name={name}
        className={cs(
          'studio-field-select',
          {
            ['studio-field-select__empty']: !value,
            ['studio-field-select__error']: isError,
          },
          className,
        )}
        value={value as string}
        disabled={!!fieldData?.disabled}
      >
        <option value="" className="studio-field-select__placeholder">
          {placeholder}
        </option>

        {options.map((op) => (
          <option disabled={!op.selectable} key={`form-render-select-${formId}-${op.value}`} value={op.value}>
            {op.label}
          </option>
        ))}
      </select>
    </NoDraggable>
  );
}

export default Select;
