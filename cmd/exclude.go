package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var excludeCmd = &cobra.Command{
	Use:     "exclude add/remove [name]",
	Aliases: []string{"e", "ex", "exc"},
	Short:   "Show, add or remove excludes",
	Long:    "Show, add or remove excludes.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("--------------------\n")
		fmt.Printf("args: %s\n", args)
		fmt.Printf("cfgFile: %s\n", cfgFile)
		destination := viper.GetString("destination")
		fmt.Printf("destination: %s\n", destination)
		fmt.Printf("excludes: %s\n", viper.GetStringSlice("excludes"))
		fmt.Printf("INOP - this command only list excludes. Update the Config File to edit them.\n")
	},
}

func init() {
	rootCmd.AddCommand(excludeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
