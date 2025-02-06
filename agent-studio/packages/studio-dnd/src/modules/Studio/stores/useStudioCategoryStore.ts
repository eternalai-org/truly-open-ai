import { create } from 'zustand';

import { COLOR_PALETTES, POPULAR } from '../constants/color-palettes';
import { DEFAULT_CATEGORY_TYPE } from '../constants/default-values';
import { StudioCategory, StudioCategoryMapValue, StudioCategoryOptionMapValue } from '../types';
import { StudioDataNode } from '../types/graph';

const DEFAULT_VALUE = {
  rootCategory: null,
  categories: [],
  categoryMap: {},
  categoryOptionMap: {},
  filters: [],
  usedKeyCollection: {},
};

type Store = {
  rootCategory: StudioCategory | null;
  setRootCategory: (category: StudioCategory | null) => void;

  categories: StudioCategory[];
  categoryMap: Record<string, StudioCategoryMapValue>;
  categoryOptionMap: Record<string, StudioCategoryOptionMapValue>;
  setCategories: (categories: StudioCategory[]) => Promise<void>;
  updateCategory: (category: StudioCategory) => void;

  updateCategoriesForEntry: (entry: StudioDataNode | null) => void;

  filters: string[];
  setFilters: (filter: string) => void;

  usedKeyCollection: Record<string, string>;
  setUsedKeyCollection: (collection: Record<string, string>) => void;

  findCategoryByOptionKey: (optionKey: string) => StudioCategory | null;

  clear: () => void;
  scanFromData: (data: StudioDataNode[]) => void;
};

let colorUsedCollection: Record<string, string> = {};

const COLOR_COMBINATION = [...POPULAR, ...COLOR_PALETTES];
const pickColorFromPalettes = (colorIndex: number, color?: string) => {
  if (color) {
    colorUsedCollection[color] = color;

    return color;
  }

  const newColor = COLOR_COMBINATION[colorIndex];

  // check if color is already used
  if (colorUsedCollection[newColor]) {
    return pickColorFromPalettes(colorIndex + 1);
  }

  colorUsedCollection[newColor] = newColor;

  return newColor;
};

const flatCategoriesByKey = (categories: StudioCategory[]) => {
  const categoryMap: Record<string, StudioCategoryMapValue> = {};
  const categoryOptionMap: Record<string, StudioCategoryOptionMapValue> = {};

  categories.forEach((item) => {
    categoryMap[item.idx] = item;

    item.options.forEach((option) => {
      categoryOptionMap[option.idx] = {
        ...option,
        parent: item,
      };
    });
  });

  return { categoryMap, categoryOptionMap };
};

