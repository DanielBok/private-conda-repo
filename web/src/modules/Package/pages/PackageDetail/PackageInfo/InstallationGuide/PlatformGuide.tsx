import { PackageSelector } from "@/features/package";
import { Icon } from "antd";
import { useSelector } from "react-redux";
import styles from "./styles.less";
import React from "react";

export default () => {
  const { latest } = useSelector(PackageSelector.packageDetail);

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
    noarch: 4
  };

  const platforms: (keyof typeof order)[] = usePlatforms()
    .map(e => {
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
    .sort((x, y) => order[x] - order[y]);

  return (
    <span className={styles.tags}>
      {platforms.map(p => (p === "noarch" ? p : <Icon type={p} key={p} />))}
    </span>
  );
};

const usePlatforms = () => {
  const { details, latest } = useSelector(PackageSelector.packageDetail);
  const platforms = details
    .filter(d => d.version === latest.version)
    .map(e => e.platform);

  if (platforms.length === 0) return latest.platforms;
  return platforms;
};
