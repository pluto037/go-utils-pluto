package system

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"runtime"
	"time"
)

func Query() string {
	cpuInfo := getCPUInfo()
	memInfo := getMemoryInfo()
	systemInfo := getSystemInfo()

	result := fmt.Sprintf("CPU芯片信息：%s\nCPU核数：%d\nCPU每个核使用率：", cpuInfo.ModelName, cpuInfo.NumCPU)
	for i, usage := range cpuInfo.Usage {
		result += fmt.Sprintf("核%d: %.2f%%", i+1, usage)
		if i != len(cpuInfo.Usage)-1 {
			result += ", "
		}
	}
	result += fmt.Sprintf("\n内存大小：%.2f GB\n内存使用率：%.2f%%\n内存剩余：%.2f GB\n系统版本：%s",
		memInfo.TotalGB, memInfo.Usage, memInfo.FreeGB, systemInfo)
	return fmt.Sprintln(result)
}

type CPUInfo struct {
	ModelName string
	NumCPU    int
	Usage     []float64
}

type MemoryInfo struct {
	TotalGB float64
	Usage   float64
	FreeGB  float64
}

func getCPUInfo() CPUInfo {
	cpuInfo, _ := cpu.Info()
	numCPU := runtime.NumCPU()
	infos, _ := cpu.Percent(time.Second, true)

	return CPUInfo{
		ModelName: cpuInfo[0].ModelName,
		NumCPU:    numCPU,
		Usage:     infos,
	}
}

func getMemoryInfo() MemoryInfo {
	memInfo, _ := mem.VirtualMemory()

	totalGB := float64(memInfo.Total) / (1024 * 1024 * 1024)
	freeGB := float64(memInfo.Available) / (1024 * 1024 * 1024)
	usedGB := totalGB - freeGB
	usage := (usedGB / totalGB) * 100

	return MemoryInfo{
		TotalGB: totalGB,
		Usage:   usage,
		FreeGB:  freeGB,
	}
}

func getSystemInfo() string {
	return fmt.Sprintf("%s %s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())
}
