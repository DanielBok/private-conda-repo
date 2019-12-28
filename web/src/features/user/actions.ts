import { createAsyncAction } from "typesafe-actions";
import * as UserType from "./types";

export const fetchUserCredentialsAsync = createAsyncAction(
  "FETCH_USER_CREDENTIALS_REQUEST",
  "FETCH_USER_CREDENTIALS_SUCCESS",
  "FETCH_USER_CREDENTIALS_FAILURE"
)<void, UserType.UserInfo, void>();
