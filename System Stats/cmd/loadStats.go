package cmd 
import (
	"fmt"
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use: "load",
	Short: "This command will show the load stats.",
	Run: func(cmd *cobra.Command, args []string){
		load1,load5,load15 := GetLoadStats()
		fmt.Println("Load1: ",load1)
		fmt.Println("Load5: ",load5)
		fmt.Println("Load15: ",load15)
	},
}

func init(){
	rootCmd.AddCommand(loadCmd)
}