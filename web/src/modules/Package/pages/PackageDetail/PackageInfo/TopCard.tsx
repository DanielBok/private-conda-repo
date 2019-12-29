import { PackageSelector } from "@/features/package";
import { Card, Icon } from "antd";
import moment from "moment";
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

function timeSinceUpload(timestamp: number) {
  const duration = moment.duration(moment().diff(moment.unix(timestamp)));

  const values = [
    ["years", duration.years()],
    ["months", duration.months()],
    ["days", duration.days()],
    ["hours", duration.hours()]
  ]
    .filter(([_, v]) => v !== 0)
    .slice(0, 2) as [string, number][];

  switch (values.length) {
    case 0:
      return "Just now";
    case 1:
      const v = values[0];
      return `${formString(v[0], v[1])} ago`;
    default:
      const [first, second] = values,
        s1 = formString(first[0], first[1]),
        s2 = formString(second[0], second[1]);
      return `${s1} and ${s2} ago`;
  }

  function formString(unit: string, value: number) {
    if (value === 1) unit = unit.slice(0, -1);
    return `${value} ${unit}`;
  }
}
