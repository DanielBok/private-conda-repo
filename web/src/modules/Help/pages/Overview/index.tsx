import Markdown from "@/components/Markdown";
import { metaInfo } from "@/features/meta/selectors";
import { Typography } from "antd";
import React from "react";
import { useSelector } from "react-redux";
import Layout from "../../Layout";

import Overview from "./overview.md";

export default () => {
  const { registry } = useSelector(metaInfo);
  return (
    <Layout>
      <Markdown
        children={Overview}
        overrides={{
          RegistryInfo: () => (
            <Typography.Paragraph code={true}>
              pcr registry set {registry}
            </Typography.Paragraph>
          ),
        }}
      />
    </Layout>
  );
};
