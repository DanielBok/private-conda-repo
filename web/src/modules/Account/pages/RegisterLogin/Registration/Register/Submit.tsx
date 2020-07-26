import { useRootSelector } from "@/infrastructure/hooks";
import { Button } from "antd";
import React from "react";
import { useRegistrationContext, useSubmit } from "./hooks";
import styles from "./styles.less";

export default () => {
  const isLoading = useRootSelector(
    (s) => s.channel.loading.validation === "REQUEST"
  );

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
