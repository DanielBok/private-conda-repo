import { createAction, createAsyncAction } from "typesafe-actions";
import * as T from "./types";

export const fetchAllPackagesAsync = createAsyncAction(
  "FETCH_ALL_PACKAGES_REQUEST",
  "FETCH_ALL_PACKAGES_SUCCESS",
  "FETCH_ALL_PACKAGES_FAILURE"
)<void, T.PackageMetaInfo[], void>();

export const fetchPackageDetail = createAsyncAction(
  "FETCH_PACKAGE_DETAIL_REQUEST",
  "FETCH_PACKAGE_DETAIL_SUCCESS",
  "FETCH_PACKAGE_DETAIL_FAILURE"
)<void, T.PackageDetail<string>, void>();

export const fetchChannelPackages = createAsyncAction(
  "FETCH_USER_PACKAGES_REQUEST",
  "FETCH_USER_PACKAGES_SUCCESS",
  "FETCH_USER_PACKAGES_FAILURE"
)<void, T.ChannelPackages<string>, void>();

export const removePackageDetail = createAsyncAction(
  "REMOVE_PACKAGE_DETAIL_REQUEST",
  "REMOVE_PACKAGE_DETAIL_SUCCESS",
  "REMOVE_PACKAGE_DETAIL_FAILURE"
)<void, T.RemovePackagePayload["package"], void>();

export const resetLoadingStore = createAction("RESET_PACKAGE_LOADING_STORE")<
  void
>();
