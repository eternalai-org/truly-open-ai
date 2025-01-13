import { XYPosition } from '@xyflow/react';

import useStudioCategoryStore from '../stores/useStudioCategoryStore';
import { FormDataMap } from '../types/base';
import { StudioCategoryOption } from '../types/category';
import { StudioDataNode } from '../types/graph';

export const getFormDataFromCategoryOption = (category: StudioCategoryOption) => {
  const categoryData = category?.data || {};

  const defaultValues = Object.keys(categoryData).reduce(
    (acc, key) => ({
      ...acc,
      [key]: categoryData[key].defaultValue,
    }),
    {},
  );

  return defaultValues;
};

export const getFieldDataFromRawData = (data: StudioDataNode[]) => {
  let formData: Record<string, FormDataMap> = {};
  const categoryMap = useStudioCategoryStore.getState().categoryMap;
  const categoryOptionMap = useStudioCategoryStore.getState().categoryOptionMap;

  data.forEach((item) => {
    const defaultValues = getFormDataFromCategoryOption(categoryOptionMap[item.idx] || categoryMap[item.idx] || {});

    formData[item.id] = {
      ...defaultValues,
      ...(item.data || {}),
    };

    if (item.children.length) {
      formData = {
        ...formData,
        ...getFieldDataFromRawData(item.children),
      };
    }
  });

  return formData;
};

export const cloneData = <T>(data: T): T => JSON.parse(JSON.stringify(data));

export const min = (...args: number[]) => Math.min(...args);
export const max = (...args: number[]) => Math.max(...args);

export const findDataById = (id: string, data: StudioDataNode[]) => {
  for (const item of data) {
    if (item.id === id) {
      return item;
    }
    if (item.children.length > 0) {
      const foundNode = findDataById(id, item.children) as StudioDataNode;
      if (foundNode) {
        return foundNode;
      }
    }
  }

  return null;
};

export const findDataByOptionKey = (optionKey: string, data: StudioDataNode[], nodeId?: string) => {
  let matchedNode;
  if (nodeId) {
    matchedNode = findDataById(nodeId, data);

    if (!matchedNode) {
      return [];
    }
  }

  const dataShouldBeFind = matchedNode ? [matchedNode] : data;

  const returnVal: StudioDataNode[] = [];
  dataShouldBeFind.forEach((dataNode) => {
    if (dataNode.idx === optionKey) {
      returnVal.push(dataNode);
    }
    if (dataNode.children.length) {
      dataNode.children.forEach((child) => {
        returnVal.push(...findDataByOptionKey(optionKey, [child], child.id));
      });
    }
  });

  return returnVal;
};

export const findDataByCategoryKey = (categoryKey: string, data: StudioDataNode[], nodeId?: string) => {
  const categoryMap = useStudioCategoryStore.getState().categoryMap;

  let matchedNode;
  if (nodeId) {
    matchedNode = findDataById(nodeId, data);

    if (!matchedNode) {
      return [];
    }
  }

  const dataShouldBeFind = matchedNode ? [matchedNode] : data;

  const returnVal: StudioDataNode[] = [];

  const categoryOptionKeys = categoryMap[categoryKey].options.map((option) => option.idx);
  categoryOptionKeys.forEach((optionKey) => {
    dataShouldBeFind.forEach((dataNode) => {
      if (dataNode.idx === optionKey) {
        returnVal.push(dataNode);
      }
      if (dataNode.children.length) {
        dataNode.children.forEach((child) => {
          returnVal.push(...findDataByOptionKey(optionKey, [child], child.id));
        });
      }
    });
  });

  return returnVal;
};

export const createNodeData = (
  id: string,
  option: StudioCategoryOption,
  children: StudioDataNode[] = [],
  data: FormDataMap = {},
  position: XYPosition = {
    x: 0,
    y: 0,
  },
): StudioDataNode => {
  return {
    id: id,
    idx: option.idx,
    title: option.title || 'Untitled',
    children: [...children],
    data: {
      ...data,
    },
    rect: {
      position: position,
    },
  } satisfies StudioDataNode;
};
