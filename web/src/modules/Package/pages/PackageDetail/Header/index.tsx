import { PackageSelector } from "@/features/package";
import React from "react";
import { useSelector } from "react-redux";
import { usePackageContext } from "../hooks";
import styles from "./styles.less";

export default () => {
  const { channel, pkg } = usePackageContext();
  const { summary } = useSelector(PackageSelector.packageDetail).latest;

  return (
    <div className={styles.title}>
      <div>
        <span className={styles.channel}>{channel}</span> / {pkg}
      </div>
      {summary && <div className={styles.subtitle}>{summary}</div>}
    </div>
  );
};
