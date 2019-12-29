import api, { ThunkFunction, ThunkFunctionAsync } from "@/infrastructure/api";
import * as UserAction from "./actions";

import { UserStorage } from "./localstorage";
import * as UserType from "./types";

/**
 * Creates user in the backend server
 */
export const createUser = (
  username: string,
  password: string
): ThunkFunctionAsync => async (dispatch, getState) => {
  if (getState().user.loading === "REQUEST") return;

  const payload: UserType.UserInfo = {
    channel: username,
    password
  };

  const { status } = await api.Post("/user", payload, {
    beforeRequest: () => dispatch(UserAction.createUserAsync.request())
  });

  if (status === 200) {
    dispatch(UserAction.createUserAsync.success(payload));
    UserStorage.save(payload);
  } else {
    dispatch(UserAction.createUserAsync.failure());
    UserStorage.clear();
  }
};

/**
 * Loads user details from local storage
 */
export const loadUser = (): ThunkFunctionAsync => async dispatch => {
  const user = UserStorage.load();
  if (user) {
    await dispatch(validateUser(user.channel, user.password));
  }
};

export const validateUser = (
  username: string,
  password: string
): ThunkFunctionAsync<boolean> => async (dispatch, getState) => {
  if (getState().user.loading === "REQUEST") return false;
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
    return true;
  } else {
    dispatch(UserAction.fetchUserCredentialsAsync.failure());
    return false;
  }
};

export const logout = (): ThunkFunction => dispatch => {
  UserStorage.clear();
  dispatch(UserAction.logoutUser());
};

/**
 * Checks if username is available from the backend
 * @param username name to check
 */
export const isUsernameAvailable = async (username: string) => {
  const { status, data } = await api.Post<string>("/user/check", {
    channel: username,
    password: ""
  });

  return status === 400 && data.trim() === "record not found";
};
