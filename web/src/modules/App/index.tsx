import ErrorBoundary from "@/components/ErrorBoundary";
import React from "react";
import Layout from "./Layout";
import RouteSwitch from "./RouteSwitch";

export default () => (
  <Layout>
    <ErrorBoundary>
      <RouteSwitch />
    </ErrorBoundary>
  </Layout>
);
