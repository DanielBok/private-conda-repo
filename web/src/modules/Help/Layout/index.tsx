import { Col, Row } from "antd";
import React, { FC } from "react";
import styles from "./styles.less";

import Menu from "./Menu";

const Layout: FC = ({ children }) => (
  <div className={styles.main}>
    <Row gutter={24}>
      <Col span={4}>
        <Menu />
      </Col>
      <Col span={20}>{children}</Col>
    </Row>
  </div>
);

export default Layout;
