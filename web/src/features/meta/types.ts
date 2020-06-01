export type Store = MetaInfo & {
  loading: LoadingState;
};

export type MetaInfo = {
  indexer: "shell" | "docker";
  image: string;
  registry: string;
  repository: string;
};
