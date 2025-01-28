"use client";

import { extendTheme, type ThemeConfig } from "@chakra-ui/react";

const config: ThemeConfig = {
  initialColorMode: "light",
  useSystemColorMode: false,
};

const breakpoints = {
  base: "0px",
  sm: "480px",
  md: "768px",
  lg: "992px",
  xl: "1280px",
  "2xl": "1536px",
  max: "1701px",
  "2k": "2048px",
};

const chakraThemes = extendTheme({
  config,
  breakpoints,
});

export default chakraThemes;
