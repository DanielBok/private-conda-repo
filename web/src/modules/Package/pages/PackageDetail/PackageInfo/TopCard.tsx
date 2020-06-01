import { PackageSelector } from "@/features/package";
import { timeSinceUpload } from "@/libs/date";
import { Card } from "antd";
import ProfileOutlined from "@ant-design/icons/ProfileOutlined";
import DownloadOutlined from "@ant-design/icons/DownloadOutlined";
import CalendarOutlined from "@ant-design/icons/CalendarOutlined";
import HomeOutlined from "@ant-design/icons/HomeOutlined";
import CodeOutlined from "@ant-design/icons/CodeOutlined";
import FileWordOutlined from "@ant-design/icons/FileWordOutlined";
import React from "react";
import { useSelector } from "react-redux";
import styles from "./styles.less";

export default () => {
  const details = useSelector(PackageSelector.packageDetail);
  const { devUrl, docUrl, home, license, timestamp } = details.latest;
  const downloads = details.details.reduce((acc, e) => acc + e.count, 0);

  return (
    <Card className={styles.mainCard}>
      {license && (
        <div className={styles.topCard}>
          <ProfileOutlined />
          <span>License: </span>
          <a
            href={`https://opensource.org/search/node/${license}`}
            rel="noopener noreferrer"
            target="_blank"
          >
            {license}
          </a>
        </div>
      )}

      {([
        [home, "Home", <HomeOutlined />],
        [devUrl, "Development", <CodeOutlined />],
        [docUrl, "Documentation", <FileWordOutlined />],
      ] as [string | null, string, JSX.Element][]).map(
        ([link, title, icon]) =>
          link && (
            <div className={styles.topCard} key={title}>
              {icon}
              <span>{title}: </span>
              <a href={link} rel="noopener noreferrer" target="_blank">
                {link}
              </a>
            </div>
          )
      )}

      <div className={styles.topCard}>
        <DownloadOutlined />
        <span>{downloads} total downloads</span>
      </div>

      <div className={styles.topCard}>
        <CalendarOutlined />
        <span>Last Upload: {timeSinceUpload(timestamp)}</span>
      </div>
    </Card>
  );
};
