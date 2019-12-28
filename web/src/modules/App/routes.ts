import homeRoutes from "@/modules/Home";

const routeMap: Record<string, ModuleRoutes> = {
  "/": {
    clusterName: "Home",
    routes: homeRoutes
  }
};

export default routeMap;
