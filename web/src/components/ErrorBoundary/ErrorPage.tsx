import { PkgAction } from "@/features/package";
import { Button, Typography } from "antd";
import { push } from "connected-react-router";
import React, { useEffect } from "react";
import { useDispatch } from "react-redux";
import { withRouter } from "react-router-dom";
import NotFound from "./NotFound.png";
import styles from "./styles.less";

const ErrorPage = () => {
  const dispatch = useDispatch();

  useEffect(resets);

  return (
    <div className={styles.errorPage}>
      <img src={NotFound} alt="Page not found" />
      <h2>The page you're looking for does not exist</h2>
      <Typography.Text>
        Do you have permissions to view this page?
      </Typography.Text>

      <Button type="link" onClick={() => goto("/")}>
        Return home
      </Button>
    </div>
  );

  function goto(path: string) {
    dispatch(push(path));
  }

  function resets() {
    dispatch(PkgAction.resetLoadingStore());
  }
};

export default withRouter(ErrorPage);
