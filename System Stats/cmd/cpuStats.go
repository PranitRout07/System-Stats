package cmd 

import (
	"fmt"
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat"
)
var cpuCmd = &cobra.Command{
	Use: "cpu",
	Short: "This command will scrape the CPU usage.",
	Run: func(cmd *cobra.Command, args []string){
		cpuUsage := stat.Mean(GetCpuUsage(),nil)
		fmt.Println("CPU Usage : ",cpuUsage)
	},
}

func init(){
	rootCmd.AddCommand(cpuCmd)
}