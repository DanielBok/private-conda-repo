import AllActions from "@/infrastructure/rootAction";
import produce from "immer";
import { getType } from "typesafe-actions";

import * as UserType from "./types";
import * as UserAction from "./actions";

const defaultState: UserType.Store = {
  username: "",
  password: "",
  validated: false,
  loading: "SUCCESS"
};

export default (state = defaultState, action: AllActions) =>
  produce(state, draft => {
    switch (action.type) {
      case getType(UserAction.createUserAsync.request):
      case getType(UserAction.fetchUserCredentialsAsync.request):
        draft.loading = "REQUEST";
        break;

      case getType(UserAction.createUserAsync.failure):
      case getType(UserAction.fetchUserCredentialsAsync.failure):
        draft.validated = false;
        draft.loading = "FAILURE";
        break;

      case getType(UserAction.createUserAsync.success):
      case getType(UserAction.fetchUserCredentialsAsync.success):
        draft.validated = true;
        draft.loading = "SUCCESS";
        draft.username = action.payload.channel;
        draft.password = action.payload.password;
        break;

      case getType(UserAction.logoutUser):
        draft.validated = false;
        draft.loading = "SUCCESS";
        draft.username = "";
        draft.password = "";
        break;
    }
  });
