import api, { ThunkFunctionAsync } from "@/infrastructure/api";
import { notification } from "antd";
import * as PackageAction from "./actions";
import * as PackageType from "./types";

/**
 * Fetches all package meta information
 */
export const fetchAllPackages = (): ThunkFunctionAsync => async (
  dispatch,
  getState
) => {
  if (getState().package.loading === "REQUEST") return;

  const { data, status } = await api.Get<PackageType.PackageMetaInfo[]>("/p", {
    beforeRequest: () =>
      dispatch(PackageAction.fetchAllPackagesAsync.request()),
    onError: e => {
      notification.error({
        message: `Could not retrieve package data. Reason: ${e.data}`,
        duration: 8
      });
    }
  });

  if (status === 200) {
    dispatch(PackageAction.fetchAllPackagesAsync.success(data));
  }
};
