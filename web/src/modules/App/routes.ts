import accountRoutes from "@/modules/Account";
import helpRoutes from "@/modules/Help";
import homeRoutes from "@/modules/Home";
import packageRoutes from "@/modules/Package";
import uploadRoutes from "@/modules/UploadPage";

const routeMap: Record<string, ModuleRoutes> = {
  "/": {
    clusterName: "Home",
    routes: homeRoutes,
  },
  "/account": {
    clusterName: "Account",
    routes: accountRoutes,
  },
  "/help": {
    clusterName: "Help",
    routes: helpRoutes,
  },
  "/p": {
    clusterName: "Package",
    routes: packageRoutes,
  },
  "/upload": {
    clusterName: "Upload",
    routes: uploadRoutes,
  },
};

export default routeMap;
