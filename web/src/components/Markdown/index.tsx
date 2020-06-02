import { Typography } from "antd";
import cx from "classnames";
import MarkdownToJS, { MarkdownOptions } from "markdown-to-jsx";
import React, { FC } from "react";

import styles from "./styles.less";

interface IProps extends MarkdownOptions {
  className?: string;
}

const Markdown: FC<IProps> = ({
  className,
  overrides = {},
  children,
  ...options
}) => {
  overrides = {
    ...{
      h1: {
        component: Typography.Title,
        props: { level: 2 },
      },
      h2: {
        component: Typography.Title,
        props: { level: 3 },
      },
      div: {
        component: Typography.Paragraph,
      },
    },
    ...overrides,
  };

  return (
    <div className={cx(styles.defaultContainer, className)}>
      <MarkdownToJS options={{ overrides, ...options }} children={children} />
    </div>
  );
};

export default Markdown;
