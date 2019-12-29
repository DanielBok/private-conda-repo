import { Form, Icon, Input } from "antd";
import React from "react";
import { useSearchContext } from "../hooks";
import styles from "./styles.less";

export default () => {
  const { search, setSearch } = useSearchContext();

  return (
    <Form.Item className={styles.searchBox}>
      <Input
        value={search}
        onChange={e => setSearch(e.target.value)}
        placeholder="Search Private Conda Repo"
        size="large"
        addonAfter={<Icon type="search" />}
      />
    </Form.Item>
  );
};
