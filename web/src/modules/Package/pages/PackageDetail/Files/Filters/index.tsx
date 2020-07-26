import { PkgSelector } from "@/features/package";
import FilterOutlined from "@ant-design/icons/FilterOutlined";
import { Col, Form, Row, Select } from "antd";
import React from "react";
import { useSelector } from "react-redux";
import { useFileContext } from "../hooks";
import styles from "./styles.less";

const { Option } = Select;

export default () => {
  const { filters, setFilters } = useFileContext();
  const { details } = useSelector(PkgSelector.packageDetail);

  const versions = ["All", ...new Set(details.map((e) => e.version))];
  const platforms = ["All", ...new Set(details.map((e) => e.platform))];

  return (
    <fieldset className={styles.filterFieldset}>
      <legend>
        <FilterOutlined /> Filters
      </legend>
      <Row gutter={12}>
        <Col xs={24} md={12}>
          <Form.Item label="Platform">
            <Select
              className={styles.select}
              value={filters.platform}
              onChange={(platform: string) => setFilters({ platform })}
            >
              {platforms.map((p, k) => (
                <Option value={p} key={k}>
                  {p}
                </Option>
              ))}
            </Select>
          </Form.Item>
        </Col>
        <Col xs={24} md={12}>
          <Form.Item label="Version">
            <Select
              className={styles.select}
              value={filters.version}
              onChange={(version: string) => setFilters({ version })}
            >
              {versions.map((p, k) => (
                <Option value={p} key={k}>
                  {p}
                </Option>
              ))}
            </Select>
          </Form.Item>
        </Col>
      </Row>
    </fieldset>
  );
};
