import { PackageType } from "@/features/package";
import { UserType } from "@/features/user";
import { RouterState } from "connected-react-router";

export type RootState = {
  user: UserType.Store;
  package: PackageType.Store;
  router: RouterState;
};
