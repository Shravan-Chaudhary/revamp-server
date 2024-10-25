package health

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemHealth struct {
    CPUUsage    float64 `json:"cpuUsage"`
    TotalMemory string  `json:"totalMemory"`
    FreeMemory  string  `json:"freeMemory"`
}

type MemoryUsage struct {
    HeapTotal string `json:"heapTotal"`
    HeapUsed  string `json:"heapUsed"`
}

type ApplicationHealth struct {
    Environment  string      `json:"environment"`
    UpTime      string      `json:"upTime"`
    MemoryUsage MemoryUsage `json:"memoryUsage"`
}

type HealthData struct {
    Application ApplicationHealth `json:"application"`
    System      SystemHealth     `json:"system"`
    Timestamp   int64           `json:"timestamp"`
}

var startTime time.Time

func init() {
    startTime = time.Now()
}

func getSystemHealth() (SystemHealth, error) {
    v, err := mem.VirtualMemory()
    if err != nil {
        return SystemHealth{}, err
    }

    cpuPercent, err := cpu.Percent(0, false)
    if err != nil {
        return SystemHealth{}, err
    }

    return SystemHealth{
        CPUUsage:    cpuPercent[0],
        TotalMemory: fmt.Sprintf("%.0f MB", float64(v.Total)/(1024*1024)),
        FreeMemory:  fmt.Sprintf("%.0f MB", float64(v.Free)/(1024*1024)),
    }, nil
}

func getApplicationHealth(env string) ApplicationHealth {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)

    return ApplicationHealth{
        Environment: env,
        UpTime:     fmt.Sprintf("%.2f Seconds", time.Since(startTime).Seconds()),
        MemoryUsage: MemoryUsage{
            HeapTotal: fmt.Sprintf("%.0f MB", float64(m.HeapSys)/(1024*1024)),
            HeapUsed:  fmt.Sprintf("%.0f MB", float64(m.HeapAlloc)/(1024*1024)),
        },
    }
}

// HealthCheckHandler returns a Gin handler function for health checks
func HealthCheck(env string) (HealthData, error) {
        sysHealth, err := getSystemHealth()
        if err != nil {
            return HealthData{}, err
        }

        healthData := HealthData{
            Application: getApplicationHealth(env),
            System:     sysHealth,
            Timestamp:  time.Now().UnixMilli(),
        }

		return healthData, nil
}