package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:     "list [name]",
	Aliases: []string{"l", "li", "lis", "ls"},
	Short:   "List backups or filter by name.",
	Long:    "List all backups or backups for specified name.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("args: %s\n", args)
		fmt.Printf("cfgFile: %s\n", cfgFile)
		fmt.Printf("destination: %s\n", viper.GetString("destination"))
		fmt.Printf("excludes: %s\n", viper.GetStringSlice("excludes"))
		fmt.Printf(" - List Command INOP\n")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
