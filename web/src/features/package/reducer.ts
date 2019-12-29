import AllActions from "@/infrastructure/rootAction";
import produce from "immer";
import { getType } from "typesafe-actions";
import * as PackageType from "./types";
import * as PackageAction from "./actions";

const defaultState: PackageType.Store = {
  packages: [],
  loading: {
    details: "SUCCESS",
    packages: "SUCCESS"
  },
  selected: {
    channel: "",
    package: "",
    details: [],
    latestVersion: ""
  }
};

export default (state = defaultState, action: AllActions) =>
  produce(state, draft => {
    switch (action.type) {
      case getType(PackageAction.fetchAllPackagesAsync.request):
        draft.loading.packages = "REQUEST";
        break;

      case getType(PackageAction.fetchPackageDetail.request):
        draft.loading.details = "REQUEST";
        break;

      case getType(PackageAction.fetchAllPackagesAsync.failure):
        draft.loading.packages = "FAILURE";
        break;

      case getType(PackageAction.fetchPackageDetail.failure):
        draft.loading.details = "FAILURE";
        break;

      case getType(PackageAction.fetchAllPackagesAsync.success):
        draft.packages = action.payload;
        draft.loading.packages = "SUCCESS";
        break;

      case getType(PackageAction.fetchPackageDetail.success): {
        draft.loading.details = "SUCCESS";
        const details = action.payload;

        draft.selected = {
          details,
          channel: details[0].channel,
          package: details[0].package,
          latestVersion: details.reduce(
            (v, { version }) => (v > version ? v : version),
            details[0].version
          )
        };
        break;
      }
    }
  });
