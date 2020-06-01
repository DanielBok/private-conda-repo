import { Form, Input } from "antd";
import React from "react";
import * as CONST from "./constants";
import { useFormContext } from "./hooks";

export default () => {
  const {
    getFieldError,
    getFieldValue,
    isFieldTouched,
  } = useFormContext().form;
  const status = isFieldTouched(CONST.CONFIRM)
    ? getFieldError(CONST.CONFIRM) === undefined
      ? "success"
      : "error"
    : "";

  return (
    <Form.Item
      name={CONST.CONFIRM}
      validateStatus={status}
      hasFeedback
      rules={[
        {
          validator(_, value: string, callback) {
            if (value && value !== getFieldValue(CONST.PASSWORD)) {
              callback("Passwords do not match");
              return;
            }
            callback();
          },
        },
      ]}
    >
      <Input placeholder="Confirm Password" type="password" />
    </Form.Item>
  );
};
