import { Form, Input } from "antd";
import React from "react";
import Errors from "./Errors";
import { useRegistrationContext, useStatus, useSubmit } from "./hooks";

export default () => {
  const {
    state: { email, disabled },
    dispatch,
  } = useRegistrationContext();

  const status = useStatus("email");
  const submit = useSubmit();

  return (
    <Form.Item
      validateStatus={status}
      hasFeedback
      help={<Errors field="email" />}
    >
      <Input
        value={email}
        placeholder="Email"
        type="email"
        onChange={(e) =>
          dispatch({
            type: "SET_EMAIL",
            payload: { email: e.target.value },
          })
        }
        onKeyPress={(e) => {
          if (e.key === "Enter" && !disabled) submit();
        }}
      />
    </Form.Item>
  );
};
