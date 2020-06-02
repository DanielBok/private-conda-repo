import ErrorPage from "@/components/ErrorBoundary/ErrorPage";
import React from "react";
import { Route, Switch } from "react-router";
import { oc } from "ts-optchain";

import routeMap from "./routes";

const moduleMap = Object.entries(routeMap).reduce(
  (acc, [prefix, { routes, ...rest }]) => ({
    ...acc,
    [prefix]: {
      ...rest,
      routes: routes.map(({ path, exact, ...r }) => ({
        ...r,
        exact: oc(exact)(true),
        path: (prefix + path).replace("//{2,}/g", "'/"),
      })),
    },
  }),
  {} as Record<string, ModuleRoutes>
);

const routeList = Object.values(moduleMap).flatMap((module) =>
  module.routes.map((r, i) => (
    <Route path={r.path} exact={r.exact} component={r.component} key={i} />
  ))
);

export default () => (
  <Switch>
    {routeList}
    <Route path="/not-found" component={() => <ErrorPage />} />
    <Route component={() => <ErrorPage />} />
  </Switch>
);
