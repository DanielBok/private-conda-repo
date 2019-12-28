import { routerMiddleware } from "connected-react-router";
import { createBrowserHistory } from "history";
import { applyMiddleware, createStore } from "redux";
import { composeWithDevTools } from "redux-devtools-extension";
import thunk from "redux-thunk";

import createReducer from "./rootReducer";

export const history = createBrowserHistory();

function configureStore() {
  const middleware = composeWithDevTools(
    applyMiddleware(thunk, routerMiddleware(history))
  );

  return createStore(createReducer(history), middleware);
}

export default configureStore();
