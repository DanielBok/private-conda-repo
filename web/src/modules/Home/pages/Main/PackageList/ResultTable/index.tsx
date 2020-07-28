import { PkgSelector, PkgType } from "@/features/package";
import { Col, List, Row } from "antd";
import Fuse from "fuse.js";
import sortBy from "lodash/sortBy";
import React from "react";
import { useSelector } from "react-redux";
import { useSearchContext } from "../hooks";
import Link from "./Link";
import styles from "./styles.less";
import Tag from "./Tag";

export default () => {
  const list = useResultList();

  const pagination =
    list.length <= 10
      ? false
      : {
          pageSize: 10,
        };

  return (
    <List
      itemLayout="vertical"
      size="large"
      pagination={pagination}
      dataSource={list}
      header={<Header />}
      bordered={true}
      className={styles.listMain}
      renderItem={(item, i) => {
        const className = i % 2 === 1 ? styles.alternate : undefined;

        return (
          <List.Item key={i} className={className}>
            <Row gutter={12}>
              <Col span={18}>
                <List.Item.Meta
                  title={<Link name={item.name} channel={item.channel} />}
                  description={item.description}
                />
              </Col>
              <Col span={6}>
                {item.platforms.map((p) => (
                  <Tag platform={p} key={p} />
                ))}
              </Col>
            </Row>
          </List.Item>
        );
      }}
    />
  );
};

const Header = () => (
  <Row gutter={12} className={styles.header}>
    <Col span={18}>Package (owner / package)</Col>
    <Col span={6}>Platforms</Col>
  </Row>
);

const useResultList = () => {
  const { search } = useSearchContext();
  const packages = useSelector(PkgSelector.packageMeta);

  if (search.length > 0) {
    const keys: {
      name: keyof PkgType.PackageMetaInfo;
      weight: number;
    }[] = [
      { name: "name", weight: 0.6 },
      { name: "description", weight: 0.25 },
      { name: "channel", weight: 0.1 },
      { name: "summary", weight: 0.05 },
    ];

    return new Fuse(packages, { threshold: 0.2, keys })
      .search(search)
      .map((e) => e.item);
  } else {
    return sortBy(packages, (e) => [e.channel, e.name]);
  }
};
