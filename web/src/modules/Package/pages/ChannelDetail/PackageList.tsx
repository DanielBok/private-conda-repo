import { PkgSelector } from "@/features/package";
import { timeSinceUpload } from "@/libs/date";
import Logo from "@/resource/conda.svg";
import QuestionCircleOutlined from "@ant-design/icons/QuestionCircleOutlined";
import { Card, List } from "antd";
import React from "react";
import { useSelector } from "react-redux";
import { Link } from "react-router-dom";
import styles from "./styles.less";

export default () => {
  const { channel, packages } = useSelector(PkgSelector.channelPackages);

  return (
    <Card
      title={
        <span className={styles.title}>
          <a
            className={styles.iconLink}
            href="https://docs.anaconda.com/anaconda-cloud/user-guide/tasks/work-with-packages/"
          >
            <QuestionCircleOutlined />
          </a>
          Packages
        </span>
      }
      className={styles.card}
    >
      <List
        itemLayout="horizontal"
        dataSource={packages}
        renderItem={(item) => (
          <List.Item className={styles.listItem}>
            <img src={Logo} alt="" />
            <Link to={`/p/${channel}/${item.name}`}>{item.name}</Link>
            <span>Updated {timeSinceUpload(item.timestamp)}</span>
          </List.Item>
        )}
      />
    </Card>
  );
};
