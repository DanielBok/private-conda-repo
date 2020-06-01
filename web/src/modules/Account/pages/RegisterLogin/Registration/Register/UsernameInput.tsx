import { UserApi } from "@/features/user";
import { Form, Input } from "antd";
import React from "react";
import { useDebouncedCallback } from "use-debounce";
import * as CONST from "./constants";
import { useFormContext } from "./hooks";

export default () => {
  const {
    form: { getFieldError },
    validateStatus,
  } = useFormContext();
  const [validator] = useUserValidator();

  return (
    <Form.Item
      name={CONST.USERNAME}
      validateStatus={
        getFieldError(CONST.USERNAME) !== undefined ? "error" : validateStatus
      }
      hasFeedback
      validateFirst
      rules={[
        { required: true, message: "Username is required" },
        { min: 4, message: "Username must have at least 4 character" },
        { validator },
      ]}
    >
      <Input placeholder="Username" />
    </Form.Item>
  );
};

const useUserValidator = () => {
  const { setValidateStatus } = useFormContext();

  return useDebouncedCallback(
    async (_, value: string, callback: (error?: string) => void) => {
      if (!value || value.length < 4) return;
      setValidateStatus("validating");
      const available = await UserApi.isUsernameAvailable(value);
      if (!available) {
        setValidateStatus("error");
        callback(`'${value}' is already taken.`);
        return;
      }
      setValidateStatus("success");
      callback();
    },
    500
  );
};
