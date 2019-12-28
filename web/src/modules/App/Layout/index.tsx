import { Layout } from "antd";
import React, { FC } from "react";
import Footer from "./Footer";
import Header from "./Header";

import styles from "./styles.less";

const { Content } = Layout;

const AppLayout: FC = ({ children }) => (
  <Layout>
    <Header />
    <Content className={styles.content}>
      <main className={styles.body}>{children}</main>
    </Content>
    <Footer />
  </Layout>
);

export default AppLayout;
