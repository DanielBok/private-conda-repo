import { Form, Input } from "antd";
import React from "react";
import * as CONST from "./constants";
import { useFormContext } from "./hooks";

export default () => {
  const {
    form: { getFieldDecorator, getFieldError },
    validateStatus
  } = useFormContext();

  return (
    <Form.Item
      validateStatus={
        getFieldError(CONST.EMAIL) !== undefined ? "error" : validateStatus
      }
      hasFeedback
    >
      {getFieldDecorator(CONST.EMAIL, {
        validateFirst: true,
        rules: [
          { required: true, message: "Email is required" },
          { type: "email", message: "Invalid email" }
        ]
      })(<Input placeholder="Email" type="email" />)}
    </Form.Item>
  );
};
