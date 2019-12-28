import { UserAction } from "@/features/user";

import { ActionType } from "typesafe-actions";

type AllActions = ActionType<typeof UserAction>;

export default AllActions;
