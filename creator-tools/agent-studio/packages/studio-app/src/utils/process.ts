export const BEGIN_OF_GREETING = "You are ";
export const END_OF_GREETING = "!";

export const getAgentNameFromPersonality = (personality: string) => {
  const nameEnd = personality.indexOf(END_OF_GREETING);

  let name = personality.substring(0, nameEnd);
  name = name.replace(BEGIN_OF_GREETING, "");

  return name;
};
