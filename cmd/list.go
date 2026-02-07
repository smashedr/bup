package cmd

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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
			listDir(destination, "All Backups")
		} else {
			for i := 0; i < len(args); i++ {
				path := filepath.Join(destination, args[i])
				listDir(path, args[i])
			}
		}
	},
}

func listDir(path, header string) {
	log.Infof("Destination: %s", path)
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Errorf("Error: %v", err)
		return
	}
	log.Debugf("entries: %v", entries)
	if len(entries) == 0 {
		log.Warnf("No backups for: %v", header)
		return
	}
	var rows [][]string
	for _, e := range entries {
		rows = append(rows, []string{e.Name()})
	}
	log.Debugf("rows: %v", rows)
	renderTable(rows, header)
}

func renderTable(rows [][]string, headers ...string) {
	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("6")).
		Bold(true).
		Align(lipgloss.Center)
	borderStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("4")).
		Bold(true)
	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(borderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return headerStyle
			}
			return lipgloss.NewStyle()
		}).
		Headers(headers...).
		Rows(rows...)
	fmt.Println(t)
}

func init() {
	rootCmd.AddCommand(listCmd)
}
