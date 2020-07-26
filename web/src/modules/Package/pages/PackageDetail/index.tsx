import { PkgApi } from "@/features/package";
import { RootState } from "@/infrastructure/rootState";
import { Tabs } from "antd";
import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { Redirect, RouteComponentProps, withRouter } from "react-router-dom";
import Files from "./Files";
import Header from "./Header";
import { PackageContext } from "./hooks";
import PackageInfo from "./PackageInfo";
import styles from "./styles.less";
import { MatchParams } from "./types";

const { TabPane } = Tabs;

type Props = RouteComponentProps<MatchParams>;

const PackageDetail = ({
  match: {
    params: { channel, pkg },
  },
}: Props) => {
  const dispatch = useDispatch();
  const [tab, setTab] = useState<"conda" | "files">("conda");

  useEffect(() => {
    dispatch(PkgApi.fetchPackageDetail(channel, pkg));
  }, [channel, pkg, dispatch]);

  const failure = useSelector(
    (s: RootState) => s.pkg.loading.details === "FAILURE"
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
          onChange={(e) => setTab(e as typeof tab)}
          className={styles.tabBar}
          tabBarStyle={{
            backgroundColor: "rgba(63,165,39,.3)",
          }}
        >
          <TabPane tab="Conda" key="conda">
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
