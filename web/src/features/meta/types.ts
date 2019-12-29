export type Store = MetaInfo & {
  loading: LoadingState;
};

export type MetaInfo = {
  image: string;
  registry: string;
  repository: string;
};
