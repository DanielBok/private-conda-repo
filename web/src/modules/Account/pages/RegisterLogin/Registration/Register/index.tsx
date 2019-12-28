import { Form, Typography } from "antd";
import { WrappedFormUtils } from "antd/es/form/Form";
import React from "react";
import ConfirmPasswordInput from "./ConfirmPasswordInput";
import * as CONST from "./constants";
import PasswordInput from "./PasswordInput";
import styles from "./styles.less";
import Submit from "./Submit";
import UsernameInput from "./UsernameInput";

type Props = {
  form: WrappedFormUtils;
};

const LoginForm = ({ form }: Props) => {
  const { getFieldDecorator, getFieldValue } = form;

  return (
    <div>
      <Typography.Paragraph className={styles.welcome}>
        New to Private Conda Repo? Register a channel for yourself!
      </Typography.Paragraph>
      <UsernameInput decorator={getFieldDecorator} />
      <PasswordInput form={form} />
      <ConfirmPasswordInput form={form} />
      <Submit form={form} onClick={submit} />
    </div>
  );

  function submit() {
    const username = getFieldValue(CONST.USERNAME);
    const password = getFieldValue(CONST.PASSWORD);
    console.log(username, password);
  }
};

export default Form.create({
  name: "login-form"
})(LoginForm);
