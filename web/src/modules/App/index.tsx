import ErrorBoundary from "@/components/ErrorBoundary";
import { Switch } from "react-router-dom";
import React from "react";
import Layout from "./Layout";

export default () => {
  return (
    <Layout>
      <ErrorBoundary>
        <Switch></Switch>
      </ErrorBoundary>
    </Layout>
  );
};
