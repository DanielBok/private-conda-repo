import QuestionCircleFilled from "@ant-design/icons/QuestionCircleFilled";
import { Card, Tooltip } from "antd";
import React from "react";
import CommandGuide from "./CommandGuide";
import PlatformGuide from "./PlatformGuide";
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
        <QuestionCircleFilled />
      </Tooltip>
    </div>

    <PlatformGuide />
    <CommandGuide />
  </Card>
);
