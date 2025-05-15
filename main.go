package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func statsHandler(w http.ResponseWriter, r *http.Request) {
	cpuInfos, _ := cpu.Info()
	cpuPercentages, _ := cpu.Percent(time.Second, false)
	virtualMemory, _ := mem.VirtualMemory()
	hostInfo, _ := host.Info()
	diskStat, _ := disk.Usage("/")

	type CPUInfo struct {
		CPUPercent float64 `json:"cpuPercent"`
		Cores      int32   `json:"cores"`
		ModelName  string  `json:"modelName"`
		MHZ        float64 `json:"mhz"`
	}

	type MemInfo struct {
		AvailableBytes uint64  `json:"availableBytes"`
		TotalBytes     uint64  `json:"totalBytes"`
		UsedBytes      uint64  `json:"usedBytes"`
		UsedPercent    float64 `json:"usedPercent"`
	}

	type HostInfo struct {
		HostName             string `json:"hostName"`
		Uptime               uint64 `json:"uptime"`
		BootTime             uint64 `json:"bootTime"`
		Processes            uint64 `json:"processes"`
		OS                   string `json:"os"`
		Platform             string `json:"platform"`
		PlatformFamily       string `json:"platformFamily"`
		PlatformVersion      string `json:"platformVersion"`
		KernelVersion        string `json:"kernelVersion"`
		KernelArch           string `json:"kernelArch"`
		VirtualizationSystem string `json:"virtualizationSystem"`
		VirtualizationRole   string `json:"virtualizationRole"`
	}

	type DiskInfo struct {
		FSType      string  `json:"fsType"`
		TotalBytes  uint64  `json:"totalBytes"`
		FreeBytes   uint64  `json:"freeBytes"`
		UsedBytes   uint64  `json:"usedBytes"`
		UsedPercent float64 `json:"usedPercent"`
	}

	type StatsResponse struct {
		CPUInfo  CPUInfo  `json:"cpuInfo"`
		MemInfo  MemInfo  `json:"memInfo"`
		HostInfo HostInfo `json:"hostInfo"`
		Disk     DiskInfo `json:"diskInfo"`
	}

	response := StatsResponse{
		CPUInfo: CPUInfo{
			CPUPercent: cpuPercentages[0],
			Cores:      cpuInfos[0].Cores,
			ModelName:  cpuInfos[0].ModelName,
			MHZ:        cpuInfos[0].Mhz,
		},
		MemInfo: MemInfo{
			AvailableBytes: virtualMemory.Available,
			TotalBytes:     virtualMemory.Total,
			UsedBytes:      virtualMemory.Used,
			UsedPercent:    virtualMemory.UsedPercent,
		},
		HostInfo: HostInfo{
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
		Disk: DiskInfo{
			FSType:      diskStat.Fstype,
			TotalBytes:  diskStat.Total,
			FreeBytes:   diskStat.Free,
			UsedBytes:   diskStat.Used,
			UsedPercent: diskStat.UsedPercent,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	prettyData, err := json.MarshalIndent(response, "", "   ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(string(prettyData))
}

func main() {
	http.HandleFunc("/api/stats", statsHandler)

	fs := http.FileServer(http.Dir("frontend/dist"))
	http.Handle("/", fs)

	fmt.Println("Serving on http://localhost:8080")
	slog.Error("error starting server: ", http.ListenAndServe(":8080", nil))
}
