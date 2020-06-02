import AllActions from "@/infrastructure/rootAction";
import produce from "immer";
import { getType } from "typesafe-actions";
import * as MetaAction from "./actions";
import * as MetaType from "./types";

const defaultState: MetaType.Store = {
  indexer: "shell",
  image: "",
  loading: "SUCCESS",
  registry: "",
  repository: "",
};

export default (state = defaultState, action: AllActions) =>
  produce(state, (draft) => {
    switch (action.type) {
      case getType(MetaAction.fetchMetaInfoAsync.request):
        draft.loading = "REQUEST";
        break;

      case getType(MetaAction.fetchMetaInfoAsync.failure):
        draft.loading = "FAILURE";
        break;

      case getType(MetaAction.fetchMetaInfoAsync.success):
        draft.loading = "SUCCESS";
        draft.indexer = action.payload.indexer;
        draft.image = action.payload.image;
        draft.registry = action.payload.registry;
        draft.repository = action.payload.repository;
        break;
    }
  });
