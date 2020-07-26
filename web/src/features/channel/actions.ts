import { createAction, createAsyncAction } from "typesafe-actions";
import * as T from "./types";

export const createChannelAsync = createAsyncAction(
  "CREATE_CHANNEL_REQUEST",
  "CREATE_CHANNEL_SUCCESS",
  "CREATE_CHANNEL_FAILURE"
)<void, T.Channel, void>();

export const fetchChannelCredentialsAsync = createAsyncAction(
  "FETCH_CHANNEL_CREDENTIALS_REQUEST",
  "FETCH_CHANNEL_CREDENTIALS_SUCCESS",
  "FETCH_CHANNEL_CREDENTIALS_FAILURE"
)<void, T.Channel, void>();

export const updateForm = createAction("UPDATE_CHANNEL_FORM")<
  DeepPartial<T.RegistrationForm>
>();

export const resetForm = createAction("RESET_CHANNEL_FORM")<void>();

export const logout = createAction("LOGOUT_CHANNEL")<void>();
