import React from "react";

import Markdown from "@/components/Markdown";
import Layout from "../../Layout";

import Overview from "./overview.md";

export default () => (
  <Layout>
    <Markdown children={Overview} />
  </Layout>
);
