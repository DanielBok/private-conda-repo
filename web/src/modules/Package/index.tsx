import PackageDetail from "./pages/PackageDetail";
import ChannelDetail from "./pages/ChannelDetail";

export default [
  {
    component: ChannelDetail,
    path: "/:channel",
    title: "Channel Detail"
  },
  {
    component: PackageDetail,
    path: "/:channel/:pkg",
    title: "Package Detail"
  }
] as ModuleRoute[];
