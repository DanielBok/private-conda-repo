import { PackageType } from "@/features/package";

export type DataRow = {
  key: number;
  name: string;
  uploaded: string;
  downloads: number;
  channel: string;
  package: PackageType.RemovePackagePayload["package"];
};
