import { MetaType } from "@/features/meta";
import { PkgType } from "@/features/package";
import { ChnType } from "@/features/channel";
import { RouterState } from "connected-react-router";

export type RootState = {
  meta: MetaType.Store;
  pkg: PkgType.Store;
  channel: ChnType.Store;
  router: RouterState;
};
