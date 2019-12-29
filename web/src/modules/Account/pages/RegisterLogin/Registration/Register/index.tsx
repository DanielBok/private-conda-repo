import { Form, Typography } from "antd";
import { WrappedFormUtils } from "antd/es/form/Form";
import React, { useState } from "react";
import ConfirmPasswordInput from "./ConfirmPasswordInput";
import { FormContext } from "./hooks";
import PasswordInput from "./PasswordInput";
import styles from "./styles.less";
import Submit from "./Submit";
import { ValidateStatus } from "./types";
import UsernameInput from "./UsernameInput";

type Props = {
  form: WrappedFormUtils;
};

const LoginForm = ({ form }: Props) => {
  const [validateStatus, setValidateStatus] = useState<ValidateStatus>("");

  return (
    <FormContext.Provider value={{ form, validateStatus, setValidateStatus }}>
      <Typography.Paragraph className={styles.welcome}>
        New to Private Conda Repo? Register a channel for yourself!
      </Typography.Paragraph>
      <UsernameInput />
      <PasswordInput />
      <ConfirmPasswordInput />
      <Submit />
    </FormContext.Provider>
  );
};

export default Form.create({
  name: "login-form"
})(LoginForm);
