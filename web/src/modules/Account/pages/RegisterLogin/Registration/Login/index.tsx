import { Typography } from "antd";
import React, { FC } from "react";
import { LoginContext } from "./hooks";
import PasswordInput from "./PasswordInput";
import { useLoginReducer } from "./reducer";
import styles from "./styles.less";
import Submit from "./Submit";
import UsernameInput from "./UsernameInput";

const LoginForm: FC = () => {
  const [state, dispatch] = useLoginReducer();

  return (
    <LoginContext.Provider value={{ state, dispatch }}>
      <Typography.Paragraph>Already a member? Sign in!</Typography.Paragraph>
      <UsernameInput />
      <PasswordInput />
      <Submit />
      {!state.valid && (
        <span className={styles.error}>User credentials are invalid</span>
      )}
    </LoginContext.Provider>
  );
};

export default LoginForm;
