import { UserApi } from "@/features/user";
import { Form, Input } from "antd";
import React, { useEffect } from "react";
import { useDebouncedCallback } from "use-debounce";
import Errors from "./Errors";
import { useRegistrationContext, useStatus, useSubmit } from "./hooks";

const UsernameInput = () => {
  const {
    state: { username, disabled },
    dispatch,
  } = useRegistrationContext();

  const check = useUsernameAvailableCheck();
  const status = useStatus("username");
  const submit = useSubmit();

  useEffect(() => {
    check(username);
  }, [check, username]);

  return (
    <Form.Item
      validateStatus={status}
      hasFeedback
      help={<Errors field="username" />}
    >
      <Input
        autoFocus
        value={username}
        placeholder="Username / Channel"
        onChange={(e) =>
          dispatch({
            type: "SET_USERNAME",
            payload: { username: e.target.value },
          })
        }
        onKeyPress={(e) => {
          if (e.key === "Enter" && !disabled) submit();
        }}
      />
    </Form.Item>
  );
};

const useUsernameAvailableCheck = () => {
  const { dispatch } = useRegistrationContext();
  const [cb] = useDebouncedCallback(async (username: string) => {
    if (username.length < 2) return;

    const available = await UserApi.isUsernameAvailable(username);
    dispatch({ type: "USERNAME_AVAILABILITY", payload: { available } });
  }, 500);

  return cb;
};

export default UsernameInput;
