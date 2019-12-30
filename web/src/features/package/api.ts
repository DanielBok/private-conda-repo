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
  if (getState().package.loading.packages === "REQUEST") return;

  const { data, status } = await api.Get<PackageType.PackageMetaInfo[]>("p", {
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

/**
 * From the specified channel, fetch all details about the package
 *
 * @param channel channel/user name
 * @param pkg package name
 */
export const fetchPackageDetail = (
  channel: string,
  pkg: string
): ThunkFunctionAsync<LoadingState> => async (dispatch, getState) => {
  if (getState().package.loading.details === "REQUEST") return "REQUEST";

  const { data, status } = await api.Get<PackageType.PackageDetail<string>>(
    `p/${channel}/${pkg}`,
    {
      beforeRequest: () => dispatch(PackageAction.fetchPackageDetail.request()),
      onError: () => dispatch(PackageAction.fetchPackageDetail.failure())
    }
  );
  if (status !== 200) {
    return "FAILURE";
  }

  dispatch(PackageAction.fetchPackageDetail.success(data));
  return "SUCCESS";
};

/**
 * Removes the package from the channel. The user must be signed in and must be the owner of the channel
 * for operation to succeed
 *
 * @param channel name of the channel
 * @param detail package details
 */
export const removePackage = (
  channel: string,
  detail: PackageType.RemovePackagePayload["package"]
): ThunkFunctionAsync => async (dispatch, getState) => {
  const {
    user,
    package: { loading }
  } = getState();
  if (loading.packages === "REQUEST") return;

  if (!user.validated || user.username !== channel) return;

  const payload: PackageType.RemovePackagePayload = {
    channel,
    password: user.password,
    package: detail
  };
  const { status } = await api.Delete<void>("p", payload, {
    beforeRequest: () => dispatch(PackageAction.removePackageDetail.request()),
    onError: () => dispatch(PackageAction.removePackageDetail.failure())
  });

  if (status === 200) {
    dispatch(PackageAction.removePackageDetail.success(payload.package));
  }
};
