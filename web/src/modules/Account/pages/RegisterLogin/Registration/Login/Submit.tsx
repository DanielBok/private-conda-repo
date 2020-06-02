import { Button } from "antd";
import React from "react";
import { useLoginContext, useSubmit } from "./hooks";
import styles from "./styles.less";

export default () => {
  const { disabled } = useLoginContext().state;
  const submit = useSubmit();

  return (
    <Button
      type="primary"
      block
      size="large"
      className={styles.submitButton}
      onClick={submit}
      disabled={disabled}
    >
      Submit
    </Button>
  );
};
