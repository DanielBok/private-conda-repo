import { Moment } from "moment";

export type Store = {
  packages: PackageMetaInfo[];
  loading: {
    packages: LoadingState;
    details: LoadingState;
    channelPackages: LoadingState;
  };
  packageDetail: PackageDetail<Moment>;
  channelPackages: ChannelPackages<Moment>;
};

export type Platform = "noarch" | "win-64" | "osx-64" | "linux-64";

export type ChannelPackages<T extends string | Moment> = {
  channel: string;
  email: string;
  joinDate: T;
  packages: PackageMetaInfo[];
};

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

export type PackageDetail<T extends string | Moment> = {
  channel: string;
  package: string;
  details: PackageCountInfo<T>[];
  latest: PackageMetaInfo;
};

export type PackageCountInfo<T extends string | Moment> = {
  channelId: number;
  package: string;
  buildString: string;
  buildNumber: number;
  version: string;
  platform: string;
  count: number;
  uploadDate: T;
};

export type RemovePackagePayload = {
  channel: string;
  password: string;
  package: {
    name: string;
    version: string;
    buildString: string;
    buildNumber: number;
    platform: string;
  };
};
