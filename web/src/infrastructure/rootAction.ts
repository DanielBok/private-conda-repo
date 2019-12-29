import { MetaAction } from "@/features/meta";
import { PackageAction } from "@/features/package";
import { UserAction } from "@/features/user";

import { ActionType } from "typesafe-actions";

type AllActions =
  | ActionType<typeof MetaAction>
  | ActionType<typeof UserAction>
  | ActionType<typeof PackageAction>;

export default AllActions;
