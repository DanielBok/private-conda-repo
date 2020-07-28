import Markdown from "@/components/Markdown";
import { useRootSelector } from "@/infrastructure/hooks";
import Layout from "@/modules/Help/Layout";
import { Divider } from "antd";
import { reduce } from "lodash";
import React from "react";
import { CopyBlock, dracula } from "react-code-blocks";
import Content from "./upload.md";

const useOverrides = () => {
  const registry = useRootSelector((s) => s.meta.registry);

  return {
    Divider,
    BasicRequest: function BasicRequest() {
      return (
        <div>
          <CopyBlock
            language="typescript"
            text={`
const data = new FormData();
data.append("channel", channel);
data.append("password", password);
data.append("file", file as File);
data.append("fixes", "no-abi"); // comma-separated
axios
  .post("@url", data)
  .then((resp) => alert(resp.data))
  .catch((e) => console.error(e));
  `
              .trim()
              .replace("@url", registry)}
            showLineNumbers={true}
            theme={dracula}
            wrapLines={true}
            codeBlock
          />
        </div>
      );
    },
  };
};

const Upload = () => {
  const overrides = useOverrides();
  const { protocol, host } = window.location;

  const content = reduce(
    {
      "@link": `${protocol}//${host}/upload`,
    },
    (acc, search, replace) => acc.replace(search, replace),
    Content
  );

  return (
    <Layout>
      <Markdown children={content} overrides={overrides} />
    </Layout>
  );
};

export default Upload;
