import accountRoutes from "@/modules/Account";
import homeRoutes from "@/modules/Home";
import packageRoutes from "@/modules/Package";

const routeMap: Record<string, ModuleRoutes> = {
  "/": {
    clusterName: "Home",
    routes: homeRoutes
  },
  "/account": {
    clusterName: "Account",
    routes: accountRoutes
  },
  "/p": {
    clusterName: "Package",
    routes: packageRoutes
  }
};

export default routeMap;
