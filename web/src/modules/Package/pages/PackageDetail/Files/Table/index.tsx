import { MetaSelector } from "@/features/meta";
import { PkgSelector } from "@/features/package";
import { timeSinceUpload } from "@/libs/date";
import CalendarOutlined from "@ant-design/icons/CalendarOutlined";
import { Table } from "antd";
import { ColumnProps } from "antd/es/table";
import React from "react";
import { useSelector } from "react-redux";
import { useFileContext } from "../hooks";
import DeleteAction from "./DeleteAction";
import styles from "./styles.less";
import { DataRow } from "./types";

const pageSize = 20;

export default () => {
  const columns = useColumns();
  const data = useDataSource();

  const pagination =
    data.length > pageSize
      ? {
          style: { marginRight: 20 },
          pageSize,
        }
      : false;

  return (
    <Table
      className={styles.table}
      dataSource={data}
      columns={columns}
      pagination={pagination}
    />
  );
};

const useColumns = (): ColumnProps<DataRow>[] => {
  const { repository } = useSelector(MetaSelector.metaInfo);

  return [
    {
      title: "Name",
      dataIndex: "name",
      render: (text: string) => {
        const [, ...fileParts] = text.split("/");
        const filename = fileParts.join("/");

        const link = `${repository}/${text}`;
        return (
          <a className={styles.downloadLink} href={link}>
            {filename}
          </a>
        );
      },
    },
    {
      title: "Uploaded",
      dataIndex: "uploaded",
      render: (text) => (
        <>
          <CalendarOutlined />
          {text}
        </>
      ),
    },
    {
      title: "Downloads",
      dataIndex: "downloads",
      render: (text) => <b>{text}</b>,
    },
    {
      title: "Action",
      key: "action",
      render: (_, r) => (
        <DeleteAction channel={r.channel} package={r.package} />
      ),
    },
  ];
};

const useDataSource = () => {
  const { filters } = useFileContext();
  const { details, channel } = useSelector(PkgSelector.packageDetail);
  return details
    .filter((d) => {
      if (filters.version !== "All" && d.version !== filters.version)
        return false;
      return !(filters.platform !== "All" && d.platform !== filters.platform);
    })
    .map(
      (d, i) =>
        ({
          key: i,
          name: `${channel}/${d.platform}/${d.package}-${d.version}-${d.buildString}_${d.buildNumber}.tar.bz2`,
          uploaded: timeSinceUpload(d.uploadDate),
          downloads: d.count,
          channel,
          package: {
            name: d.package,
            version: d.version,
            platform: d.platform,
            buildNumber: d.buildNumber,
            buildString: d.buildString,
          },
        } as DataRow)
    );
};
