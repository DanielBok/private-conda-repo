import PackageDetail from "./pages/PackageDetail";

export default [
  {
    component: PackageDetail,
    path: "/:channel/:pkg",
    title: "Package Detail"
  }
] as ModuleRoute[];
