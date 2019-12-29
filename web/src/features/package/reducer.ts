import AllActions from "@/infrastructure/rootAction";
import produce from "immer";
import { getType } from "typesafe-actions";
import * as PackageType from "./types";
import * as PackageAction from "./actions";

const defaultState: PackageType.Store = {
  packages: [],
  loading: "SUCCESS"
};

export default (state = defaultState, action: AllActions) =>
  produce(state, draft => {
    switch (action.type) {
      case getType(PackageAction.fetchAllPackagesAsync.request):
        draft.loading = "REQUEST";
        break;

      case getType(PackageAction.fetchAllPackagesAsync.failure):
        draft.loading = "FAILURE";
        break;

      case getType(PackageAction.fetchAllPackagesAsync.success):
        draft.packages = action.payload;
        draft.loading = "SUCCESS";
        break;
    }
  });
