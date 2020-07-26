import AllActions from "@/infrastructure/rootAction";
import produce from "immer";
import { getType } from "typesafe-actions";
import { merge } from "lodash";

import * as A from "./actions";
import * as T from "./types";

const defaultState: T.Store = {
  channel: "",
  password: "",
  validated: false,
  form: {
    channel: "",
    confirm: "",
    password: "",
    email: "",
    errors: {
      channel: "",
      confirm: "",
      password: "",
      email: "",
    },
    pristine: {
      channel: true,
      confirm: true,
      password: true,
      email: true,
    },
  },
  loading: {
    availableCheck: "SUCCESS",
    validation: "SUCCESS",
  },
};

export default (state = defaultState, action: AllActions) =>
  produce(state, (draft) => {
    switch (action.type) {
      case getType(A.createChannelAsync.request):
      case getType(A.fetchChannelCredentialsAsync.request):
        draft.loading.validation = "REQUEST";
        break;

      case getType(A.createChannelAsync.failure):
      case getType(A.fetchChannelCredentialsAsync.failure):
        draft.validated = false;
        draft.loading.validation = "FAILURE";
        break;

      case getType(A.createChannelAsync.success):
      case getType(A.fetchChannelCredentialsAsync.success):
        draft.validated = true;
        draft.loading.validation = "SUCCESS";
        draft.channel = action.payload.channel;
        draft.password = action.payload.password;
        break;

      case getType(A.logout):
        draft.validated = false;
        draft.loading.validation = "SUCCESS";
        draft.channel = "";
        draft.password = "";
        break;

      case getType(A.updateForm): {
        const p = action.payload;
        merge(draft.form, p);

        if (p.channel) draft.form.pristine.channel = false;
        if (p.email) draft.form.pristine.email = false;
        if (p.password) draft.form.pristine.password = false;
        if (p.confirm) draft.form.pristine.confirm = false;

        const { pristine } = draft.form;

        // only run checks once it is non-pristine
        if (!pristine.channel) {
          const { channel } = draft.form;
          if (channel === "")
            draft.form.errors.channel = "Username is required";
          else if (channel.length < 2)
            draft.form.errors.channel =
              "Username must be at least 2 characters long";
          else draft.form.errors.channel = "";
        }

        if (p.email) draft.form.pristine.email = false;
        if (!pristine.email) {
          const { email } = draft.form;
          if (email === "") {
            draft.form.errors.email = "Email is required";
          } else if (!/^[^@\s]+@[^@\s]+$/.test(email)) {
            draft.form.errors.email = "Email doesn't seem to be valid";
          } else {
            draft.form.errors.email = "";
          }
        }

        if (p.password) draft.form.pristine.password = false;
        if (!pristine.password) {
          const { password } = draft.form;
          if (password === "")
            draft.form.errors.password = "Password is required";
          else if (password.length < 4)
            draft.form.errors.password =
              "Password must be at least 4 characters long";
          else draft.form.errors.password = "";
        }

        // both password and confirm must have been edited
        if (!pristine.confirm && !pristine.password) {
          const { password, confirm } = draft.form;
          draft.form.errors.confirm =
            password !== confirm ? "passwords do not match" : "";
        }

        break;
      }

      case getType(A.resetForm):
        draft.form = {
          password: "",
          email: "",
          channel: "",
          confirm: "",
          errors: {
            password: "",
            email: "",
            channel: "",
            confirm: "",
          },
          pristine: {
            password: true,
            email: true,
            channel: true,
            confirm: true,
          },
        };
    }
  });
