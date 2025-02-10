export const prettyPrintBytes = (size: number): string => {
  const units = ["B", "KB", "MB"];
  let s = size;
  for (const unit of units) {
    if (s < 1000) {
      return s.toFixed(0) + unit;
    }
    s /= 1024;
  }
  return s.toFixed(0) + "GB";
};
