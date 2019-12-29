import { useSelector } from "react-redux";
import { RootState } from "./rootState";

export const useRouter = () => useSelector((state: RootState) => state.router);
