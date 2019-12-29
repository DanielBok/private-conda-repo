import { Button } from "antd";
import React from "react";
import * as CONST from "./constants";
import { useLoginContext } from "./hooks";
import styles from "./styles.less";

export default () => {
  const {
    form: { getFieldsError: E, isFieldTouched: T },
    submit
  } = useLoginContext();
  const touched = T(CONST.USERNAME) && T(CONST.PASSWORD);

  const disabled = Object.values(E()).reduce(
    (acc, e) => acc || e !== undefined,
    !touched
  );

  return (
    <Button
      disabled={disabled}
      type="primary"
      block
      size="large"
      className={styles.submitButton}
      onClick={submit}
    >
      Submit
    </Button>
  );
};
