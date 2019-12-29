import { MetaType } from "@/features/meta";
import { PackageType } from "@/features/package";
import { UserType } from "@/features/user";
import { RouterState } from "connected-react-router";

export type RootState = {
  meta: MetaType.Store;
  package: PackageType.Store;
  user: UserType.Store;
  router: RouterState;
};
