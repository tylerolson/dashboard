package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"math"
	"net/http"
	"os"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v4/common"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/sensors"
)

type Config map[string]bool

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
	CpuStat  CpuStat                   `json:"cpuStat"`
	DiskStat DiskStat                  `json:"diskStat"`
	MemStat  MemStat                   `json:"memStat"`
	HostInfo HostInfo                  `json:"hostInfo"`
	TempStat []sensors.TemperatureStat `json:"tempStat"`
}

func fetchStats(ctx context.Context, config Config) StatsResponse {
	cpuInfos, _ := cpu.InfoWithContext(ctx)
	coreCount, _ := cpu.CountsWithContext(ctx, false)
	threadCount, _ := cpu.CountsWithContext(ctx, true)
	cpuPercentages, _ := cpu.PercentWithContext(ctx, time.Second, false)
	diskStat, _ := disk.UsageWithContext(ctx, "/")
	virtualMemory, _ := mem.VirtualMemoryWithContext(ctx)
	hostInfo, _ := host.InfoWithContext(ctx)

	tempStat, _ := sensors.TemperaturesWithContext(ctx)

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

	tempToUse := make([]sensors.TemperatureStat, 0)
	for _, v := range tempStat {
		if config[v.SensorKey] {
			tempToUse = append(tempToUse, v)
		}
	}

	slices.SortFunc(tempToUse, func(a, b sensors.TemperatureStat) int {
		return strings.Compare(strings.ToLower(a.SensorKey), strings.ToLower(b.SensorKey))
	})

	response.TempStat = tempToUse

	return response
}

func getStatsHandler(ctx context.Context, config Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := fetchStats(ctx, config)

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func configNotExist() bool {
	_, err := os.Stat("config.json")
	return os.IsNotExist(err)
}

func readConfig() (Config, error) {
	fileData, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	config := make(Config)

	err = json.Unmarshal(fileData, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func createConfig(ctx context.Context) error {
	temps, _ := sensors.TemperaturesWithContext(ctx)

	sort.Slice(temps, func(i, j int) bool {
		return temps[i].SensorKey < temps[j].SensorKey
	})

	configMap := make(Config)
	for _, value := range temps {
		configMap[value.SensorKey] = false
	}

	jsonData, err := json.MarshalIndent(configMap, "", "  ")
	if err != nil {
		return err
	}

	filePath := "config.json"
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Set port
	PORT := os.Getenv("PORT")
	if PORT == "" {
		slog.Info("Could not find PORT, setting to 80")
		PORT = "80"
	} else {
		slog.Info("Found PORT " + PORT)
	}

	// Get host env variables
	env := common.EnvMap{}
	if proc := os.Getenv("HOST_PROC"); proc != "" {
		env[common.HostProcEnvKey] = proc
		slog.Info("Found HOST_PROC", "key", proc)
	}
	if sys := os.Getenv("HOST_SYS"); sys != "" {
		env[common.HostSysEnvKey] = sys
		slog.Info("Found HOST_SYS", "key", sys)
	}
	if etc := os.Getenv("HOST_ETC"); etc != "" {
		env[common.HostEtcEnvKey] = etc
		slog.Info("Found HOST_ETC", "key", etc)
	}
	ctx := context.WithValue(context.Background(), common.EnvKey, env)

	if configNotExist() {
		slog.Info("Config does not exist, creating it now")

		err := createConfig(ctx)
		if err != nil {
			slog.Error("error creating config", "error", err)
		}

		slog.Info("Created config")
	}

	config, err := readConfig()
	if err != nil {
		slog.Error("error reading config", "error", err)
	}

	slog.Info("Found config")

	http.HandleFunc("/api/stats", getStatsHandler(ctx, config))

	fs := http.FileServer(http.Dir("frontend/dist"))
	http.Handle("/", fs)

	slog.Info("Starting server", "port", PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
