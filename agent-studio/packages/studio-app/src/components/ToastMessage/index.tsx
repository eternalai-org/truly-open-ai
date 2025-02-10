import React from "react";
import s from "./styles.module.scss";

interface IProps {
  message: string;
}

const ToastMessage: React.FC<IProps> = ({
  message,
}: IProps): React.ReactElement => {
  return <div className={s.wrapper}>{message}</div>;
};

export default ToastMessage;
