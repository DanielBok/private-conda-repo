import { PackageSelector } from "@/features/package";
import { timeSinceUpload } from "@/libs/date";
import { Card, Icon } from "antd";
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
          <Icon type="profile" />
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
        [home, "Home", "home"],
        [devUrl, "Development", "code"],
        [docUrl, "Documentation", "file-word"]
      ] as [string | null, string, string][]).map(
        ([link, title, icon]) =>
          link && (
            <div className={styles.topCard} key={icon}>
              <Icon type={icon} />
              <span>{title}: </span>
              <a href={link} rel="noopener noreferrer" target="_blank">
                {link}
              </a>
            </div>
          )
      )}

      <div className={styles.topCard}>
        <Icon type="download" />
        <span>{downloads} total downloads</span>
      </div>

      <div className={styles.topCard}>
        <Icon type="calendar" />
        <span>Last Upload: {timeSinceUpload(timestamp)}</span>
      </div>
    </Card>
  );
};
