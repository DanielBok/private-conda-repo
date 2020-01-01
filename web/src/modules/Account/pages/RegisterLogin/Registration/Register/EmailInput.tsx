import { Form, Input } from "antd";
import React from "react";
import * as CONST from "./constants";
import { useFormContext } from "./hooks";

export default () => {
  const {
    getFieldDecorator,
    getFieldError,
    isFieldTouched
  } = useFormContext().form;

  const status = isFieldTouched(CONST.EMAIL)
    ? getFieldError(CONST.EMAIL) === undefined
      ? "success"
      : "error"
    : "";

  return (
    <Form.Item validateStatus={status} hasFeedback>
      {getFieldDecorator(CONST.EMAIL, {
        rules: [
          { required: true, message: "Email is required" },
          { type: "email", message: "Invalid email" }
        ]
      })(<Input placeholder="Email" type="email" />)}
    </Form.Item>
  );
};
