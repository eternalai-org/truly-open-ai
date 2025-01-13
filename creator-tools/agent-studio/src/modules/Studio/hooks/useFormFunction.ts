import { useMemo } from 'react';

import useStudioCategoryStore from '../stores/useStudioCategoryStore';
import { StudioFormFieldValidate } from '../types';

/**
 * Hook to get form validation functions for a specific field
 * @param {string} idx - Field identifier
 * @returns {Object} Object containing validation functions
 * @returns {Function} .onFieldValidate - Function to validate field value
 */
export const useFormFunction = (idx: string) => {
  const categoryOptionMap = useStudioCategoryStore((state) => state.categoryOptionMap);

  const memorizedFuncs = useMemo(() => {
    return {
      // onFormChange: categoryMap[key].onFormChange,
      // onFormValidate: categoryMap[key].onFormValidate,
      // onFieldChange: categoryMap[key].onFieldChange,
      onFieldValidate: categoryOptionMap[idx]?.onFieldValidate || ((() => true) as StudioFormFieldValidate),
    };
  }, [idx, categoryOptionMap]);

  return memorizedFuncs;
};
