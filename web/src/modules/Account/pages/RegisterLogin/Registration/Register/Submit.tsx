import { UserSelector } from "@/features/user";
import { Button } from "antd";
import React from "react";
import { useSelector } from "react-redux";
import { useRegistrationContext, useSubmit } from "./hooks";
import styles from "./styles.less";

export default () => {
  const isLoading = useSelector(UserSelector.userInfo).loading === "REQUEST";

  const { disabled } = useRegistrationContext().state;
  const submit = useSubmit();

  return (
    <Button
      disabled={disabled}
      type="primary"
      block
      size="large"
      className={styles.submitButton}
      onClick={submit}
      loading={isLoading}
    >
      Submit
    </Button>
  );
};
