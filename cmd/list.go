package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var listCmd = &cobra.Command{
	Use:     "list [name]",
	Aliases: []string{"l", "li", "lis", "ls"},
	Short:   "List backups for a given name",
	Long:    "List all backups or backups for specified name.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("--------------------\n")
		//fmt.Printf("args: %s\n", args)
		destination := viper.GetString("destination")
		fmt.Printf("Destination: %s\n", destination)

		if len(args) == 0 {
			fmt.Println("--- All Backups ---")
			listDir(destination)
		} else {
			for i := 0; i < len(args); i++ {
				path := filepath.Join(destination, args[i])
				fmt.Printf("--- %s ---\n", args[i])
				listDir(path)
			}
		}
	},
}

func listDir(path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("Name Not Found")
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}
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
