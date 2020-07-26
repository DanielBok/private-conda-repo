import api, { ThunkFunctionAsync } from "@/infrastructure/api";
import { notification } from "antd";
import { push } from "connected-react-router";
import * as A from "./actions";
import * as T from "./types";

/**
 * Fetches all package meta information
 */
export const fetchAllPackages = (): ThunkFunctionAsync => async (
  dispatch,
  getState
) => {
  if (getState().pkg.loading.packages === "REQUEST") return;

  const { data, status } = await api.Get<T.PackageMetaInfo[]>("p", {
    beforeRequest: () => dispatch(A.fetchAllPackagesAsync.request()),
    onError: (e) => {
      notification.error({
        message: `Could not retrieve package data. Reason: ${e.data}`,
        duration: 8,
      });
    },
  });

  if (status === 200) {
    dispatch(A.fetchAllPackagesAsync.success(data));
  }
};

/**
 * From the specified channel, fetch all details about the package
 *
 * @param channel channel name
 * @param pkg package name
 */
export const fetchPackageDetail = (
  channel: string,
  pkg: string
): ThunkFunctionAsync => async (dispatch, getState) => {
  if (getState().pkg.loading.details === "REQUEST") return;

  const { data, status } = await api.Get<T.PackageDetail<string>>(
    `p/${channel}/${pkg}`,
    {
      beforeRequest: () => dispatch(A.fetchPackageDetail.request()),
      onError: () => dispatch(A.fetchPackageDetail.failure()),
    }
  );
  if (status === 200) dispatch(A.fetchPackageDetail.success(data));
};

/**
 * Fetches all packages in channel
 * @param channel channel name
 */
export const fetchChannelPackages = (
  channel: string
): ThunkFunctionAsync => async (dispatch, getState) => {
  if (getState().pkg.loading.channelPackages === "REQUEST") return;

  const { status: status1, data: packages } = await api.Get<
    T.PackageMetaInfo[]
  >(`p/${channel}`, {
    beforeRequest: () => dispatch(A.fetchChannelPackages.request()),
  });

  if (status1 !== 200) return;

  const { status: status2, data: channelData } = await api.Get<
    Omit<T.ChannelPackages<string>, "packages">
  >(`channel/${channel}`, {
    onError: () => dispatch(A.fetchChannelPackages.failure()),
  });

  if (status2 === 200) {
    dispatch(A.fetchChannelPackages.success({ ...channelData, packages }));
  }
};

/**
 * Removes the package from the channel. The channel must be signed in and must be the owner of the channel
 * for operation to succeed
 *
 * @param channel name of the channel
 * @param detail package details
 */
export const removePackage = (
  channel: string,
  detail: T.RemovePackagePayload["package"]
): ThunkFunctionAsync => async (dispatch, getState) => {
  const {
    channel: chn,
    pkg: { loading },
  } = getState();

  if (loading.details === "REQUEST") return;
  if (!chn.validated || chn.channel !== channel) return;

  const payload: T.RemovePackagePayload = {
    channel,
    password: chn.password,
    package: detail,
  };

  const { status } = await api.Delete<void>("p", payload, {
    beforeRequest: () => dispatch(A.removePackageDetail.request()),
    onError: () => dispatch(A.removePackageDetail.failure()),
  });

  if (status === 200) {
    dispatch(A.removePackageDetail.success(payload.package));
    dispatch(push("/"));
  }
};
