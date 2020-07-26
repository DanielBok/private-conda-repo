import { ChnApi } from "@/features/channel";
import { useRootSelector } from "@/infrastructure/hooks";
import { Button } from "antd";
import React from "react";
import { useDispatch } from "react-redux";
import styles from "./styles.less";
import { useDisabled } from "./utils";

export default () => {
  const dispatch = useDispatch();
  const disabled = useDisabled();
  const isLoading = useRootSelector(
    (s) => s.channel.loading.validation === "REQUEST"
  );

  return (
    <Button
      disabled={disabled}
      type="primary"
      block
      size="large"
      className={styles.submitButton}
      onClick={() => dispatch(ChnApi.createChannel())}
      loading={isLoading}
    >
      Submit
    </Button>
  );
};
