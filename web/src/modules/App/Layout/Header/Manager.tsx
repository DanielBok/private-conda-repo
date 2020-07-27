import { ChnApi, ChnSelector } from "@/features/channel";
import UserOutlined from "@ant-design/icons/UserOutlined";
import { Menu } from "antd";
import React from "react";
import { useDispatch, useSelector } from "react-redux";
import { Link } from "react-router-dom";
import styles from "./styles.less";

export default () => {
  const user = useSelector(ChnSelector.channelInfo);
  const dispatch = useDispatch();
  const validated = user.validated;

  return (
    <Menu mode="horizontal" selectedKeys={[""]}>
      <Menu.Item key="help">
        <Link to="/help">Help</Link>
      </Menu.Item>
      {validated ? (
        <Menu.SubMenu
          title={
            <span className={styles.user}>
              <UserOutlined />
              {user.channel}
            </span>
          }
        >
          <Menu.Item key="Upload">
            <Link to="/upload">Upload Package</Link>
          </Menu.Item>
          <Menu.Item key="logout" onClick={() => dispatch(ChnApi.logout())}>
            <Link to="/">Logout</Link>
          </Menu.Item>
        </Menu.SubMenu>
      ) : (
        <Menu.Item key="login-register">
          <Link to="/account">Login / Register</Link>
        </Menu.Item>
      )}
    </Menu>
  );
};
