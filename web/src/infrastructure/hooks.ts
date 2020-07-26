import { isEqual } from "lodash";
import { useSelector } from "react-redux";
import { RootState } from "./rootState";

export const useRouter = () => useSelector((state: RootState) => state.router);

export const useRootSelector = <T = unknown>(
  selector: (state: RootState) => T
) => useSelector(selector, isEqual);
