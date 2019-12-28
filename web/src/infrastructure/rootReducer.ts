import { connectRouter } from "connected-react-router";
import { RootState } from "./rootState";
import { combineReducers } from "redux";
import { History } from "history";

import UserReducer from "@/features/user/reducer";

export default (history: History) =>
  combineReducers<RootState>({
    router: connectRouter(history),
    user: UserReducer
  });
