import { useLoginContext } from "./hooks";
import { Form, Input } from "antd";
import React from "react";
import * as CONST from "./constants";

export default () => {
  const {
    form: { getFieldDecorator, getFieldError, isFieldTouched },
    submit,
    valid
  } = useLoginContext();

  return (
    <Form.Item hasFeedback validateStatus={status()}>
      {getFieldDecorator(CONST.USERNAME, {
        rules: [
          {
            required: true,
            message: "username is required"
          },
          {
            min: 4,
            message: "username must be at least 4 characters long"
          }
        ]
      })(
        <Input
          placeholder="Username / Channel"
          onKeyPress={e => {
            if (e.key === "Enter") submit();
          }}
        />
      )}
    </Form.Item>
  );

  function status() {
    if (!isFieldTouched(CONST.USERNAME)) return "";
    if (!valid || getFieldError(CONST.USERNAME) !== undefined) return "error";
    return "success";
  }
};
