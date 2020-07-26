import { ChnAction, ChnApi } from "@/features/channel";
import { Form, Input } from "antd";
import React from "react";
import { useDispatch } from "react-redux";
import Errors from "./Errors";
import { useDetails, useDisabled } from "./utils";

const PasswordInput = () => {
  const dispatch = useDispatch();
  const disabled = useDisabled();
  const [password, errors, status] = useDetails("password");

  return (
    <Form.Item
      validateStatus={status}
      hasFeedback
      help={<Errors errors={errors} />}
    >
      <Input
        value={password}
        placeholder="Password"
        type="password"
        onChange={(e) => {
          const password = e.target.value.trim().toLowerCase();
          dispatch(ChnAction.updateForm({ password }));
        }}
        onKeyPress={(e) => {
          if (e.key === "Enter" && !disabled) dispatch(ChnApi.createChannel());
        }}
      />
    </Form.Item>
  );
};

export default PasswordInput;
