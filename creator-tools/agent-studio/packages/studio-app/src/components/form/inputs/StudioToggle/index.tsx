import { memo } from "react";
import cn from "classnames";

import styles from "./styles.module.scss";

type Props = {
  active?: boolean;
  onToggle?: (active: boolean) => void;
};

const StudioToggle = ({ active, onToggle }: Props) => {
  return (
    <div
      className={cn(styles.toggle, {
        [styles.toggle__active]: active,
      })}
      onClick={() => onToggle?.(!active)}
    >
      <div className={styles.toggle_circle} />
    </div>
  );
};

export default memo(StudioToggle);
