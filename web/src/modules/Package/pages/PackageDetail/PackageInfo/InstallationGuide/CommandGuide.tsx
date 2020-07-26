import { MetaSelector } from "@/features/meta";
import { PkgSelector } from "@/features/package";
import { Typography } from "antd";
import React from "react";
import { useSelector } from "react-redux";
import styles from "./styles.less";

export default () => {
  const { channel, package: pkg } = useSelector(PkgSelector.packageDetail);
  const { repository } = useSelector(MetaSelector.metaInfo);

  return (
    <div className={styles.command}>
      <div>To install this package with conda run:</div>
      <Typography.Text code={true} copyable={true}>
        conda install -c {repository}/{channel} {pkg}
      </Typography.Text>
    </div>
  );
};
