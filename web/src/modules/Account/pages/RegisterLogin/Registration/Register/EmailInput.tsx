import { Form, Input } from "antd";
import React from "react";
import * as CONST from "./constants";
import { useFormContext } from "./hooks";

export default () => {
  const { getFieldError, isFieldTouched } = useFormContext().form;

  const status = isFieldTouched(CONST.EMAIL)
    ? getFieldError(CONST.EMAIL) === undefined
      ? "success"
      : "error"
    : "";

  return (
    <Form.Item
      name={CONST.EMAIL}
      validateStatus={status}
      hasFeedback
      rules={[
        { required: true, message: "Email is required" },
        { type: "email", message: "Invalid email" },
      ]}
    >
      <Input placeholder="Email" type="email" />
    </Form.Item>
  );
};
