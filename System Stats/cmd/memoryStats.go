package cmd 
import (
	"fmt"
	"github.com/spf13/cobra"
)

var memoryCmd = &cobra.Command{
	Use: "mem",
	Short: "This command will scrape the memory usage stats.",
	Run: func(cmd *cobra.Command, args []string){
		memoryTotal,memoryAvailable,memoryUsed,memoryFree,memoryUsedPercent := GetMemoryUsageStats()
		fmt.Println("Total Memory: ",memoryTotal)
		fmt.Println("Available Memory: ",memoryAvailable)
		fmt.Println("Used Memory: ",memoryUsed)
		fmt.Println("Free Memory: ",memoryFree)
		fmt.Println("Used Memory Percent: ",memoryUsedPercent)
	},
}

func init(){
	rootCmd.AddCommand(memoryCmd)
}