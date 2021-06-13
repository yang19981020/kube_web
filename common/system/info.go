package system

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"time"
)

func GetCpuPercent() float64 {
	percent, _:= cpu.Percent(time.Second, false)
	return percent[0]
}
func GetMemPercent()float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}