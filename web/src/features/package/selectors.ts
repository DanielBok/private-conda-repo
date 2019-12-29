import { RootState } from "@/infrastructure/rootState";

export const packageMeta = (state: RootState) => state.package.packages;
export const packageDetail = (state: RootState) => state.package.packageDetail;
