export type Store = {
  packages: PackageMetaInfo[];
  loading: LoadingState;
};

type Platform = "noarch" | "win-64" | "osx-64" | "linux-64";

export type PackageMetaInfo = {
  channel: string;
  platforms: Platform[];
  version: string | null;
  description: string | null;
  dev_url: string | null;
  doc_url: string | null;
  home: string | null;
  license: string | null;
  summary: string | null;
  timestamp: number;
  name: string;
};
