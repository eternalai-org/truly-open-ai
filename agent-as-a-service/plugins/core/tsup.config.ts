import { defineConfig } from "tsup";

export default defineConfig({
  entry: ["src/index.ts"],
  format: ["esm", "cjs"],
  target: "node18",
  platform: "node",
  dts: true,
  clean: true,
  external: [],
});
