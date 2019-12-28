declare module "*.less";
declare module "*.png";

type LoadingState = "REQUEST" | "SUCCESS" | "FAILURE";

type ModuleRoute = {
  component:
    | React.ComponentType<RouteComponentProps<any>>
    | React.ComponentType<any>;
  path: string;
  title: string;
  exact?: boolean;
};

type ModuleRoutes = {
  clusterName: string;
  routes: ModuleRoute[];
};
