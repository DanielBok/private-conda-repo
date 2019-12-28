import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
import * as ApiTypes from "./types";

const apiUrl = process.env.REACT_APP_API_URL;

export class API {
  private client: AxiosInstance;

  constructor(config?: AxiosRequestConfig) {
    const { baseURL = apiUrl, ...rest } = config || {};

    this.client = axios.create({ baseURL, ...rest });
  }

  async Get<R>(endpoint: string, config?: ApiTypes.RequestConfig<R>) {
    return await this.execute<R>("GET", endpoint, {}, config);
  }

  async Post<R>(
    endpoint: string,
    payload: any,
    config?: ApiTypes.RequestConfig<R>
  ) {
    return await this.execute<R>("POST", endpoint, payload, config);
  }

  async Put<R>(
    endpoint: string,
    payload: any,
    config?: ApiTypes.RequestConfig<R>
  ) {
    return await this.execute<R>("PUT", endpoint, payload, config);
  }

  async Delete<R>(endpoint: string, config?: ApiTypes.RequestConfig<R>) {
    return await this.execute<R>("DELETE", endpoint, {}, config);
  }

  private async execute<R>(
    method: "GET" | "PUT" | "POST" | "DELETE",
    endpoint: string,
    payload?: any,
    config?: ApiTypes.RequestConfig<R>
  ): Promise<AxiosResponse<R>> {
    const {
      onError = [],
      onSuccess = [],
      beforeRequest = [],
      afterResponse = [],
      returnErrorResponse = true,
      ...req
    } = config || {};

    runFunctionHandlers(beforeRequest);

    try {
      const result = await this.axiosCall<R>(method, endpoint, payload, req);

      runSuccessHandlers(result.data, onSuccess);
      return result;
    } catch (e) {
      runErrorHandlers(e.response, onError);
      if (returnErrorResponse) return e.response;

      throw e;
    } finally {
      // Run after response function only if there are no errors
      runFunctionHandlers(afterResponse);
    }
  }

  private async axiosCall<R>(
    method: "GET" | "PUT" | "POST" | "DELETE",
    endpoint: string,
    payload: any,
    config: AxiosRequestConfig
  ): Promise<AxiosResponse<R>> {
    endpoint = endpoint.replace(/^\/*/, "");

    switch (method) {
      case "GET":
        return await this.client.get<R>(endpoint, config);
      case "PUT":
        return await this.client.put<R>(endpoint, payload, config);
      case "POST":
        return await this.client.post<R>(endpoint, payload, config);
      case "DELETE":
        return await this.client.delete<R>(endpoint, config);
    }
  }
}

function runFunctionHandlers(
  funcs?: ApiTypes.BeforeRequestFunction | ApiTypes.BeforeRequestFunction[]
) {
  if (funcs === undefined) return;

  if (typeof funcs === "function") {
    funcs();
  } else {
    funcs.forEach(f => runFunctionHandlers(f));
  }
}

function runSuccessHandlers<R>(
  data: R,
  funcs?: ApiTypes.SuccessFunction<R> | ApiTypes.SuccessFunction<R>[]
) {
  if (funcs === undefined) return;

  if (typeof funcs === "function") {
    funcs(data);
  } else {
    funcs.forEach(f => runSuccessHandlers(data, f));
  }
}

function runErrorHandlers(
  error: AxiosResponse,
  funcs?: ApiTypes.ErrorFunction | ApiTypes.ErrorFunction[]
) {
  if (funcs === undefined) return;

  if (typeof funcs === "function") {
    funcs(error);
  } else {
    funcs.forEach(f => runErrorHandlers(error, f));
  }
}

export default new API();

export type ThunkFunction = ApiTypes.ThunkFunction;
export type ThunkFunctionAsync = ApiTypes.ThunkFunctionAsync;
