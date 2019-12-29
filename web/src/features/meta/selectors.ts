import { RootState } from "@/infrastructure/rootState";
import { MetaInfo } from "./types";

export const metaInfo = ({ meta: { loading, ...rest } }: RootState): MetaInfo =>
  rest;
