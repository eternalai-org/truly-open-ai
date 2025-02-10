import styles from "./styles.module.scss";

const LoadingRow = () => (
  <div className={styles.lds_ellipsis}>
    <div></div>
    <div></div>
    <div></div>
    <div></div>
  </div>
);

export default LoadingRow;
