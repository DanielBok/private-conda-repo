import ErrorBoundary from "@/components/ErrorBoundary";
import RouteSwitch from "./RouteSwitch";
import React from "react";
import Layout from "./Layout";

export default () => (
  <Layout>
    <ErrorBoundary>
      <RouteSwitch />
    </ErrorBoundary>
  </Layout>
);
