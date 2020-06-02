import React, { useContext } from "react";
import { ContextType } from "./types";

export const FileContext = React.createContext<ContextType>({
  filters: {
    platform: "All",
    version: "All,",
  },
  setFilters: () => {},
  isAdmin: false,
});

export const useFileContext = () => useContext(FileContext);
