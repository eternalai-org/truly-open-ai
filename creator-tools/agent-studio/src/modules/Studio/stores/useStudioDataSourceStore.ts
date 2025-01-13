import { create } from 'zustand';

import { DataSource } from '../types/data-source';

type Store = {
  dataSource: Record<string, DataSource[]>;
  setDataSource: (data: Record<string, DataSource[]>) => void;
};

const useStudioDataSourceStore = create<Store>((set, get) => ({
  dataSource: {},
  setDataSource: (data) => {
    // processing input data
    const inputSource = Object.keys(data).reduce(
      (acc, key) => ({
        ...acc,
        [key]: data[key].map((item) => ({
          ...item,
          selectable: item.selectable ?? true,
        })),
      }),
      {},
    );
    set({
      dataSource: {
        ...get().dataSource,
        ...inputSource,
      },
    });
  },
}));

export default useStudioDataSourceStore;
