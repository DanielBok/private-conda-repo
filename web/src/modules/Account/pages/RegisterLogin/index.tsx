import { Col, Row } from "antd";
import React from "react";

import Description from "./Description";
import Registration from "./Registration";
import styles from "./styles.less";

export default () => (
  <div className={styles.container}>
    <Row className={styles.background}>
      <Col md={12} xs={24}>
        <Description />
      </Col>
      <Col md={12} xs={24}>
        <Registration />
      </Col>
    </Row>
  </div>
);
