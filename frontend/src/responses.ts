export interface StatsResponse {
  cpu_info: CpuInfo[];
  cpu_percent: number[];
  disk: Disk;
  host_info: HostInfo;
  memory: Memory;
}

export interface CpuInfo {
  cpu: number;
  vendorId: string;
  family: string;
  model: string;
  stepping: number;
  physicalId: string;
  coreId: string;
  cores: number;
  modelName: string;
  mhz: number;
  cacheSize: number;
  flags: any;
  microcode: string;
}

export interface Disk {
  path: string;
  fstype: string;
  total: number;
  free: number;
  used: number;
  usedPercent: number;
  inodesTotal: number;
  inodesUsed: number;
  inodesFree: number;
  inodesUsedPercent: number;
}

export interface HostInfo {
  hostname: string;
  uptime: number;
  bootTime: number;
  procs: number;
  os: string;
  platform: string;
  platformFamily: string;
  platformVersion: string;
  kernelVersion: string;
  kernelArch: string;
  virtualizationSystem: string;
  virtualizationRole: string;
  hostId: string;
}

export interface Memory {
  available: number;
  total: number;
  used: number;
  used_percent: number;
}