const useStudioCategoryStore = create<Store>((set, get) => ({
  ...DEFAULT_VALUE,

  setRootCategory: (category) => {
    set({ rootCategory: category });
  },

  setCategories: async (categories) => {
    colorUsedCollection = {};
    const pipeData = (categories || [])
      .map((item) => {
        const options = item.options
          .map((option) => {
            if (option.color) {
              colorUsedCollection[option.color] = option.color;
            }

            return {
              ...option,
              zIndex: option.zIndex ?? 0,
              color: option.color || item.color,
              order: option.order ?? Number.MAX_SAFE_INTEGER,
              type: option.type ?? DEFAULT_CATEGORY_TYPE,
              multipleChoice: option.multipleChoice ?? true,
              onLinkValidate: option.onLinkValidate || item.onLinkValidate,
              onSnapValidate: option.onSnapValidate || item.onSnapValidate,
              onSplitValidate: option.onSplitValidate || item.onSplitValidate,
              onMergeValidate: option.onMergeValidate || item.onMergeValidate,
              onDropInValidate: option.onDropInValidate || item.onDropInValidate,
              onDropOutValidate: option.onDropOutValidate || item.onDropOutValidate,
              onAddValidate: option.onAddValidate || item.onAddValidate,
            };
          })
          .sort((a, b) => a.order - b.order);

        if (item.color) {
          colorUsedCollection[item.color] = item.color;
        }

        return {
          ...item,
          options,
          order: item.order ?? Number.MAX_SAFE_INTEGER,
          multipleOption: item.multipleOption ?? true,
        };
      })
      .map((item) => {
        // generate color for category
        return {
          ...item,
          color: pickColorFromPalettes(0, item.color),
        };
      })
      .map((item) => {
        // generate color for option category
        const options = item.options.map((option) => {
          return {
            ...option,
            color: pickColorFromPalettes(0, option.color || item.color),
          };
        });

        return {
          ...item,
          options,
        };
      })
      .sort((a, b) => a.order - b.order) as StudioCategory[];

    const rootCategory = pipeData?.find((item) => item.isRoot);

    set({
      rootCategory,
      categories: pipeData,
      ...flatCategoriesByKey(pipeData),
    });
  },

  setFilters: (filter) => {
    const filters = get().filters;
    const matchedFilter = filters.find((item) => item === filter);
    if (matchedFilter) {
      set({ filters: filters.filter((item) => item !== filter) });
    } else {
      set({ filters: [...filters, filter] });
    }
  },

  updateCategory: (category) => {
    // just update category for render
    const { categories } = get();
    const newCategories = categories.map((item) => {
      if (item.idx === category.idx) {
        return category;
      }

      return item;
    });

    set({ categories: newCategories });
  },

  updateCategoriesForEntry: (entry) => {
    const { rootCategory, categoryMap, categories } = get();

    if (rootCategory) {
      if (entry) {
        const newCategories = categories.map((item) => {
          if (item.idx === rootCategory.idx) {
            return {
              ...item,
              disabled: true,
            };
          }

          return {
            ...item,
            disabled: categoryMap[item.idx]?.disabled ?? false,
          };
        });

        set({ categories: newCategories });
      } else {
        const newCategories = categories.map((item) => {
          if (item.idx === rootCategory.idx) {
            return {
              ...item,
              disabled: false,
            };
          } else {
            return {
              ...item,
              disabled: true,
            };
          }
        });

        set({ categories: newCategories });
      }
    }
  },

  setUsedKeyCollection: (collection) => {
    // const { usedKeyCollection } = get();
    // set({ usedKeyCollection: { ...usedKeyCollection, ...collection } });
    // const { usedKeyCollection } = get();
    set({ usedKeyCollection: collection });
  },

  findCategoryByOptionKey: (optionKey: string) => {
    const { categoryOptionMap } = get();

    return categoryOptionMap[optionKey]?.parent;
  },

  clear: () => set(DEFAULT_VALUE),

  scanFromData: (data) => {
    try {
      const { categoryMap, categoryOptionMap } = get();
      const addNewCategories: Record<string, StudioCategory> = {};
      const addNewOptions: Record<string, StudioCategoryOptionMapValue> = {};

      const addNewCategoryAndOption = (optionIdx: string, categoryIdx?: string) => {
        if (categoryIdx) {
          const existedCategory = categoryMap[categoryIdx];
          if (!existedCategory) {
            // add new category
            addNewCategories[categoryIdx] = {
              idx: categoryIdx,
              title: '',
              options: [],
              disabled: true,
              hidden: true,
            } satisfies StudioCategory;
          }

          if (optionIdx) {
            const parentCategory = categoryMap[categoryIdx] || addNewCategories[categoryIdx];
            const existedOption = categoryOptionMap[optionIdx];
            if (!existedOption) {
              // add new option
              addNewOptions[optionIdx] = {
                idx: optionIdx,
                title: '',
                disabled: true,
                hidden: true,
                parent: parentCategory,
              } satisfies StudioCategoryOptionMapValue;

              parentCategory.options.push(addNewOptions[optionIdx]);
            }
          }
        }
      };

      const runLoopData = (dataNodes: StudioDataNode[]) => {
        dataNodes.forEach((item) => {
          addNewCategoryAndOption(item.idx, item.categoryIdx);

          if (item.children.length) {
            runLoopData(item.children);
          }
        });
      };

      if (data) {
        runLoopData(data);

        set({
          categories: [...get().categories, ...Object.values(addNewCategories)],
          categoryMap: { ...categoryMap, ...addNewCategories },
          categoryOptionMap: { ...categoryOptionMap, ...addNewOptions },
        });
      }
    } catch (error) {
      console.error('Error while scanning data', error);
    }
  },
}));

export default useStudioCategoryStore;
