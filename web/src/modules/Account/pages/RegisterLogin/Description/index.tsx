import React from "react";
import Logo from "./logo.png";
import styles from "./styles.less";

export default () => (
  <div className={styles.container}>
    <img className={styles.icon} src={Logo} alt="PCR Logo" />
    <div className={styles.title}>
      Private Conda <br />
      <span>Repository</span>
    </div>

    <div className={styles.subtitle}>
      Where your in-house packages are shared
    </div>
  </div>
);
