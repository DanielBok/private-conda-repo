import React, { useState } from "react";
import { SearchContext } from "./hooks";
import ResultTable from "./ResultTable";
import SearchBar from "./SearchBar";

export default () => {
  const [search, setSearch] = useState<string>("");

  return (
    <SearchContext.Provider value={{ search, setSearch }}>
      <SearchBar />
      <ResultTable />
    </SearchContext.Provider>
  );
};
