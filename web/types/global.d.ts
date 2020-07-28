declare module "*.less";
declare module "*.png";
declare module "*.svg";
declare module "*.md" {
  /**
   * Replaces text in a string, using a regular expression or search string.
   * @param searchValue A string to search for.
   * @param replaceValue A string containing the text to replace for every successful match of searchValue in this string.
   */
  function replace(searchValue: string | RegExp, replaceValue: string): this;
}

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

type DeepPartial<T> = {
  [K in keyof T]?: T[K] extends Array<infer R>
    ? Array<DeepPartial<R>>
    : DeepPartial<T[K]>;
};
