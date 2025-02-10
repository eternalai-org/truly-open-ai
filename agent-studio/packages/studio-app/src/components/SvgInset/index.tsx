import { ReactSVG } from "react-svg";
import React from "react";
import px2rem from "../../types/px2rem";

type IProps = {
  svgUrl: string;
  className?: string;
  size?: number;
  height?: number;
  fullWidth?: boolean;
  onClick?: () => void;
  style?: React.CSSProperties;
};

const SvgInset: React.FC<IProps> = ({
  svgUrl,
  className,
  size,
  height,
  onClick,
  fullWidth = false,
  style,
}) => {
  return (
    <ReactSVG
      onClick={onClick}
      className={className}
      src={svgUrl}
      style={style}
      beforeInjection={(svg): void => {
        if (size) {
          svg.setAttribute("height", `100%`);
          svg.setAttribute("width", `${fullWidth ? "100%" : px2rem(size)}`);
          svg.style.maxWidth = `${size}`;
          svg.style.minHeight = `100%`;
          svg.style.width = `${fullWidth ? "100%" : px2rem(size)}`;
          svg.style.height = height ? px2rem(height) : `100%`;
        }
      }}
    />
  );
};

export default SvgInset;
