package cmd

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var excludeCmd = &cobra.Command{
	Use:     "exclude add/remove name",
	Aliases: []string{"e", "ex", "exc"},
	Short:   "Show, add or remove excludes",
	Long:    "Show, add or remove excludes.",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("excludeCmd:", "args", args)

		destination := viper.GetString("destination")
		fmt.Printf("destination: %s\n", destination)
		fmt.Printf("excludes: %s\n", viper.GetStringSlice("excludes"))

		log.Fatal("INOP - this command only list excludes. Update the Config File to edit them.")
	},
}

func init() {
	rootCmd.AddCommand(excludeCmd)
}
