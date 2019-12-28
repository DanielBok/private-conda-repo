import { Layout, Typography } from "antd";
import React from "react";

import styles from "./styles.less";

const { Title, Paragraph } = Typography;

export default () => {
  return (
    <Layout.Footer className={styles.footer}>
      <Title level={3} className={styles.subtitle}>
        Private Conda Repository
      </Title>
      <Paragraph className={styles.text}>
        Private Conda Repository © {getCopyRightYear()}
      </Paragraph>
    </Layout.Footer>
  );
};

function getCopyRightYear() {
  const currentYear = new Date().getFullYear();

  return currentYear === 2019
    ? currentYear.toString()
    : `2019 - ${currentYear}`;
}
