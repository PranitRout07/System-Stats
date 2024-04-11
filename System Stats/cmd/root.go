package cmd 

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"
)

var rootCmd = &cobra.Command{
	Use: "ScrapeStats",
	Short: "This app will collect the system stats.",
	Long: "This cli app collects the system statistics like CPU stats , Memory Usage stats , Disk stats",
}

func Execute(){
	if err:= rootCmd.Execute(); err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
}