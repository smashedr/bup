package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var infoCmd = &cobra.Command{
	Use:     "info [name]",
	Aliases: []string{"i", "in", "inf"},
	Short:   "Show information about application",
	Long:    "Show information about application.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("--------------------\n")
		fmt.Printf("args: %s\n", args)
		fmt.Printf("cfgFile: %s\n", cfgFile)
		destination := viper.GetString("destination")
		fmt.Printf("destination: %s\n", destination)
		fmt.Printf("excludes: %s\n", viper.GetStringSlice("excludes"))
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
