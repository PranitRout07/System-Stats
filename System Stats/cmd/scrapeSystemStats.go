package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func GetCpuUsage() []float64{
	usage, err := cpu.Percent(time.Second, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return usage
}

func GetDiskUsage() (uint64 ,uint64 ,float64){
	usage,err := disk.Usage("/")
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	return usage.Total,usage.Free,usage.UsedPercent
}

func GetMemoryUsageStats()(uint64,uint64,uint64,uint64,float64){
	stats,err := mem.VirtualMemory()
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	return stats.Total,stats.Available,stats.Used,stats.Free,stats.UsedPercent
}

func GetLoadStats() (float64,float64,float64){
	stats,err := load.Avg()
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	return stats.Load1,stats.Load5,stats.Load15
}