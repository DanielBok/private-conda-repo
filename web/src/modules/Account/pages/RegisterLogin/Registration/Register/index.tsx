import { ChnAction } from "@/features/channel";
import { Typography } from "antd";
import React, { useEffect } from "react";
import { useDispatch } from "react-redux";
import ChannelInput from "./ChannelInput";
import ConfirmInput from "./ConfirmInput";
import EmailInput from "./EmailInput";
import PasswordInput from "./PasswordInput";
import styles from "./styles.less";
import Submit from "./Submit";

const RegistrationForm = () => {
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(ChnAction.resetForm());
    // eslint-disable-next-line
  }, []);

  return (
    <>
      <Typography.Paragraph className={styles.welcome}>
        New to Private Conda Repo? Register a channel for yourself!
      </Typography.Paragraph>
      <ChannelInput />
      <PasswordInput />
      <ConfirmInput />
      <EmailInput />
      <Submit />
    </>
  );
};

export default RegistrationForm;
