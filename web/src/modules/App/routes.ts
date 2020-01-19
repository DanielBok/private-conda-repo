import accountRoutes from "@/modules/Account";
import helpRoutes from "@/modules/Help";
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
  "/help": {
    clusterName: "Help",
    routes: helpRoutes
  },
  "/p": {
    clusterName: "Package",
    routes: packageRoutes
  }
};

export default routeMap;
