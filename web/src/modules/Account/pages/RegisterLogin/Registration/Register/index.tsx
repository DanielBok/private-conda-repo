import { Typography } from "antd";
import React, { FC } from "react";
import ConfirmInput from "./ConfirmInput";
import EmailInput from "./EmailInput";
import { RegistrationContext } from "./hooks";
import PasswordInput from "./PasswordInput";
import { useRegistrationReducer } from "./reducer";
import styles from "./styles.less";
import Submit from "./Submit";
import UsernameInput from "./UsernameInput";

const RegistrationForm: FC = () => {
  const [state, dispatch] = useRegistrationReducer();

  return (
    <RegistrationContext.Provider value={{ state, dispatch }}>
      <Typography.Paragraph className={styles.welcome}>
        New to Private Conda Repo? Register a channel for yourself!
      </Typography.Paragraph>
      <UsernameInput />
      <PasswordInput />
      <ConfirmInput />
      <EmailInput />
      <Submit />
    </RegistrationContext.Provider>
  );
};

export default RegistrationForm;
