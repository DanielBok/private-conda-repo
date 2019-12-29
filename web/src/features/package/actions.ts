import { createAsyncAction } from "typesafe-actions";
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
)<void, PackageType.PackageDetail[], void>();
