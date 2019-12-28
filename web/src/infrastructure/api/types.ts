import { RootState } from "@/infrastructure/rootState";
import { AxiosRequestConfig, AxiosResponse } from "axios";
import { Action } from "redux";
import { ThunkDispatch } from "redux-thunk";

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

export type ThunkFunction<T = void> = (
  dispatch: ThunkDispatch<RootState, void, Action>,
  getState: () => RootState
) => T;

export type ThunkFunctionAsync<T = void> = (
  dispatch: ThunkDispatch<RootState, void, Action>,
  getState: () => RootState
) => Promise<T>;
