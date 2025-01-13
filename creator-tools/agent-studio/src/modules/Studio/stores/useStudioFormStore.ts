import { create } from 'zustand';

import { FormDataMap } from '../types/base';

const DEFAULT_VALUE = {
  formMap: {},
};

type Store = {
  formMap: Record<string, FormDataMap>;
  initDataForms: (data: Record<string, FormDataMap>) => void;
  addForm: (id: string, data: FormDataMap) => void;
  editForm: (id: string, data: FormDataMap) => void;
  setFormFields: (id: string, fields: Partial<FormDataMap>) => void;
  removeForm: (id: string) => void;
  getFormById: (id: string) => FormDataMap | undefined;
  resetFormById: (id: string) => void;
  clear: () => void;
};

const useStudioFormStore = create<Store>((set, get) => ({
  ...DEFAULT_VALUE,

  initDataForms: (data) => {
    set({ formMap: data });
  },
  addForm: (id, data) => {
    set((state) => ({
      formMap: {
        ...state.formMap,
        [id]: data,
      },
    }));
  },
  editForm: (id, data) => {
    set((state) => ({
      formMap: {
        ...state.formMap,
        [id]: data,
      },
    }));
  },
  setFormFields: (id, fields) => {
    set((state) => ({
      formMap: {
        ...state.formMap,
        [id]: {
          ...state.formMap[id],
          ...fields,
        },
      },
    }));
  },
  removeForm: (id) => {
    set((state) => {
      const formMap = { ...state.formMap };
      delete formMap[id];

      return { formMap };
    });
  },
  getFormById: (id) => get().formMap[id],
  resetFormById: (id) => {
    set((state) => ({
      formMap: {
        ...state.formMap,
        [id]: {},
      },
    }));
  },
  clear: () => set(DEFAULT_VALUE),
}));

export default useStudioFormStore;
