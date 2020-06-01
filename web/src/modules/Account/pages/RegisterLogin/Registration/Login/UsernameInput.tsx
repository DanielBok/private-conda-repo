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
      name={CONST.USERNAME}
      hasFeedback
      validateStatus={status()}
      rules={[
        {
          required: true,
          message: "username is required",
        },
        {
          min: 2,
          message: "username must be at least 2 characters long",
        },
      ]}
    >
      <Input
        placeholder="Username / Channel"
        onKeyPress={(e) => {
          if (e.key === "Enter") submit();
        }}
      />
    </Form.Item>
  );

  function status() {
    if (!isFieldTouched(CONST.USERNAME)) return "";
    if (!valid || getFieldError(CONST.USERNAME) !== undefined) return "error";
    return "success";
  }
};
