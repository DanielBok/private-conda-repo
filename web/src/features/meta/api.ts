import api, { ThunkFunctionAsync } from "@/infrastructure/api";
import { notification } from "antd";
import * as MetaAction from "./actions";
import * as MetaType from "./types";

/**
 * Fetches application meta information
 */
export const fetchMetaInfo = (): ThunkFunctionAsync => async (
  dispatch,
  getState
) => {
  if (getState().meta.loading === "REQUEST") return;

  const { data, status } = await api.Get<MetaType.MetaInfo>("/meta", {
    beforeRequest: () => dispatch(MetaAction.fetchMetaInfoAsync.request()),
    onError: (e) => {
      notification.error({
        message: `Could not retrieve meta information data. Reason: ${e.data}`,
        duration: 8,
      });
    },
  });

  if (status === 200) {
    dispatch(MetaAction.fetchMetaInfoAsync.success(data));
  }
};
