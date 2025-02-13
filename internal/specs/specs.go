
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

func GetHostInfo() (string, error) {
	out := ""
	hostInfo, err := host.Info()
	if err != nil {
		return out, err
	}

	osys := strings.Title(hostInfo.OS)
	platform := strings.Title(hostInfo.Platform)
	arch := hostInfo.KernelArch
	out += fmt.Sprintf(`<li hx-swap-oob="innerHTML:#os">OS: %s %s %s</li>`, platform, osys, arch)

	uptime := hostInfo.Uptime
	duration := time.Duration(uptime) * time.Second
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60
	out += fmt.Sprintf(`<li hx-swap-oob="innerHTML:#uptime">Uptime: %d hours, %d minutes</li>`, hours, minutes)

	return out, nil
}

func GetCpuInfo() (string, error) {
	out := ""
	cpuInfo, err := cpu.Info()
	if err != nil {
		return out, err
	}

	out += fmt.Sprintf(`<li hx-swap-oob="innerHTML:#cpu">CPU: %s</li>`, cpuInfo[0].ModelName)

	return out, nil
}

func GetMemInfo() (string, error) {
	out := ""
	mem, err := mem.VirtualMemory()
	if err != nil {
		return out, err
	}

	used := mem.Used / factor
	total := mem.Total / factor
	out += fmt.Sprintf(`<li hx-swap-oob="innerHTML:#mem">Memory: %dMB / %dMB</li>`, used, total)
	
	return out, nil
}
