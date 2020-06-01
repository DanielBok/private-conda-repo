import { Form, Input } from "antd";
import React from "react";
import * as CONST from "./constants";
import { useLoginContext } from "./hooks";

export default () => {
  const {
    form: { getFieldError, isFieldTouched },
    submit,
    valid,
  } = useLoginContext();

  return (
    <Form.Item
      name={CONST.PASSWORD}
      hasFeedback
      validateStatus={status()}
      rules={[
        { required: true, message: "password is required" },
        { min: 4, message: "password needs to have at least 4 characters" },
      ]}
    >
      <Input
        placeholder="Password"
        type="password"
        onKeyPress={(e) => {
          if (e.key === "Enter") submit();
        }}
      />
    </Form.Item>
  );

  function status() {
    if (!isFieldTouched(CONST.PASSWORD)) return "";
    if (!valid || getFieldError(CONST.PASSWORD) !== undefined) return "error";
    return "success";
  }
};
