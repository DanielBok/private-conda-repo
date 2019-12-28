import { Form, Input } from "antd";
import { WrappedFormUtils } from "antd/es/form/Form";
import React from "react";
import * as CONST from "./constants";

type Props = {
  form: WrappedFormUtils;
};

export default ({ form }: Props) => {
  const { getFieldDecorator, isFieldTouched, validateFields } = form;
  return (
    <Form.Item>
      {getFieldDecorator(CONST.PASSWORD, {
        rules: [
          { required: true, message: "Password cannot be empty" },
          {
            validator(_, value, callback) {
              if (value && isFieldTouched(CONST.PASSWORD)) {
                validateFields([CONST.CONFIRM], { force: true });
              }
              callback();
            }
          }
        ]
      })(<Input placeholder="Password" type="password" />)}
    </Form.Item>
  );
};
