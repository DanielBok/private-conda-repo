import { ChnAction, ChnApi } from "@/features/channel";
import { Form, Input } from "antd";
import React from "react";
import { useDispatch } from "react-redux";
import Errors from "./Errors";
import { useDetails, useDisabled } from "./utils";

const ConfirmInput = () => {
  const dispatch = useDispatch();
  const disabled = useDisabled();
  const [confirm, errors, status] = useDetails("confirm");

  return (
    <Form.Item
      validateStatus={status}
      hasFeedback
      help={<Errors errors={errors} />}
    >
      <Input
        value={confirm}
        placeholder="Confirm Password"
        type="password"
        onChange={(e) => {
          const confirm = e.target.value.trim().toLowerCase();
          dispatch(ChnAction.updateForm({ confirm }));
        }}
        onKeyPress={(e) => {
          if (e.key === "Enter" && !disabled) dispatch(ChnApi.createChannel());
        }}
      />
    </Form.Item>
  );
};

export default ConfirmInput;
