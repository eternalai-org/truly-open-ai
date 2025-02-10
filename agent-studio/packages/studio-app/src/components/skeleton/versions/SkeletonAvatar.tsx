import React from "react";
import Skeleton from "react-loading-skeleton";
import "react-loading-skeleton/dist/skeleton.css";
import { BASE_COLOR, HIGHLIGHT_COLOR } from "./constants";

interface SkeletonPropsAvatar {
  width?: string;
  height?: string;
}

const SkeletonAvatar = ({
  width = "44px",
  height = "44px",
}: SkeletonPropsAvatar) => {
  return (
    <Skeleton
      baseColor={BASE_COLOR}
      highlightColor={HIGHLIGHT_COLOR}
      width={width}
      height={height}
      borderRadius="100%"
    />
  );
};

export default SkeletonAvatar;
