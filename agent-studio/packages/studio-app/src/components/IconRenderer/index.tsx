import React, { FunctionComponent } from "react";

type Props = {
  icon: string | React.ReactNode | FunctionComponent;
};

const IconRenderer = (props: Props) => {
  const { icon } = props;

  if (typeof icon === "string") {
    return <img width={20} height={20} src={icon} alt="" />;
  }

  if (typeof icon === "function") {
    return icon(props) as any;
  }

  return icon as any;
};

export default IconRenderer;
