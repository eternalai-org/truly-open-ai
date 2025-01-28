import useStudioCategoryStore from '../stores/useStudioCategoryStore';

export const getProcessedCategory = () => {
  return useStudioCategoryStore.getState().categories;
};
