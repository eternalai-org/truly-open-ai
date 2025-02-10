import { BASE_COLOR, HIGHLIGHT_COLOR } from "./constants";
import React from "react";
import Skeleton from "react-loading-skeleton";
import "react-loading-skeleton/dist/skeleton.css";

const SkeletonLongText = ({ height }: { height?: string }) => {
  return (
    <Skeleton
      baseColor={BASE_COLOR}
      highlightColor={HIGHLIGHT_COLOR}
      width="200px"
      height={height}
    />
  );
};

export default SkeletonLongText;
