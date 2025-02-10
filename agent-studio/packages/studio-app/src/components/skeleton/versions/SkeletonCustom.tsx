import { BASE_COLOR, HIGHLIGHT_COLOR } from "./constants";
import React from "react";
import Skeleton from "react-loading-skeleton";
import "react-loading-skeleton/dist/skeleton.css";

const SkeletonCustom = (props: any) => {
  return (
    <Skeleton
      baseColor={BASE_COLOR}
      highlightColor={HIGHLIGHT_COLOR}
      {...props}
    />
  );
};

export default SkeletonCustom;
