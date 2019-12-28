import { push } from "connected-react-router";
import React from "react";
import styles from "./styles.less";
import { useDispatch } from "react-redux";
import { withRouter } from "react-router-dom";

const ErrorPage = () => {
  const dispatch = useDispatch();

  return (
    <div className={styles.errorPage}>
      <h1 className={styles.errorTitle}>
        The page you're looking for does not exist
      </h1>
      <div className={styles.errorButton} onClick={() => goto("/")}>
        Take me back to Homepage
      </div>
    </div>
  );

  function goto(path: string) {
    dispatch(push(path));
  }
};

export default withRouter(ErrorPage);
