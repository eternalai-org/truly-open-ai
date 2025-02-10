import { findDataByCategoryKey } from "@agent-studio/studio-dnd";
import {
  AGENT_CATEGORY_KEY,
  PERSONALITY_CATEGORY_KEY,
} from "../constants/category-keys";
import { CATEGORY_OPTION_KEYS } from "../constants/category-option-keys";
import _ from "lodash";

export const BEGIN_OF_GREETING = "You are ";
export const END_OF_GREETING = "!";

export const getAgentNameFromPersonality = (personality: string) => {
  const nameEnd = personality.indexOf(END_OF_GREETING);

  let name = personality.substring(0, nameEnd);
  name = name.replace(BEGIN_OF_GREETING, "");

  return name;
};

export const findById = (data: any[], id: string): any => {
  let result;

  const iter = (item: any): boolean => {
    if (item.id === id) {
      result = item;
      return true;
    }
    return _.some(item.children, iter);
  };
  _.some(data, iter);
  return result;
};

export const findParentById = (
  data: any[],
  id: string,
  parent: any = null
): any => {
  for (const item of data) {
    if (item.id === id) {
      return parent;
    }
    if (item.children) {
      const found = findParentById(item.children, id, item);
      if (found) {
        return found;
      }
    }
  }
  return null;
};

export const findByCategoryIdx = (data: any[], categoryIdx: string): any => {
  let result;
  const iter = (item: any): boolean => {
    if (item.categoryIdx === categoryIdx) {
      result = item;
      return true;
    }
    return _.some(item.children, iter);
  };
  _.some(data, iter);
  return result;
};

export const findPersonalityObj = (data: any, parentObj?: any): any => {
  if (!parentObj) {
    const foundMainPersonality = findDataByCategoryKey(
      AGENT_CATEGORY_KEY,
      data
    );

    return foundMainPersonality[0]?.children.find(
      (child: any) => child.categoryIdx === PERSONALITY_CATEGORY_KEY
    );
  }

  return parentObj?.idx === CATEGORY_OPTION_KEYS.agent.agent_new
    ? // find children of parentObj that has idx = "personality_customize"
      parentObj?.children.find(
        (child: any) => child.categoryIdx === PERSONALITY_CATEGORY_KEY
      )
    : parentObj;
};
