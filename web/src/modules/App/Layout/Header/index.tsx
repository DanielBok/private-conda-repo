import Logo from "@/resource/conda.svg";
import { Layout } from "antd";
import { push } from "connected-react-router";
import React from "react";
import { useDispatch } from "react-redux";
import UserManager from "./Manager";

import styles from "./styles.less";

export default () => {
  const dispatch = useDispatch();

  return (
    <Layout.Header className={styles.header}>
      <div className={styles.title} onClick={() => dispatch(push("/"))}>
        <img src={Logo} alt="PCR Logo" className={styles.logo} />
        <span className={styles.text}>
          Private Conda <span className={styles.nonBold}>Repository</span>
        </span>
      </div>
      <UserManager />
    </Layout.Header>
  );
};
