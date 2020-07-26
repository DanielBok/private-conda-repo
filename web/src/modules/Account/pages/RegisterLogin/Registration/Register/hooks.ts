import { ChnApi } from "@/features/channel";
import { ThunkDispatchAsync } from "@/infrastructure/api";
import { push } from "connected-react-router";
import React, { useContext } from "react";
import { useDispatch } from "react-redux";
import { Fields, State, useRegistrationReducer } from "./reducer";

export const RegistrationContext = React.createContext<{
  state: State;
  dispatch: ReturnType<typeof useRegistrationReducer>[1];
}>({
  state: {} as any,
  dispatch: (_) => {},
});

export const useRegistrationContext = () => useContext(RegistrationContext);

export const useStatus = (field: keyof Fields) => {
  const { pristine, errors, confirm } = useRegistrationContext().state;

  if (pristine[field]) return "";
  if (field === "confirm" && confirm === "") return "";

  if (errors[field].length > 0) return "error";
  return "success";
};

export const useSubmit = () => {
  const dispatch = useDispatch() as ThunkDispatchAsync;
  const { username, password, email } = useRegistrationContext().state;

  return async () => {
    await dispatch(ChnApi.createChannel(username, password, email));
    dispatch(push("/"));
  };
};
