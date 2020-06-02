import { Form, Input } from "antd";
import React from "react";
import Errors from "./Errors";
import { useRegistrationContext, useStatus, useSubmit } from "./hooks";

const PasswordInput = () => {
  const {
    state: { password, disabled },
    dispatch,
  } = useRegistrationContext();

  const status = useStatus("password");
  const submit = useSubmit();

  return (
    <Form.Item
      validateStatus={status}
      hasFeedback
      help={<Errors field="password" />}
    >
      <Input
        value={password}
        placeholder="Password"
        type="password"
        onChange={(e) =>
          dispatch({
            type: "SET_PASSWORD",
            payload: { password: e.target.value },
          })
        }
        onKeyPress={(e) => {
          if (e.key === "Enter" && !disabled) submit();
        }}
      />
    </Form.Item>
  );
};

export default PasswordInput;
