import { UserApi, UserSelector } from "@/features/user";
import { Icon, Layout, Menu } from "antd";
import React from "react";
import { useDispatch, useSelector } from "react-redux";
import Logo from "./logo.png";

import styles from "./styles.less";

export default () => {
  return (
    <Layout.Header className={styles.header}>
      <div className={styles.title}>
        <img src={Logo} alt="PCR Logo" className={styles.logo} />
        <span className={styles.text}>
          Private Conda <span className={styles.nonBold}>Repository</span>
        </span>
      </div>
      <UserManager />
    </Layout.Header>
  );
};

const UserManager = () => {
  const dispatch = useDispatch();
  const user = useSelector(UserSelector.userInfo);
  if (!user.validated) return null;

  return (
    <Menu mode="horizontal">
      <Menu.SubMenu
        title={
          <span className={styles.user}>
            <Icon type="user" className={styles.userLogo} />
            {user.username}
          </span>
        }
      >
        <Menu.Item onClick={() => dispatch(UserApi.logout())}>Logout</Menu.Item>
      </Menu.SubMenu>
    </Menu>
  );
};
