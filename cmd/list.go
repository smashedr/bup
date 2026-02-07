package cmd

import (
	"fmt"
	"github.com/charmbracelet/log"
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
		log.Debug("listCmd:", "args", args)

		destination := viper.GetString("destination")

		if len(args) == 0 {
			fmt.Printf("Destination: %s\n", destination)
			fmt.Println("--- All Backups ---")
			listDir(destination)
		} else {
			for i := 0; i < len(args); i++ {
				path := filepath.Join(destination, args[i])
				fmt.Printf("Destination: %s\n", path)
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
}
