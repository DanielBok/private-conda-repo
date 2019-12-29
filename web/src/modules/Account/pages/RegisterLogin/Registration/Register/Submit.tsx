import { UserApi } from "@/features/user";
import { Button } from "antd";
import React from "react";
import { useDispatch } from "react-redux";
import * as CONST from "./constants";
import { useFormContext } from "./hooks";
import styles from "./styles.less";

export default () => {
  const dispatch = useDispatch();

  const {
    form: { getFieldsError: E, isFieldTouched: T, getFieldValue: G },
    validateStatus
  } = useFormContext();
  const touched = T(CONST.USERNAME) && T(CONST.PASSWORD) && T(CONST.CONFIRM);

  const disabled =
    !touched ||
    Object.values(E()).reduce(
      (acc, e) => acc || e !== undefined,
      validateStatus !== "success"
    );

  return (
    <Button
      disabled={disabled}
      type="primary"
      block
      size="large"
      className={styles.submitButton}
      onClick={() => {
        const username = G(CONST.USERNAME);
        const password = G(CONST.PASSWORD);

        dispatch(UserApi.createUser(username, password));
      }}
    >
      Submit
    </Button>
  );
};
