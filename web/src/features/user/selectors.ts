import { RootState } from "@/infrastructure/rootState";

export const userInfo = (state: RootState) => state.user;
export const userValidated = (state: RootState) => state.user.validated;
