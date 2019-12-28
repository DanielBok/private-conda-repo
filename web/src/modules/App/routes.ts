import homeRoutes from "@/modules/Home";
import accountRoutes from "@/modules/Account";

const routeMap: Record<string, ModuleRoutes> = {
  "/": {
    clusterName: "Home",
    routes: homeRoutes
  },
  "/account": {
    clusterName: "Account",
    routes: accountRoutes
  }
};

export default routeMap;
