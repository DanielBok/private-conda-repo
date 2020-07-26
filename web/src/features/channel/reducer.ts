import AllActions from "@/infrastructure/rootAction";
import produce from "immer";
import { getType } from "typesafe-actions";
import * as A from "./actions";
import * as T from "./types";

const defaultState: T.Store = {
  channel: "",
  password: "",
  validated: false,
  form: {
    channel: "",
    confirm: "",
    password: "",
  },
  loading: {
    availableCheck: "SUCCESS",
    validation: "SUCCESS",
  },
};

export default (state = defaultState, action: AllActions) =>
  produce(state, (draft) => {
    switch (action.type) {
      case getType(A.createChannelAsync.request):
      case getType(A.fetchChannelCredentialsAsync.request):
        draft.loading.validation = "REQUEST";
        break;

      case getType(A.createChannelAsync.failure):
      case getType(A.fetchChannelCredentialsAsync.failure):
        draft.validated = false;
        draft.loading.validation = "FAILURE";
        break;

      case getType(A.createChannelAsync.success):
      case getType(A.fetchChannelCredentialsAsync.success):
        draft.validated = true;
        draft.loading.validation = "SUCCESS";
        draft.channel = action.payload.channel;
        draft.password = action.payload.password;
        break;

      case getType(A.logout):
        draft.validated = false;
        draft.loading.validation = "SUCCESS";
        draft.channel = "";
        draft.password = "";
        break;
    }
  });
