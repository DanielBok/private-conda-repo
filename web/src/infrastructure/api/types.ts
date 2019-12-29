import { AxiosRequestConfig, AxiosResponse } from "axios";

export type BeforeRequestFunction = () => void;
export type AfterResponseFunction = BeforeRequestFunction;
export type SuccessFunction<R> = (data: R) => void;
export type ErrorFunction = (e: AxiosResponse) => void;

export interface RequestConfig<R> extends AxiosRequestConfig {
  afterResponse?: AfterResponseFunction | AfterResponseFunction[];
  beforeRequest?: BeforeRequestFunction | BeforeRequestFunction[];
  onError?: ErrorFunction | ErrorFunction[];
  onSuccess?: SuccessFunction<R>;
  returnErrorResponse?: boolean;
}
