import { UserApi } from "@/features/user";
import { Form, Input } from "antd";
import { WrappedFormUtils } from "antd/es/form/Form";
import React from "react";
import * as CONST from "./constants";

type Props = {
  decorator: WrappedFormUtils["getFieldDecorator"];
};

export default ({ decorator }: Props) => (
  <Form.Item>
    {decorator(CONST.USERNAME, {
      rules: [
        { min: 4, message: "Username is at least 4 character long" },
        { required: true, message: "Username is required" },
        {
          async validator(_, value: string, callback) {
            if (!value || value.length < 4) {
              return;
            }

            const available = await UserApi.isUsernameAvailable(value);
            if (!available) {
              callback(`${value} is already taken.`);
              return;
            }

            callback();
          }
        }
      ]
    })(<Input placeholder="Username" />)}
  </Form.Item>
);
