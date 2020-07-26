import { PkgSelector } from "@/features/package";
import React from "react";
import { useSelector } from "react-redux";
import { Link } from "react-router-dom";
import { usePackageContext } from "../hooks";
import styles from "./styles.less";

export default () => {
  const { channel, pkg } = usePackageContext();
  const { summary } = useSelector(PkgSelector.packageDetail).latest;

  return (
    <div className={styles.title}>
      <div>
        <Link to={`/p/${channel}`}>
          <span className={styles.channel}>{channel}</span>
        </Link>{" "}
        / {pkg}
      </div>
      {summary && <div className={styles.subtitle}>{summary}</div>}
    </div>
  );
};
