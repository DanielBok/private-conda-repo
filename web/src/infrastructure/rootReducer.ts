import MetaReducer from "@/features/meta/reducer";
import PackageReducer from "@/features/package/reducer";
import ChannelReducer from "@/features/channel/reducer";
import { connectRouter } from "connected-react-router";
import { History } from "history";
import { combineReducers } from "redux";
import { RootState } from "./rootState";

export default (history: History) =>
  combineReducers<RootState>({
    meta: MetaReducer,
    pkg: PackageReducer,
    router: connectRouter(history),
    channel: ChannelReducer,
  });
