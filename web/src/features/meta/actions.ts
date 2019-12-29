import { createAsyncAction } from "typesafe-actions";
import * as MetaType from "./types";

export const fetchMetaInfoAsync = createAsyncAction(
  "FETCH_META_INFO_REQUEST",
  "FETCH_META_INFO_SUCCESS",
  "FETCH_META_INFO_FAILURE"
)<void, MetaType.MetaInfo, void>();
