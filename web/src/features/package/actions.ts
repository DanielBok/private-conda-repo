import { createAction, createAsyncAction } from "typesafe-actions";
import * as PackageType from "./types";

export const fetchAllPackagesAsync = createAsyncAction(
  "FETCH_ALL_PACKAGES_REQUEST",
  "FETCH_ALL_PACKAGES_SUCCESS",
  "FETCH_ALL_PACKAGES_FAILURE"
)<void, PackageType.PackageMetaInfo[], void>();

export const fetchPackageDetail = createAsyncAction(
  "FETCH_PACKAGE_DETAIL_REQUEST",
  "FETCH_PACKAGE_DETAIL_SUCCESS",
  "FETCH_PACKAGE_DETAIL_FAILURE"
)<void, PackageType.PackageDetail<string>, void>();

export const fetchUserPackages = createAsyncAction(
  "FETCH_USER_PACKAGES_REQUEST",
  "FETCH_USER_PACKAGES_SUCCESS",
  "FETCH_USER_PACKAGES_FAILURE"
)<void, PackageType.ChannelPackages<string>, void>();

export const removePackageDetail = createAsyncAction(
  "REMOVE_PACKAGE_DETAIL_REQUEST",
  "REMOVE_PACKAGE_DETAIL_SUCCESS",
  "REMOVE_PACKAGE_DETAIL_FAILURE"
)<void, PackageType.RemovePackagePayload["package"], void>();

export const resetLoadingStore = createAction("RESET_PACKAGE_LOADING_STORE")<
  void
>();
