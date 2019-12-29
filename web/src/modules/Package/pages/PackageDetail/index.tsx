import { PackageApi } from "@/features/package";
import { Tabs } from "antd";
import React, { useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import { RouteComponentProps, withRouter } from "react-router-dom";
import Header from "./Header";
import { PackageContext } from "./hooks";
import styles from "./styles.less";
import { MatchParams } from "./types";

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
            Content of Tab Pane 1
          </TabPane>
          <TabPane tab="Files" key="files">
            Content of Tab Pane 2
          </TabPane>
        </Tabs>
      </div>
    </PackageContext.Provider>
  );
};

export default withRouter(PackageDetail);
