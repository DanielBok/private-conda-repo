import AllActions from "@/infrastructure/rootAction";
import produce from "immer";
import moment from "moment";
import { getType } from "typesafe-actions";
import * as PackageAction from "./actions";
import * as PackageType from "./types";

const defaultState: PackageType.Store = {
  packages: [],
  loading: {
    details: "SUCCESS",
    packages: "SUCCESS"
  },
  packageDetail: {
    channel: "",
    package: "",
    details: [],
    latest: {
      channel: "",
      platforms: [],
      version: "",
      description: "",
      devUrl: "",
      docUrl: "",
      home: "",
      license: "",
      summary: "",
      timestamp: 0,
      name: ""
    }
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
        const { details, ...rest } = action.payload;

        draft.packageDetail = {
          details: details
            .map(({ uploadDate, ...rest }) => ({
              ...rest,
              uploadDate: moment(uploadDate)
            }))
            .sort((x, y) => (x.uploadDate.isAfter(y.uploadDate) ? -1 : 1)),
          ...rest
        };
        break;
      }
    }
  });
