import { PropsWithChildren } from "react";

export type StudioFieldBase = PropsWithChildren & {
  label?: string;
  tooltip?: string;
  errorMessage?: string;
  action?: React.ReactNode;
};
