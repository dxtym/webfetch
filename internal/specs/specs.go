package specs

import (
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

const (
	bytes  = 1024
	factor = bytes * bytes
)

type HostInfo struct {
	OS     string
	Uptime string
}

func GetHostInfo() (HostInfo, error) {
	var out HostInfo
	hostInfo, err := host.Info()
	if err != nil {
		return out, err
	}

	osys := strings.Title(hostInfo.OS)
	platform := strings.Title(hostInfo.Platform)
	arch := hostInfo.KernelArch
	out.OS = fmt.Sprintf("%s %s %s", platform, osys, arch)

	uptime := hostInfo.Uptime
	duration := time.Duration(uptime) * time.Second
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60
	out.Uptime = fmt.Sprintf("%d hours, %d minutes", hours, minutes)

	return out, nil
}

type CpuInfo struct {
	Cpu   string
}

func GetCpuInfo() (CpuInfo, error) {
	var out CpuInfo
	cpuInfo, err := cpu.Info()
	if err != nil {
		return out, err
	}

	out.Cpu = cpuInfo[0].ModelName

	return out, nil
}

type MemInfo struct {
	Mem string
}

func GetMemInfo() (MemInfo, error) {
	var out MemInfo
	mem, err := mem.VirtualMemory()
	if err != nil {
		return out, err
	}

	used := mem.Used / factor
	total := mem.Total / factor
	out.Mem = fmt.Sprintf("%dMB / %dMB", used, total)
	
	return out, nil
}
