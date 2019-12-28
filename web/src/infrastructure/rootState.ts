import { RouterState } from "connected-react-router";
import { UserType } from "@/features/user";

export type RootState = {
  user: UserType.Store;
  router: RouterState;
};
