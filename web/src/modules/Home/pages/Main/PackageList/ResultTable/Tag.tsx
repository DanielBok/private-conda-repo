import { PackageType } from "@/features/package";
import React from "react";
import { Tag } from "antd";

type Props = {
  platform: PackageType.Platform;
};

export default ({ platform }: Props) => {
  return <Tag color={color()}>{platform}</Tag>;

  function color() {
    switch (platform) {
      case "linux-64":
        return "volcano";
      case "noarch":
        return "gold";
      case "osx-64":
        return "magenta";
      case "win-64":
        return "blue";
      default:
        return "grey";
    }
  }
};
