import { SystemStyleObject } from "@chakra-ui/react";

export const TextStyleMap: Record<string, SystemStyleObject> = {
  FORM_HEADING_STYLE: {
    fontSize: "14px",
    lineHeight: "calc(20 / 14)",
    fontWeight: "600",
    fontFamily: "var(--font-SFProDisplay)",
    color: "#ffffff",
  },

  FORM_HEADING_TAG_STYLE: {
    fontSize: "14px",
    lineHeight: "calc(20 / 14)",
    fontWeight: "600",
    fontFamily: "var(--font-SFProDisplay)",
    color: "#ffffff",
    backgroundColor: "#e18326",
    borderRadius: "999px",
    padding: "4px 8px",
    width: "fit-content",
  },

  LABEL_STYLE: {
    fontSize: "18px",
    lineHeight: "calc(28 / 18)",
    fontWeight: "500",
    fontFamily: "var(--font-SFProDisplay)",
    color: "white",
  },

  SIDEBAR_CATEGORY_OPTION_LABEL_STYLE: {
    fontSize: "18px",
    lineHeight: "calc(28 / 18)",
    fontWeight: "500",
    fontFamily: "var(--font-SFProDisplay)",
    color: "white",
  },

  INPUT_STYLE: {
    fontSize: "18px",
    lineHeight: "calc(28 / 18)",
    fontWeight: "500",
    fontFamily: "var(--font-SFProDisplay)",
    color: "black",
  },

  INPUT_PLACEHOLDER_STYLE: {
    color: "#777777",
    fontSize: "18px",
    lineHeight: "calc(28 / 18)",
    fontWeight: "500",
    fontFamily: "var(--font-SFProDisplay)",
  },

  TEXTAREA_INPUT_STYLE: {
    fontSize: "14px",
    lineHeight: "calc(20 / 14)",
    fontWeight: "500",
    fontFamily: "var(--font-SFProDisplay)",
    color: "black",
  },

  TEXTAREA_PLACEHOLDER_STYLE: {
    color: "#777777",
    fontSize: "14px",
    lineHeight: "calc(20 / 14)",
    fontWeight: "500",
    fontFamily: "var(--font-SFProDisplay)",
  },

  TOOLTIP_STYLE: {
    fontSize: "14px",
    lineHeight: "calc(20 / 14)",
    fontWeight: "500",
    fontFamily: "var(--font-SFProDisplay)",
  },

  ERROR_STYLE: {
    height: "12px",
    fontSize: "12px",
    lineHeight: "1",
    fontWeight: "500",
    fontFamily: "var(--font-SFProDisplay)",
  },
};
