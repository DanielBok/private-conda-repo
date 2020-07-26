import { PkgType } from "@/features/package";
import { Tag } from "antd";
import React from "react";

type Props = {
  platform: PkgType.Platform;
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
