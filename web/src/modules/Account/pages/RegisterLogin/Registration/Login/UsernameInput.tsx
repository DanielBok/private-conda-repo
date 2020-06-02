import Errors from "./Errors";
import { Form, Input } from "antd";
import React, { FC } from "react";
import { useLoginContext, useStatus, useSubmit } from "./hooks";

const UsernameInput: FC = () => {
  const {
    state: { username, disabled },
    dispatch,
  } = useLoginContext();
  const submit = useSubmit();
  const status = useStatus("username");

  return (
    <Form.Item
      hasFeedback
      validateStatus={status}
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

export default UsernameInput;
