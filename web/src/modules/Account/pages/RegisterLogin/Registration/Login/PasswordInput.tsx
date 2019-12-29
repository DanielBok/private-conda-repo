import { Form, Input } from "antd";
import React from "react";
import * as CONST from "./constants";
import { useLoginContext } from "./hooks";

export default () => {
  const {
    form: { getFieldDecorator, getFieldError, isFieldTouched },
    submit,
    valid
  } = useLoginContext();

  return (
    <Form.Item hasFeedback validateStatus={status()}>
      {getFieldDecorator(CONST.PASSWORD, {
        rules: [
          {
            required: true,
            message: "password is required"
          }
        ]
      })(
        <Input
          placeholder="Password"
          type="password"
          onKeyPress={e => {
            if (e.key === "Enter") submit();
          }}
        />
      )}
    </Form.Item>
  );

  function status() {
    if (!isFieldTouched(CONST.PASSWORD)) return "";
    if (!valid || getFieldError(CONST.PASSWORD) !== undefined) return "error";
    return "success";
  }
};
