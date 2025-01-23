import globals from "globals";
import js from "@eslint/js";
import ts from "@typescript-eslint/eslint-plugin";
import tsParser from "@typescript-eslint/parser";
import react from "eslint-plugin-react";

/** @type {import('eslint').Linter.FlatConfig[]} */
export default [
  {
    files: ["**/*.{js,jsx,ts,tsx,mjs,cjs}"],
    ignorePatterns: [
      "node_modules",
      "**/dist/**",
      "scripts/**",
      "client/browser/**"
    ],
    settings: {
      "react": {
        "version": "detect"
      }
    },
    languageOptions: {
      parser: tsParser,
      parserOptions: {
        ecmaVersion: "latest",
        sourceType: "module",
        project: "./tsconfig.json",
      },
      globals: {
        ...globals.browser,
        ...globals.node,
      },
    },
    plugins: {
      "@typescript-eslint": ts,
      react,
    },
    "extends": [
      "eslint:recommended",
      "plugin:@typescript-eslint/recommended",
      "plugin:react/recommended"
    ],
    rules: {
      ...js.configs.recommended.rules,
      ...ts.configs.recommended.rules,

      // Custom rules
      "object-curly-spacing": ["error", "always"],
      "indent": ["error", 2],
      "quotes": ["error", "double"],
      "semi": ["error", "always"],
      "@typescript-eslint/no-non-null-asserted-optional-chain": "off",
      "@typescript-eslint/no-empty-object-type": "off",
      "import/no-anonymous-default-export": "off",
      "@typescript-eslint/no-explicit-any": "off",
      "no-useless-catch": "off",
      "@typescript-eslint/no-unsafe-function-type": "off",
      "no-constant-binary-expression": "off",
      "react/react-in-jsx-scope": "off"
    },
  },
];