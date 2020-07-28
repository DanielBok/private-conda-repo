import Markdown from "@/components/Markdown";
import { useRootSelector } from "@/infrastructure/hooks";
import Layout from "@/modules/Help/Layout";
import React from "react";
import Content from "./overview.md";

const Overview = () => {
  const registry = useRootSelector((s) => s.meta.registry);
  const content = Content.replace("@registry", registry);

  return (
    <Layout>
      <Markdown children={content} />
    </Layout>
  );
};

export default Overview;
