import { UserApi } from "@/features/user";
import { ThunkDispatchAsync } from "@/infrastructure/api";
import { Form, Typography } from "antd";
import { WrappedFormUtils } from "antd/lib/form/Form";
import { push } from "connected-react-router";
import React, { useState } from "react";
import { useDispatch } from "react-redux";
import * as CONST from "./constants";
import { LoginContext } from "./hooks";
import PasswordInput from "./PasswordInput";

import styles from "./styles.less";
import Submit from "./Submit";
import UsernameInput from "./UsernameInput";

type LoginFormProps = {
  form: WrappedFormUtils;
};

const LoginForm = ({ form }: LoginFormProps) => {
  const { getFieldValue } = form;
  const dispatch = useDispatch() as ThunkDispatchAsync;
  const [isValid, setIsValid] = useState(true);

  return (
    <LoginContext.Provider value={{ form, submit, valid: isValid }}>
      <Typography.Paragraph>Already a member? Sign in!</Typography.Paragraph>
      <UsernameInput />
      <PasswordInput />
      <Submit />
      {!isValid && (
        <span className={styles.error}>User Credentials are invalid</span>
      )}
    </LoginContext.Provider>
  );

  async function submit() {
    const username = getFieldValue(CONST.USERNAME);
    const password = getFieldValue(CONST.PASSWORD);

    const valid = await dispatch(UserApi.validateUser(username, password));
    if (valid) {
      dispatch(push("/"));
    } else {
      setIsValid(false);
    }
  }
};

export default Form.create({
  name: "login-form"
})(LoginForm);
