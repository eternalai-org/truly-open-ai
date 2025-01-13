export const removeItemFromArray = <T>(array: T[], item: T) => {
  return array.filter((i) => JSON.stringify(i) !== JSON.stringify(item));
};

export const removeItemsFromArray = <T>(array: T[], items: T[]) => {
  return array.filter((i) => !items.includes(i));
};

export const noUndefinedElement = <T>(array: (T | undefined)[]): T[] => {
  return array.filter((i) => i !== undefined) as T[];
};

export const noNullElement = <T>(array: (T | null)[]): T[] => {
  return array.filter((i) => i !== null) as T[];
};

export const interactionsByKeys = <T>(array1: T[], array2: T[], key: keyof T) => {
  return array1.filter((item) => array2.some((i) => i[key] === item[key]));
};
