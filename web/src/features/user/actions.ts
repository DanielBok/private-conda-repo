import { createAction, createAsyncAction } from "typesafe-actions";
import * as UserType from "./types";

export const createUserAsync = createAsyncAction(
  "CREATE_USER_REQUEST",
  "CREATE_USER_SUCCESS",
  "CREATE_USER_FAILURE"
)<void, UserType.UserInfo, void>();

export const fetchUserCredentialsAsync = createAsyncAction(
  "FETCH_USER_CREDENTIALS_REQUEST",
  "FETCH_USER_CREDENTIALS_SUCCESS",
  "FETCH_USER_CREDENTIALS_FAILURE"
)<void, UserType.UserInfo, void>();

export const logoutUser = createAction("LOGOUT_USER")<void>();
