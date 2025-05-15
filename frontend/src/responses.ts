export interface CpuInfo {
  cpuPercent: number;
  cores: number;
  modelName: string;
  mhz: number;
}

export interface MemInfo {
  availableBytes: number;
  totalBytes: number;
  usedBytes: number;
  usedPercent: number;
}

export interface HostInfo {
  hostName: string;
  uptime: number;
  bootTime: number;
  processes: number;
  os: string;
  platform: string;
  platformFamily: string;
  platformVersion: string;
  kernelVersion: string;
  kernelArch: string;
  virtualizationSystem: string;
  virtualizationRole: string;
}

export interface DiskInfo {
  fsType: string;
  totalBytes: number;
  freeBytes: number;
  usedBytes: number;
  usedPercent: number;
}

export interface StatsResponse {
  cpuInfo: CpuInfo;
  memInfo: MemInfo;
  hostInfo: HostInfo;
  diskInfo: DiskInfo;
}
