import { Form, Typography } from "antd";
import React, { FC, useState } from "react";
import ConfirmPasswordInput from "./ConfirmPasswordInput";
import EmailInput from "./EmailInput";
import { FormContext } from "./hooks";
import PasswordInput from "./PasswordInput";
import styles from "./styles.less";
import Submit from "./Submit";
import { ValidateStatus } from "./types";
import UsernameInput from "./UsernameInput";

const RegistrationForm: FC = () => {
  const [form] = Form.useForm();
  const [validateStatus, setValidateStatus] = useState<ValidateStatus>("");

  return (
    <Form name="registration-form" form={form}>
      <FormContext.Provider value={{ form, validateStatus, setValidateStatus }}>
        <Typography.Paragraph className={styles.welcome}>
          New to Private Conda Repo? Register a channel for yourself!
        </Typography.Paragraph>
        <UsernameInput />
        <PasswordInput />
        <ConfirmPasswordInput />
        <EmailInput />
        <Submit />
      </FormContext.Provider>
    </Form>
  );
};

export default RegistrationForm;
