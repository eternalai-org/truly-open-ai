export const getHourFromSecond = (second: number) => {
  return second / (60 * 60);
};

export const getSecondsFromHour = (hour: number) => {
  return Math.floor(hour * 60 * 60);
};

export const sleep = (second: number) => {
  return new Promise((resolve) => setTimeout(resolve, second * 1000));
};
