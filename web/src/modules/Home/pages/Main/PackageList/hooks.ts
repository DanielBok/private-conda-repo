import React, { useContext } from "react";

export const SearchContext = React.createContext<{
  search: string;
  setSearch: (v: string) => void;
}>({
  search: "",
  setSearch: () => {},
});

export const useSearchContext = () => useContext(SearchContext);
