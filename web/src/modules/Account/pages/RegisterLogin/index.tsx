import { UserSelector } from "@/features/user";
import { Col, Row } from "antd";
import { push } from "connected-react-router";
import React from "react";
import { useDispatch, useSelector } from "react-redux";

import Description from "./Description";
import Registration from "./Registration";
import styles from "./styles.less";

export default () => {
  const dispatch = useDispatch();
  if (useSelector(UserSelector.userValidated)) {
    dispatch(push("/"));
  }

  return (
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
};
