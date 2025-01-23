import { memo } from "react";
import styles from "./styles.module.scss";
import { TextStyleMap } from "../../styles";

type Props = {
  value: string;
  onChange: (value: string) => void;
};

const StudioDateTimeInput = ({ value, onChange }: Props) => {
  return (
    <input
      type="datetime-local"
      value={value}
      onChange={(e) => onChange(e.target.value)}
      className={styles.input}
      style={{
        ...(TextStyleMap.INPUT_STYLE as any),
      }}
    />
  );
};

export default memo(StudioDateTimeInput);
