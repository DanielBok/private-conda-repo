import { useRouter } from "@/infrastructure/hooks";
import React from "react";
import { Link } from "react-router-dom";
import cx from "classnames";

import styles from "./styles.less";

const routes = {
  "Getting Started": {
    Overview: "/"
  }
};

export default () => {
  const pathname = useRouter().location.pathname.replace(/^\/help/, "");

  return (
    <>
      {Object.entries(routes).map(([title, links], i) => (
        <div key={i}>
          <div className={styles.title}>{title}</div>
          {Object.entries(links).map(([name, link], j) => (
            <Link
              key={j}
              to={"/help" + link}
              className={cx(styles.link, pathname === link && styles.selected)}
            >
              {name}
            </Link>
          ))}
        </div>
      ))}
    </>
  );
};
