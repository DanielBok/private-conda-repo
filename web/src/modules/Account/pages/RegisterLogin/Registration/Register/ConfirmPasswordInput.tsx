import { Form, Input } from "antd";
import React from "react";
import * as CONST from "./constants";
import { useFormContext } from "./hooks";

export default () => {
  const {
    getFieldDecorator,
    getFieldError,
    getFieldValue,
    isFieldTouched
  } = useFormContext().form;
  const status = isFieldTouched(CONST.CONFIRM)
    ? getFieldError(CONST.CONFIRM) === undefined
      ? "success"
      : "error"
    : "";

  return (
    <Form.Item validateStatus={status} hasFeedback>
      {getFieldDecorator(CONST.CONFIRM, {
        rules: [
          {
            validator(_, value: string, callback) {
              if (value && value !== getFieldValue(CONST.PASSWORD)) {
                callback("Passwords do not match");
                return;
              }
              callback();
            }
          }
        ]
      })(<Input placeholder="Password" type="password" />)}
    </Form.Item>
  );
};
