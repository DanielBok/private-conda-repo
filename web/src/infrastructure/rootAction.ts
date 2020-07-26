import { MetaAction } from "@/features/meta";
import { PkgAction } from "@/features/package";
import { ChnAction } from "@/features/channel";

import { ActionType } from "typesafe-actions";

type AllActions =
  | ActionType<typeof MetaAction>
  | ActionType<typeof ChnAction>
  | ActionType<typeof PkgAction>;

export default AllActions;
