import { UserApi } from "@/features/user";
import { Form, Input } from "antd";
import debounce from "lodash/debounce";
import React from "react";
import * as CONST from "./constants";
import { useFormContext } from "./hooks";

export default () => {
  const {
    form: { getFieldDecorator, getFieldError },
    validateStatus,
    setValidateStatus
  } = useFormContext();

  return (
    <Form.Item
      validateStatus={
        getFieldError(CONST.USERNAME) !== undefined ? "error" : validateStatus
      }
      hasFeedback
    >
      {getFieldDecorator(CONST.USERNAME, {
        validateFirst: true,
        rules: [
          { required: true, message: "Username is required" },
          { min: 4, message: "Username is at least 4 character long" },
          {
            validator: debounce(
              async (
                _: string,
                value: string,
                callback: (error?: string) => void
              ) => {
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
              300
            )
          }
        ]
      })(<Input placeholder="Username" />)}
    </Form.Item>
  );
};
