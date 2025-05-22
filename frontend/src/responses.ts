export interface CpuStat {
  usedPercent: number;
}

export interface DiskStat {
  usedPercent: number;
  totalGbs: number;
  usedGbs: number;
}

export interface MemStat {
  usedPercent: number;
  totalGbs: number;
  usedGbs: number;
}

export interface HostInfo {
  cpuCores: number;
  cpuThreads: number;
  cpuName: string;
  cpuMhz: number;
  fsType: string;
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

export interface StatsResponse {
  cpuStat: CpuStat;
  diskStat: DiskStat;
  memStat: MemStat;
  hostInfo: HostInfo;
}
