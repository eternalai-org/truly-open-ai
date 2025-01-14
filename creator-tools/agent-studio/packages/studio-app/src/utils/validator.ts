export const isEmptyString = (value?: string | null | unknown) => {
  return typeof value === "string" && value?.trim() === "";
};
