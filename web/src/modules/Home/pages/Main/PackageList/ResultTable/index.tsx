import { PackageSelector, PackageType } from "@/features/package";
import { Col, List, Row } from "antd";
import Fuse from "fuse.js";
import React from "react";
import { useSelector } from "react-redux";
import { useSearchContext } from "../hooks";
import Tag from "./Tag";
import styles from "./styles.less";
import sortBy from "lodash/sortBy";

export default () => {
  const list = useResultList();

  const pagination =
    list.length <= 10
      ? false
      : {
          pageSize: 10
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
                  title={`${item.channel} / ${item.name}`}
                  description={item.description}
                />
              </Col>
              <Col span={6}>
                {item.platforms.map(p => (
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
  const packages = useSelector(PackageSelector.packageMeta);
  const fuse = useFuse(packages);

  if (search.length > 0) {
    return fuse.search(search).map(e => e.item);
  } else {
    return sortBy(packages, e => [e.channel, e.name]);
  }
};

const useFuse = (packages: PackageType.PackageMetaInfo[]) => {
  const keys: {
    name: keyof PackageType.PackageMetaInfo;
    weight: number;
  }[] = [
    { name: "name", weight: 0.6 },
    { name: "description", weight: 0.25 },
    { name: "channel", weight: 0.1 },
    { name: "summary", weight: 0.05 }
  ];

  return new Fuse(packages, {
    shouldSort: true,
    includeScore: true,
    threshold: 0.2,
    tokenize: true,
    location: 0,
    distance: 100,
    maxPatternLength: 32,
    minMatchCharLength: 1,
    keys
  });
};
