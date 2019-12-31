import AllActions from "@/infrastructure/rootAction";
import produce from "immer";
import moment from "moment";
import { getType } from "typesafe-actions";
import * as Action from "./actions";
import * as PackageType from "./types";

const defaultState: PackageType.Store = {
  packages: [],
  loading: {
    details: "SUCCESS",
    packages: "SUCCESS",
    channelPackages: "SUCCESS"
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
  },
  channelPackages: {
    channel: "",
    email: "",
    joinDate: moment(),
    packages: []
  }
};

export default (state = defaultState, action: AllActions) =>
  produce(state, draft => {
    switch (action.type) {
      case getType(Action.fetchAllPackagesAsync.request):
        draft.loading.packages = "REQUEST";
        break;

      case getType(Action.removePackageDetail.request):
      case getType(Action.fetchPackageDetail.request):
        draft.loading.details = "REQUEST";
        break;

      case getType(Action.fetchAllPackagesAsync.failure):
        draft.loading.packages = "FAILURE";
        break;

      case getType(Action.removePackageDetail.failure):
      case getType(Action.fetchPackageDetail.failure):
        draft.loading.details = "FAILURE";
        break;

      case getType(Action.fetchAllPackagesAsync.success):
        draft.packages = action.payload;
        draft.loading.packages = "SUCCESS";
        break;

      case getType(Action.fetchPackageDetail.success): {
        draft.loading.details = "SUCCESS";
        const { details, ...rest } = action.payload;

        draft.packageDetail = {
          details: details
            .map(({ uploadDate, ...rest }) => ({
              ...rest,
              uploadDate: moment.utc(uploadDate)
            }))
            .sort((x, y) => (x.uploadDate.isAfter(y.uploadDate) ? -1 : 1)),
          ...rest
        };
        break;
      }

      case getType(Action.removePackageDetail.success): {
        const p = action.payload;

        // remove same item from list of package details
        draft.packageDetail.details = draft.packageDetail.details.filter(
          d =>
            ![
              p.version === d.version,
              p.buildNumber === d.buildNumber,
              p.buildString === d.buildString,
              p.platform === d.platform,
              p.name === d.package
            ].reduce((a, e) => a && e, true)
        );

        draft.loading.details = "SUCCESS";
        break;
      }

      case getType(Action.fetchUserPackages.request):
        draft.loading.channelPackages = "REQUEST";
        break;

      case getType(Action.fetchUserPackages.failure):
        draft.loading.channelPackages = "FAILURE";
        break;

      case getType(Action.fetchUserPackages.success): {
        draft.loading.channelPackages = "SUCCESS";
        const { joinDate, ...rest } = action.payload;
        draft.channelPackages = {
          ...rest,
          joinDate: moment.utc(joinDate)
        };
        break;
      }
      case getType(Action.resetLoadingStore):
        draft.loading = {
          channelPackages: "SUCCESS",
          details: "SUCCESS",
          packages: "SUCCESS"
        };
        break;
    }
  });
