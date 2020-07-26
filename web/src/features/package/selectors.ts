import { RootState } from "@/infrastructure/rootState";

export const packageMeta = (state: RootState) => state.pkg.packages;
export const packageDetail = (state: RootState) => state.pkg.packageDetail;
export const isAdmin = (state: RootState) =>
  state.channel.validated &&
  state.channel.channel === state.pkg.packageDetail.channel;

export const channelPackages = (state: RootState) =>
  state.pkg.channelPackages;
