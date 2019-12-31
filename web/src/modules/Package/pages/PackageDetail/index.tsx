import { PackageApi } from "@/features/package";
import { RootState } from "@/infrastructure/rootState";
import { Tabs } from "antd";
import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { Redirect, RouteComponentProps, withRouter } from "react-router-dom";
import Header from "./Header";
import { PackageContext } from "./hooks";
import PackageInfo from "./PackageInfo";
import styles from "./styles.less";
import { MatchParams } from "./types";
import Files from "./Files";

const { TabPane } = Tabs;

type Props = RouteComponentProps<MatchParams>;

const PackageDetail = ({
  match: {
    params: { channel, pkg }
  }
}: Props) => {
  const dispatch = useDispatch();
  const [tab, setTab] = useState<"conda" | "files">("conda");

  useEffect(() => {
    dispatch(PackageApi.fetchPackageDetail(channel, pkg));
  }, [channel, pkg, dispatch]);

  const failure = useSelector(
    (s: RootState) => s.package.loading.details === "FAILURE"
  );

  if (failure) {
    return <Redirect to="/not-found" />;
  }

  return (
    <PackageContext.Provider value={{ channel, pkg }}>
      <div className={styles.container}>
        <Header />
        <Tabs
          activeKey={tab}
          onChange={e => setTab(e as typeof tab)}
          className={styles.tabBar}
        >
          <TabPane tab={<span>Conda</span>} key="conda">
            <PackageInfo />
          </TabPane>
          <TabPane tab="Files" key="files">
            <Files />
          </TabPane>
        </Tabs>
      </div>
    </PackageContext.Provider>
  );
};

export default withRouter(PackageDetail);
