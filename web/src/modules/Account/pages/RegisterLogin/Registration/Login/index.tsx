import { UserApi } from "@/features/user";
import { ThunkDispatchAsync } from "@/infrastructure/api";
import { Form, Typography } from "antd";
import { push } from "connected-react-router";
import React, { FC, useState } from "react";
import { useDispatch } from "react-redux";
import * as CONST from "./constants";
import { LoginContext } from "./hooks";
import PasswordInput from "./PasswordInput";

import styles from "./styles.less";
import Submit from "./Submit";
import UsernameInput from "./UsernameInput";

const LoginForm: FC = () => {
  const [form] = Form.useForm();
  const dispatch = useDispatch() as ThunkDispatchAsync;
  const [isValid, setIsValid] = useState(true);

  return (
    <Form name="login-form" form={form}>
      <LoginContext.Provider value={{ form, submit, valid: isValid }}>
        <Typography.Paragraph>Already a member? Sign in!</Typography.Paragraph>
        <UsernameInput />
        <PasswordInput />
        <Submit />
        {!isValid && (
          <span className={styles.error}>User Credentials are invalid</span>
        )}
      </LoginContext.Provider>
    </Form>
  );

  async function submit() {
    const username = form.getFieldValue(CONST.USERNAME);
    const password = form.getFieldValue(CONST.PASSWORD);

    const valid = await dispatch(UserApi.validateUser(username, password));
    if (valid) {
      dispatch(push("/"));
    } else {
      setIsValid(false);
    }
  }
};

export default LoginForm;
