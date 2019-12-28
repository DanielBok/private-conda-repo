import React from "react";
import { Layout } from "antd";

import styles from "./styles.less";
import Logo from "./logo.png";

export default () => {
  return (
    <Layout.Header className={styles.header}>
      <div className={styles.title}>
        <img src={Logo} alt="PCR Logo" className={styles.logo} />
        <span className={styles.text}>
          Private Conda <span className={styles.nonBold}>Repository</span>
        </span>
      </div>
    </Layout.Header>
  );
};
