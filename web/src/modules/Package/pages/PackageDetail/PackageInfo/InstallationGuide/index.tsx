import { Card, Icon, Tooltip } from "antd";
import React from "react";
import PlatformGuide from "./PlatformGuide";
import CommandGuide from "./CommandGuide";
import styles from "./styles.less";

export default () => (
  <Card className={`${styles.mainCard} ${styles.installGuide}`}>
    <div className={styles.title}>Installers</div>
    <div className={styles.subtitle}>
      conda install{" "}
      <Tooltip
        placement="bottom"
        title="Learn more about the conda install process"
        overlayClassName={styles.tooltip}
      >
        <Icon type="question-circle" theme="filled" />
      </Tooltip>
    </div>

    <PlatformGuide />
    <CommandGuide />
  </Card>
);
