import { RootState } from "@/infrastructure/rootState";

export const channelInfo = (state: RootState) => state.channel;
export const channelValidated = (state: RootState) => state.channel.validated;
