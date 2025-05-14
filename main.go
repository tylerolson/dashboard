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
	cpuInfo, _ := cpu.Info()
	cpuPercent, _ := cpu.Percent(time.Second, false)
	memoryStat, _ := mem.VirtualMemory()
	hostInfo, _ := host.Info()
	diskStat, _ := disk.Usage("/")

	data := map[string]any{
		"cpu_info":    cpuInfo,
		"cpu_percent": cpuPercent,
		"memory": map[string]any{
			"total":        memoryStat.Total,
			"available":    memoryStat.Available,
			"used":         memoryStat.Used,
			"used_percent": memoryStat.UsedPercent,
		},
		"host_info": hostInfo,
		"disk":      diskStat,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/api/stats", statsHandler)

	fs := http.FileServer(http.Dir("frontend/dist"))
	http.Handle("/", fs)

	fmt.Println("Serving on http://localhost:8080")
	slog.Error("error starting server: ", http.ListenAndServe(":8080", nil))
}
