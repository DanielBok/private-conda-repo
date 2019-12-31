import { RootState } from "@/infrastructure/rootState";

export const packageMeta = (state: RootState) => state.package.packages;
export const packageDetail = (state: RootState) => state.package.packageDetail;
export const isUserChannelAdmin = (state: RootState) =>
  state.user.validated &&
  state.user.username === state.package.packageDetail.channel;

export const channelPackages = (state: RootState) =>
  state.package.channelPackages;
