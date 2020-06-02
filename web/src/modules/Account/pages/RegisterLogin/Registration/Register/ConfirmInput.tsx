import { Form, Input } from "antd";
import React from "react";
import Errors from "./Errors";
import { useRegistrationContext, useStatus, useSubmit } from "./hooks";

const ConfirmInput = () => {
  const {
    state: { confirm, disabled },
    dispatch,
  } = useRegistrationContext();

  const status = useStatus("confirm");
  const submit = useSubmit();
  return (
    <Form.Item
      validateStatus={status}
      hasFeedback
      help={<Errors field="confirm" />}
    >
      <Input
        value={confirm}
        placeholder="Confirm Password"
        type="password"
        onChange={(e) =>
          dispatch({
            type: "SET_CONFIRM",
            payload: { confirm: e.target.value },
          })
        }
        onKeyPress={(e) => {
          if (e.key === "Enter" && !disabled) submit();
        }}
      />
    </Form.Item>
  );
};

export default ConfirmInput;