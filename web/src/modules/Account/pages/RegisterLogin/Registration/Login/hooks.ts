import { ChnApi } from "@/features/channel";
import { ThunkDispatchAsync } from "@/infrastructure/api";
import { push } from "connected-react-router";
import React, { useContext } from "react";
import { useDispatch } from "react-redux";
import { Credential, State, useLoginReducer } from "./reducer";

export const LoginContext = React.createContext<{
  state: State;
  dispatch: ReturnType<typeof useLoginReducer>[1];
}>({
  state: {} as any,
  dispatch: (_) => {},
});

export const useLoginContext = () => useContext(LoginContext);

export const useStatus = (field: keyof Credential) => {
  const { pristine, errors, valid } = useLoginContext().state;

  if (pristine[field]) return "";
  if (!valid || errors[field].length > 0) return "error";
  return "success";
};

export const useSubmit = () => {
  const thunkDispatch = useDispatch() as ThunkDispatchAsync;
  const {
    state: { username, password },
    dispatch,
  } = useLoginContext();

  return async () => {
    const valid = await thunkDispatch(ChnApi.validateChannel(username, password));
    if (valid) {
      thunkDispatch(push("/"));
    } else {
      dispatch({ type: "SET_VALID", payload: { valid: false } });
    }
  };
};
