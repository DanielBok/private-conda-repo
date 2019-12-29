import React from "react";
import { usePackageContext } from "../hooks";
import styles from "./styles.less";

export default () => {
  const { channel, pkg } = usePackageContext();

  return (
    <div className={styles.title}>
      <span className={styles.channel}>{channel}</span> / {pkg}
    </div>
  );
};
