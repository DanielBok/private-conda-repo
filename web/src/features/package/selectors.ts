import { RootState } from "@/infrastructure/rootState";

export const packageMeta = (state: RootState) => state.package.packages;
