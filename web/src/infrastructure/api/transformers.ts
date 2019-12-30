import { AxiosTransformer } from "axios";
import camelCaseKeys from "camelcase-keys";
import snakeCaseKeys from "snakecase-keys";

export const SnakeCaseKeysTransformer: AxiosTransformer = (data, headers) => {
  if (!isJson(headers)) return data;

  if (typeof data === "object" && !(data instanceof FormData)) {
    return snakeCaseKeys(data, { deep: true });
  }
  return data;
};

export const CamelCaseKeysTransformer: AxiosTransformer = (data, headers) => {
  if (!isJson(headers)) return data;

  if (typeof data === "object" && !(data instanceof FormData)) {
    if (Array.isArray(data) && typeof data[0] === "string") return data;
    return camelCaseKeys(data, { deep: true });
  }
  return data;
};

const isJson = (headers: Record<string, string>) => {
  for (const key of Object.keys(headers)) {
    if (key.toLowerCase() === "content-type") {
      return headers[key].toLowerCase().includes("application/json");
    }
  }
  return false;
};
