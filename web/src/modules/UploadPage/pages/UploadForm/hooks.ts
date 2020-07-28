import { useRootSelector } from "@/infrastructure/hooks";
import { UploadFile } from "antd/es/upload/interface";
import axios, { AxiosResponse } from "axios";

const url =
  (process.env.REACT_APP_API_URL ??
    `${window.location.protocol}//${window.location.hostname}:5060`) + "/p";

type Response = {
  buildNumber: number;
  buildString: string;
  name: string;
  platform: string;
  version: string;
};

export const useUpload = () => {
  const [
    channel,
    password,
  ] = useRootSelector(({ channel: { channel, password } }) => [
    channel,
    password,
  ]);

  return async (
    { originFileObj: file }: UploadFile,
    noAbi: boolean
  ): Promise<AxiosResponse<Response>> => {
    const data = new FormData();
    data.append("channel", channel);
    data.append("password", password);
    data.append("file", file as File);

    const fixes = [noAbi && "no-abi"].filter((e) => e).join(",");
    if (fixes) data.append("fixes", fixes);

    return await axios.post(url, data);
  };
};
