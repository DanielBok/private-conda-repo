import { PkgSelector } from "@/features/package";
import React, { useState } from "react";
import { useSelector } from "react-redux";
import Filters from "./Filters";
import { FileContext } from "./hooks";
import Table from "./Table";
import { Filter } from "./types";

export default () => {
  const [filters, setFilters] = useState<Filter>({
    platform: "All",
    version: "All",
  });
  const isAdmin = useSelector(PkgSelector.isAdmin);

  return (
    <FileContext.Provider
      value={{
        isAdmin,
        filters,
        setFilters: (f: Partial<Filter>) =>
          setFilters((prev) => ({ ...prev, ...f })),
      }}
    >
      <Filters />
      <Table />
    </FileContext.Provider>
  );
};
