import api, { ThunkFunctionAsync } from "@/infrastructure/api";
import * as UserAction from "./actions";
import * as UserType from "./types";

export const validateUser = (
  username: string,
  password: string
): ThunkFunctionAsync => async (dispatch, getState) => {
  if (getState().user.loading === "REQUEST") return;
  const payload: UserType.UserInfo = {
    channel: username,
    password
  };

  const { status } = await api.Post("/user/check", payload, {
    beforeRequest: () =>
      dispatch(UserAction.fetchUserCredentialsAsync.request())
  });

  if (status === 200) {
    dispatch(UserAction.fetchUserCredentialsAsync.success(payload));
  } else {
    dispatch(UserAction.fetchUserCredentialsAsync.failure());
  }
};

export const isUsernameAvailable = async (username: string) => {
  const { status, data } = await api.Post<string>("/user/check", {
    channel: username,
    password: ""
  });

  return status === 400 && data.trim() === "record not found";
};
