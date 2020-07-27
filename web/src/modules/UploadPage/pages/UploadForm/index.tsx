import { InboxOutlined } from "@ant-design/icons";
import { Button, Card, Checkbox, Form, message, Upload } from "antd";
import React, { useState } from "react";
import { useUpload } from "./hooks";
import styles from "./styles.less";

const FormItem = Form.Item;
const { Dragger } = Upload;

const UploadForm = () => {
  const [loading, setLoading] = useState(false);
  const [form] = Form.useForm();
  const upload = useUpload();

  return (
    <div className={styles.container}>
      <Card>
        <Form
          form={form}
          layout="horizontal"
          labelCol={{ span: 4 }}
          wrapperCol={{ xs: 20, sm: 20, md: 14 }}
          initialValues={{ abi: true }}
          onFinish={async () => {
            const { file, abi } = form.getFieldsValue();
            if (file === undefined || file.length === 0) {
              message.error("No file has been attached yet");
              return;
            }

            try {
              setLoading(true);
              const { data } = await upload(file[0], abi);
              const content: [string, string | number][] = [
                ["Name", data.name],
                ["Version", data.version],
                ["Platform", data.platform],
                ["Build String", data.buildString],
                ["Build Number", data.buildNumber],
              ];
              message.success(
                <div>
                  <p>Package uploaded successfully</p>
                  <div>Details</div>
                  <div className={styles.successMessage}>
                    {content.map(([label, value]) => (
                      <div key={label}>
                        <b>{label}: </b>
                        <span>{value}</span>
                      </div>
                    ))}
                  </div>
                </div>,
                8
              );
            } catch (e) {
              message.error("Upload was unsuccessful", 8);
            } finally {
              setLoading(false);
            }
          }}
        >
          <FormItem label="Package">
            <FormItem
              name="file"
              valuePropName="fileList"
              noStyle
              getValueFromEvent={(e) =>
                Array.isArray(e) ? e : e && e.fileList
              }
            >
              <Dragger
                accept=".tar.bz2"
                name="file"
                customRequest={({ file, onSuccess }) =>
                  setTimeout(() => onSuccess({}, file), 10)
                }
                onChange={({ file }) => {
                  form.setFieldsValue({ file: [file] });
                }}
              >
                <p>
                  <InboxOutlined />
                </p>
                <p>Click or drag file to this area to upload</p>
              </Dragger>
            </FormItem>
          </FormItem>

          <FormItem
            name="abi"
            label="Skip ABI"
            valuePropName="checked"
            help="If you don't know what this is, leave it as true"
          >
            <Checkbox />
          </FormItem>

          <Form.Item label=" " colon={false} className={styles.submitBox}>
            <Button type="primary" htmlType="submit" loading={loading}>
              Upload
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
};

export default UploadForm;
