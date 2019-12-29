import { UserAction } from "@/features/user";
import { PackageAction } from "@/features/package";

import { ActionType } from "typesafe-actions";

type AllActions =
  | ActionType<typeof UserAction>
  | ActionType<typeof PackageAction>;

export default AllActions;
