import { Form, Input } from "antd";
import { WrappedFormUtils } from "antd/es/form/Form";
import React from "react";
import * as CONST from "./constants";

type Props = {
  form: WrappedFormUtils;
};

export default ({ form }: Props) => {
  const { getFieldDecorator, getFieldValue } = form;

  return (
    <Form.Item>
      {getFieldDecorator(CONST.CONFIRM, {
        rules: [
          {
            validator(_, value: string, callback) {
              if (value && value !== getFieldValue(CONST.PASSWORD)) {
                callback("Passwords do not match");
                return;
              }
              callback();
            }
          }
        ]
      })(<Input placeholder="Password" type="password" />)}
    </Form.Item>
  );
};
