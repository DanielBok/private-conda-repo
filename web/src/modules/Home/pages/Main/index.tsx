import React from "react";
import PackageList from "./PackageList";
import styles from "./styles.less";

export default () => {
  return (
    <div className={styles.container}>
      <PackageList />
    </div>
  );
};
