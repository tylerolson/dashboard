package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/shirou/gopsutil/v4/common"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

type CpuStat struct {
	UsedPercent float64 `json:"usedPercent"`
}

type DiskStat struct {
	UsedPercent float64 `json:"usedPercent"`
	TotalGbs    float64 `json:"totalGbs"`
	UsedGbs     float64 `json:"usedGbs"`
}

type MemStat struct {
	UsedPercent float64 `json:"usedPercent"`
	TotalGbs    float64 `json:"totalGbs"`
	UsedGbs     float64 `json:"usedGbs"`
}

type HostInfo struct {
	CpuCores             int     `json:"cpuCores"`
	CpuThreads           int     `json:"cpuThreads"`
	CpuName              string  `json:"cpuName"`
	CpuMhz               float64 `json:"cpuMhz"`
	FSType               string  `json:"fsType"`
	HostName             string  `json:"hostName"`
	Uptime               uint64  `json:"uptime"`
	BootTime             uint64  `json:"bootTime"`
	Processes            uint64  `json:"processes"`
	OS                   string  `json:"os"`
	Platform             string  `json:"platform"`
	PlatformFamily       string  `json:"platformFamily"`
	PlatformVersion      string  `json:"platformVersion"`
	KernelVersion        string  `json:"kernelVersion"`
	KernelArch           string  `json:"kernelArch"`
	VirtualizationSystem string  `json:"virtualizationSystem"`
	VirtualizationRole   string  `json:"virtualizationRole"`
}

type StatsResponse struct {
	CpuStat  CpuStat  `json:"cpuStat"`
	DiskStat DiskStat `json:"diskStat"`
	MemStat  MemStat  `json:"memStat"`
	HostInfo HostInfo `json:"hostInfo"`
}

func fetchStats(ctx context.Context) StatsResponse {
	cpuInfos, _ := cpu.InfoWithContext(ctx)
	coreCount, _ := cpu.CountsWithContext(ctx, false)
	threadCount, _ := cpu.CountsWithContext(ctx, true)
	cpuPercentages, _ := cpu.PercentWithContext(ctx, time.Second, false)
	virtualMemory, _ := mem.VirtualMemoryWithContext(ctx)
	hostInfo, _ := host.InfoWithContext(ctx)
	diskStat, _ := disk.UsageWithContext(ctx, "/")

	response := StatsResponse{
		CpuStat: CpuStat{
			UsedPercent: math.Round(cpuPercentages[0]*100) / 100,
		},
		DiskStat: DiskStat{
			UsedPercent: math.Round(diskStat.UsedPercent*100) / 100,
			TotalGbs:    math.Round(float64(diskStat.Total)/1e9*100) / 100,
			UsedGbs:     math.Round(float64(diskStat.Used)/1e9*100) / 100,
		},
		MemStat: MemStat{
			UsedPercent: math.Round(virtualMemory.UsedPercent*100) / 100,
			TotalGbs:    math.Round(float64(virtualMemory.Total)/1e9*100) / 100,
			UsedGbs:     math.Round(float64(virtualMemory.Used)/1e9*100) / 100,
		},
		HostInfo: HostInfo{
			CpuCores:             coreCount,
			CpuThreads:           threadCount,
			CpuName:              cpuInfos[0].ModelName,
			CpuMhz:               cpuInfos[0].Mhz,
			FSType:               diskStat.Fstype,
			HostName:             hostInfo.Hostname,
			Uptime:               hostInfo.Uptime,
			BootTime:             hostInfo.BootTime,
			Processes:            hostInfo.Procs,
			OS:                   hostInfo.OS,
			Platform:             hostInfo.Platform,
			PlatformFamily:       hostInfo.PlatformFamily,
			PlatformVersion:      hostInfo.PlatformVersion,
			KernelVersion:        hostInfo.KernelVersion,
			KernelArch:           hostInfo.KernelArch,
			VirtualizationSystem: hostInfo.VirtualizationSystem,
			VirtualizationRole:   hostInfo.VirtualizationRole,
		},
	}

	return response
}

func getStatsHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := fetchStats(ctx)

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "80"
	}

	env := common.EnvMap{}
	if proc := os.Getenv("HOST_PROC"); proc != "" {
		env[common.HostProcEnvKey] = proc
	}
	if sys := os.Getenv("HOST_SYS"); sys != "" {
		env[common.HostSysEnvKey] = sys
	}
	ctx := context.WithValue(context.Background(), common.EnvKey, env)

	http.HandleFunc("/api/stats", getStatsHandler(ctx))

	fs := http.FileServer(http.Dir("frontend/dist"))
	http.Handle("/", fs)

	slog.Info("Starting server", "port", PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
