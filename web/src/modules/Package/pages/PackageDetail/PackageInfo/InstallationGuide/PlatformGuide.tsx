import { PkgSelector } from "@/features/package";
import AndroidOutlined from "@ant-design/icons/AndroidOutlined";
import AppleOutlined from "@ant-design/icons/AppleOutlined";
import WindowsOutlined from "@ant-design/icons/WindowsOutlined";
import React from "react";
import { useSelector } from "react-redux";
import styles from "./styles.less";

export default () => {
  const { latest } = useSelector(PkgSelector.packageDetail);

  return (
    <div className={styles.platform}>
      <PlatformTags />
      {latest.version}
    </div>
  );
};

const PlatformTags = () => {
  const order = {
    windows: 2,
    apple: 1,
    android: 3,
    noarch: 4,
  };

  const platforms = usePlatforms()
    .map((e) => {
      switch (e) {
        case "win-64":
          return "windows";
        case "osx-64":
          return "apple";
        case "linux-64":
          return "android";
        default:
          return "noarch";
      }
    })
    .sort((x, y) => order[x] - order[y])
    .map((e, i) => {
      switch (e) {
        case "windows":
          return <WindowsOutlined key={i} />;
        case "apple":
          return <AppleOutlined key={i} />;
        case "android":
          return <AndroidOutlined key={i} />;
        default:
          return <span key={i}>noarch</span>;
      }
    });

  return <span className={styles.tags}>{platforms}</span>;
};

const usePlatforms = () => {
  const { details, latest } = useSelector(PkgSelector.packageDetail);
  const platforms = details
    .filter((d) => d.version === latest.version)
    .map((e) => e.platform);

  if (platforms.length === 0) return latest.platforms;
  return platforms;
};
