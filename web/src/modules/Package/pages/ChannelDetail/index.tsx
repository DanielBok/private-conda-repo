import { PkgApi } from "@/features/package";
import { RootState } from "@/infrastructure/rootState";
import { Col, Row } from "antd";
import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { Redirect, RouteComponentProps, withRouter } from "react-router-dom";
import PackageList from "./PackageList";
import Profile from "./Profile";
import styles from "./styles.less";

type Props = RouteComponentProps<{ channel: string }>;

const ChannelDetail = ({
  match: {
    params: { channel },
  },
}: Props) => {
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(PkgApi.fetchChannelPackages(channel));
  }, [channel, dispatch]);

  const failure = useSelector(
    (s: RootState) => s.pkg.loading.channelPackages === "FAILURE"
  );

  if (failure) {
    return <Redirect to="/not-found" />;
  }

  return (
    <div className={styles.container}>
      <Row gutter={12}>
        <Col xs={24} md={12}>
          <Profile />
        </Col>
        <Col xs={24} md={12}>
          <PackageList />
        </Col>
      </Row>
    </div>
  );
};

export default withRouter(ChannelDetail);
