import MetaReducer from "@/features/meta/reducer";
import PackageReducer from "@/features/package/reducer";
import UserReducer from "@/features/user/reducer";
import { connectRouter } from "connected-react-router";
import { History } from "history";
import { combineReducers } from "redux";
import { RootState } from "./rootState";

export default (history: History) =>
  combineReducers<RootState>({
    meta: MetaReducer,
    package: PackageReducer,
    router: connectRouter(history),
    user: UserReducer,
  });
