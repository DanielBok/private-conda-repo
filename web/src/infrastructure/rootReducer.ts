import { connectRouter } from "connected-react-router";
import { RootState } from "./rootState";
import { combineReducers } from "redux";
import { History } from "history";

export default (history: History) =>
  combineReducers<RootState>({
    router: connectRouter(history)
  });
