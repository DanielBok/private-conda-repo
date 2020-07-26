import api, { ThunkFunction, ThunkFunctionAsync } from "@/infrastructure/api";
import { notification } from "antd";
import * as A from "./actions";
import { ChannelStorage } from "./localstorage";
import * as T from "./types";

/**
 * Creates channel in the backend server
 */
export const createChannel = (): ThunkFunctionAsync => async (
  dispatch,
  getState
) => {
  const {
    form: { errors, pristine, channel, password, confirm, email },
    loading: { validation },
  } = getState().channel;

  const invalid =
    confirm !== password ||
    Object.values(pristine).some((e) => e) ||
    Object.values(errors).some((e) => e.length > 0);
  if (validation === "REQUEST" || invalid) return;

  const payload: T.Channel = { channel, password };

  const { status } = await api.Post(
    "/channel",
    { ...payload, email },
    {
      beforeRequest: () => dispatch(A.createChannelAsync.request()),
    }
  );

  if (status === 200) {
    dispatch(A.createChannelAsync.success(payload));
    ChannelStorage.save(payload);
    notification.success({
      message: `Channel: ${payload.channel} created. Looking forward to your contributions!`,
    });
  } else {
    dispatch(A.createChannelAsync.failure());
    ChannelStorage.clear();
  }
};

/**
 * Loads channel details from local storage
 */
export const loadChannel = (): ThunkFunctionAsync => async (dispatch) => {
  const channel = ChannelStorage.load();
  if (channel) {
    await dispatch(validateChannel(channel.channel, channel.password));
  }
};

/**
 * Checks if the channel is valid. Used for logging in
 */
export const validateChannel = (
  channel: string,
  password: string
): ThunkFunctionAsync<boolean> => async (dispatch, getState) => {
  if (getState().channel.loading.validation === "REQUEST") return false;
  const payload: T.Channel = { channel, password };

  const { status } = await api.Post("/channel/check", payload, {
    beforeRequest: () => dispatch(A.fetchChannelCredentialsAsync.request()),
  });

  if (status === 200) {
    dispatch(A.fetchChannelCredentialsAsync.success(payload));
    ChannelStorage.save(payload);
    return true;
  } else {
    dispatch(A.fetchChannelCredentialsAsync.failure());
    return false;
  }
};

/**
 * Logs the channel out
 */
export const logout = (): ThunkFunction => (dispatch) => {
  ChannelStorage.clear();
  dispatch(A.logout());
};

/**
 * Checks if channel is available from the backend
 * @param channel name to check
 */
export const isChannelAvailable = (
  channel: string
): ThunkFunctionAsync => async (dispatch) => {
  const { status, data } = await api.Post<string>("/channel/check", {
    channel,
    password: "",
  });

  const available = status === 400 && data.trim() === "record not found";
  if (!available) {
    dispatch(A.updateForm({ errors: { channel: "channel is not available" } }));
  } else {
  }
};
