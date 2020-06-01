import React, { useContext } from "react";
import * as Types from "./types";

export const PackageContext = React.createContext<Types.MatchParams>({
  channel: "",
  pkg: "",
});

export const usePackageContext = () => useContext(PackageContext);
