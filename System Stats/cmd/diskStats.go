package cmd 

import (
	"fmt"
	"github.com/spf13/cobra"
)

var diskCmd = &cobra.Command{
	Use: "disk",
	Short: "This command will show the disk stats.",
	Run: func(cmd *cobra.Command, args []string){
		diskUsageTotal, diskUsageFree , diskUsagePercent := GetDiskUsage()
		fmt.Println("Total Disk Usage: ",diskUsageTotal)
		fmt.Println("Free Disk: ",diskUsageFree)
		fmt.Println("Disk Used Percent: ",diskUsagePercent)
	},
}

func init(){
	rootCmd.AddCommand(diskCmd)
}