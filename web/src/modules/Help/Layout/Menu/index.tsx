import { useRouter } from "@/infrastructure/hooks";
import cx from "classnames";
import React from "react";
import { Link } from "react-router-dom";

import styles from "./styles.less";

const routes = {
  "Getting Started": {
    Overview: "/",
    Upload: "/upload",
  },
};

const Menu = () => {
  const pathname = useRouter().location.pathname.replace(/^\/help/, "");

  return (
    <div className={styles.groupContainer}>
      {Object.entries(routes).map(([title, links], i) => (
        <div key={i}>
          <div className={styles.title}>{title}</div>
          {Object.entries(links).map(([name, link], j) => (
            <div key={j} className={styles.linkContainer}>
              <Link
                to={"/help" + link}
                className={cx(
                  styles.link,
                  pathname === link && styles.selected
                )}
              >
                {name}
              </Link>
            </div>
          ))}
        </div>
      ))}
    </div>
  );
};

export default Menu;
