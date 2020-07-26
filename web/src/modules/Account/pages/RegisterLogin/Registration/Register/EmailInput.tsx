import { ChnAction, ChnApi } from "@/features/channel";
import { Form, Input } from "antd";
import React from "react";
import { useDispatch } from "react-redux";
import Errors from "./Errors";
import { useDetails, useDisabled } from "./utils";

const EmailInput = () => {
  const dispatch = useDispatch();
  const disabled = useDisabled();
  const [email, errors, status] = useDetails("email");

  return (
    <Form.Item
      validateStatus={status}
      hasFeedback
      help={<Errors errors={errors} />}
    >
      <Input
        value={email}
        placeholder="Email"
        type="email"
        onChange={(e) => {
          const email = e.target.value.trim().toLowerCase();
          dispatch(ChnAction.updateForm({ email }));
        }}
        onKeyPress={(e) => {
          if (e.key === "Enter" && !disabled) dispatch(ChnApi.createChannel());
        }}
      />
    </Form.Item>
  );
};

export default EmailInput;
