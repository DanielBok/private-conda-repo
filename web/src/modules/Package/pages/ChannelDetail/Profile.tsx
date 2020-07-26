import { PkgSelector } from "@/features/package";
import { Card } from "antd";
import React from "react";
import { useSelector } from "react-redux";
import styles from "./styles.less";

export default () => {
  const { channel, email, joinDate } = useSelector(
    PkgSelector.channelPackages
  );

  const subject = encodeURI(`Hello ${channel}`);

  return (
    <Card
      title={<span className={styles.title}>Profile</span>}
      className={styles.card}
    >
      <div className={styles.profileDiv}>
        <span className={styles.channelName}>{channel}</span>
      </div>
      <div className={styles.profileDiv}>
        <b>Contributor</b> since {joinDate.format("MMM DD, YYYY")}
      </div>
      <div className={styles.profileDiv}>
        <span>Email: </span>{" "}
        <a href={`mailto:${email}?subject=${subject}`} target="_top">
          {email}
        </a>
      </div>
    </Card>
  );
};
