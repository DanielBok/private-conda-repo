import { Form, Input } from "antd";
import React, { FC } from "react";
import { useLoginContext, useStatus, useSubmit } from "./hooks";
import Errors from "./Errors";

const PasswordInput: FC = () => {
  const {
    state: { password, disabled },
    dispatch,
  } = useLoginContext();
  const submit = useSubmit();
  const status = useStatus("password");

  return (
    <Form.Item
      hasFeedback
      validateStatus={status}
      help={<Errors field="password" />}
    >
      <Input
        placeholder="Password"
        type="password"
        value={password}
        onChange={(e) => {
          dispatch({
            type: "SET_PASSWORD",
            payload: { password: e.target.value },
          });
        }}
        onKeyPress={(e) => {
          if (e.key === "Enter" && !disabled) submit();
        }}
      />
    </Form.Item>
  );
};

export default PasswordInput;
