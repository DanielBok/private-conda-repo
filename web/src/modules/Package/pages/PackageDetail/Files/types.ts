export type ContextType = {
  isAdmin: boolean;
  filters: Filter;
  setFilters: (f: Partial<Filter>) => void;
};

export type Filter = {
  platform: "All" | string;
  version: "All" | string;
};
