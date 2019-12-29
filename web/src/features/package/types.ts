export type Store = {
  packages: PackageMetaInfo[];
  loading: {
    packages: LoadingState;
    details: LoadingState;
  };
  selected: {
    channel: string;
    package: string;
    details: PackageDetail[];
    latestVersion: string;
  };
};

export type Platform = "noarch" | "win-64" | "osx-64" | "linux-64";

export type PackageMetaInfo = {
  channel: string;
  platforms: Platform[];
  version: string | null;
  description: string | null;
  devUrl: string | null;
  docUrl: string | null;
  home: string | null;
  license: string | null;
  summary: string | null;
  timestamp: number;
  name: string;
};

export type PackageDetail = {
  channel: string;
  package: string;
  buildString: string;
  buildNumber: number;
  version: string;
  platform: string;
  count: number;
  uploadDate: string;
};
