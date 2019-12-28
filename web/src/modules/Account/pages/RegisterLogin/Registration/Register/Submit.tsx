import { Button } from "antd";
import { WrappedFormUtils } from "antd/es/form/Form";
import React from "react";
import * as CONST from "./constants";
import styles from "./styles.less";

type Props = {
  form: WrappedFormUtils;
  onClick: () => void;
};

export default ({ form, onClick }: Props) => {
  const { getFieldsError, isFieldTouched: T } = form;
  const touched = T(CONST.USERNAME) && T(CONST.PASSWORD) && T(CONST.CONFIRM);

  const disabled =
    !touched ||
    Object.values(getFieldsError()).reduce(
      (acc, e) => acc || e !== undefined,
      false as boolean
    );

  return (
    <Button
      disabled={disabled}
      type="primary"
      block
      className={styles.submitButton}
      onClick={onClick}
    >
      Submit
    </Button>
  );
};
