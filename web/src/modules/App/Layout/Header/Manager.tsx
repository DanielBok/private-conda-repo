import { ChnApi, ChnSelector } from "@/features/channel";
import UserOutlined from "@ant-design/icons/UserOutlined";
import { Menu } from "antd";
import { push } from "connected-react-router";
import React from "react";
import { useDispatch, useSelector } from "react-redux";
import styles from "./styles.less";

export default () => {
  const user = useSelector(ChnSelector.channelInfo);
  const dispatch = useDispatch();
  const validated = user.validated;

  return (
    <Menu mode="horizontal" selectedKeys={[""]}>
      <Menu.Item
        key="help"
        onClick={() => {
          dispatch(push("/help"));
        }}
      >
        Help
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
          <Menu.Item
            key="logout"
            onClick={() => {
              dispatch(ChnApi.logout());
              dispatch(push("/"));
            }}
          >
            Logout
          </Menu.Item>
        </Menu.SubMenu>
      ) : (
        <Menu.Item
          key="login-register"
          onClick={() => dispatch(push("/account"))}
        >
          Login / Register
        </Menu.Item>
      )}
    </Menu>
  );
};
